package order

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Order struct {
	Recipes []string `yaml:"recipes"`
}

func Load(path string) (*Order, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := yaml.Unmarshal(data, &order); err != nil {
		return nil, err
	}

	return &order, nil
}
