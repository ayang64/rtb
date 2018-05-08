package main

import (
	"fmt"
	"github.com/ayang64/rtb"
	"reflect"
	"strings"
)

type Table struct {
	Name    string
	Columns []Column
}

func (t *Table) String() string {
	if t == nil {
		return "*nil*"
	}

	// determine max lenght of field name.
	max := func() int {
		m := -1
		for _, col := range t.Columns {
			if len(col.Name) > m {
				m = len(col.Name)
			}
		}
		return m
	}()

	rc := fmt.Sprintf("drop table if exists %q cascade;\n", t.Name) + fmt.Sprintf("create table %q (\n", t.Name) + fmt.Sprintf("\t%-*.*s   serial,\n", max, max, "id")

	for i, col := range t.Columns {
		rc += fmt.Sprintf("\t%-*.*s   %s", max, max, col.Name, col.Type)
		if i < len(t.Columns)-1 {
			rc += fmt.Sprintf(",")
		}
		rc += fmt.Sprintf("\n")
	}
	rc += fmt.Sprintf(");\n")

	return rc
}

type Column struct {
	Name string
	Type string
}

func GenerateTable(name string, v interface{}) (*Table, error) {

	typeMap := map[string]string{
		"int":    "integer",
		"string": "text",
	}

	t := reflect.TypeOf(v)

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("type %q is not a struct", t.Name())
	}

	var cols []Column

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag

		r := func() string {
			if val, exists := tag.Lookup("col"); exists == true {
				return strings.Split(val, ",")[0]
			}

			if val, exists := tag.Lookup("json"); exists == true {
				return strings.Split(val, ",")[0]
			}

			return "*nothing*"
		}()

		colType := func() string {
			if val, exists := tag.Lookup("col"); exists == true {
				r := strings.Split(val, ",")
				if len(r) > 1 {
					return r[1]
				}
			}
			rc, exists := typeMap[t.Field(i).Type.Name()]
			if exists == true {
				return rc
			}
			return "*unknown*"
		}()

		cols = append(cols, Column{Name: r, Type: colType})
	}

	return &Table{Name: name, Columns: cols}, nil
}

func main() {
	br, _ := GenerateTable("bid_request", rtb.BidRequest{})

	fmt.Printf("%s", br)
	GenerateTable("offer", rtb.Offer{})
}
