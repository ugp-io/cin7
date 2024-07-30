package cin7

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type PurchasingServiceOp struct {
	client *Client
}

type PurchasingService interface {
	BrowsePurchase(ctx context.Context, req BrowsePurchaseRequest) (*PurchaseListResponse, error)
	ReadPurchase(ctx context.Context, req ReadPurchaseRequest) (*PurchaseResponse, error)
	ReadPurchaseOrder(ctx context.Context, req ReadPurchaseOrderRequest) (*PurchaseOrderResponse, error)
	CreatePurchase(ctx context.Context, req CreatePurchase) (*PurchaseResponse, error)
	CreatePurchaseOrder(ctx context.Context, req CreatePurchaseOrder) (*PurchaseOrderResponse, error)
}

func (s *PurchasingServiceOp) BrowsePurchase(ctx context.Context, req BrowsePurchaseRequest) (*PurchaseListResponse, error) {

	var reqResponse []byte
	urlBuild := []string{}

	if req.Page != nil {
		urlBuild = append(urlBuild, "Page="+url.QueryEscape(*req.Page))
	}

	if req.Limit != nil {
		urlBuild = append(urlBuild, "Limit="+url.QueryEscape(*req.Limit))
	}

	if req.Search != nil {
		urlBuild = append(urlBuild, "Search="+url.QueryEscape(*req.Search))
	}

	if req.RequiredBy != nil {
		urlBuild = append(urlBuild, "RequiredBy="+url.QueryEscape(*req.RequiredBy))
	}

	if req.UpdatedSince != nil {
		urlBuild = append(urlBuild, "UpdatedSince="+url.QueryEscape(*req.UpdatedSince))
	}

	if req.OrderStatus != nil {
		urlBuild = append(urlBuild, "OrderStatus="+url.QueryEscape(*req.OrderStatus))
	}

	if req.RestockReceivedStatus != nil {
		urlBuild = append(urlBuild, "RestockReceivedStatus="+url.QueryEscape(*req.RestockReceivedStatus))
	}

	if req.InvoiceStatus != nil {
		urlBuild = append(urlBuild, "InvoiceStatus="+url.QueryEscape(*req.InvoiceStatus))
	}

	if req.CreditNoteStatus != nil {
		urlBuild = append(urlBuild, "CreditNoteStatus="+url.QueryEscape(*req.CreditNoteStatus))
	}

	if req.UnstockStatus != nil {
		urlBuild = append(urlBuild, "UnstockStatus="+url.QueryEscape(*req.UnstockStatus))
	}

	if req.Status != nil {
		urlBuild = append(urlBuild, "Status="+url.QueryEscape(*req.Status))
	}

	if req.DropShipTaskID != nil {
		urlBuild = append(urlBuild, "DropShipTaskID="+url.QueryEscape(*req.DropShipTaskID))
	}

	purchasingURL := requestURL + `purchaseList?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", purchasingURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response PurchaseListResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	bolB, _ := json.Marshal(response)
	fmt.Println(string(bolB))

	return &response, nil
}

func (s *PurchasingServiceOp) ReadPurchase(ctx context.Context, req ReadPurchaseRequest) (*PurchaseResponse, error) {

	var reqResponse []byte
	urlBuild := []string{
		"ID=" + req.ID,
	}

	if req.ID != "" {
		urlBuild = append(urlBuild, "ID="+url.QueryEscape(req.ID))
	}

	if req.CombineAdditionalCharges != nil {
		urlBuild = append(urlBuild, "CombineAdditionalCharges="+url.QueryEscape(*req.CombineAdditionalCharges))
	}

	purchasingURL := requestURL + `purchase?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", purchasingURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response PurchaseResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *PurchasingServiceOp) CreatePurchase(ctx context.Context, req CreatePurchase) (*PurchaseResponse, error) {

	var reqResponse []byte

	productURL := requestURL + `purchase`

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("POST", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response PurchaseResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *PurchasingServiceOp) ReadPurchaseOrder(ctx context.Context, req ReadPurchaseOrderRequest) (*PurchaseOrderResponse, error) {

	var reqResponse []byte
	urlBuild := []string{
		"TaskID=" + req.TaskID,
	}

	if req.TaskID != "" {
		urlBuild = append(urlBuild, "TaskID="+url.QueryEscape(req.TaskID))
	}

	if req.CombineAdditionalCharges != nil {
		urlBuild = append(urlBuild, "CombineAdditionalCharges="+url.QueryEscape(*req.CombineAdditionalCharges))
	}

	purchasingURL := requestURL + `purchase/order?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", purchasingURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response PurchaseOrderResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *PurchasingServiceOp) CreatePurchaseOrder(ctx context.Context, req CreatePurchaseOrder) (*PurchaseOrderResponse, error) {

	var reqResponse []byte

	productURL := requestURL + `purchase/order`

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("POST", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response PurchaseOrderResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
