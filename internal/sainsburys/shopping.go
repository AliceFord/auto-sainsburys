package sainsburys

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/AliceFord/auto-sainsburys/internal/plan"
)

const (
	storeNumber = "0600"
	slotBooked  = false
)

func (c *Client) CheckAuth() error {
	// Check auth using https://www.sainsburys.co.uk/groceries-api/gol-services/customer/v1/customer/profile

	req, err := http.NewRequest(http.MethodGet, "https://www.sainsburys.co.uk/groceries-api/gol-services/customer/v1/customer/profile", nil)
	if err != nil {
		return err
	}

	// Add the cookies to the request
	for _, cookie := range c.Cookies {
		req.AddCookie(&http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}

	// Set WC_AUTHENTICATION cookie in the request header
	wcAuthToken, err := GetWCAuthToken(c.Cookies)
	if err != nil {
		return err
	}

	req.Header.Set("wcauthtoken", wcAuthToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) SearchProduct(productUid string) (Product, error) {
	values := url.Values{}
	values.Set("filter[keyword]", productUid)
	values.Set("page_number", "1")
	values.Set("page_size", "24")

	endpoint := "https://www.sainsburys.co.uk/groceries-api/gol-services/product/v1/product?" + values.Encode()

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return Product{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := c.Client.Do(req)
	if err != nil {
		return Product{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Product{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result productSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Product{}, err
	}

	// Search for product with matching productUid
	for _, product := range result.Products {
		if product.ProductUID == fmt.Sprintf("%s", productUid) {
			return product, nil
		}
	}

	return Product{}, fmt.Errorf("product with UID %s not found", productUid)
}

func (c *Client) ValidatePlan(p *plan.Plan) {
	for i := range p.Items {
		item := &p.Items[i]

		product, err := c.SearchProduct(item.ProductUID)
		if err != nil {
			item.ValidationStatus = plan.ValidationInvalid
			item.ValidationReason = fmt.Sprintf("error searching product: %v", err)
			continue
		}

		if !product.IsAvailable {
			item.ValidationStatus = plan.ValidationInvalid
			item.ValidationReason = "product is not available"
			continue
		}

		item.ValidationStatus = plan.ValidationValid
		item.ValidationReason = ""
	}
}

func (c *Client) AddItem(productUid string, quantity int) error {
	pickTime := time.Now().AddDate(0, 0, 1).Format("2006-01-02T15:04:05.000Z")

	body := AddItemRequest{
		ProductUID:          productUid,
		Quantity:            quantity,
		UOM:                 "ea",
		SelectedCatchweight: "",
	}

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	params := url.Values{}
	params.Set("store_number", storeNumber)
	params.Set("pick_time", pickTime)
	params.Set("slot_booked", strconv.FormatBool(slotBooked))

	endpoint := "https://www.sainsburys.co.uk/groceries-api/gol-services/basket/v2/basket/item?" + params.Encode()

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(data))
	if err != nil {
		return err
	}

	for _, cookie := range c.Cookies {
		req.AddCookie(&http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}

	wcAuthToken, err := GetWCAuthToken(c.Cookies)
	if err != nil {
		return err
	}

	req.Header.Set("wcauthtoken", wcAuthToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)

		return fmt.Errorf(
			"unexpected status code %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	return nil
}

func (c *Client) AddPlan(p *plan.Plan) error {
	for _, item := range p.Items {
		if item.ValidationStatus != plan.ValidationValid {
			return fmt.Errorf("cannot add invalid item %s to basket: %s", item.Name, item.ValidationReason)
		}

		if item.OrderQuantity <= 0 {
			return fmt.Errorf("cannot add item %s with non-positive order quantity %d to basket", item.Name, item.OrderQuantity)
		}

		err := c.AddItem(item.ProductUID, item.OrderQuantity)
		if err != nil {
			return fmt.Errorf("failed to add item %s to basket: %v", item.Name, err)
		}
	}

	return nil
}
