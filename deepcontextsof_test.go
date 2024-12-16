package jsonld_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-jsonld"
)

func TestDeepContextsOf(t *testing.T) {

	tests := []struct{
		Value any
		Expected []jsonld.Context
	}{
		{
			Value: struct{
			}{},
		},



		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"http://ns.example/id/42"`
			}{},
			Expected: []jsonld.Context{
				jsonld.Context{
					NameSpace: `http://ns.example/id/42`,
				},
			},
		},



		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"http://ns.example/id/42"`

				Something struct{
					NameSpace jsonld.NameSpace `jsonld:"http://example.com/ns"`
				} `json:"something,omitempty"`
			}{},
			Expected: []jsonld.Context{
				jsonld.Context{
					NameSpace: `http://ns.example/id/42`,
					Names: []string{"something"},
				},
				jsonld.Context{
					NameSpace: `http://example.com/ns`,
				},
			},
		},



		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"http://ns.example/id/42"`

				Something any
			}{
				Something: struct{
					NameSpace jsonld.NameSpace `jsonld:"http://example.com/ns"`
				}{},
			},
			Expected: []jsonld.Context{
				jsonld.Context{
					NameSpace: `http://ns.example/id/42`,
					Names: []string{"something"},
				},
				jsonld.Context{
					NameSpace: `http://example.com/ns`,
				},
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonld.DeepContextsOf(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual contexts are not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%#v", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%#v", len(actual), actual)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}
	}
}
