const yargs = require('yargs');
const config = require('./config');

const get = () => {
  return yargs
    .command(
      config.cli.commands.list.name,
      config.cli.commands.list.description,
      {
        task: config.cli.flags.task,
      },
    )
    .command(
      config.cli.commands.create.name,
      config.cli.commands.create.description,
      {
        description: config.cli.flags.description,
      },
    )
    .command(
      config.cli.commands.update.name,
      config.cli.commands.update.description,
      {
        complete: config.cli.flags.complete,
        description: config.cli.flags.description,
      },
    )
    .help()
    .argv;
}

module.exports = {
  get,
}