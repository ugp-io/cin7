package cin7

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	APIKey          string
	APIAccountID    string
	Products        ProductsService
	Sales           SalesService
	Purchasing      PurchasingService
	StockAdjustment StockAdjustmentService
	Location        LocationService
}

func NewClient(apiKey, apiAccountID string) *Client {

	c := &Client{
		APIKey:       apiKey,
		APIAccountID: apiAccountID,
	}

	c.Products = &ProductsServiceOp{client: c}
	c.Sales = &SalesServiceOp{client: c}
	c.Purchasing = &PurchasingServiceOp{client: c}
	c.StockAdjustment = &StockAdjustmentServiceOp{client: c}
	c.Location = &LocationServiceOp{client: c}

	return c

}

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

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var errResp map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &errResp); err != nil {
			return fmt.Errorf("request failed with status %d: %s", res.StatusCode, string(bodyBytes))
		}
		return fmt.Errorf("request failed with status %d: %v", res.StatusCode, errResp)
	}

	return nil
}
