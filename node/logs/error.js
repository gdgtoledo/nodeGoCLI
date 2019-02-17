const colors = require('colors');

const errorLabel = colors.bgRed.black.bold('  ERROR  ');

const error = (message) => {
  console.log(`${errorLabel} ${colors.red(message)}`);
};

exports.error = error;
