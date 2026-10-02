package product

import (
	"os"

	"gopkg.in/yaml.v2"
)

func Load(path string) (Catalogue, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var catalogue Catalogue
	if err := yaml.Unmarshal(data, &catalogue); err != nil {
		return nil, err
	}

	return catalogue, nil
}
