export interface Template {
  name: string;
  desc: string;
  parameters: Parameter[];
  scripts: {
    before?: string[];
    after?:  string[];
  };
}

export interface Parameter {
  key: string
  type: 'string'|'int'|'float'|'select'
  tip?: string
  options? : string[] 
  default?: string | number
  placeholder?: string
  min?: number
  max?: number
  // 前端使用的值
  value?: string | number
}