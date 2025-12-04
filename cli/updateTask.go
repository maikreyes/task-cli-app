package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func updateCommand(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: Update <Id> <Description>")
		return
	}

	idStr := args[0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID", idStr)
		return
	}

	description := strings.Join(args[1:], "")

	err = update(id, description)
	if err != nil {
		fmt.Println("Error updating task")
		return
	}

	fmt.Printf("Task %d update to: %s\n", id, description)
}

func update(id int, description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	updated := false

	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			updated = true
			break
		}
	}

	if !updated {
		return fmt.Errorf("sorry, task %d not found", id)
	}

	return saveTask(tasks)
}
