package jsonld

import (
	"reflect"
)

func contextOfStruct(value any) (Context, error) {
	if nil == value {
		var nada Context
		return nada, errNilValue
	}

	var reflectedType reflect.Type
	{
		reflectedType  = reflect.TypeOf(value)
		if nil == reflectedType {
			var nada Context
			return nada, errNilReflectedType
		}

		if reflect.Struct != reflectedType.Kind() {
			var nada Context
			return nada, errNotStruct
		}
	}

	var reflectedValue reflect.Value = reflect.ValueOf(value)

	var context Context
	{
		var limit int = reflectedType.NumField()

		for index:=0; index<limit; index++ {

			var structField reflect.StructField = reflectedType.Field(index)
			var reflectedFieldValue reflect.Value = reflectedValue.Field(index)

			switch reflectedFieldValue.Interface().(type) {
			case NameSpace:
				value, found := structField.Tag.Lookup(structTagName)
				if found {
					context.NameSpace = value
				}

			case Prefix:
				value, found := structField.Tag.Lookup(structTagName)
				if found {
					context.Prefix = value
				}

			default:
				name := structFieldName(structField)

				context.Names = append(context.Names, name)
			}
		}
	}

	return context, nil
}
