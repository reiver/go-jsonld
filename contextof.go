package jsonld

import (
	"reflect"

	"github.com/reiver/go-erorr"
)

// ContextOf returns the JSON-LD context.
//
// For example:
//
//	type MyStruct struct {
//		NameSpace jsonld.NameSpace `jsonld:"http://example.com/ns#"`
//		Preix     jsonld.Prefix    `jsonld:"ex"`
//		
//		Apple  string `json:"apple"`
//		Banana int    `json:"banana"`
//		Cherry bool   `json:"cherry"`
//	}
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
