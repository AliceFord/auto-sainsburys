package product

type Catalogue map[string]Product

type Product struct {
	Name       string            `yaml:"name"`
	Sainsburys SainsburysProduct `yaml:"sainsburys"`
}

type SainsburysProduct struct {
	ProductName string  `yaml:"product_name"`
	Quantity    float64 `yaml:"quantity"`
	Unit        string  `yaml:"unit"`
}
