package formly

import (
	"database/sql"
	"testing"

	"github.com/mohuk/genie/models"
)

func TestHasMaxLength(t *testing.T) {
	cases := []struct {
		Expected bool
		Input    string
	}{
		{Expected: true, Input: "15"},
		{Expected: false, Input: "-15"},
	}

	for _, c := range cases {
		_, valid := hasValidMaxLength(c.Input)
		if valid != c.Expected {
			t.Fail()
		}
	}
}

func TestMapColumn(t *testing.T) {
	colName := "colName"
	length := "15"
	nullable := "YES"
	someT := "text"
	isDefault := "NULL"
	c := models.Column{
		Name: sql.NullString{
			String: colName,
		},
		Length: sql.NullString{
			String: length,
		},
		Type: sql.NullString{
			String: someT,
		},
		Nullable: sql.NullString{
			String: nullable,
		},
		Default: sql.NullString{
			String: isDefault,
		},
	}
	m := NewFormlyMapper().MapColumn(c)
	if m.TemplateOps.Type != someT {
		t.Fail()
		return
	}
	if m.TemplateOps.Max != 15 {
		t.Fail()
		return
	}
	if !m.TemplateOps.Required {
		t.Fail()
		return
	}
	if m.Default {
		t.Fail()
		return
	}
}
