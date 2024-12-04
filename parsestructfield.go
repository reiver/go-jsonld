package jsonld

import (
	"reflect"
	"strings"
)

func parseStructField(structField reflect.StructField) (name string, omitEmpty bool) {

	name = structField.Name

	tagvalue, found := structField.Tag.Lookup(structTagName)
	if !found {
		return
	}

	a := strings.Split(tagvalue, ",")
	if len(a) <= 0 {
		return
	}

	if candidateName := a[0]; "" != candidateName {
//@TODO: should validate this name.
		name = candidateName
	}

	for _, token := range a[1:] {
		switch token {
		case "omitempty":
			omitEmpty = true
		}
	}

	name = strings.TrimSpace(name)
	return
}

func structFieldName(structField reflect.StructField) (name string) {
	name, _ = parseStructField(structField)
	return name
}
