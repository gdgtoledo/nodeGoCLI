const { error } = require('../logs/error');
const tasks = require('../tasks/get');
const icon = require('../task/icon');

const list = async () => {
  const LIST_ERROR = 'Sorry! the tasks list is empty';
  const tasksToDo = await tasks.get();
  const hasTasks = tasksToDo.length > 0;

  if (!hasTasks) {
    process.stdout.write('\n');
    error(LIST_ERROR);
    process.stdout.write('\n');
    return;
  }

  process.stdout.write('\n');

  tasksToDo.forEach((taskToDo) => {
    const taskToDoIcon = icon.get(taskToDo.isComplete);
    const listTemplate = `${taskToDoIcon}  ${taskToDo.description}`;
    console.log(listTemplate);
  });

  process.stdout.write('\n');
};

exports.list = list;
