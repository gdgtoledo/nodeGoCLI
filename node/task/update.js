const tasks = require('../tasks');
const { getIndex } = require('./get');
const icon = require('./icon');
const { success } = require('../logs/success');
const { error } = require('../logs/error');

const updateTasksToDo = (task = { index: -1, isComplete: false }, tasks) => {
  const NO_TASK_ERROR = 'Error: the task that you are trying to update doesnt exist in database';
  const hasTaskToUpdate = task.index > 0;

  if (!hasTaskToUpdate) {
    error(NO_TASK_ERROR);
    process.exit(1);
  }

  tasks[task.index].isComplete = task.isComplete;
  return tasks;
};

const update = async (description, isComplete) => {
  let tasksToDo = await tasks.get();
  const taskIndexToUpdate = getIndex(description, tasksToDo);

  process.stdout.write('\n');

  tasksToDo = updateTasksToDo({ index: taskIndexToUpdate, isComplete }, [...tasksToDo]);
  await tasks.save(tasksToDo);

  const taskIcon = icon.get(tasksToDo[taskIndexToUpdate].isComplete);
  const UPDATE_SUCCESS = `"${description}" has being updated to state ${taskIcon}`;

  success(UPDATE_SUCCESS);
};

exports.update = update;
