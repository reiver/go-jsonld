package jsonld

import (
	"reflect"
	"sort"

	"github.com/reiver/go-erorr"
)

func marshalOneMap(value any) ([]byte, error) {
	result, err := nakedMarshalOneMap(value)
	if nil != err {
		return nil, err
	}

	var bytes [] byte

	bytes = append(bytes, '{')
	bytes = append(bytes, result...)
	bytes = append(bytes, '}')

	return bytes, nil
}

func nakedMarshalOneMap(value any) ([]byte, error) {
	var reflectedValue reflect.Value = reflect.ValueOf(value)

	{
		var kind reflect.Kind = reflectedValue.Kind()

		if reflect.Map != kind {
			return nil, erorr.Errorf("jsonld: cannot marshal something of type %T, expected a map", value)
		}
	}

	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	{
		var reflectedKeys []reflect.Value = reflectedValue.MapKeys()

		{
			var fn = func(index1, index2 int) bool {
				value1 := reflectedKeys[index1]
				value2 := reflectedKeys[index2]

				return value1.String() < value2.String()
			}

			sort.Slice(reflectedKeys, fn)
		}

		for index, reflectedKey := range reflectedKeys {
			var reflectedDatum reflect.Value = reflectedValue.MapIndex(reflectedKey)

			if 0 < index {
				bytes = append(bytes, ',')
			}

			if reflect.String != reflectedDatum.Kind() {
				return nil, erorr.Errorf("jsonld: map key must be of type 'string', cannot be of type %T", reflectedDatum.Interface())
			}

			{
				result, err := marshalOne(reflectedKey.Interface())
				if nil != err {
					return nil, err
				}

				bytes = append(bytes, result...)

			}

			bytes = append(bytes, ':')

			{
				result, err := marshalOne(reflectedDatum.Interface())
				if nil != err {
					return nil, err
				}

				bytes = append(bytes, result...)
			}
		}
	}

	return bytes, nil
}
