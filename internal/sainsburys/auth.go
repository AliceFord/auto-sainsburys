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

	ctx, err := pw.Chromium.LaunchPersistentContext(
		".auto-sainsburys-browser",
		playwright.BrowserTypeLaunchPersistentContextOptions{
			Headless: playwright.Bool(false),
		},
	)
	if err != nil {
		return nil, err
	}
	defer ctx.Close()

	pages := ctx.Pages()

	var page playwright.Page

	if len(pages) > 0 {
		page = pages[0]
	} else {
		page, err = ctx.NewPage()
		if err != nil {
			return nil, err
		}
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

	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')

	// Explicitly persist cookies + localStorage + IndexedDB.
	_, err = ctx.StorageState(
		playwright.BrowserContextStorageStateOptions{
			Path: playwright.String(".auto-sainsburys-browser/storageState.json"),
		},
	)
	if err != nil {
		_ = ctx.Close()
		return nil, fmt.Errorf("save authentication state: %w", err)
	}

	// Close gracefully so Chromium can flush the persistent profile.
	if err := ctx.Close(); err != nil {
		return nil, fmt.Errorf("close browser context: %w", err)
	}

	fmt.Println("Authentication state saved.")

	return cookies, nil
}
