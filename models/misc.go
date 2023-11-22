package models

type IP struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`

	ASN         int    `json:"asn"`
	AS          string `json:"as"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

type URL struct {
	Schema string `json:"schema"`
	Addr   string `json:"addr"`
	Fqdn   string `json:"fqdn"`
	Domain string `json:"domain"`
	TLD    string `json:"tld"`
}
