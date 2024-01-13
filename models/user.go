package models

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
