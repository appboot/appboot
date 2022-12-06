import { Parameter } from "./appboot";

export function decodeParams(params: Parameter[]): Parameter[] {
  const result: Parameter[] = [];
  params.forEach((param: Parameter) => {
    param.value = param.default;
    result.push(param);
  });
  return result;
}

export const encodeParams = (params: Parameter[]): string => {
  const obj: any = {};
  for (let j = 0, len = params.length; j < len; j++) {
    const param = params[j];
    let value = param.value;
    if (param.type !== 'string') {
      value = value ? value.toString(): "";
    }
    obj[param.key] = value;
  }
  return JSON.stringify(obj);
};
