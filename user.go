package urlquery

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func GetUser() (*PublicUserInfo, error) {
	return NewDefaultRequest().GetUser()
}

func (a apiRequest) GetUser() (*PublicUserInfo, error) {
	var usr PublicUserInfo

	apiurl := fmt.Sprintf("https://%s/public/v1/user", a.server)
	data, err := apiRequestHandle("GET", apiurl, nil, a.key)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &usr)
	return &usr, nil
}

func SetUserNotifyWebhook(webhook string) error {
	return NewDefaultRequest().SetUserNotifyWebhook(webhook)
}

func (a apiRequest) SetUserNotifyWebhook(webhook string) error {

	apiurl := fmt.Sprintf("https://%s/public/v1/user/notify?webhook.url=%s", a.server, url.QueryEscape(webhook))
	_, err := apiRequestHandle("PATCH", apiurl, nil, a.key)
	if err != nil {
		return nil, err
	}

	return err
}

func SetUserNotifyWebhookEnable(enabled bool) error {
	return NewDefaultRequest().SetUserNotifyWebhookEnable(enabled)
}

func (a apiRequest) SetUserNotifyWebhookEnable(enabled string) error {

	apiurl := fmt.Sprintf("https://%s/public/v1/user/notify?webhook.enabled=%t", a.server, enabled)
	_, err := apiRequestHandle("PATCH", apiurl, nil, a.key)
	if err != nil {
		return nil, err
	}

	return err
}
