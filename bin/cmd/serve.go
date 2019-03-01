package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohuk/genie/dbase"
	"github.com/mohuk/genie/formly"
	"github.com/mohuk/genie/handlers"
	"github.com/mohuk/genie/manager"
	"github.com/urfave/cli"
)

var (
	password = os.Getenv("DB_PWD")
	user     = os.Getenv("DB_USER")
	portEnv  = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	database = os.Getenv("DB_NAME")
	port     int
	httpPort = 3000
)

// Serve ...
var Serve = cli.Command{
	Name:        "serve",
	Description: "spins up an http server",
	Before: func(c *cli.Context) error {
		if password == "" {
			password = "10Pearls!"
		}
		if user == "" {
			user = "sa"
		}
		if database == "" {
			database = "WideWorldImporters"
		}
		if host == "" {
			host = "localhost"
		}
		if portEnv == "" {
			port = 1401
		} else {
			intPort, err := strconv.Atoi(portEnv)
			if err != nil {
				return cli.NewExitError("invalid port", 1)
			}
			port = intPort
		}
		return nil
	},
	Action: func(c *cli.Context) error {

		store := dbase.NewStore(host, port, user, password)
		gm := manager.NewGenieManager(store, formly.NewFormlyMapper())

		API := mux.NewRouter().PathPrefix("/api").Subrouter()

		API.HandleFunc("/db", handlers.WithLogger(handlers.GetDatabases(gm))).Methods("GET")
		API.HandleFunc("/db/{dbname}/tables", handlers.WithLogger(handlers.GetTables(gm))).Methods("GET")
		API.HandleFunc("/db/{dbname}/tables/{tableId}", handlers.WithLogger(handlers.GetColumns(gm))).Methods("GET")

		return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), API)
	},
}
