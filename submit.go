package urlquery

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Submit(submit SubmitJob) (*QueuedJob, error) {
	return NewDefaultRequest().Submit(submit)
}

func (a apiRequest) Submit(submit SubmitJob) (*QueuedJob, error) {
	var j QueuedJob

	url := fmt.Sprintf("https://%s/public/v1/submit/url", a.server)
	data, err := apiRequestHandle("POST", url, strings.NewReader(submit.String()), a.key)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &j)
	return &j, nil
}

func GetQueueStatus(queue_id string) (*QueuedJob, error) {
	return NewDefaultRequest().GetQueueStatus(queue_id)
}

func (a apiRequest) GetQueueStatus(queue_id string) (*QueuedJob, error) {
	var j QueuedJob

	url := fmt.Sprintf("https://%s/public/v1/submit/status/%s", a.server, queue_id)
	data, err := apiRequestHandle("GET", url, nil, a.key)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &j)
	return &j, err
}
