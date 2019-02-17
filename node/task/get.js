const { error } = require('../logs/error');

const getIndex = (description = '', tasksToDo = []) => {
  const FLAG_ERROR = 'Error: task is not a string';
  const taskDescriptionIsAString = typeof description === 'string';
  const hasTaskToDo = tasksToDo.length > 0;

  if (!taskDescriptionIsAString) {
    error(FLAG_ERROR);
    process.stdout.write('\n');
    process.exit(1);
  }

  if (!hasTaskToDo) {
    return -1;
  }

  return tasksToDo.findIndex((taskToDo) => taskToDo.description === description);
};

module.exports = {
  getIndex: getIndex,
};
