package product

type Catalogue map[string]Product

type Product struct {
	Name       string            `yaml:"name"`
	Sainsburys SainsburysProduct `yaml:"sainsburys"`
}

type SainsburysProduct struct {
	SainId   string  `yaml:"sain_id"`
	SKU      string  `yaml:"sku"`
	Quantity float64 `yaml:"quantity"`
	Unit     string  `yaml:"unit"`
}
