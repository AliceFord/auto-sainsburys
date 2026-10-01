package cli

import (
	"github.com/AliceFord/auto-sainsburys/internal/sainsburys"
	"github.com/spf13/cobra"
)

func newAuthCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "auth",
	}

	cmd.AddCommand(
		&cobra.Command{
			Use:   "login",
			Short: "Log into Sainsbury's",
			RunE:  runAuthLogin,
		},
		&cobra.Command{
			Use:   "status",
			Short: "Check Sainsbury's authentication",
			RunE:  runAuthStatus,
		},
		&cobra.Command{
			Use:   "logout",
			Short: "Delete the saved Sainsbury's session",
			RunE:  runAuthLogout,
		},
	)

	return cmd
}

func runAuthLogin(cmd *cobra.Command, args []string) error {
	// Start Playwright, login, get session cookies & delivery info, save locally

	cookies, err := sainsburys.AuthLogin()
	if err != nil {
		return err
	}

	return sainsburys.WriteSessionToFile(cookies)
}

func runAuthStatus(cmd *cobra.Command, args []string) error {
	// Check if session cookies & delivery info are saved locally, and if they are try access something to see if they work

	cookies, err := sainsburys.ReadSessionFromFile()
	if err != nil {
		return err
	}

	client := sainsburys.NewClient(cookies)
	return client.CheckAuth()
}

func runAuthLogout(cmd *cobra.Command, args []string) error {
	// Delete the saved session cookies & delivery info
	return sainsburys.DeleteSessionFile()
}
