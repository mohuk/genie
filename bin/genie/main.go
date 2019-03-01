package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohuk/genie/dbase"
	"github.com/mohuk/genie/formly"
	"github.com/mohuk/genie/handlers"
	"github.com/mohuk/genie/manager"
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

func main() {

	store := dbase.NewStore(host, port, user, password)
	mapper := formly.NewFormlyMapper()
	gm := manager.NewGenieManager(store, mapper)

	APIRouter := mux.NewRouter().PathPrefix("/api").Subrouter()

	APIRouter.HandleFunc("/db", handlers.WithLogger(handlers.GetDatabases(gm))).Methods("GET")
	APIRouter.HandleFunc("/db/{dbname}/tables", handlers.GetTables(gm)).Methods("GET")
	APIRouter.HandleFunc("/db/{dbname}/tables/{tableId}", handlers.GetColumns(gm)).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), APIRouter))
}

func init() {
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
			log.Fatal("invalid port")
		}
		port = intPort
	}
}
