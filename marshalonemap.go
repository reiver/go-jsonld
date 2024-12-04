package jsonld

import (
	"reflect"

	"github.com/reiver/go-erorr"
)

func marshalOneMap(value any) ([]byte, error) {
	var reflectedValue reflect.Value = reflect.ValueOf(value)

	{
		var kind reflect.Kind = reflectedValue.Kind()

		if reflect.Map != kind {
			return nil, erorr.Errorf("jsonld: cannot marshal something of type %T, expected a map", value)
		}
	}

	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	bytes = append(bytes, '{')

	{
		var reflectedKeys []reflect.Value = reflectedValue.MapKeys()

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

	bytes = append(bytes, '}')

	return bytes, nil
}
