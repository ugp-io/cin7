package cin7

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"
)

type SupplierServiceOp struct {
	client *Client
}

type SupplierService interface {
	BrowseSuppliers(ctx context.Context, req BrowseSupplierRequest) (*SupplierResponse, error)
}

func (s *SupplierServiceOp) BrowseSuppliers(ctx context.Context, req BrowseSupplierRequest) (*SupplierResponse, error) {

	var reqResponse []byte
	var urlBuild []string

	if req.Page != nil {
		urlBuild = append(urlBuild, "Page="+url.QueryEscape(*req.Page))
	}

	if req.Limit != nil {
		urlBuild = append(urlBuild, "Limit="+url.QueryEscape(*req.Limit))
	}

	if req.ID != nil {
		urlBuild = append(urlBuild, "ID="+url.QueryEscape(*req.ID))
	}

	if req.Name != nil {
		urlBuild = append(urlBuild, "Name="+url.QueryEscape(*req.Name))
	}

	supplierURL := requestURL + `supplier?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", supplierURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response SupplierResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
