package models

import "encoding/json"

// Verdict can have the following values
//    malware, phishing, fraud, suspicoius

type ReputationResult struct {
	Url     string `json:"url"`
	Verdict string `json:"verdict"`
}

func (r ReputationResult) String() string {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
