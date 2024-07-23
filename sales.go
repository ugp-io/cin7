package cin7

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type SalesServiceOp struct {
	client *Client
}

type SalesService interface {
	ReadSale(ctx context.Context, req ReadSaleRequest) (*SaleResponse, error)
	ReadSaleOrder(ctx context.Context, req ReadSaleOrderRequest) (*SaleOrderResponse, error)
	ReadSaleQuote(ctx context.Context, req ReadSaleQuoteRequest) (*SaleQuoteResponse, error)
	CreateSale(ctx context.Context, req CreateSale) (*SaleResponse, error)
	CreateSaleOrder(ctx context.Context, req CreateSaleOrder) (*SaleOrderResponse, error)
	CreateSaleQuote(ctx context.Context, req CreateSaleQuote) (*SaleQuoteResponse, error)
}

func (s *SalesServiceOp) ReadSale(ctx context.Context, req ReadSaleRequest) (*SaleResponse, error) {

	var reqResponse []byte
	urlBuild := []string{
		"ID=" + req.ID,
	}

	if req.CombineAdditionalCharges != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("CombineAdditionalCharges=%v", *req.CombineAdditionalCharges))
	}

	if req.HideInventoryMovements != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("HideInventoryMovements=%v", *req.HideInventoryMovements))
	}

	if req.IncludeTransactions != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("IncludeTransactions=%v", *req.IncludeTransactions))
	}

	if req.CountryFormat != nil {
		urlBuild = append(urlBuild, "CountryFormat="+*req.CountryFormat)
	}

	salesURL := url + `sale?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", salesURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response SaleResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *SalesServiceOp) ReadSaleOrder(ctx context.Context, req ReadSaleOrderRequest) (*SaleOrderResponse, error) {

	var reqResponse []byte
	urlBuild := []string{
		"SaleID=" + req.SaleID,
	}

	if req.CombineAdditionalCharges != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("CombineAdditionalCharges=%v", *req.CombineAdditionalCharges))
	}

	if req.IncludeProductInfo != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("IncludeProductInfo=%v", *req.IncludeProductInfo))
	}

	salesURL := url + `sale/order?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", salesURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response SaleOrderResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return nil, err
	}

	return &response, nil
}

func (s *SalesServiceOp) ReadSaleQuote(ctx context.Context, req ReadSaleQuoteRequest) (*SaleQuoteResponse, error) {

	var reqResponse []byte
	urlBuild := []string{
		"SaleID=" + req.SaleID,
	}

	if req.CombineAdditionalCharges != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("CombineAdditionalCharges=%v", *req.CombineAdditionalCharges))
	}

	if req.IncludeProductInfo != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("IncludeProductInfo=%v", *req.IncludeProductInfo))
	}

	salesURL := url + `sale/quote?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", salesURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response SaleQuoteResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *SalesServiceOp) CreateSale(ctx context.Context, req CreateSale) (*SaleResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	productURL := url + `sale?` + strings.Join(urlBuild, "&")

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("POST", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response SaleResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *SalesServiceOp) CreateSaleOrder(ctx context.Context, req CreateSaleOrder) (*SaleOrderResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	productURL := url + `sale/order?` + strings.Join(urlBuild, "&")

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("POST", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response SaleOrderResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *SalesServiceOp) CreateSaleQuote(ctx context.Context, req CreateSaleQuote) (*SaleQuoteResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	productURL := url + `sale/quote?` + strings.Join(urlBuild, "&")

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("POST", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response SaleQuoteResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
