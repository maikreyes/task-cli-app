package cli

import "fmt"

func list(args []string) {

	if len(args) < 1 {
		listAllTask()
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading taks", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	status := args[0]

	for i, t := range tasks {
		if t.Status == status {
			fmt.Printf("%d - %s [%s]\n", tasks[i].Id, tasks[i].Description, tasks[i].Status)
		}
	}

}

func listAllTask() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for _, task := range tasks {
		fmt.Printf("%d - %s [%s]\n", task.Id, task.Description, task.Status)
	}
}
