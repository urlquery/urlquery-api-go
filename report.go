package urlquery

import (
	"encoding/json"
	"fmt"
)

func GetReport(report_id string) (*Report, error) {
	return NewDefaultRequest().GetReport(report_id)
}

func (api apiRequest) GetReport(report_id string) (*Report, error) {
	var reply Report

	url := fmt.Sprintf("https://%s/public/v1/report/%s", api.server, report_id)

	data, err := apiRequestHandle("GET", url, nil, api.key)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &reply)
	return &reply, err
}
