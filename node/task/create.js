const tasks = require('../tasks');
const { Task } = require('./Task');
const { success } = require('../logs/success');

const create = async (description, isComplete) => {
  const CREATE_SUCCESS = `"${description}" has being added to your todo list`;

  process.stdout.write('\n');

  const tasksToDo = await tasks.get();
  const newTaskToDo = new Task(description, isComplete);

  tasksToDo.push(newTaskToDo);
  await tasks.save(tasksToDo);
  success(CREATE_SUCCESS);
};

exports.create = create;
