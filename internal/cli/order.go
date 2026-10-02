package cli

import (
	"github.com/AliceFord/auto-sainsburys/internal/plan"
	"github.com/AliceFord/auto-sainsburys/internal/sainsburys"
	"github.com/spf13/cobra"
)

func newOrderCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "order",
		Short: "Order the validated plan",
		Args:  cobra.NoArgs,
		RunE:  runOrder,
	}
}

func runOrder(cmd *cobra.Command, args []string) error {
	// Load the plan
	p, err := plan.Load()
	if err != nil {
		return err
	}

	// Load Sainsbury's session cookies
	c, err := sainsburys.ReadSessionFromFile()
	if err != nil {
		return err
	}

	// Create a Sainsbury's client with the session cookies
	client := sainsburys.NewClient(c)

	// Add the plan
	return client.AddPlan(p)
}
