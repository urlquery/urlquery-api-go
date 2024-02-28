package urlquery

import (
	"fmt"
	"net/url"
	"time"
)

func NewSearchParams(query string) *searchParams {
	s := searchParams{params: url.Values{}}
	s.params.Set("query", query)
	return &s
}

type searchParams struct {
	params url.Values
}

func (s *searchParams) Query(q string) {
	s.params.Set("query", q)
}

func (s *searchParams) Offset(offset int) {
	s.params.Set("offset", fmt.Sprintf("%d", offset))
}

func (s *searchParams) Limit(limit int) {
	s.params.Set("limit", fmt.Sprintf("%d", limit))
}

func (s *searchParams) StartTime(start time.Time) {
	s.params.Set("start_date", start.Format("2006-01-02"))
}
func (s *searchParams) StartEnd(end time.Time) {
	s.params.Set("end_date", end.Format("2006-01-02"))
}

func (p *searchParams) Encode() string {
	return p.params.Encode()
}

//-------------------------------
// Search Reports
//-------------------------------

func Search(queryParams searchParams) (*SearchResponse, error) {
	return DefaultClient.Search(queryParams)
}

func (c *Client) Search(queryParams searchParams) (*SearchResponse, error) {

	endpoint := fmt.Sprintf("/public/v1/search/reports/?%s", queryParams.Encode())
	resp, err := c.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	data := new(SearchResponse)
	return data, DecodeResponse(resp, data)
}
