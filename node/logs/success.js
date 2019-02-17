const colors = require('colors');

const successLabel = colors.bgGreen.black.bold(' SUCCESS ');

const success = (message) => {
  console.log(`${successLabel} ${colors.green(message)}`);
};

exports.success = success;
