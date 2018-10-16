const commands = {
  create: {
    name:  'create',
    description:'Create a todo task',
  },
  list : {
    name: 'list',
    description: 'Print in console all todo tasks',
  },
  update : {
    name: 'update',
    description: 'Update the complete state for a task',
  }
};

const flags = {
  task: {
    alias: 't',
    default: 'all',
    desc: 'List a task',
    type: 'string',
  },
  complete: {
    default: true,
    alias: 'c',
    desc: 'Set complete task state',
  },
  description: {
    demand: true,
    alias: 'd',
    desc: 'Description for task',
    type: 'string',
  },
};

const cli = {
  commands,
  flags,
};

const FILE_TO_WRITE_TASKS_TO_DO = 'tasks.json';

const files = {
  tasks: FILE_TO_WRITE_TASKS_TO_DO,
};

module.exports = {
  cli,
  files,
};