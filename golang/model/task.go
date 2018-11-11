package model

import (
	"log"
)

// Task interface representing task actions
type Task interface {
	Create(description string, complete bool) (TaskModel, error)
	GetDescription() string
	IsComplete() bool
}

// TaskModel implementation for a Task
type TaskModel struct {
	Complete    bool
	Description string
	store       taskStore
}

// Create creates a task
func Create(description string, complete bool) (TaskModel, error) {
	task := TaskModel{
		Complete:    complete,
		Description: description,
		store:       taskStoreImpl{},
	}

	return task.store.create(task)
}

// taskStore interface for the task store
type taskStore interface {
	create(task TaskModel) (TaskModel, error)
	//get(description string) (TaskModel, error)
	//list() ([]TaskModel, error)
	//update(t Task) (TaskModel, error)
}

type taskStoreImpl struct {
}

// create creates a task
func (t taskStoreImpl) create(task TaskModel) (TaskModel, error) {
	log.Println(`A task: ` + task.Description + ` has being added`)
	return task, nil
}
