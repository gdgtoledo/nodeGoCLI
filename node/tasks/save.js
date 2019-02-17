const fs = require('fs');

const config = require('../config');
const { success } = require('../logs/success');
const { error } = require('../logs/error');

const save = async (tasksToDo) => {
  const tasksToDoFilePathToStoreThem = `./tasks/${config.files.tasks}`;
  const tasksToDoStringed = JSON.stringify(tasksToDo);
  const SAVED_SUCCESS = `Tasks have being saved`;

  fs.writeFile(tasksToDoFilePathToStoreThem, tasksToDoStringed, (err) => {
    if (err) {
      error(err);
      process.exit(1);
    } else {
      process.stdout.write('\n');
      success(SAVED_SUCCESS);
    }
  });
};

exports.save = save;
