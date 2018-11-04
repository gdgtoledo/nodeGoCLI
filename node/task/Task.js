const colors = require('colors');
class Task {
  constructor(description, isComplete = false) {
    this.description = description;
    this.isComplete = isComplete;
  }
}

exports.Task = Task;
