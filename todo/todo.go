package todo

import (
	"fmt"
	"github.com/devmaufh/todo-console-app/utils"
	"os"
	"text/tabwriter"
)

var Tasks []Task

// Run initialize the to-do application.
func Run() {
	fmt.Println(intro())
	for {
		input := utils.ReadConsoleInput(menu())
		matchInputWithMenu(input)
	}
}

// intro return the intro message.
func intro() string {
	return `
Welcome to 2Do application
-----------------------------
Here you can manage your tasks in a simple, funny and console way xD.`
}

// menu return the menu options as string.
func menu() string {
	return `
Menu
[0] : Exit.
[1] : List tasks.
[2] : Create Task.
[3] : Update Task.
[4] : Set task status.
[5] : Delete Task.
`
}

// matchInputWithMenu process the input of the user.
func matchInputWithMenu(input string) {
	switch input {
	case "0":
		os.Exit(1)
	case "1":
		listTasks()
		break
	case "2":
		addTask()
		break
	}
}

// addTask add a new task to the Tasks collection.
func addTask() {
	name := utils.ReadConsoleInput("Task name: ")
	description := utils.ReadConsoleInput("Task description")
	task := InitializeTask(name, description)
	Tasks = append(Tasks, task)
}

// listTasks List all the tasks stored.
func listTasks() {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	println("List of tasks")
	fmt.Fprintln(w, "Code\tName\tStatus\tCreation date\t")
	for _, task := range Tasks {
		println()
		fmt.Fprintln(w, task.Code+"\t"+task.Name+"\t"+task.Status.String()+"\t"+task.CreatedAt+"\t")
	}
	w.Flush()
}
