export function decodeParams(params) {
  var result = [];
  params.forEach((param) => {
    param.value = param.default;
    result.push(param);
  });
  return result;
}

export function encodeParams(params) {
  var obj = {};
  for (var j = 0, len = params.length; j < len; j++) {
    const param = params[j];
    var value = param.value;
    if (param.type != "string") {
      value = value.toString();
    }
    obj[param.key] = value;
  }
  return JSON.stringify(obj);
}
