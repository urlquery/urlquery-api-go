package urlquery

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func ReputationCheck(query string) (*ReputationResult, error) {
	return NewDefaultRequest().ReputationCheck(query)
}

func (api apiRequest) ReputationCheck(query string) (*ReputationResult, error) {
	var r ReputationResult

	url := fmt.Sprintf("https://%s/public/v1/reputation/check/?query=%s", api.server, url.QueryEscape(query))
	data, err := apiRequestHandle("GET", url, nil, api.key)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &r)
	return &r, err
}
