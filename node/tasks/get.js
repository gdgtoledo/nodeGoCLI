const config = require('../config');

const get = async () => {
  try {
    return require(`./${config.files.tasks}`);
  } catch (error) {
    return [];
  }
};

module.exports = {
  get: get,
};
