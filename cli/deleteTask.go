package cli

import (
	"fmt"
	"slices"
	"strconv"
)

func deleteTask(args []string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	idStr := args[0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID", idStr)
		return err
	}

	index := -1
	for i, task := range tasks {
		if task.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("Task with ID %d not found\n", id)
		return fmt.Errorf("task not found")
	}

	tasks = slices.Delete(tasks, index, index+1)

	return saveTask(tasks)

}
