package httpio

import (
	"github.com/mohuk/genie/models"
)

// Database database response struct
type Database struct {
	Name string `json:"name,omitempty"`
}

type Table struct {
	Catalog string
	Schema  string
	Name    string
	Type    string
}

// Column is a column
type Column struct {
	Name     string
	Type     string
	Nullable string
	Default  string
	Length   string
}

func UnmarshallDB(database models.Database) Database {
	return Database{Name: database.DBName.String}
}

func UnmarshallTable(t models.Table) Table {
	return Table{Name: t.Name.String, Catalog: t.Catalog.String, Schema: t.Schema.String, Type: t.Type.String}
}
