package sainsburys

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mxschmitt/playwright-go"
)

func AuthLogin() ([]playwright.Cookie, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		return nil, err
	}

	ctx, err := browser.NewContext()
	if err != nil {
		return nil, err
	}

	page, err := ctx.NewPage()
	if err != nil {
		return nil, err
	}

	_, err = page.Goto(
		"https://www.sainsburys.co.uk/gol-ui/oauth/login",
		playwright.PageGotoOptions{
			WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		},
	)
	if err != nil {
		return nil, err
	}

	fmt.Println("Press enter in the terminal once login is complete...")

	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')

	cookies, err := ctx.Cookies()
	if err != nil {
		return nil, err
	}

	return cookies, nil
}
