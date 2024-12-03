package jsonld

import (
	"reflect"
)

func structFieldName(structField reflect.StructField) (name string) {

	var found bool
	name, found = structField.Tag.Lookup(structTagName)
	if !found {
		name = structField.Name
	}

	return name
}
