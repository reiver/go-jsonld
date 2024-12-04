package jsonld

import (
	"fmt"
	"reflect"
)

func forStructFields(fn func(string,any,bool)error, strct any) error {
	if nil == fn {
		return errNilFunc
	}

	var reflectedStructValue reflect.Value = reflect.ValueOf(strct)
	var reflectedStructType  reflect.Type  = reflect.TypeOf(strct)

	if reflect.Struct != reflectedStructValue.Kind() {
		panic(fmt.Sprintf("jsonld: not struct — is actually a %T", strct))
	}

	var limit int = reflectedStructValue.NumField()

	for index:=0; index < limit; index++ {
		var reflectedStructFieldValue reflect.Value = reflectedStructValue.Field(index)
		var reflectedStructFieldType reflect.StructField = reflectedStructType.Field(index)

		if !reflectedStructFieldType.IsExported() {
			continue
		}

		var value any = reflectedStructFieldValue.Interface()

		switch value.(type) {
		case NameSpace:
			continue
		case Prefix:
			continue
		}

		var structField reflect.StructField = reflectedStructType.Field(index)

		if !structField.IsExported() {
			continue
		}

		var name string
		var omitEmpty bool
		name, omitEmpty = parseStructField(structField)

		err := fn(name, value, omitEmpty)
		if nil != err {
			return err
		}
	}

	return nil
}
