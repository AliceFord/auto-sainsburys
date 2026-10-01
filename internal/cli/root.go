package cli

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	root := &cobra.Command{
		Use:   "auto-sains",
		Short: "Build recipes and add them to a Sainsbury's order",
	}

	root.AddCommand(
		newAuthCommand(),
	)

	return root
}
