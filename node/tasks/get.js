const colors = require('colors');
const config = require('../config');

const get = async () => {
  NO_TASKS_ERROR = 'Error: There arent tasks in the database';
  try {
    return require(`./${config.files.tasks}`);
  } catch (error) {
    return [];
  }
};

exports.get = get;
