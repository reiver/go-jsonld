package jsonld

import (
	"reflect"
)

func forStructFields(fn func(string,any), strct any) {
	if nil == fn {
		return
	}

	var reflectedStructValue reflect.Value = reflect.ValueOf(strct)
	var reflectedStructType  reflect.Type  = reflect.TypeOf(strct)

	if reflect.Struct != reflectedStructValue.Kind() {
		panic("jsonld: not struct")
	}

	var limit int = reflectedStructValue.NumField()

	for index:=0; index < limit; index++ {
		var reflectedFieldValue reflect.Value = reflectedStructValue.Field(index)
		var value any = reflectedFieldValue.Interface()

		switch value.(type) {
		case NameSpace:
			continue
		case Prefix:
			continue
		}

		var structField reflect.StructField = reflectedStructType.Field(index)
		var name string = structFieldName(structField)

		fn(name, value)
	}
}
