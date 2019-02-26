package models

import "database/sql"

type TableForm struct {
	TableName string     `json:"tableName,omitempty"`
	Template  []Template `json:"formlyTemplate,omitempty"`
}
type Template struct {
	Key         string       `json:"key,omitempty"`
	Type        string       `json:"type,omitempty"`
	TemplateOps TemplateOpts `json:"templateOptions,omitempty"`
}
type TemplateOpts struct {
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

// Column is a column
type Column struct {
	Name     sql.NullString
	Type     sql.NullString
	Nullable sql.NullString
	Default  sql.NullString
	Length   sql.NullString
}

// Table is a table
type Table struct {
	Catalog sql.NullString
	Schema  sql.NullString
	Name    sql.NullString
	Type    sql.NullString
}

type Database struct {
	DBName sql.NullString `json:"dbName,omitempty"`
}
