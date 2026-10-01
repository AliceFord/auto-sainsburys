package sainsburys

import (
	"net/http"

	"github.com/mxschmitt/playwright-go"
)

type Client struct {
	Cookies []playwright.Cookie
	Client  *http.Client
}

func NewClient(cookies []playwright.Cookie) *Client {
	return &Client{
		Cookies: cookies,
		Client:  http.DefaultClient,
	}
}
