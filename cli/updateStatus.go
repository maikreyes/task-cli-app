package cli

import (
	"fmt"
	"strconv"
)

func updateStatus(args []string, index int) {

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

			switch index {
			case 0:
				tasks[i].Status = "done"
			case 1:
				tasks[i].Status = "in-progress"
			}
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
