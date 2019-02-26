package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/mohuk/genie/dbase"
	"github.com/mohuk/genie/models"
)

func GetColumns(store dbase.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbname := mux.Vars(r)["dbname"]
		tableName := mux.Vars(r)["tableId"]
		columns, err := store.GetColumns(dbname, tableName)
		if err != nil {
			boom.Internal(w, err)
			return
		}
		tf := models.TableForm{TableName: tableName, Template: []models.Template{}}
		for _, col := range columns {
			tf.Template = append(tf.Template, models.Template{
				Key:  col.Name.String,
				Type: "input",
				TemplateOps: models.TemplateOpts{
					Type:        col.Type.String,
					PlaceHolder: fmt.Sprintf("Enter %s...", col.Name.String),
				},
			})
		}
		err = json.NewEncoder(w).Encode(tf)
		if err != nil {
			boom.Internal(w, err)
			return
		}

	}
}
