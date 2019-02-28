package manager

import (
	"github.com/mohuk/genie/dbase"
	"github.com/mohuk/genie/httpio"
	"github.com/mohuk/genie/models"
)

type GenieManager interface {
	GetColumns(dbname, tableName string) (*models.TableForm, error)
	GetDatabases() ([]httpio.Database, error)
	GetTables(dbname string) ([]httpio.Table, error)
}

type genieManager struct {
	store dbase.Store
}

func NewGenieManager(s dbase.Store) GenieManager {
	return &genieManager{store: s}
}
