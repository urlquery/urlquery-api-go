package urlquery

import "encoding/json"

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

type QueuedJob struct {
	QueueID  string `json:"queue_id"`
	ReportID string `json:"report_id"`
	Status   string `json:"status"`

	Url URL `json:"url"`
	Ip  IP  `json:"ip"`

	UserAgent string `json:"useragent"`
	Referer   string `json:"referer"`
	ExitNode  string `json:"exit_node"`
	Owner     string `json:"owner"`
	Access    string `json:"access"`
	RunTime   int    `json:"run_time"`
}

func (j QueuedJob) String() string {
	b, _ := json.MarshalIndent(j, "", "  ")
	return string(b)
}

func (j QueuedJob) Bytes() []byte {
	b, _ := json.MarshalIndent(j, "", "  ")
	return b
}

type SubmitJob struct {
	Url  string            `json:"url"`
	Tags []string          `json:"tags"`
	Meta map[string]string `json:"meta"`

	UserAgent string `json:"useragent"`
	Referer   string `json:"referer"`
	ExitNode  string `json:"exit_node"`
	Access    string `json:"access"` // public, protected, private
}

func (j SubmitJob) String() string {
	b, _ := json.MarshalIndent(j, "", "  ")

	return string(b)
}

func (j SubmitJob) Bytes() []byte {
	b, _ := json.MarshalIndent(j, "", "  ")

	return b
}

// RFC3339 - is the default time format with RFC3339Nano used sometimes (http transactions)
type Report struct {
	ReportOverview

	FileDetections []FileObservation `json:"files"`

	Sensors struct {
		NetworkSensors  []IDSSensor      `json:"ids"`
		AnalyzerSensors []AnalyzerSensor `json:"analyzer"`
		UrlQueryAlerts  []UrlqueryAlert  `json:"urlquery"`
	} `json:"sensors"`

	Javascript       JavaScriptCode    `json:"javascript"`
	HttpTransactions []HttpTransaction `json:"http"`
}

type ReportOverview struct {
	ID      string   `json:"report_id"`
	Version int      `json:"version"`
	Status  string   `json:"status"`
	Tags    []string `json:"tags"`
	Date    string   `json:"date"` // RFC3339 -  "2006-01-02T15:04:05Z07:00"

	Url URL `json:"url"`
	Ip  IP  `json:"ip"`

	Final struct {
		Url   URL    `json:"url"`
		Title string `json:"title"`
	} `json:"final"`

	Submit struct {
		Tags []string          `json:"tags"`
		Meta map[string]string `json:"meta"`
	} `json:"submit"`

	// Settings the report was run  with
	ReportSettings struct {
		UserAgent string            `json:"useragent"`
		Referer   string            `json:"referer"`
		Cookies   map[string]string `json:"cookies"` // Cookie[<domain>]<cookie string>
		Access    string            `json:"access"`
		ExitNode  string            `json:"exit_node"`
	} `json:"settings"`

	Stats struct {
		AlertCount struct {
			Ids      int `json:"ids"`
			Urlquery int `json:"urlquery"`
			Analyzer int `json:"analyzer"`
		} `json:"alert_count"`
	} `json:"stats"`

	Summary []ReportSummary `json:"summary"`
}

type ReportSummary struct {
	Fqdn string `json:"fqdn"`

	Ip IP `json:"ip"`

	DomainRegistered string `json:"domain_registered"`
	DomainRank       int    `json:"domain_rank"`
	FirstSeen        string `json:"first_seen"`
	LastSeen         string `json:"last_seen"`

	AlertCount   int `json:"alert_count"`
	RequestCount int `json:"request_count"`
	ReceivedData int `json:"received_data"`
	SentData     int `json:"sent_data"`

	Comment string   `json:"comment"`
	Tags    []string `json:"tags"`
}

type FileAnalyzer struct {
	SensorName  string `json:"sensor_name"`
	Description string `json:"description"`

	ScanDate string  `json:"scan_date"`
	Trigger  string  `json:"trigger"`
	Verdict  string  `json:"verdict"`
	Comment  string  `json:"comment"`
	Link     *string `json:"link"`
}

type FileObservation struct {
	Md5    string `json:"md5"`
	Sha1   string `json:"sha1"`
	Sha256 string `json:"sha256"`
	Sha512 string `json:"sha512"`
	Magic  string `json:"magic"`
	Size   int    `json:"size"`

	Url URL `json:"url"`
	Ip  IP  `json:"ip"`

	Alerts struct {
		AnalyzerAlerts []AnalyzerAlert `json:"analyzer"`
	} `json:"alerts"`
}

type JavaScriptCode struct {
	Script []JSSourceCode `json:"script"`
	Eval   []JSCode       `json:"eval"`
	Write  []JSCode       `json:"write"`
}

type JSSourceCode struct {
	Url URL `json:"url"`
	Ip  IP  `json:"ip"`

	IntroductionType string `json:"introduction_type"`
	IsInline         bool   `json:"is_inline"`
	JSCode
}

type JSCode struct {
	Md5    string `json:"md5"`
	Sha1   string `json:"sha1"`
	Sha256 string `json:"sha256"`
	Sha512 string `json:"sha512"`
	Size   int    `json:"size"`
	Data   string `json:"data"`

	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
	TimesSeen int    `json:"times_seen"`

	Alerts Alerts `json:"alerts"`
}

