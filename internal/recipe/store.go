package recipe

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func Load(path string) (*Recipe, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var recipe Recipe
	if err := yaml.Unmarshal(data, &recipe); err != nil {
		return nil, err
	}

	return &recipe, nil
}

func LoadAll(dir string) ([]Recipe, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var recipes []Recipe
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		recipe, err := Load(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, *recipe)
	}

	return recipes, nil
}

func LoadMany(dir string, names []string) ([]Recipe, error) {
	var recipes []Recipe

	for _, name := range names {
		recipe, err := Load(filepath.Join(dir, name+".yaml"))
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, *recipe)
	}

	return recipes, nil
}
