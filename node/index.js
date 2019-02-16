#!/usr/bin/env node

const colors = require('colors');
const args = require('./args');
const config = require('./config');
const tasks = require('./tasks');
const task = require('./task');

const argv = args.get();
const command = argv._[0];

const flags = {
  description: argv.description,
  complete: argv.complete,
  task: argv.task,
};

const inputs = {
  command,
  flags,
};

const start = async (userInputs, configuration) => {
  switch (userInputs.command) {
    case configuration.cli.commands.create.name:
      await task.create(userInputs.flags.description, userInputs.flags.complete);
      break;

    case configuration.cli.commands.list.name:
      await tasks.list();
      break;

    case configuration.cli.commands.update.name:
      await task.update(userInputs.flags.description, userInputs.flags.complete);
      break;

    case configuration.cli.commands.delete.name:
      await task.remove(userInputs.flags.description);
      break;

    default:
      console.log('Command is not in command list, try to execute node index --help').red;
      break;
  }
};

start(inputs, config);
