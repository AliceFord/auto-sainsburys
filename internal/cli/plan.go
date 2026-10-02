package cli

import (
	"github.com/AliceFord/auto-sainsburys/internal/order"
	"github.com/AliceFord/auto-sainsburys/internal/plan"
	"github.com/AliceFord/auto-sainsburys/internal/product"
	"github.com/AliceFord/auto-sainsburys/internal/recipe"
	"github.com/spf13/cobra"
)

func newPlanCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "plan",
		Short: "Output plan for current Sainsbury's order",
		Args:  cobra.NoArgs,
		RunE:  runPlan,
	}
}

func runPlan(cmd *cobra.Command, args []string) error {
	// Load the current order
	o, err := order.Load("data/order.yaml")
	if err != nil {
		return err
	}

	// Load specified recipes
	recipes, err := recipe.LoadMany("data/recipes", o.Recipes)
	if err != nil {
		return err
	}

	// Load item -> sainsburys mappings
	catalogue, err := product.Load("data/products.yaml")
	if err != nil {
		return err
	}

	_, err = plan.Build(catalogue, recipes)
	if err != nil {
		return err
	}

	return nil
}
