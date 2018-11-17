const colors = require('colors');
const tasks = require('../tasks');
const { getIndex } = require('./get');
const icon = require('./icon');

const updateTasksToDo = (task = { index: -1, isComplete: false }, tasks) => {
  const NO_TASK_ERROR = 'Error: the task that you are trying to update doesnt exist in database';

  if (task.index < 0) {
    throw new Error(NO_TASK_ERROR.red);
  }

  tasks[task.index].isComplete = task.isComplete;
  return tasks;
};

const consoleSuccessMessage = (description, taskIcon) => {
  const SUCCESS = `A task: "${description}" has being updated to state ${taskIcon}`;
  console.log(SUCCESS.green);
};

const update = async (description, isComplete) => {
  let tasksToDo = await tasks.get();
  const taskIndexToUpdate = getIndex(description, tasksToDo);
  tasksToDo = updateTasksToDo({ index: taskIndexToUpdate, isComplete }, [...tasksToDo]);
  await tasks.post(tasksToDo);
  const taskIcon = icon.get(tasksToDo[taskIndexToUpdate].isComplete);
  consoleSuccessMessage(description, taskIcon);
};

exports.update = update;
