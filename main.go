package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	debug    = flag.Bool("debug", false, "enable debugging")
	password = flag.String("password", "10Pearls!", "the database password")
	port     = flag.Int("port", 1433, "the database port")
	host     = flag.String("host", "localhost", "the database host")
	user     = flag.String("user", "sa", "the database user")
	database = flag.String("database", "WideWorldImporters", "the database name")
)

func main() {
	flag.Parse()

	db := &MSSqlDatabase{
		Host:         *host,
		Port:         *port,
		User:         *user,
		Password:     *password,
		DatabaseName: *database,
	}

	err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}

	tables, err := db.ListTables()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tables)
	// columns, err := db.ListColumns(tables[0].Name.String)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Conn.Close()
}
