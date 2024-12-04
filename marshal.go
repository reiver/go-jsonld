package jsonld

import (
	gobytes "bytes"

	"github.com/reiver/go-json"
)

var emptyJSON []byte = []byte(`{}`)

// Marshal returns the (merged) JSON-LD encoding of a series of values.
//
// Example usage:
//
//	bytes, err := jsonld.Marshal(activitypub, activitystreams, security, toot, schema)
func Marshal(values ...any) ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	bytes = append(bytes, '{')

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
					val, err = json.Marshal(value)
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
