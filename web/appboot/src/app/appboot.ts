export interface TemplateConfig {
  name: string;
  desc: string;
  parameters: Parameter[];
  scripts: {
    before?: string[];
    after?:  string[];
  };
}

export interface TemplateGroup {
  id: string;
  desc: string;
  templates: Template[]
}

export interface Template {
  id: string;
  desc: string;
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