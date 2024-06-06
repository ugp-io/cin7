package cin7

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
		urlBuild = append(urlBuild, "Page="+*req.Page)
	}

	if req.Limit != nil {
		urlBuild = append(urlBuild, "Limit="+*req.Limit)
	}

	if req.Search != nil {
		urlBuild = append(urlBuild, "Search="+*req.Search)
	}

	if req.RequiredBy != nil {
		urlBuild = append(urlBuild, "RequiredBy="+*req.RequiredBy)
	}

	if req.UpdatedSince != nil {
		urlBuild = append(urlBuild, "UpdatedSince="+*req.UpdatedSince)
	}

	if req.OrderStatus != nil {
		urlBuild = append(urlBuild, "OrderStatus="+*req.OrderStatus)
	}

	if req.RestockReceivedStatus != nil {
		urlBuild = append(urlBuild, "RestockReceivedStatus="+*req.RestockReceivedStatus)
	}

	if req.InvoiceStatus != nil {
		urlBuild = append(urlBuild, "InvoiceStatus="+*req.InvoiceStatus)
	}

	if req.CreditNoteStatus != nil {
		urlBuild = append(urlBuild, "CreditNoteStatus="+*req.CreditNoteStatus)
	}

	if req.UnstockStatus != nil {
		urlBuild = append(urlBuild, "UnstockStatus="+*req.UnstockStatus)
	}

	if req.Status != nil {
		urlBuild = append(urlBuild, "Status="+*req.Status)
	}

	if req.DropShipTaskID != nil {
		urlBuild = append(urlBuild, "DropShipTaskID="+*req.DropShipTaskID)
	}

	purchasingURL := url + `purchaseList?` + strings.Join(urlBuild, "&")

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

	if req.CombineAdditionalCharges != nil {
		urlBuild = append(urlBuild, "CombineAdditionalCharges="+*req.CombineAdditionalCharges)
	}

	purchasingURL := url + `purchase?` + strings.Join(urlBuild, "&")

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
	var urlBuild []string

	productURL := url + `purchase?` + strings.Join(urlBuild, "&")

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

	if req.CombineAdditionalCharges != nil {
		urlBuild = append(urlBuild, "CombineAdditionalCharges="+*req.CombineAdditionalCharges)
	}

	purchasingURL := url + `purchase/order?` + strings.Join(urlBuild, "&")

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
	var urlBuild []string

	productURL := url + `purchase/order?` + strings.Join(urlBuild, "&")

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
