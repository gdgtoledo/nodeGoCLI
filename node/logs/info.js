const colors = require('colors');

const infoLabel = colors.bgBlue.black.bold('  INFO   ');

const info = (message) => {
  console.log(`${infoLabel} ${colors.blue(message)}`);
};

exports.info = info;
