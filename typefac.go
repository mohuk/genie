package main

type tableForm struct {
	TableName string     `json:"tableName,omitempty"`
	Template  []template `json:"formlyTemplate,omitempty"`
}
type template struct {
	Key         string       `json:"key,omitempty"`
	Type        string       `json:"type,omitempty"`
	TemplateOps templateOpts `json:"templateOptions,omitempty"`
}
type templateOpts struct {
	Type        string `json:"type,omitempty"`
	Label       string `json:"label,omitempty"`
	PlaceHolder string `json:"placeholder,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

var typeMaps = map[string]string{
	"nvarchar":  "text",
	"datetime2": "date",
	"bigint":    "number",
	"geography": "number",
	"decimal":   "number",
	"date":      "date",
	"bit":       "checkbox",
	"varbinary": "checkbox",
	"int":       "number",
}

// Type ...
func Type(s string) string {
	if v, ok := typeMaps[s]; ok {
		return v
	}
	return "text"
}
