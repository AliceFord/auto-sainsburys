package plan

type Plan struct {
	Recipes []string
	Items   []PlanItem
}

type PlanItem struct {
	Product string
	Name    string

	RequiredQuantity float64
	RequiredUnit     string

	SainsburysName string
	OrderQuantity  int

	Included bool
}
