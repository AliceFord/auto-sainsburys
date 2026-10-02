package cli

import (
	"github.com/AliceFord/auto-sainsburys/internal/plan"
	"github.com/AliceFord/auto-sainsburys/internal/sainsburys"
	"github.com/spf13/cobra"
)

func newValidateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "validate",
		Short: "Validate the current Sainsbury's plan",
		Args:  cobra.NoArgs,
		RunE:  runValidate,
	}
}

func runValidate(cmd *cobra.Command, args []string) error {
	// Load the current order
	p, err := plan.Load()
	if err != nil {
		return err
	}

	sainsburys.NewClient(nil).ValidatePlan(p)

	for _, item := range p.Items {
		if item.ValidationStatus == plan.ValidationInvalid {
			cmd.Printf("Validation failed for %s: %s\n", item.Name, item.ValidationReason)
		}
	}

	// Write the validated order
	return plan.Save(p)
}
