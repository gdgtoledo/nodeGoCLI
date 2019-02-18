const getIndex = (description = '', tasksToDo = []) => {
  const hasTaskToDo = tasksToDo.length > 0;

  if (!hasTaskToDo) {
    return -1;
  }

  return tasksToDo.findIndex((taskToDo) => taskToDo.description === description);
};

module.exports = {
  getIndex: getIndex,
};
