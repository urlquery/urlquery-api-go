package urlquery

import (
	"fmt"
	"net/url"
)

//-------------------------------
// CheckReputation
//-------------------------------

func ReputationCheck(query string) (*ReputationResult, error) {
	return DefaultClient.ReputationCheck(query)
}

func (c *Client) ReputationCheck(query string) (*ReputationResult, error) {

	endpoint := fmt.Sprintf("/public/v1/reputation/check/?query=%s", url.QueryEscape(query))
	resp, err := c.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var data ReputationResult
	return &data, DecodeResponse(resp, &data)
}
