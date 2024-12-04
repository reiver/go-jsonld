package jsonld

import (
	gobytes "bytes"
	"reflect"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-json"
)

var emptyJSON []byte = []byte(`{}`)

// Marshal returns the (merged) JSON-LD encoding of a series of values.
//
// Example usage:
//
//	bytes, err := jsonld.Marshal(activitypub, activitystreams, security, toot, schema)
func Marshal(values ...any) ([]byte, error) {
	for index, value := range values {
		var reflectedType reflect.Type = reflect.TypeOf(value)
		var kind reflect.Kind = reflectedType.Kind()
		if reflect.Struct != kind  {
			return nil, erorr.Errorf("jsonld: cannot marshal value №%d of type %T — type must be struct", 1+index, value)
		}
	}

	var contexts []Context
	{
		for _, value := range values {
			context, err := ContextOf(value)
			if nil != err {
				return nil, err
			}

			contexts = append(contexts, context)
		}
	}

	return marshal(contexts, values...)
}

func marshal(contexts []Context, values ...any) ([]byte, error) {
	if 1 == len(values) {
		var value any = values[0]
		if isSimpleType(value) {
			return marshalSimpleType(value)
		}
	}

	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	bytes = append(bytes, '{')

	var hasContext bool = false
	if 0 < len(contexts) {
		result, err := MarshalContexts(contexts...)
		if nil != err {
			return nil, err
		}

		if !gobytes.Equal(emptyJSON, result) {
			bytes = append(bytes, json.MarshalString("@context")...)
			bytes = append(bytes, ':')
			bytes = append(bytes, result...)
			hasContext = true
		}
	}

	var comma bool = false
	if hasContext {
		comma = true
	}

	{
		for _, value := range values {
			err := forStructFields(func(name string, value any, omitEmpty bool)error{
				if omitEmpty && isSimpleEmpty(value) {
					return nil
				}

				var val []byte
				{
					var err error
					val, err = marshalOne(value)
					if omitEmpty {
						switch err.(type) {
						case EmptyError:
							return nil
						}
					}
					if nil != err {
						return err
					}
				}

				if comma {
					bytes = append(bytes, ',')
				}
				comma = true

				bytes = append(bytes, json.MarshalString(name)...)
				bytes = append(bytes, ':')
				bytes = append(bytes, val...)

				return nil
			}, value)
			if nil != err {
				return nil, err
			}
		}
	}

	bytes = append(bytes, '}')

	return bytes, nil
}
