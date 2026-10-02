package product

type Catalogue map[string]Product

type Product struct {
	Name       string            `yaml:"name"`
	Sainsburys SainsburysProduct `yaml:"sainsburys"`
}

type SainsburysProduct struct {
	ProductUid int     `yaml:"product_uid"`
	Quantity   float64 `yaml:"quantity"`
	Unit       string  `yaml:"unit"`
}
