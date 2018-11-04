const get = (isComplete) => {
  let icon = '';

  if (isComplete) {
    icon = '☑'.green;
    return icon;
  }

  icon = '□';
  return icon;
};

exports.get = get;
