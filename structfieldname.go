package jsonld

import (
	"reflect"
	"strings"
)

func structFieldName(structField reflect.StructField) (name string) {

	tagvalue, found := structField.Tag.Lookup(structTagName)
	if !found {
		name = structField.Name
		return
	}

	{
		a := strings.Split(tagvalue, ",")
		if 1 <= len(a) {
			name = a[0]
			return
		}
	}

	return name
}
