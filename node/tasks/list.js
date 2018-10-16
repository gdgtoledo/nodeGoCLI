const tasks = require('./get');
const colors = require('colors');

const LIST_ERROR = 'Sorry! the tasks list is empty';

const list = async () => {
  const taskToDo = await tasks.get();
  const hasTasks = taskToDo.length > 0;

  if (!hasTasks) {
    console.log(LIST_ERROR.red);
    return;
  }

  taskToDo.forEach((taskToDo, index) => {
    console.log(`${index}. ${taskToDo.description}`.green);
  });
};

exports.list = list;