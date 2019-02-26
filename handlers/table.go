package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mohuk/genie/httpio"

	"github.com/darahayes/go-boom"

	"github.com/gorilla/mux"

	"github.com/mohuk/genie/dbase"
)

// GetTables ...
func GetTables(store dbase.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbname := mux.Vars(r)["dbname"]
		tables, err := store.GetTables(dbname)
		if err != nil {
			boom.Internal(w, err)
			return
		}
		var resp []httpio.Table
		for _, t := range tables {
			resp = append(resp, httpio.UnmarshallTable(t))
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			boom.Internal(w, err)
			return
		}
	}
}
