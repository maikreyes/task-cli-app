package cli

import (
	"fmt"
	"strconv"
)

func updateDoneStatus(args []string) {

	idStr := args[0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID", idStr)
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		return
	}
	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Status = "done"
			break
		}
	}

	err = saveTask(tasks)
	if err != nil {
		fmt.Println("Error updating task")
		return
	}

	fmt.Printf("Task with id %s are updated.\n", idStr)

}

func updateInProgressStatus(args []string) {

	idStr := args[0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID", idStr)
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		return
	}
	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Status = "in-progress"
			break
		}
	}

	err = saveTask(tasks)
	if err != nil {
		fmt.Println("Error updating task")
		return
	}

	fmt.Printf("Task with id %s are updated.\n", idStr)

}
