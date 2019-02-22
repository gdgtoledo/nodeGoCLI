package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
	config "github.com/gdgtoledo/nodeGoCLI/golang/config"
)

// Task interface representing task actions
type Task interface {
	GetDescription() string
	IsComplete() bool
	ToString() string
}

// TaskModel implementation for a Task
type TaskModel struct {
	Complete    bool   `json:"complete"`
	Description string `json:"description"`
}

// GetDescription returns the description of the task
func (t TaskModel) GetDescription() string {
	return t.Description
}

// IsComplete returns the description of the task
func (t TaskModel) IsComplete() bool {
	return t.Complete
}

// ToString returns the string representation  of the task
func (t TaskModel) ToString() string {
	var completedMessage = "✅"

	completed := t.IsComplete()
	if !completed {
		completedMessage = "🕗"
	}

	return completedMessage + " " + t.GetDescription()
}

// Create creates a task
func Create(description string, complete bool) error {
	task := TaskModel{
		Complete:    complete,
		Description: description,
	}

	store := taskStoreImpl{}

	return store.create(task)
}

// List lists all tasks in the data store
func List() error {
	store := taskStoreImpl{}

	tasks, err := store.list()
	if err != nil {
		log.Println(color.RedString("Error listing tasks"))
		return err
	}

	for _, t := range tasks {
		fmt.Println(color.YellowString(t.ToString()))
	}

	return nil
}

// Remove removes a task
func Remove(description string) error {
	task := TaskModel{
		Description: description,
	}

	store := taskStoreImpl{}

	return store.remove(task)
}

// Update updates a task
func Update(description string, complete bool) error {
	task := TaskModel{
		Complete:    complete,
		Description: description,
	}

	store := taskStoreImpl{}

	return store.update(task)
}

// taskStore interface for the task store
type taskStore interface {
	create(task TaskModel) error
	//get(description string) (TaskModel, error)
	list() ([]TaskModel, error)
	remove(task TaskModel) error
	update(task TaskModel) error
}

type taskStoreImpl struct {
}

// create creates a task
func (t taskStoreImpl) create(task TaskModel) error {
	tasks, _ := readTasksFromFile()

	for _, t := range tasks {
		if t.GetDescription() == task.GetDescription() {
			return errors.New(
				`The task: ` + task.Description + ` could not be created: Already exists`)
		}
	}

	tasks = append(tasks, task)

	saveTasksToFile(tasks)

	log.Println(color.GreenString(`The task: ` + task.Description + ` has being added`))

	return nil
}

func (t taskStoreImpl) list() ([]TaskModel, error) {
	tasks, _ := readTasksFromFile()

	return tasks, nil
}

// remove removed a task
func (t taskStoreImpl) remove(task TaskModel) error {
	tasks, _ := readTasksFromFile()

	removed := false

	for i, t := range tasks {
		if t.GetDescription() == task.GetDescription() {
			t.Complete = task.IsComplete()

			tasks = append(tasks[:i], tasks[i+1:]...)

			removed = true

			break
		}
	}

	if !removed {
		return errors.New(`The task: ` + task.Description + ` could not be deleted: Not found`)
	}

	saveTasksToFile(tasks)

	log.Println(color.GreenString(`The task: ` + task.Description + ` has being removed`))
	return nil
}

// update updated a task
func (t taskStoreImpl) update(task TaskModel) error {
	tasks, _ := readTasksFromFile()

	found := false

	for i, t := range tasks {
		if t.GetDescription() == task.GetDescription() {
			t.Complete = task.IsComplete()

			tasks[i] = t

			found = true

			break
		}
	}

	if !found {
		return errors.New(`The task: ` + task.Description + ` could not be updated: Not Found`)
	}

	saveTasksToFile(tasks)

	log.Println(color.GreenString(`The task: ` + task.Description + ` has being updated`))
	return nil
}

func getTasksFile() (*os.File, error) {
	var tasksFilePath = config.TasksFileName()

	return os.Open(tasksFilePath)
}

func readTasksFromFile() ([]TaskModel, error) {
	var tasks []TaskModel

	tasksFile, err := getTasksFile()
	if err != nil {
		return tasks, err
	}

	log.Println(color.GreenString("Successfully Opened " + tasksFile.Name()))

	defer tasksFile.Close()

	byteValue, _ := ioutil.ReadAll(tasksFile)

	json.Unmarshal([]byte(byteValue), &tasks)

	return tasks, nil
}

func saveTasksToFile(tasks []TaskModel) error {
	tasksJSON, _ := json.Marshal(tasks)

	err := ioutil.WriteFile(config.TasksFileName(), tasksJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}
