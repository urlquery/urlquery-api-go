package models

import "encoding/json"

type SearchResponse struct {
	Query     string           `json:"query"`
	TotalHits int              `json:"total_hits"`
	TimeUsed  string           `json:"timeused"`
	Limit     int              `json:"limit"`
	Offset    int              `json:"offset"`
	Reports   []ReportOverview `json:"reports"`
}

func (sr *SearchResponse) Bytes() []byte {
	data, _ := json.Marshal(sr)
	return data
}

func (sr *SearchResponse) String() string {
	return string(sr.Bytes())
}
