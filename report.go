package urlquery

import (
	"fmt"
)

//-------------------------------
// GetReport
//-------------------------------

func GetReport(report_id string) (*Report, error) {
	return DefaultClient.GetReport(report_id)
}

func (c *Client) GetReport(report_id string) (*Report, error) {
	endpoint := fmt.Sprintf("/public/v1/report/%s", report_id)
	resp, err := c.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	data := new(Report)
	return data, DecodeResponse(resp, data)
}

//-------------------------------
// DeleteReport
//-------------------------------

func DeleteReport(report_id string) error {
	return DefaultClient.DeleteReport(report_id)
}

func (c *Client) DeleteReport(report_id string) error {

	endpoint := fmt.Sprintf("/public/v1/report/%s", report_id)
	resp, err := c.DoRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	return DecodeResponse(resp, nil)
}
