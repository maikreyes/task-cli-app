package cli

import (
	"fmt"
	"time"
)

func AddTask(args []string) (Task, error) {

	tasks, err := loadTasks()
	if err != nil {
		return Task{}, err
	}

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].Id + 1
	}

	T := Task{
		Id:          newID,
		Description: args[0],
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, T)

	err = saveTask(tasks)
	if err != nil {
		return Task{}, err
	}

	fmt.Printf("Congratulations, yuo are create: %s  with this status: [%s]\n", T.Description, T.Status)

	return T, nil

}
