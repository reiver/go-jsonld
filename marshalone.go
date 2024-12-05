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

	var result []byte
	{
		var err error

		switch kind {
		case reflect.Array,reflect.Slice:
			return marshalOneSlice(reflectedValue.Interface())
		case reflect.Map,reflect.Struct:
			result, err = nakedMarshalOne(reflectedValue.Interface())

		default:
			return nil, erorr.Errorf("jsonld: cannot marshal something of type %T", value)
		}
		if nil != err {
			return nil, err
		}
	}

	var bytes []byte

	bytes = append(bytes, '{')
	bytes = append(bytes, result...)
	bytes = append(bytes, '}')

	return bytes, nil
}

func nakedMarshalOne(value any) ([]byte, error) {
	var reflectedValue reflect.Value = reflect.ValueOf(value)
	var kind reflect.Kind = reflectedValue.Kind()

	switch kind {
	case reflect.Map:
		return nakedMarshalOneMap(reflectedValue.Interface())
	case reflect.Struct:
		return nakedMarshalOneStruct(reflectedValue.Interface())

	default:
		return nil, erorr.Errorf("jsonld: cannot marshal something of type %T", value)
	}
}
