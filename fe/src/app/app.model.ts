export interface Database {
  name: string;
}

export interface Table {
  Name: string;
  catalog: string;
  schema: string;
  type: string;
}

export interface ColumnForm {
  tableName?: string;
  formlyTemplate: FormlyTemplate[];
}

export interface FormlyTemplate {
  key: string;
  templateOptions: TemplateOptions;
  type: string;
}

interface TemplateOptions {
  placeholder: string;
  type: string;
}
