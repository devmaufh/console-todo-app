package todo

import (
	"github.com/devmaufh/todo-console-app/utils"
	"time"
)

type Task struct {
	Code        string
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
	Status      Status
}

// InitializeTask Create a new task object.
func InitializeTask(name string, description string) Task {
	return Task{
		Code:        utils.GenerateUuid(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		Status:      Pending,
	}
}
