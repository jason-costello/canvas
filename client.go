package canvas

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	authToken  string
	observerID string
}


func New(u *url.URL, authToken, observerID string) *Client {
	jar, _ := cookiejar.New(nil)

	hc := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           jar,
		Timeout:       10 * time.Second,
	}

	return &Client{
		httpClient: hc,
		baseURL:    u,
		authToken:  authToken,
		observerID: observerID,
	}

}


func (c *Client) GetBaseURL() *url.URL{
	return c.baseURL
}

func (c *Client) GetHttpClient() *http.Client{
	return c.httpClient
}
