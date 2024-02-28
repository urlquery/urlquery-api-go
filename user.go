package urlquery

import (
	"fmt"
	"net/url"
)

//-------------------------------
// GetUser
//-------------------------------

func GetUser() (*PublicUserInfo, error) {
	return DefaultClient.GetUser()
}

func (c *Client) GetUser() (*PublicUserInfo, error) {

	endpoint := "/public/v1/user"
	resp, err := c.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	reply := new(PublicUserInfo)
	return reply, DecodeResponse(resp, reply)
}

//-------------------------------
// SetUserNotify
//-------------------------------

func SetUserNotifyWebhook(webhook string) error {
	return DefaultClient.SetUserNotifyWebhook(webhook)
}

func (c *Client) SetUserNotifyWebhook(webhook string) error {

	endpoint := fmt.Sprintf("/public/v1/user/notify?webhook.url=%s", url.QueryEscape(webhook))
	resp, err := c.DoRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	return DecodeResponse(resp, nil)
}

func SetUserNotifyWebhookEnable(enabled bool) error {
	return DefaultClient.SetUserNotifyWebhookEnable(enabled)
}

func (c *Client) SetUserNotifyWebhookEnable(enabled bool) error {

	endpoint := fmt.Sprintf("/public/v1/user/notify?webhook.enabled=%t", enabled)

	resp, err := c.DoRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	return DecodeResponse(resp, nil)
}
