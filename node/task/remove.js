const tasks = require('../tasks');
const { getIndex } = require('./get');

const remove = async (description) => {
  const SUCCESS = `A task: "${description}" has being deleted`;
  const tasksToDo = await tasks.get();
  const taskIndexToRemove = getIndex(description, tasksToDo);
  tasksToDo.splice(taskIndexToRemove, 1);
  await tasks.post(tasksToDo);
  console.log(SUCCESS.green);
};

exports.remove = remove;
