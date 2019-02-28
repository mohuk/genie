package manager

import (
	"fmt"

	"github.com/mohuk/genie/errors"
	"github.com/mohuk/genie/models"
)

// GetColumns ...
func (g *genieManager) GetColumns(dbname, tableName string) (*models.TableForm, error) {

	columns, err := g.store.GetColumns(dbname, tableName)
	if err != nil {
		return nil, err
	}
	if len(columns) == 0 {
		return nil, errors.NewErrNoRows(fmt.Sprintf("table %s does not exist", tableName))
	}
	tf := models.TableForm{
		TableName: tableName,
		Template:  make([]models.Template, len(columns)),
	}
	for j, col := range columns {
		tf.Template[j] = models.Template{
			Key:  col.Name.String,
			Type: "input",
			TemplateOps: models.TemplateOpts{
				Type:        models.Type(col.Type.String),
				PlaceHolder: fmt.Sprintf("Enter %s...", col.Name.String),
			},
		}
	}
	return &tf, nil
}
