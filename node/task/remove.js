const tasks = require('../tasks');
const { getIndex } = require('./get');
const { success } = require('../logs/success');
const { error } = require('../logs/error');

const remove = async (description) => {
  const REMOVE_SUCCESS = `"${description}" has being deleted to your todo list`;
  const NO_TASKS_ERROR = 'Error: the task that you are trying to remove doesnt exist in database';
  const tasksToDo = await tasks.get();
  const taskIndexToRemove = getIndex(description, tasksToDo);
  const hasTaskToRemove = taskIndexToRemove > -1;

  process.stdout.write('\n');

  if (!hasTaskToRemove) {
    error(NO_TASKS_ERROR);
    process.stdout.write('\n');
    process.exit(1);
  }

  tasksToDo.splice(taskIndexToRemove, 1);
  await tasks.save(tasksToDo);
  success(REMOVE_SUCCESS);
};

exports.remove = remove;
