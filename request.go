package cin7

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	APIAccountID     string
	APIKey           string
	AlternateAPIKeys []string
	Products         ProductsService
	Sales            SalesService
	Purchasing       PurchasingService
	StockAdjustment  StockAdjustmentService
	Location         LocationService
	Supplier         SupplierService
}

func NewClient(apiAccountID, apiKey string, AlternateAPIKeys []string) *Client {

	c := &Client{
		APIKey:           apiKey,
		APIAccountID:     apiAccountID,
		AlternateAPIKeys: AlternateAPIKeys,
	}

	c.Products = &ProductsServiceOp{client: c}
	c.Sales = &SalesServiceOp{client: c}
	c.Purchasing = &PurchasingServiceOp{client: c}
	c.StockAdjustment = &StockAdjustmentServiceOp{client: c}
	c.Location = &LocationServiceOp{client: c}
	c.Supplier = &SupplierServiceOp{client: c}

	return c

}

var requestTime time.Time

func (c *Client) Request(method string, url string, bodyJSON io.Reader, response *[]byte) error {

	httpReq, errNewRequest := http.NewRequest(method, url, bodyJSON)
	if errNewRequest != nil {
		return errNewRequest
	}

	// Content Type
	httpReq.Header.Add("content-type", "application/json")
	httpReq.Header.Add("api-auth-accountid", c.APIAccountID)
	httpReq.Header.Add("api-auth-applicationkey", c.APIKey)

	client := &http.Client{}
	res, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	*response = bodyBytes

	if res.StatusCode == 503 {
		if requestTime.IsZero() {
			requestTime = time.Now()
		}
		if res, err := c.retryRequest(httpReq, response); err != nil {
			return err
		} else if res.StatusCode == 503 {
			waitTime := ((time.Second * 60) - time.Since(requestTime)) + time.Second
			// fmt.Printf("\t\t\tRetrying request after %.1f seconds: %v %v %v\n", waitTime.Seconds(), time.Since(requestTime), waitTime, exponentialBackoff(6))
			time.Sleep(waitTime)
			requestTime = time.Time{}
			// fmt.Println("\t\t\t60 seconds Done:", requestTime)
			return c.Request(method, url, bodyJSON, response)
		}
		// if res, err := c.retryRequest(method, url, bodyJSON, response, c.APIKeyBackup); err != nil {
		// 	return err
		// } else if res.StatusCode == 503 {
		// 	if res, err := c.retryRequest(method, url, bodyJSON, response, c.APIKeyBackup2); err != nil {
		// 		return err
		// 	} else if res.StatusCode == 503 {
		// 		fmt.Println("\t\t\tRetrying request after 60 seconds")
		// 		time.Sleep(exponentialBackoff(6))
		// 		return c.Request(method, url, bodyJSON, response)
		// 	} else {
		// 		fmt.Println("\t\tRetrying request Second Backup")
		// 	}
		// } else {
		// 	fmt.Println("\tRetrying request First Backup")
		// }
	} else if res.StatusCode >= 400 && res.StatusCode < 500 {
		return fmt.Errorf(string(bodyBytes))
	} else if res.StatusCode < 200 || res.StatusCode >= 300 {
		var errResp map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &errResp); err != nil {
			return fmt.Errorf("request failed with status %d: %s", res.StatusCode, string(bodyBytes))
		}
		return fmt.Errorf("request failed with status %d: %v", res.StatusCode, errResp)
	} else {
		// fmt.Println("First request Success")
	}

	return nil
}

// func (c *Client) retryRequest(method string, url string, bodyJSON io.Reader, response *[]byte, newAPIKey string) (*http.Response, error) {
func (c *Client) retryRequest(httpReq *http.Request, response *[]byte) (*http.Response, error) {

	// wordNumber := map[int]string{
	// 	0: "First",
	// 	1: "Second",
	// 	2: "Third",
	// 	3: "Fourth",
	// 	4: "Fifth",
	// }

	for index, newAPIKey := range c.AlternateAPIKeys {

		// Content Type
		// fmt.Printf("%vRetrying request %v Backup\n", strings.Repeat("\t", index+1), wordNumber[index])
		httpReq.Header.Del("api-auth-applicationkey")
		httpReq.Header.Add("api-auth-applicationkey", newAPIKey)

		client := &http.Client{}
		res, err := client.Do(httpReq)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if res.StatusCode == 503 {
			if index == len(c.AlternateAPIKeys)-1 {
				return res, nil
			}
			continue
		} else if res.StatusCode >= 400 && res.StatusCode < 500 {
			return nil, fmt.Errorf(string(bodyBytes))
		} else if res.StatusCode < 200 || res.StatusCode >= 300 {
			var errResp map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &errResp); err != nil {
				return nil, fmt.Errorf("request failed with status %d: %s", res.StatusCode, string(bodyBytes))
			}
			return nil, fmt.Errorf("request failed with status %d: %v", res.StatusCode, errResp)
		} else {
			*response = bodyBytes
			return res, nil
		}
	}
	return nil, nil
}
