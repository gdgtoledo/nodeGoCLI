package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	config "github.com/gdgtoledo/nodeGoCLI/golang/config"
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

	tasks = append(tasks, task)

	saveTasksToFile(tasks)

	return task, nil
}

func (t taskStoreImpl) list() ([]TaskModel, error) {
	tasks, _ := readTasksFromFile()

	return tasks, nil
}

// update updated a task
func (t taskStoreImpl) update(task TaskModel) (TaskModel, error) {
	tasks, _ := readTasksFromFile()

	for i, t := range tasks {
		if t.GetDescription() == task.GetDescription() {
			t.Complete = task.IsComplete()

			tasks[i] = t
			break
		}
	}

	saveTasksToFile(tasks)

	log.Println(`A task: ` + task.Description + ` has being updated`)
	return task, nil
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

	log.Println("Successfully Opened " + tasksFile.Name())

	defer tasksFile.Close()

	byteValue, _ := ioutil.ReadAll(tasksFile)

	json.Unmarshal([]byte(byteValue), &tasks)

	return tasks, nil
}

func saveTasksToFile(tasks []TaskModel) error {
	tasksJSON, _ := json.Marshal(tasks)

	for _, t := range tasks {
		log.Println(t)
	}

	err := ioutil.WriteFile(config.TasksFileName(), tasksJSON, 0644)
	if err != nil {
		return err
	}

	log.Println("Tasks saved successfully")
	return nil
}
