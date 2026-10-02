package plan

import (
	"encoding/json"
	"os"
)

func Save(plan *Plan) error {
	data, err := json.MarshalIndent(plan, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("plan.json", data, 0600)
}

func Load() (*Plan, error) {
	data, err := os.ReadFile("plan.json")
	if err != nil {
		return nil, err
	}

	var plan Plan
	if err := json.Unmarshal(data, &plan); err != nil {
		return nil, err
	}

	return &plan, nil
}

func Clear() error {
	return os.Remove("plan.json")
}
