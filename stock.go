package cin7

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
)

type StockAdjustmentServiceOp struct {
	client *Client
}

type StockAdjustmentService interface {
	ReadStockAdjustment(ctx context.Context, taskID string) (*StockAdjustmentResponse, error)
	CreateStockAdjustments(ctx context.Context, req CreateStockAdjustment) (*StockAdjustmentResponse, error)
	UpdateStockAdjustments(ctx context.Context, req CreateStockAdjustment) (*StockAdjustmentResponse, error)

	ReadStockTake(ctx context.Context, taskID string) (*StockTakeResponse, error)
	CreateStockTake(ctx context.Context, req CreateStockTake) (*StockTakeResponse, error)
	UpdateStockTake(ctx context.Context, req CreateStockTake) (*StockTakeResponse, error)

	BrowseStockTransfer(ctx context.Context, req BrowseStockTransfer) (*StockTransferResponse, error)
	ReadStockTransfer(ctx context.Context, taskID string) (*StockTransfer, error)
	CreateStockTransfer(ctx context.Context, req CreateStockTransfer) (*StockTransfer, error)
}

func (s *StockAdjustmentServiceOp) ReadStockAdjustment(ctx context.Context, taskID string) (*StockAdjustmentResponse, error) {

	var reqResponse []byte

	productURL := url + `stockadjustment?TaskID=` + taskID

	errRequest := s.client.Request("GET", productURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response StockAdjustmentResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *StockAdjustmentServiceOp) CreateStockAdjustments(ctx context.Context, req CreateStockAdjustment) (*StockAdjustmentResponse, error) {

	var reqResponse []byte

	productURL := url + `stockadjustment`

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("POST", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response StockAdjustmentResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *StockAdjustmentServiceOp) UpdateStockAdjustments(ctx context.Context, req CreateStockAdjustment) (*StockAdjustmentResponse, error) {

	var reqResponse []byte

	productURL := url + `stockadjustment`

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("PUT", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response StockAdjustmentResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *StockAdjustmentServiceOp) ReadStockTake(ctx context.Context, taskID string) (*StockTakeResponse, error) {

	var reqResponse []byte

	productURL := url + `stocktake?TaskID=` + taskID

	errRequest := s.client.Request("GET", productURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response StockTakeResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *StockAdjustmentServiceOp) CreateStockTake(ctx context.Context, req CreateStockTake) (*StockTakeResponse, error) {

	var reqResponse []byte

	productURL := url + `stocktake`

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("POST", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response StockTakeResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *StockAdjustmentServiceOp) UpdateStockTake(ctx context.Context, req CreateStockTake) (*StockTakeResponse, error) {

	var reqResponse []byte

	productURL := url + `stocktake`

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("PUT", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response StockTakeResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *StockAdjustmentServiceOp) ReadStockTransfer(ctx context.Context, taskID string) (*StockTransfer, error) {

	var reqResponse []byte

	productURL := url + `stockTransfer?TaskID=` + taskID

	errRequest := s.client.Request("GET", productURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response StockTransfer
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *StockAdjustmentServiceOp) CreateStockTransfer(ctx context.Context, req CreateStockTransfer) (*StockTransfer, error) {

	var reqResponse []byte

	productURL := url + `stockTransfer`

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("POST", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response StockTransfer
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *StockAdjustmentServiceOp) BrowseStockTransfer(ctx context.Context, req BrowseStockTransfer) (*StockTransferResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	if req.Page != nil {
		urlBuild = append(urlBuild, "Page="+*req.Page)
	}

	if req.Limit != nil {
		urlBuild = append(urlBuild, "Limit="+*req.Limit)
	}

	if req.Status != nil {
		urlBuild = append(urlBuild, "Status="+*req.Status)
	}

	if req.Search != nil {
		urlBuild = append(urlBuild, "Search="+*req.Search)
	}

	productURL := url + `stockTransfer?` + strings.Join(urlBuild, "&")

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("GET", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response StockTransferResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
