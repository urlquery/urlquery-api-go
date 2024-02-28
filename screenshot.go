package urlquery

import (
	"fmt"
	"io"
)

//-------------------------------
// GetScreenshot
//-------------------------------

func GetScreenshot(report_id string) ([]byte, error) {
	return DefaultClient.GetScreenshot(report_id)
}

func (c *Client) GetScreenshot(report_id string) ([]byte, error) {
	endpoint := fmt.Sprintf("/public/v1/report/%s/screenshot", report_id)
	resp, err := c.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	err = handleResponseError(resp)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return data, err
}

//-------------------------------
// GetDomainGraph
//-------------------------------

func GetDomainGraph(report_id string) ([]byte, error) {
	return DefaultClient.GetDomainGraph(report_id)
}

func (c *Client) GetDomainGraph(report_id string) ([]byte, error) {
	endpoint := fmt.Sprintf("/public/v1/report/%s/domain_graph", report_id)
	resp, err := c.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	err = handleResponseError(resp)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return data, err
}
