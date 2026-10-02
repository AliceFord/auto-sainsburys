package plan

import (
	"fmt"
	"math"

	"github.com/AliceFord/auto-sainsburys/internal/product"
	"github.com/AliceFord/auto-sainsburys/internal/recipe"
)

type requirement struct {
	Product  string
	Quantity float64
	Unit     string
}

func Build(catalogue product.Catalogue, recipes []recipe.Recipe) (*Plan, error) {
	requirements := make(map[string]requirement)

	p := &Plan{}

	// Iterate over each recipe and its ingredients
	for _, r := range recipes {
		for _, ing := range r.Ingredients {
			existing, found := requirements[ing.Product]

			if !found {
				// If the product is not already in the requirements, add it
				requirements[ing.Product] = requirement{
					Product:  ing.Product,
					Quantity: ing.Quantity,
					Unit:     ing.Unit,
				}
				continue
			}

			existing.Quantity += ing.Quantity
			requirements[ing.Product] = existing
		}
	}

	// Turn requirements into products to order
	for _, req := range requirements {
		prod, found := catalogue[req.Product]
		if !found {
			return nil, fmt.Errorf("no product mapping for %q", req.Product)
		}

		sainsburys := prod.Sainsburys

		orderQuantity := int(math.Ceil(req.Quantity / sainsburys.Quantity))

		p.Items = append(p.Items, PlanItem{
			Product:          req.Product,
			Name:             prod.Name,
			Unit:             req.Unit,
			OrderQuantity:    orderQuantity,
			SainsburysUid:    sainsburys.ProductUid,
			ProductUnits:     sainsburys.Quantity,
			ValidationStatus: ValidationPending,
		})
	}

	fmt.Printf("plan: %+v\n", p)

	return p, nil
}
