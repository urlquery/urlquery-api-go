package models

type UrlqueryAlert struct {
	SensorName string   `json:"sensor_name"`
	Alert      string   `json:"alert"`
	Verdict    string   `json:"verdict"`
	Severity   string   `json:"severity"`
	Comment    string   `json:"comment"`
	Tags       []string `json:"tags"`
}

// --- IDS ---

type IDSSensor struct {
	SensorName  string `json:"sensor_name"`
	Description string `json:"description"`

	Alerts []IDSAlert `json:"alerts"`
}

type IDSAlert struct {
	SensorName string `json:"sensor_name"`
	Date       string `json:"date"`
	Timestamp  int    `json:"timestamp"`
	IpDst      IP     `json:"ip_dst"`
	IpSrc      IP     `json:"ip_src"`
	Severity   string `json:"severity"`
	Alert      string `json:"alert"`
}

// --- Analyzer ---

type AnalyzerSensor struct {
	SensorName  string `json:"sensor_name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Link        string `json:"link"`

	Alerts []AnalyzerAlert `json:"alerts"`
}

type AnalyzerAlert struct {
	SensorName  string `json:"sensor_name"`
	SensorType  string `json:"sensor_type"`
	Description string `json:"description"`

	ScanDate     string             `json:"scan_date"`
	Alert        string             `json:"alert"`
	Trigger      string             `json:"trigger"`
	Verdict      string             `json:"verdict"`
	Severity     string             `json:"severity"`
	Comment      string             `json:"comment"`
	ResourceLink *string            `json:"link"`
	Meta         *map[string]string `json:"meta"`
}
