# Golang

<img width="30%" src="../images/gopher.png" />

visit the oficial [Golang oficial page](page https://golang.org)

## Commands

### List tasks

List all your tasks

```gdg-tasks list```

### Create task

Create a single task. By default the task is created with false complete state

`gdg-tasks create -d "Example for task description"`

### Update task

Update the complete state for a task

`gdg-tasks update -d "Example for task description" -c`

### Remove task

Remove a single task

`gdg-tasks remove -d "Example for task description"`


## Options

### Description

String  description for a task

`--description -d <String>`

### Complete

Boolean state (complete/not complete) for a task

`--complete -c <Boolean>`