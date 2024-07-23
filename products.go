package cin7

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type ProductsServiceOp struct {
	client *Client
}

type ProductsService interface {
	BrowseProducts(ctx context.Context, req BrowseProductRequest) (*ProductResponse, error)
	BrowseProductsAvailability(ctx context.Context, req BrowseProductAvailabilityRequest) (*ProductAvailabilityResponse, error)
	CreateProduct(ctx context.Context, req CreateProduct) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, req CreateProduct) (*ProductResponse, error)
}

func (s *ProductsServiceOp) BrowseProducts(ctx context.Context, req BrowseProductRequest) (*ProductResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	if req.ID != nil {
		urlBuild = append(urlBuild, "ID="+*req.ID)
	}

	if req.Page != nil {
		urlBuild = append(urlBuild, "Page="+*req.Page)
	}

	if req.Limit != nil {
		urlBuild = append(urlBuild, "Limit="+*req.Limit)
	}

	if req.Name != nil {
		urlBuild = append(urlBuild, "Name="+*req.Name)
	}

	if req.Sku != nil {
		urlBuild = append(urlBuild, "Sku="+*req.Sku)
	}

	if req.ModifiedSince != nil {
		urlBuild = append(urlBuild, "ModifiedSince="+*req.ModifiedSince)
	}

	if req.IncludeDeprecated != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("IncludeDeprecated=%v", *req.IncludeDeprecated))
	}

	if req.IncludeBOM != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("IncludeBOM=%v", *req.IncludeBOM))
	}

	if req.IncludeSuppliers != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("IncludeSuppliers=%v", *req.IncludeSuppliers))
	}

	if req.IncludeMovements != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("IncludeMovements=%v", *req.IncludeMovements))
	}

	if req.IncludeAttachments != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("IncludeAttachments=%v", *req.IncludeAttachments))
	}

	if req.IncludeReorderLevels != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("IncludeReorderLevels=%v", *req.IncludeReorderLevels))
	}

	if req.IncludeCustomPrices != nil {
		urlBuild = append(urlBuild, fmt.Sprintf("IncludeCustomPrices=%v", *req.IncludeCustomPrices))
	}
	productURL := url + `product?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", productURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response ProductResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ProductsServiceOp) CreateProduct(ctx context.Context, req CreateProduct) (*ProductResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	productURL := url + `product?` + strings.Join(urlBuild, "&")

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("POST", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response ProductResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ProductsServiceOp) UpdateProduct(ctx context.Context, req CreateProduct) (*ProductResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	productURL := url + `product?` + strings.Join(urlBuild, "&")

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("PUT", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response ProductResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ProductsServiceOp) BrowseProductsAvailability(ctx context.Context, req BrowseProductAvailabilityRequest) (*ProductAvailabilityResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	if req.Page != nil {
		urlBuild = append(urlBuild, "Page="+*req.Page)
	}

	if req.Limit != nil {
		urlBuild = append(urlBuild, "Limit="+*req.Limit)
	}

	if req.ID != nil {
		urlBuild = append(urlBuild, "ID="+*req.ID)
	}

	if req.Name != nil {
		urlBuild = append(urlBuild, "Name="+*req.Name)
	}

	if req.Sku != nil {
		urlBuild = append(urlBuild, "Sku="+*req.Sku)
	}

	if req.Location != nil {
		urlBuild = append(urlBuild, "Location="+*req.Location)
	}

	if req.Batch != nil {
		urlBuild = append(urlBuild, "Batch="+*req.Batch)
	}

	if req.Category != nil {
		urlBuild = append(urlBuild, "Category="+*req.Category)
	}

	productURL := url + `ref/productavailability?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", productURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response ProductAvailabilityResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
