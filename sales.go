package cin7

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type SalesServiceOp struct {
	client *Client
}

type SalesService interface {
	BrowseSale(ctx context.Context, req BrowseSaleRequest) (*SaleListResponse, error)
	ReadSale(ctx context.Context, req ReadSaleRequest) (*SaleResponse, error)
	ReadSaleOrder(ctx context.Context, req ReadSaleOrderRequest) (*SaleOrderResponse, error)
	ReadSaleQuote(ctx context.Context, req ReadSaleQuoteRequest) (*SaleQuoteResponse, error)
	CreateSale(ctx context.Context, req CreateSale) (*SaleResponse, error)
	CreateSaleOrder(ctx context.Context, req CreateSaleOrder) (*SaleOrderResponse, error)
	CreateSaleQuote(ctx context.Context, req CreateSaleQuote) (*SaleQuoteResponse, error)
}

func (s *SalesServiceOp) BrowseSale(ctx context.Context, req BrowseSaleRequest) (*SaleListResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	if req.Page != nil {
		urlBuild = append(urlBuild, "Page="+url.QueryEscape(*req.Page))
	}

	if req.Limit != nil {
		urlBuild = append(urlBuild, "Limit="+url.QueryEscape(*req.Limit))
	}

	if req.Search != nil {
		urlBuild = append(urlBuild, "Search="+url.QueryEscape(*req.Search))
	}

	if req.CreatedSince != nil {
		urlBuild = append(urlBuild, "CreatedSince="+url.QueryEscape(*req.CreatedSince))
	}

	if req.UpdatedSince != nil {
		urlBuild = append(urlBuild, "UpdatedSince="+url.QueryEscape(*req.UpdatedSince))
	}

	if req.ShipBy != nil {
		urlBuild = append(urlBuild, "ShipBy="+url.QueryEscape(*req.ShipBy))
	}

	if req.QuoteStatus != nil {
		urlBuild = append(urlBuild, "QuoteStatus="+url.QueryEscape(*req.QuoteStatus))
	}

	if req.OrderStatus != nil {
		urlBuild = append(urlBuild, "OrderStatus="+url.QueryEscape(*req.OrderStatus))
	}

	if req.CombinedPickStatus != nil {
		urlBuild = append(urlBuild, "CombinedPickStatus="+url.QueryEscape(*req.CombinedPickStatus))
	}

	if req.CombinedPackStatus != nil {
		urlBuild = append(urlBuild, "CombinedPackStatus="+url.QueryEscape(*req.CombinedPackStatus))
	}

	if req.CombinedShippingStatus != nil {
		urlBuild = append(urlBuild, "CombinedShippingStatus="+url.QueryEscape(*req.CombinedShippingStatus))
	}

	if req.CombinedInvoiceStatus != nil {
		urlBuild = append(urlBuild, "CombinedInvoiceStatus="+url.QueryEscape(*req.CombinedInvoiceStatus))
	}

	if req.CreditNoteStatus != nil {
		urlBuild = append(urlBuild, "CreditNoteStatus="+url.QueryEscape(*req.CreditNoteStatus))
	}

	if req.ExternalID != nil {
		urlBuild = append(urlBuild, "ExternalID="+url.QueryEscape(*req.ExternalID))
	}

	if req.Status != nil {
		urlBuild = append(urlBuild, "Status="+url.QueryEscape(*req.Status))
	}

	if req.ReadyForShipping != nil {
		urlBuild = append(urlBuild, "ReadyForShipping="+url.QueryEscape(*req.ReadyForShipping))
	}

	if req.OrderLocationID != nil {
		urlBuild = append(urlBuild, "OrderLocationID="+url.QueryEscape(*req.OrderLocationID))
	}

	salesURL := requestURL + `saleList?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", salesURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response SaleListResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *SalesServiceOp) ReadSale(ctx context.Context, req ReadSaleRequest) (*SaleResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	if req.ID != "" {
		urlBuild = append(urlBuild, "ID="+url.QueryEscape(req.ID))
	}

	if req.CombineAdditionalCharges != nil {
		urlBuild = append(urlBuild, "CombineAdditionalCharges="+url.QueryEscape(*req.CombineAdditionalCharges))
	}

	if req.HideInventoryMovements != nil {
		urlBuild = append(urlBuild, "HideInventoryMovements="+url.QueryEscape(*req.HideInventoryMovements))
	}

	if req.IncludeTransactions != nil {
		urlBuild = append(urlBuild, "IncludeTransactions="+url.QueryEscape(*req.IncludeTransactions))
	}

	if req.CountryFormat != nil {
		urlBuild = append(urlBuild, "CountryFormat="+url.QueryEscape(*req.CountryFormat))
	}

	salesURL := requestURL + `sale?` + strings.Join(urlBuild, "&")

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
	var urlBuild []string

	if req.SaleID != "" {
		urlBuild = append(urlBuild, "SaleID="+url.QueryEscape(req.SaleID))
	}

	if req.CombineAdditionalCharges != nil {
		urlBuild = append(urlBuild, "CombineAdditionalCharges="+url.QueryEscape(*req.CombineAdditionalCharges))
	}

	if req.IncludeProductInfo != nil {
		urlBuild = append(urlBuild, "IncludeProductInfo="+url.QueryEscape(*req.IncludeProductInfo))
	}

	salesURL := requestURL + `sale/order?` + strings.Join(urlBuild, "&")

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

	salesURL := requestURL + `sale/quote?` + strings.Join(urlBuild, "&")

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

	productURL := requestURL + `sale`

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

	productURL := requestURL + `sale/order`

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

	productURL := requestURL + `sale/quote`

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
