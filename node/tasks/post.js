const fs = require('fs');
const config = require('../config');

const tasksToDoFilePathToStoreThem = `./tasks/${config.files.tasks}`;
const SUCCESS = `Tasks have being saved`;

const post = async (tasksToDo) => {
  const tasksToDoStringed = JSON.stringify(tasksToDo);

  fs.writeFile(tasksToDoFilePathToStoreThem, tasksToDoStringed, (err) => {
    if (err) {
      throw new Error(err.red);
    } else {
      console.log(SUCCESS.blue);
    }
  });
};

exports.post = post;
