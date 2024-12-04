package jsonld

import (
	"reflect"

	"github.com/reiver/go-erorr"
)

func marshalOne(value any) ([]byte, error) {
	if isSimpleType(value) {
		return marshalSimpleType(value)
	}

	var reflectedValue reflect.Value = reflect.ValueOf(value)
	var kind reflect.Kind = reflectedValue.Kind()

	switch kind {
	case reflect.Array,reflect.Slice:
		return marshalOneSlice(reflectedValue.Interface())
	case reflect.Map:
		return marshalOneMap(reflectedValue.Interface())
	case reflect.Struct:
		return marshalOneStruct(reflectedValue.Interface())

	default:
		return nil, erorr.Errorf("jsonld: cannot marshal someething of type %T", value)
	}
}
