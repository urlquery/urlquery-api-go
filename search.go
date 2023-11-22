package urlquery

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func Search(query string) (*SearchResponse, error) {
	return NewDefaultRequest().Search(query)
}

func (api apiRequest) Search(query string) (*SearchResponse, error) {
	var reply SearchResponse
	var limitParam string
	var offsetParam string
	var startParam string
	var endParam string

	if api.limit > 0 {
		limitParam = fmt.Sprintf("&limit=%d", api.limit)
	}
	if api.offset > 0 {
		offsetParam = fmt.Sprintf("&offset=%d", api.offset)
	}
	if api.start != nil {
		startParam = fmt.Sprintf("&start_date=%s", api.start.Format("2006-01-02"))
	}
	if api.end != nil {
		endParam = fmt.Sprintf("&end_date=%s", api.end.Format("2006-01-02"))
	}

	url := fmt.Sprintf("https://%s/public/v1/search/reports/?query=%s%s%s%s%s", api.server, url.QueryEscape(query), limitParam, offsetParam, startParam, endParam)
	data, err := apiRequestHandle("GET", url, nil, api.key)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &reply)
	return &reply, nil
}
