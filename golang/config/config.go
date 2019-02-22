package config

import (
	"log"
	"os"

	"github.com/fatih/color"
)

const cliHomeName = ".gdg-cli"
const taskFileName = "tasks.json"

// config configuration of the application
type config struct {
	home  string
	tasks string
}

var cliHome = os.Getenv("HOME") + string(os.PathSeparator) + cliHomeName

var cliConfiguration = config{
	home:  cliHome,
	tasks: cliHome + string(os.PathSeparator) + taskFileName,
}

// TasksFileName returns the path of the tasks file
func TasksFileName() string {
	return cliConfiguration.tasks
}

// WorkingDir returns working directory of the CLI
func WorkingDir() string {
	return cliConfiguration.home
}

func checkGDGCLIHome() {
	path := cliConfiguration.home
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

func checkTasksFile() {
	path := cliConfiguration.tasks

	_, err := os.Stat(path)
	if err != nil {
		log.Println(color.MagentaString("File does not exist. Creating it now"))
		tasksFile, _ := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)

		defer tasksFile.Close()
	}
}

func init() {
	checkGDGCLIHome()

	checkTasksFile()
}
