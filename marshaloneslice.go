package jsonld

import (
	"reflect"

	"github.com/reiver/go-erorr"
)

func marshalOneSlice(value any) ([]byte, error) {
	var reflectedValue reflect.Value = reflect.ValueOf(value)

	{
		var kind reflect.Kind = reflectedValue.Kind()

		if reflect.Slice != kind && reflect.Array != kind {
			return nil, erorr.Errorf("jsonld: cannot marshal something of type %T, expected a slice or an array", value)
		}
	}

	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	bytes = append(bytes, '[')

	{
		var limit int = reflectedValue.Len()

		for index:=0; index<limit; index++ {

			var reflectedDatum reflect.Value = reflectedValue.Index(index)

			if 0 < index {
				bytes = append(bytes, ',')
			}

			{
				result, err := marshalOne(reflectedDatum.Interface())
				if nil != err {
					return nil, err
				}

				bytes = append(bytes, result...)
			}
		}
	}

	bytes = append(bytes, ']')

	return bytes, nil
}
