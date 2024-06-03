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
	ReadPurchase(ctx context.Context, req ReadPurchaseRequest) (*PurchaseResponse, error)
	ReadPurchaseOrder(ctx context.Context, req ReadPurchaseOrderRequest) (*PurchaseOrderResponse, error)
	CreatePurchase(ctx context.Context, req CreatePurchase) (*PurchaseResponse, error)
	CreatePurchaseOrder(ctx context.Context, req CreatePurchaseOrder) (*PurchaseOrderResponse, error)
}

func (s *PurchasingServiceOp) ReadPurchase(ctx context.Context, req ReadPurchaseRequest) (*PurchaseResponse, error) {

	var reqResponse []byte
	urlBuild := []string{
		"ID=" + req.ID,
	}

	if req.CombineAdditionalCharges != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("CombineAdditionalCharges=%t", *req.CombineAdditionalCharges))
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
		urlBuild = append(urlBuild, fmt.Sprintf("CombineAdditionalCharges=%t", *req.CombineAdditionalCharges))
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
