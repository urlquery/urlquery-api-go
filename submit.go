package urlquery

import (
	"fmt"
	"strings"
)

//-------------------------------
// Submit
//-------------------------------

func Submit(submit SubmitJob) (*QueuedJob, error) {
	return DefaultClient.Submit(submit)
}

func (c *Client) Submit(submit SubmitJob) (*QueuedJob, error) {

	endpoint := "/public/v1/submit/url"
	resp, err := c.DoRequest("POST", endpoint, strings.NewReader(submit.String()))
	if err != nil {
		return nil, err
	}

	data := new(QueuedJob)
	return data, DecodeResponse(resp, data)
}

//-------------------------------
// GetQueueStatus
//-------------------------------

func GetQueueStatus(queue_id string) (*QueuedJob, error) {
	return DefaultClient.GetQueueStatus(queue_id)
}

func (c *Client) GetQueueStatus(queue_id string) (*QueuedJob, error) {

	endpoint := fmt.Sprintf("/public/v1/submit/status/%s", queue_id)
	resp, err := c.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	data := new(QueuedJob)
	return data, DecodeResponse(resp, data)
}
