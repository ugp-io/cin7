package cin7

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
	"strings"
)

type ProductsServiceOp struct {
	client *Client
}

type ProductsService interface {
	BrowseProducts(ctx context.Context, req BrowseProductRequest) (*ProductResponse, error)
	BrowseProductsAvailability(ctx context.Context, req BrowseProductAvailabilityRequest) (*ProductAvailabilityResponse, error)
	CreateProduct(ctx context.Context, req CreateProduct) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, req EditProduct) (*ProductResponse, error)

	BrowseProductFamilies(ctx context.Context, req BrowseProductFamilyRequest) (*ProductFamilyResponse, error)
	UpdateProductFamily(ctx context.Context, req EditProductFamily) (*ProductFamilyResponse, error)
	CreateProductFamily(ctx context.Context, req CreateProductFamily) (*ProductFamilyResponse, error)
}

func (s *ProductsServiceOp) BrowseProducts(ctx context.Context, req BrowseProductRequest) (*ProductResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	if req.ID != nil {
		urlBuild = append(urlBuild, "ID="+url.QueryEscape(*req.ID))
	}

	if req.Page != nil {
		urlBuild = append(urlBuild, "Page="+url.QueryEscape(*req.Page))
	}

	if req.Limit != nil {
		urlBuild = append(urlBuild, "Limit="+url.QueryEscape(*req.Limit))
	}

	if req.Name != nil {
		urlBuild = append(urlBuild, "Name="+url.QueryEscape(*req.Name))
	}

	if req.Sku != nil {
		urlBuild = append(urlBuild, "Sku="+url.QueryEscape(*req.Sku))
	}

	if req.ModifiedSince != nil {
		urlBuild = append(urlBuild, "ModifiedSince="+url.QueryEscape(*req.ModifiedSince))
	}

	if req.IncludeDeprecated != nil {
		urlBuild = append(urlBuild, "IncludeDeprecated="+url.QueryEscape(*req.IncludeDeprecated))
	}

	if req.IncludeBOM != nil {
		urlBuild = append(urlBuild, "IncludeBOM="+url.QueryEscape(*req.IncludeBOM))
	}

	if req.IncludeSuppliers != nil {
		urlBuild = append(urlBuild, "IncludeSuppliers="+url.QueryEscape(*req.IncludeSuppliers))
	}

	if req.IncludeMovements != nil {
		urlBuild = append(urlBuild, "IncludeMovements="+url.QueryEscape(*req.IncludeMovements))
	}

	if req.IncludeAttachments != nil {
		urlBuild = append(urlBuild, "IncludeAttachments="+url.QueryEscape(*req.IncludeAttachments))
	}

	if req.IncludeReorderLevels != nil {
		urlBuild = append(urlBuild, "IncludeReorderLevels="+url.QueryEscape(*req.IncludeReorderLevels))
	}

	if req.IncludeCustomPrices != nil {
		urlBuild = append(urlBuild, "IncludeCustomPrices="+url.QueryEscape(*req.IncludeCustomPrices))
	}

	productURL := requestURL + `product?` + strings.Join(urlBuild, "&")

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

	productURL := requestURL + `product`

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

func (s *ProductsServiceOp) UpdateProduct(ctx context.Context, req EditProduct) (*ProductResponse, error) {

	var reqResponse []byte

	productURL := requestURL + `product`

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

	if req.ID != nil {
		urlBuild = append(urlBuild, "ID="+url.QueryEscape(*req.ID))
	}

	if req.Page != nil {
		urlBuild = append(urlBuild, "Page="+url.QueryEscape(*req.Page))
	}

	if req.Limit != nil {
		urlBuild = append(urlBuild, "Limit="+url.QueryEscape(*req.Limit))
	}

	if req.Name != nil {
		urlBuild = append(urlBuild, "Name="+url.QueryEscape(*req.Name))
	}

	if req.Sku != nil {
		urlBuild = append(urlBuild, "Sku="+url.QueryEscape(*req.Sku))
	}

	if req.Location != nil {
		urlBuild = append(urlBuild, "Location="+url.QueryEscape(*req.Location))
	}

	if req.Batch != nil {
		urlBuild = append(urlBuild, "Batch="+url.QueryEscape(*req.Batch))
	}

	if req.Category != nil {
		urlBuild = append(urlBuild, "Category="+url.QueryEscape(*req.Category))
	}

	productURL := requestURL + `ref/productavailability?` + strings.Join(urlBuild, "&")

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

func (s *ProductsServiceOp) BrowseProductFamilies(ctx context.Context, req BrowseProductFamilyRequest) (*ProductFamilyResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	if req.ID != nil && *req.ID != "" {
		urlBuild = append(urlBuild, "ID="+url.QueryEscape(*req.ID))
	}

	if req.Sku != nil && *req.Sku != "" {
		urlBuild = append(urlBuild, "Sku="+url.QueryEscape(*req.Sku))
	}

	if req.Page != nil {
		urlBuild = append(urlBuild, "Page="+url.QueryEscape(*req.Page))
	}

	if req.Limit != nil {
		urlBuild = append(urlBuild, "Limit="+url.QueryEscape(*req.Limit))
	}

	if req.Name != nil {
		urlBuild = append(urlBuild, "Name="+url.QueryEscape(*req.Name))
	}

	productURL := requestURL + `productFamily?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", productURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response ProductFamilyResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ProductsServiceOp) CreateProductFamily(ctx context.Context, req CreateProductFamily) (*ProductFamilyResponse, error) {

	var reqResponse []byte

	productURL := requestURL + `productFamily`

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("POST", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response ProductFamilyResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ProductsServiceOp) UpdateProductFamily(ctx context.Context, req EditProductFamily) (*ProductFamilyResponse, error) {

	var reqResponse []byte

	productURL := requestURL + `productFamily`

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	errRequest := s.client.Request("PUT", productURL, bytes.NewBuffer(reqBody), &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response ProductFamilyResponse
	err = json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
