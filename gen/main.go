package main

import (
	"fmt"
	"github.com/ayang64/rtb"
	"log"
	"reflect"
	"strings"
)

type Column struct {
	Name string
	Type string
}

func GenerateTable(name, v interface{}) error {
	typeMap := map[string]string{
		"int":    "integer",
		"string": "text",
	}

	t := reflect.TypeOf(v)

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("type %q is not a struct", t.Name())
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

	// determine max lenght of field name.
	max := func() int {
		m := -1
		for i := range cols {
			if len(cols[i].Name) > m {
				m = len(cols[i].Name)
			}
		}
		return m
	}()

	fmt.Printf("drop table if exists %q cascade;\n", name)
	fmt.Printf("create table %q (\n", name)
	fmt.Printf("\t%-*.*s   serial,\n", max, max, "id")
	for i := range cols {
		fmt.Printf("\t%-*.*s   %s", max, max, cols[i].Name, cols[i].Type)
		if i < len(cols)-1 {
			fmt.Printf(",")
		}
		fmt.Printf("\n")
	}
	fmt.Printf(");\n")

	return nil
}

func main() {
	if err := GenerateTable("bid_request", []int{1}[0]); err != nil {
		log.Fatalf("%v", err)
	}
	GenerateTable("bid_request", rtb.BidRequest{})
	GenerateTable("offer", rtb.Offer{})
}
