package sainsburys

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/AliceFord/auto-sainsburys/internal/plan"
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

func (c *Client) SearchProduct(productUid int) (Product, error) {
	values := url.Values{}
	values.Set("filter[keyword]", strconv.Itoa(productUid))
	values.Set("page_number", "1")
	values.Set("page_size", "24")

	endpoint := "https://www.sainsburys.co.uk/groceries-api/gol-services/product/v1/products?" + values.Encode()

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
		if product.ProductUID == fmt.Sprintf("%d", productUid) {
			return product, nil
		}
	}

	return Product{}, fmt.Errorf("product with UID %d not found", productUid)
}

func (c *Client) ValidatePlan(p *plan.Plan) {
	for _, item := range p.Items {
		product, err := c.SearchProduct(item.SainsburysUid)
		if err != nil {
			item.ValidationStatus = plan.ValidationInvalid
			item.ValidationReason = fmt.Sprintf("error searching product: %v", err)
			continue
		}

		if !product.InStock {
			item.ValidationStatus = plan.ValidationInvalid
			item.ValidationReason = "product is out of stock"
			continue
		}

		item.ValidationStatus = plan.ValidationValid
	}
}
