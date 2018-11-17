const colors = require('colors');

const tasks = require('../tasks');
const { Task } = require('./Task.js');

const consoleSuccessMessage = (description) => {
  const SUCCESS = `A task: "${description}" has being added to list`;
  console.log(SUCCESS.green);
};

const create = async (description, isComplete) => {
  const FLAG_ERROR = 'Error: task is not a string';
  const taskDescriptionIsAString = typeof description === 'string';

  if (!taskDescriptionIsAString) {
    throw new Error(FLAG_ERROR.red);
    return;
  }

  const tasksToDo = await tasks.get();
  const newTaskToDo = new Task(description, isComplete);

  tasksToDo.push(newTaskToDo);
  await tasks.post(tasksToDo);
  consoleSuccessMessage(description);
};

exports.create = create;
