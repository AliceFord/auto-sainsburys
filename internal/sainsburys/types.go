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

type AddItemRequest struct {
	ProductUID          string `json:"product_uid"`
	Quantity            int    `json:"quantity"`
	UOM                 string `json:"uom"`
	SelectedCatchweight string `json:"selected_catchweight"`
}

type Product struct {
	ProductUID  string `json:"product_uid"`
	IsAvailable bool   `json:"is_available"`
}

type productSearchResponse struct {
	Products []Product `json:"products"`
}
