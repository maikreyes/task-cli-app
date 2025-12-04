package cli

import (
	"encoding/json"
	"os"
)

func loadTasks() ([]Task, error) {
	file, err := os.Open("data.json")
	if os.IsNotExist(err) {
		return []Task{}, nil
	}
	if err != nil {
		return nil, err
	}

	defer file.Close()
	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
