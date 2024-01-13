package urlquery

import (
	"compress/gzip"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type apiRequest struct {
	key    string
	server string

	limit  int
	offset int
	start  *time.Time
	end    *time.Time
}

func NewDefaultRequest() apiRequest {
	var r apiRequest
	r.key = defaultKey
	r.server = defaultServer
	return r
}

func NewRequest(key string) apiRequest {
	r := NewDefaultRequest()
	r.key = key
	return r
}

func getReader(resp *http.Response) io.ReadCloser {
	var reader io.ReadCloser
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, _ = gzip.NewReader(resp.Body)
	} else {
		reader = resp.Body
	}
	return reader
}

func apiRequestHandle(method string, url string, body io.Reader, apikey string) ([]byte, error) {
	var data []byte
	var err error

	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("X-APIKEY", apikey)

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	retryClient.RetryWaitMin = 5 * time.Second
	client := retryClient.StandardClient() // *http.Client
	client.Timeout = 60 * time.Second

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err == nil && resp != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.New(resp.Header.Get("x-api-error"))
		}

		data, err = io.ReadAll(getReader(resp))
	}

	return data, err
}

func (api apiRequest) WithLimit(limit int) apiRequest {
	api.limit = limit
	return api
}

func (api apiRequest) WithOffset(offset int) apiRequest {
	api.offset = offset
	return api
}

func (api apiRequest) WithStartTime(start time.Time) apiRequest {
	api.start = &start
	return api
}

func (api apiRequest) WithEndTime(end time.Time) apiRequest {
	api.end = &end
	return api
}
