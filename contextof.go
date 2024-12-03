package jsonld

import (
	"reflect"

	"github.com/reiver/go-erorr"
)

func ContextOf(value any) (Context, error) {
	if nil == value {
		var nada Context
		return nada, errNilValue
	}

	{
		var reflectedType reflect.Type = reflect.TypeOf(value)

		var kind reflect.Kind = reflectedType.Kind()

		switch kind {
		case reflect.Struct:
			return contextOfStruct(value)
		default:
			var nada Context
			return nada, erorr.Errorf("jsonld: cannot get context-of something of kind %q", kind)
		}
	}
}
