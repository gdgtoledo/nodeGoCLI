const colors = require('colors');
const fs = require('fs');

const config = require('../config');
const tasks = require('../tasks/get');
const { Task } = require('./Task.js');

const FLAG_ERROR = 'Error: task is not a string';

const create = async (description) => {
  const SUCCESS = `A task: "${description}" has being added`;
  const taskDescriptionIsAString = typeof description === 'string';

  if (!taskDescriptionIsAString){
    throw new Error(FLAG_ERROR.red);
    return;
  }

  const tasksToDo = await tasks.get();
  const newTaskToDo = new Task(description, false);

  tasksToDo.push(newTaskToDo);

  const tasksToDoFilePathToStoreThem = `./tasks/${config.files.tasks}`;
  const tasksToDoStringed= JSON.stringify(tasksToDo);

  fs.writeFile(tasksToDoFilePathToStoreThem, tasksToDoStringed, (err) => {
    if (err) {
      throw new Error(err.red);
    } else {
      console.log(SUCCESS.blue);
    }
  });
};

exports.create = create;