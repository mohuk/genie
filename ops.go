package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

// MSSqlDatabase class as a struct
type MSSqlDatabase struct {
	Host         string
	Port         int
	DatabaseName string
	User         string
	Password     string
	Conn         *sql.DB
}

// Connect to initialize database connection
func (db *MSSqlDatabase) Connect() error {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", db.Host, db.User, db.Password, db.Port, db.DatabaseName)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		return err
	}
	db.Conn = conn
	return nil
}

// Close connection to DB
func (db *MSSqlDatabase) Close() error {
	return db.Conn.Close()
}

// ListTables return the tables for database
func (db *MSSqlDatabase) ListTables() ([]Table, error) {
	qry := fmt.Sprintf("SELECT DISTINCT TABLE_NAME, TABLE_CATALOG, TABLE_SCHEMA, TABLE_TYPE FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_CATALOG = '%s' AND TABLE_TYPE = 'BASE TABLE'", db.DatabaseName)
	rows, err := db.Conn.Query(qry)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables = []Table{}
	for rows.Next() {
		var tbl = Table{}
		err := rows.Scan(&tbl.Name, &tbl.Catalog, &tbl.Schema, &tbl.Type)

		if err != nil {
			return nil, err
		}

		tables = append(tables, tbl)
	}

	return tables, nil
}

// ListColumns returns the columns for a table
func (db *MSSqlDatabase) ListColumns(table string) ([]Column, error) {
	qry := fmt.Sprintf("SELECT DISTINCT COLUMN_NAME, COLUMN_DEFAULT, IS_NULLABLE, DATA_TYPE, CHARACTER_MAXIMUM_LENGTH FROM INFORMATION_SCHEMA.COLUMNS where TABLE_NAME = '%s' AND TABLE_CATALOG = '%s'", table, db.DatabaseName)
	rows, err := db.Conn.Query(qry)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var columns = []Column{}
	for rows.Next() {
		var col = Column{}
		err := rows.Scan(&col.Name, &col.Default, &col.Nullable, &col.Type, &col.Length)

		if err != nil {
			return nil, err
		}

		columns = append(columns, col)
	}

	return columns, nil
}
