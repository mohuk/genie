package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	debug    = flag.Bool("debug", false, "enable debugging")
	password = flag.String("password", "10Pearls!", "the database password")
	port     = flag.Int("port", 1401, "the database port")
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
	defer db.Conn.Close()

	tables, err := db.ListTables()
	if err != nil {
		log.Fatal(err)
	}
	var tforms []tableForm
	for _, t := range tables {
		tf := tableForm{TableName: t.Name.String, Template: []template{}}
		cols, err := db.ListColumns(t.Name.String)
		if err != nil {
			log.Fatal(err)
		}
		for _, col := range cols {
			tf.Template = append(tf.Template, template{
				Key:  col.Name.String,
				Type: "input",
				TemplateOps: templateOpts{
					Type:        col.Type.String,
					PlaceHolder: fmt.Sprintf("Enter %s...", col.Name.String),
				},
			})
		}
		tforms = append(tforms, tf)
	}

	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(tforms, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("temp.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	println("created temp.json")

}
