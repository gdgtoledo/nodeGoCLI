const config = require('../config');

const get = async () => {
  let tasks;
  try {
    tasks =  require(`./${config.files.tasks}`);
  } catch (error) {
    tasks =  [];
  }
  return tasks;
};

exports.get = get;
