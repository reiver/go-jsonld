package jsonld

import (
	"bytes"
	"reflect"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-json"
)

var emptyJSON []byte = []byte(`{}`)

var badPrefix1 []byte = []byte(`{"@context":{},`)
var badPrefix2 []byte = []byte(`{"@context":{}}`)

// Marshal returns the (merged) JSON-LD encoding of a series of values.
//
// Example usage:
//
//	bytes, err := jsonld.Marshal(activitypub, activitystreams, security, toot, schema)
func Marshal(values ...any) ([]byte, error) {
	for index, value := range values {
		var reflectedType reflect.Type = reflect.TypeOf(value)
		var kind reflect.Kind = reflectedType.Kind()
		if reflect.Struct != kind && reflect.Map != kind {
			switch value.(type) {
			case json.Marshaler:
				// this is ok
			default:
				return nil, erorr.Errorf("jsonld: cannot marshal value №%d of type %T — type must be struct or map", 1+index, value)
			}
		}
	}

	var contexts []Context
	{
		var contextsMap map[string]*Context = map[string]*Context{}

		for _, value := range values {
			if reflect.Struct != reflect.TypeOf(value).Kind() {
				continue
			}

			ctxs, err := DeepContextsOf(value)
			if nil != err {
				return nil, err
			}

			for _, ctx := range ctxs {
				var ns string = ctx.NameSpace

				contextPointer, found := contextsMap[ns]
				if found {
					contextPointer.Names = append(contextPointer.Names)
				} else {
					contexts = append(contexts, ctx)
					contextsMap[ctx.NameSpace] = &(contexts[len(contexts)-1])
				}
			}
		}
	}
//@TODO: sort names

	var ctx = struct{
		CTX Contexts `json:"@context,omitempty"`
	}{
		CTX: Contexts(contexts),
	}

	values = append([]any{ctx}, values...)

	{
		result, err := json.MergeAndMarshal(values...)

//@TODO: should not need to do this.
		if nil == err && bytes.HasPrefix(result,badPrefix1) {
			result = result[len(badPrefix1)-1:]
			result[0] = '{'
		}
		if nil == err && bytes.HasPrefix(result,badPrefix2) {
			result = result[len(badPrefix2)-2:]
			result[0] = '{'
		}

		return result, err
	}
}
