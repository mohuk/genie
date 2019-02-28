package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/mohuk/genie/handlers"
	"github.com/mohuk/genie/manager"

	"github.com/gorilla/mux"

	"github.com/mohuk/genie/dbase"
)

var (
	debug    = flag.Bool("debug", false, "enable debugging")
	password = flag.String("password", "10Pearls!", "the database password")
	port     = flag.Int("port", 1401, "the database port")
	host     = flag.String("host", "localhost", "the database host")
	user     = flag.String("user", "sa", "the database user")
	database = flag.String("database", "WideWorldImporters", "the database name")
)

const (
	httpPort = 3000
)

func main() {
	store := dbase.NewStore(*host, *port, *user, *password)
	gm := manager.NewGenieManager(store)

	APIRouter := mux.NewRouter().PathPrefix("/api").Subrouter()

	APIRouter.HandleFunc("/db", handlers.WithLogger(handlers.GetDatabases(gm))).Methods("GET")
	APIRouter.HandleFunc("/db/{dbname}/tables", handlers.GetTables(gm)).Methods("GET")
	APIRouter.HandleFunc("/db/{dbname}/tables/{tableId}", handlers.GetColumns(gm)).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), APIRouter))
}
