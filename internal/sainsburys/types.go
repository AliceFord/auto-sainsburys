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

type Product struct {
	ProductUID string `json:"product_uid"`
	InStock    bool   `json:"in_stock"`
}

type productSearchResponse struct {
	Products []Product `json:"products"`
}
