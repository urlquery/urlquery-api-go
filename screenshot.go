package urlquery

import (
	"fmt"
)

func GetScreenshot(report_id string) ([]byte, error) {
	return NewDefaultRequest().GetScreenshot(report_id)
}

func (api apiRequest) GetScreenshot(report_id string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/public/v1/report/%s/screenshot", api.server, report_id)
	data, err := apiRequestHandle("GET", url, nil, api.key)
	return data, err
}

func GetDomainGraph(report_id string) ([]byte, error) {
	return NewDefaultRequest().GetDomainGraph(report_id)
}

func (api apiRequest) GetDomainGraph(report_id string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/public/v1/report/%s/domain_graph", api.server, report_id)
	data, err := apiRequestHandle("GET", url, nil, api.key)
	return data, err
}
