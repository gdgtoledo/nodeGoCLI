const colors = require('colors');

const getIndex = (description = '', tasksToDo = []) => {
  const NO_TASKS_ERROR = 'Error: doesnt exist tasks in database';
  const FLAG_ERROR = 'Error: task is not a string';
  const taskDescriptionIsAString = typeof description === 'string';
  const hasTaskToDo = tasksToDo.length > 0;

  if (!taskDescriptionIsAString) {
    throw new Error(FLAG_ERROR.red);
  }

  if (!hasTaskToDo) {
    console.log(NO_TASKS_ERROR.red);
    return;
  }

  return tasksToDo.findIndex((taskToDo) => taskToDo.description === description);
};

module.exports = {
  getIndex: getIndex,
};
