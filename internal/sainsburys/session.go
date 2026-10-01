package sainsburys

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/mxschmitt/playwright-go"
)

func WriteSessionToFile(cookies []playwright.Cookie) error {
	data, err := json.MarshalIndent(cookies, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("sainsburys-session.json", data, 0600)
	if err != nil {
		panic(err)
	}

	return nil
}

func ReadSessionFromFile() ([]playwright.Cookie, error) {
	data, err := os.ReadFile("sainsburys-session.json")
	if err != nil {
		return nil, err
	}

	var cookies []playwright.Cookie
	err = json.Unmarshal(data, &cookies)
	if err != nil {
		return nil, err
	}

	return cookies, nil
}

func GetWCAuthToken(cookies []playwright.Cookie) (string, error) {
	for _, cookie := range cookies {
		if strings.HasPrefix(cookie.Name, "WC_AUTHENTICATION_") {
			return cookie.Value, nil
		}
	}

	return "", errors.New("WC_AUTHENTICATION cookie not found")
}

func DeleteSessionFile() error {
	return os.Remove("sainsburys-session.json")
}
