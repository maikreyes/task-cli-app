package cli

import (
	"encoding/json"
	"os"
)

func saveTask(tasks []Task) error {
	file, err := os.Create("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	return encoder.Encode(tasks)
}
