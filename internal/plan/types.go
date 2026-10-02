package plan

const (
	ValidationPending ValidationStatus = "pending"
	ValidationValid   ValidationStatus = "valid"
	ValidationInvalid ValidationStatus = "invalid"
)

type ValidationStatus string

type Plan struct {
	Items []PlanItem
}

type PlanItem struct {
	// Code name
	Product string
	// Human-readable name
	Name string
	// Unit of measurement
	Unit string
	// How many of the given product should we order
	OrderQuantity int
	// Sainsbury's product id
	SainsId string
	// Sainsbury's product SKU
	SKU string
	// How many units of the product are contained in each Sainsbury's product
	ProductUnits float64
	// The validation status of the plan item
	ValidationStatus ValidationStatus
	// The reason for the validation status
	ValidationReason string
}