type HttpTransaction struct {
	Url                 URL    `json:"url"`
	Ip                  IP     `json:"ip"`
	IsNavigationRequest bool   `json:"is_navigation_request"`
	ResourceType        string `json:"resource_type"`
	RequestedBy         string `json:"requested_by"`

	Date      string `json:"date"`      // ISO8601 -  YYYY-MM-DDThh:mm:ss.sTZD
	Timestamp int64  `json:"timestamp"` // unix epoch

	HttpVersion   string            `json:"http_version"`
	SecurityState string            `json:"security_state"` // Can be "secure", "insecure", "broken"
	SecurityInfo  *HttpSecurityInfo `json:"security_info"`

	Request  HttpRequest  `json:"request"`
	Response HttpResponse `json:"response"`

	TotalTimeUsed int         `json:"time_used"` // Total time used by the reques/response in milliseconds
	Timings       HttpTimings `json:"timings"`
	Alerts        Alerts      `json:"alerts"`
}

type HttpRequest struct {
	Raw     string            `json:"raw"`
	Headers []HttpHeaderValue `json:"headers"`
	Cookies []HttpHeaderValue `json:"cookies"`

	Method string `json:"method"`
}

type HttpResponse struct {
	Raw     string            `json:"raw"`
	Headers []HttpHeaderValue `json:"headers"`
	Cookies []HttpHeaderValue `json:"cookies"`

	StatusCode string `json:"status_code"`
	StatusText string `json:"status_text"`

	Content HttpContent `json:"data"`
}

type HttpHeaderValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type HttpTimings struct {
	Blocked int `json:"blocked"` // Time spent in a queue waiting for a network connection. Use -1 if the timing does not apply to the current request.
	DNS     int `json:"dns"`     // DNS resolution time. The time required to resolve a host name. Use -1 if the timing does not apply to the current request.
	Connect int `json:"connect"` // Time required to create TCP connection. Use -1 if the timing does not apply to the current request.
	Send    int `json:"send"`    // Time required to send HTTP request to the server.
	Wait    int `json:"wait"`    // Waiting for a response from the server.
	Receive int `json:"receive"` // Time required to read entire response from the server (or cache).
	SSL     int `json:"ssl"`     // Time required for SSL/TLS negotiation. If this field is defined then the time is also included in the connect field (to ensure backward compatibility with HAR 1.1). Use -1 if the timing does not apply to the current request.
}

type HttpPostData struct {
	MimeType string `json:"mime_type"`

	// Note that text and params fields are mutually exclusive.
	Params []HttpPostDataParams `json:"params"`
	Text   string               `json:"text"`
}

type HttpPostDataParams struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
}

type HttpContent struct {
	Size     int    `json:"size"`
	MimeType string `json:"mime_type"`
	Magic    string `json:"magic"`
	Md5      string `json:"md5"`
	Sha1     string `json:"sha1"`
	Sha256   string `json:"sha256"`
	Sha512   string `json:"sha512"`
	Data     []byte `json:"data"`
}

type HttpCookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Path     string `json:"path"`
	Domain   string `json:"domain"`
	Expires  string `json:"expire"`
	HttpOnly bool   `json:"http_only"` // HAR - TRUE if the cookie is HTTP only
	Secure   bool   `json:"secure"`    // HAR - TRUE if the cookie was transmitted of ssl, otherwise false
}

type HttpSecurityInfo struct {
	CipherSuite   string   `json:"cipher_suite"`
	KeyGroupName  string   `json:"key_group_name"`
	SignatureName string   `json:"signature_name"`
	Protocol      string   `json:"protocol"`
	Cert          CertInfo `json:"cert"`
}

type CertInfo struct {
	Subject     SubjectInfo     `json:"subject"`
	Issuer      IssuerInfo      `json:"issuer"`
	Validity    ValidityInfo    `json:"validity"`
	Fingerprint FingerprintInfo `json:"fingerprint"`
}

type SubjectInfo struct {
	CommonName   string `json:"commonName"`
	Organization string `json:"organization"`
}
type IssuerInfo struct {
	CommonName   string `json:"commonName"`
	Organization string `json:"organization"`
}
type ValidityInfo struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
type FingerprintInfo struct {
	Sha1   string `json:"sha1"`
	Sha256 string `json:"sha256"`
}

type Alerts struct {
	IDSAlerts      []IDSAlert      `json:"ids"`
	AnalyzerAlerts []AnalyzerAlert `json:"analyzer"`
	UrlqueryAlerts []UrlqueryAlert `json:"urlquery"`
}

func (r *Report) String() string {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

func (r *Report) Bytes() []byte {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}

	return b
}

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

// User
type PublicUserInfo struct {
	ID       string `json:"id"`
	Created  string `json:"created"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Group    string `json:"group"`
	Notify   Notify `json:"notify"`
	ApiKey   string `json:"apikey"`
}

type Notify struct {
	Webhook struct {
		Enabled bool   `json:"enabled"`
		URL     string `json:"url"`
	} `json:"webhook"`
}
