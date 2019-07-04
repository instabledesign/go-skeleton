package config

import (
	"bytes"
	"fmt"
	"reflect"
	"text/tabwriter"
)

func ToString(c interface{}) string {
	b := &bytes.Buffer{}
	w := tabwriter.NewWriter(b, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	val := reflect.ValueOf(c).Elem()
	fmt.Fprint(w, "\n-----------------------------------\n---  Application configuration  ---\n")
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := val.Type().Field(i)

		fmt.Fprintf(w, "%s\t\033[0m%v\t\033[1;34m%s\033[0m \033[1;92m`%s`\033[0m\n", typeField.Name, field.Interface(), field.Type().String(), typeField.Tag)
	}
	fmt.Fprint(w, "-----------------------------------\n")
	w.Flush()

	return b.String()
}
