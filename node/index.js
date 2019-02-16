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

const start = async (inputs, config) => {
  switch (inputs.command) {
    case config.cli.commands.create.name:
      await task.create(inputs.flags.description, inputs.flags.complete);
      break;

    case config.cli.commands.list.name:
      await tasks.list();
      break;

    case config.cli.commands.update.name:
      await task.update(inputs.flags.description, inputs.flags.complete);
      break;

    case config.cli.commands.delete.name:
      await task.remove(inputs.flags.description);
      break;

    default:
      console.log('Command is not in command list, try to execute node index --help').red;
      break;
  }
};

start(inputs, config);
