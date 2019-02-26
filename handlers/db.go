package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/mohuk/genie/dbase"
	"github.com/mohuk/genie/httpio"
)

// GetDatabases ...
func GetDatabases(store dbase.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbs, err := store.GetDatabases()
		if err != nil {
			boom.Internal(w, err)
			return
		}

		var resp []httpio.Database
		for _, x := range dbs {
			resp = append(resp, httpio.UnmarshallDB(x))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			boom.Internal(w, err)
			return
		}
	}
}
