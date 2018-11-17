const { create } = require('./create');
const { update } = require('./update');
const { remove } = require('./remove');
const { getIndex } = require('./get');
const { icon } = require('./icon');
const { Task } = require('./Task');

module.exports = {
  Task,
  create,
  update,
  icon,
  remove,
  getIndex,
};
