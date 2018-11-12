package model

import (
	"log"
)

var tasks []TaskModel

// Task interface representing task actions
type Task interface {
	GetDescription() string
	IsComplete() bool
}

// TaskModel implementation for a Task
type TaskModel struct {
	Complete    bool
	Description string
}

// Create creates a task
func Create(description string, complete bool) (TaskModel, error) {
	task := TaskModel{
		Complete:    complete,
		Description: description,
	}

	store := taskStoreImpl{}

	return store.create(task)
}

// List lists all tasks in the data store
func List() ([]TaskModel, error) {
	store := taskStoreImpl{}

	return store.list()
}

// Update updates a task
func Update(description string, complete bool) (TaskModel, error) {
	task := TaskModel{
		Complete:    complete,
		Description: description,
	}

	store := taskStoreImpl{}

	return store.update(task)
}

// taskStore interface for the task store
type taskStore interface {
	create(task TaskModel) (TaskModel, error)
	//get(description string) (TaskModel, error)
	list() ([]TaskModel, error)
	update(task TaskModel) (TaskModel, error)
}

type taskStoreImpl struct {
}

// create creates a task
func (t taskStoreImpl) create(task TaskModel) (TaskModel, error) {
	log.Println(`A task: ` + task.Description + ` has being added`)

	tasks = append(tasks, task)

	return task, nil
}

func (t taskStoreImpl) list() ([]TaskModel, error) {
	return tasks, nil
}

// update updated a task
func (t taskStoreImpl) update(task TaskModel) (TaskModel, error) {
	log.Println(`A task: ` + task.Description + ` has being updated`)
	return task, nil
}
