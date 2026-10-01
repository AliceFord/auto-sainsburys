package cli

import "github.com/spf13/cobra"

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

	return nil
}

func runAuthStatus(cmd *cobra.Command, args []string) error {
	// Check if session cookies & delivery info are saved locally, and if they are try access something to see if they work

	return nil
}

func runAuthLogout(cmd *cobra.Command, args []string) error {
	// Delete the saved session cookies & delivery info

	return nil
}
