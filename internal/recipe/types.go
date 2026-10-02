package recipe

type Recipe struct {
	Name        string       `yaml:"name"`
	Ingredients []Ingredient `yaml:"ingredients"`
}

type Ingredient struct {
	Product  string  `yaml:"product"`
	Quantity float64 `yaml:"quantity"`
	Unit     string  `yaml:"unit"`
}
