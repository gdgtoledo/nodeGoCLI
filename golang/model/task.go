package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Task interface representing task actions
type Task interface {
	GetDescription() string
	IsComplete() bool
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

// Tasks struct which contains an array of tasks
type Tasks struct {
	Items []TaskModel `json:"tasks"`
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

	tasks, _ := readTasksFromFile()

	tasks.Items = append(tasks.Items, task)

	return task, nil
}

func (t taskStoreImpl) list() ([]TaskModel, error) {
	tasks, _ := readTasksFromFile()

	return tasks.Items, nil
}

// update updated a task
func (t taskStoreImpl) update(task TaskModel) (TaskModel, error) {
	log.Println(`A task: ` + task.Description + ` has being updated`)
	return task, nil
}

func checkGDGCLIHome(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

func checkTasksFile(path string) {
	_, err := os.Stat(path)
	if err != nil {
		log.Println("File does not exist. Creating it now")
		tasksFile, _ := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)

		defer tasksFile.Close()
	}
}

func readTasksFromFile() (Tasks, error) {
	var tasks Tasks

	var cliHome = os.Getenv("HOME") + string(os.PathSeparator) + ".gdg-cli"
	checkGDGCLIHome(cliHome)

	var tasksFilePath = cliHome + string(os.PathSeparator) + "tasks.json"
	checkTasksFile(tasksFilePath)

	tasksFile, err := os.Open(tasksFilePath)
	if err != nil {
		return tasks, err
	}

	log.Println("Successfully Opened tasks.json")

	defer tasksFile.Close()

	byteValue, _ := ioutil.ReadAll(tasksFile)

	json.Unmarshal([]byte(byteValue), &tasks.Items)

	return tasks, nil
}
