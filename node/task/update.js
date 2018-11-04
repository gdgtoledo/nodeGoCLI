const colors = require('colors');
const tasks = require('../tasks');
const icon = require('./icon');

const getTaskIndexByDescription = (description = '', tasksToDo = []) =>
  tasksToDo.findIndex((taskToDo) => taskToDo.description === description);

const updateTask = (task = { index: -1, isComplete: false }, tasks) => {
  const NO_TASK_ERROR = 'Error: the task that you are trying to update doesnt exist';

  if (task.index < 0) {
    throw new Error(NO_TASK_ERROR.red);
  }

  tasks[task.index].isComplete = task.isComplete;

  return tasks;
};

const update = async (description, isComplete) => {
  const NO_TASKS_ERROR = 'Error: doesnt exist tasks in database';
  const FLAG_ERROR = 'Error: task is not a string';
  const taskDescriptionIsAString = typeof description === 'string';

  if (!taskDescriptionIsAString) {
    throw new Error(FLAG_ERROR.red);
  }

  let tasksToDo = await tasks.get();
  const hasTaskToDo = tasksToDo.length > 0;

  if (!hasTaskToDo) {
    console.log(NO_TASKS_ERROR.red);
    return;
  }

  const taskIndexByDescription = getTaskIndexByDescription(description, tasksToDo);
  tasksToDo = updateTask(
    {
      index: taskIndexByDescription,
      isComplete,
    },
    [...tasksToDo]
  );
  await tasks.post(tasksToDo);
  const taskIcon = icon.get(tasksToDo[taskIndexByDescription].isComplete);
  const SUCCESS = `A task: "${description}" has being updated to state ${taskIcon}`;
  console.log(SUCCESS.blue);
};

exports.update = update;
