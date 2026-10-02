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
	SainID              string `json:"sainId"`
	SKU                 string `json:"sku"`
	UOM                 string `json:"uom"`
	Quantity            int    `json:"quantity"`
	SelectedCatchweight string `json:"selectedCatchweight"`
	StoreNumber         string `json:"storeNumber"`
	SlotBooked          bool   `json:"slotBooked"`
	PickTime            string `json:"pickTime"`
	IsBasketCreated     bool   `json:"isBasketCreated"`
}

type Product struct {
	ProductUID  string `json:"product_uid"`
	IsAvailable bool   `json:"is_available"`
}

type productSearchResponse struct {
	Products []Product `json:"products"`
}
