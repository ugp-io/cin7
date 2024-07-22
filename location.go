package cin7

import (
	"context"
	"encoding/json"
	"strings"
)

type LocationServiceOp struct {
	client *Client
}

type LocationService interface {
	BrowseLocation(ctx context.Context, req BrowseLocationRequest) (*LocationResponse, error)
}

func (s *LocationServiceOp) BrowseLocation(ctx context.Context, req BrowseLocationRequest) (*LocationResponse, error) {

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

	if req.Deprecated != nil {
		urlBuild = append(urlBuild, "Deprecated="+*req.Deprecated)
	}

	if req.Name != nil {
		urlBuild = append(urlBuild, "Name="+*req.Name)
	}

	productURL := url + `ref/location?` + strings.Join(urlBuild, "&")

	errRequest := s.client.Request("GET", productURL, nil, &reqResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	var response LocationResponse
	err := json.Unmarshal(reqResponse, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
