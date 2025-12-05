package cli

import (
	"flag"
	"fmt"
)

func Run() {

	help := flag.Bool("help", false, "Show help message")

	flag.Parse()

	if *help {
		PrintHelp()
		return
	}

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("No command provided. Use --help for usage.")
		return
	}

	command := args[0]

	switch command {
	case "add":
		AddTask(args[1:])
	case "update":
		updateCommand(args[1:])
	case "mark-done":
		updateStatus(args[1:], 0)
	case "mark-in-progress":
		updateStatus(args[1:], 1)
	case "delete":
		deleteTask(args[1:])
	case "list":
		list(args[1:])
	}

}

func PrintHelp() {
	fmt.Println(`Usage:
  task-cli [OPTIONS] COMMAND

OPTIONS:
  --help                                Show this message

COMMANDS:

  add <TaskDescription>                 	Adding a new task

  update <TaskID> <New TaskDescription> 	Updating a task
  mark-in-progress <TaskID> 			Updating a task statuts to in-progress
  mark-done <TaskID>				Updating a task status to done
  
  delete <TaskID>                       	Delete a task
  
  list 	<Status> (done, todo, in-progress)	List all task`)
}
