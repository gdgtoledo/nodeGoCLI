const colors = require('colors');
const tasks = require('../tasks/get');
const icon = require('../task/icon');

const list = async () => {
  const LIST_ERROR = 'Sorry! the tasks list is empty';
  const tasksToDo = await tasks.get();
  const hasTasks = tasksToDo.length > 0;

  if (!hasTasks) {
    console.log(LIST_ERROR.red);
    return;
  }

  tasksToDo.forEach((taskToDo) => {
    const taskToDoIcon = icon.get(taskToDo.isComplete);
    const listTemplate = `${taskToDoIcon}  ${taskToDo.description}`;
    console.log(listTemplate);
  });
};

exports.list = list;
