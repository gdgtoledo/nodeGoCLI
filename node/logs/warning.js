const colors = require('colors');

const warningLabel = colors.bgYellow.black.bold(' WARNING ');

const warning = (message) => {
  console.log(`${warningLabel} ${colors.yellow(message)}`);
};

exports.warning = warning;
