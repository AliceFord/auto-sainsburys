package sainsburys

import (
	"fmt"
	"net/http"
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
