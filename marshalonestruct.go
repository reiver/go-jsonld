package jsonld

import (
	"reflect"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-json"
)


func marshalOneStruct(value any) ([]byte, error) {
	result, err := nakedMarshalOneStruct(value)
	if nil != err {
		return nil, err
	}

	var bytes [] byte

	bytes = append(bytes, '{')
	bytes = append(bytes, result...)
	bytes = append(bytes, '}')

	return bytes, nil
}

func nakedMarshalOneStruct(value any) ([]byte, error) {

	var reflectedStructValue reflect.Value = reflect.ValueOf(value)
	var reflectedStructType  reflect.Type  = reflect.TypeOf(value)

	{
		var kind reflect.Kind = reflectedStructType.Kind()

		if reflect.Struct != kind {
			return nil, erorr.Errorf("jsonld: cannot marshal something of type %T, expected a struct", value)
		}
	}

	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	{
		var comma bool = false

		var limit int = reflectedStructValue.NumField()

		for index:=0; index<limit; index++ {

			var reflectedStructFieldValue reflect.Value       = reflectedStructValue.Field(index)
			var reflectedStructFieldType  reflect.StructField = reflectedStructType.Field(index)

			if !reflectedStructFieldType.IsExported() {
				continue
			}

			switch reflectedStructFieldValue.Interface().(type) {
			case NameSpace:
				continue
			case Prefix:
				continue
			}

			name, omitEmpty := parseStructField(reflectedStructFieldType)

			if omitEmpty && isSimpleEmpty(reflectedStructFieldValue.Interface()) {
				continue
			}

			if comma {
				bytes = append(bytes, ',')
			}

			bytes = append(bytes, json.MarshalString(name)...)

			bytes = append(bytes, ':')

			{
				result, err := marshalOne(reflectedStructFieldValue.Interface())
				if nil != err {
					return nil, erorr.Errorf("jsonld: problem marshaling field №%d (%q) of struct %T: %w", 1+index, reflectedStructFieldType.Name, value, err)
				}

				bytes = append(bytes, result...)
			}

			comma = true
		}
	}

	return bytes, nil
}
