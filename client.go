package urlquery

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
)

var (
	baseURL       = "https://api.urlquery.net"
	DefaultClient = NewClient("")
)

// Client represents the REST API client.
type Client struct {
	baseURL string
	client  *http.Client
	apiKey  string
}

func NewClient(key string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  key,
		client:  &http.Client{},
	}
}

func SetDefaultKey(apikey string) {
	DefaultClient.apiKey = apikey
}

func (c *Client) WithApiKey(key string) *Client {
	c.apiKey = key
	return c
}

func (c *Client) NewRequest(method string, path string, body io.Reader) (*http.Request, error) {
	url := c.baseURL + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("X-APIKEY", c.apiKey)

	return req, nil
}

// MakeRequest makes an HTTP request to the API.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

// DoRequest makes an HTTP request to the API.
func (c *Client) DoRequest(method string, path string, body io.Reader) (*http.Response, error) {

	req, err := c.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// DecodeResponse decodes the HTTP response body.
func DecodeResponse(resp *http.Response, target interface{}) error {

	err := handleResponseError(resp)
	if err != nil {
		return err
	}

	err = decodeBody(resp, target)
	if err != nil {
		target = nil
	}

	return err
}

func decodeBody(resp *http.Response, target interface{}) error {
	var err error

	if resp.Body == nil || target == nil {
		return nil // Nothing to decode
	}
	defer resp.Body.Close()

	// TODO: add support for decoding Content-Type
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		var reader *gzip.Reader
		reader, err = gzip.NewReader(resp.Body)
		if err == nil {
			defer reader.Close()
			err = json.NewDecoder(reader).Decode(target)
		}

	default:
		err = json.NewDecoder(resp.Body).Decode(target)
	}

	return err
}
