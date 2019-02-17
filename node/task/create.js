const colors = require('colors');

const tasks = require('../tasks');
const { Task } = require('./Task');
const { success } = require('../logs/success');
const { error } = require('../logs/error');

const create = async (description, isComplete) => {
  const NO_CORRECT_DESCRIPTION_ERROR = 'Error: task is not a string';
  const CREATE_SUCCESS = `"${description}" has being added to your todo list`;
  const taskDescriptionIsAString = typeof description === 'string';

  process.stdout.write('\n');

  if (!taskDescriptionIsAString) {
    error(NO_CORRECT_DESCRIPTION_ERROR);
    process.stdout.write('\n');
    process.exit(1);
  }

  const tasksToDo = await tasks.get();
  const newTaskToDo = new Task(description, isComplete);

  tasksToDo.push(newTaskToDo);
  await tasks.save(tasksToDo);
  success(CREATE_SUCCESS);
};

exports.create = create;
