package jsonld_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
)

func TestContexts_MarshalJSON(t * testing.T) {

	tests := []struct{
		Contexts jsonld.Contexts
		Expected []byte
	}{
		{
			Expected: []byte(`{}`),
		},



		{
			Contexts: []jsonld.Context{
				jsonld.Context{
					NameSpace: "http://example.com/ns#",
				},
			},
			Expected: []byte(`["http://example.com/ns#"]`),
		},
		{
			Contexts: []jsonld.Context{
				jsonld.Context{
					NameSpace: "http://example.com/ns#",
				},
				jsonld.Context{
					NameSpace: "http://id.example/ns/1#",
				},
			},
			Expected: []byte(`["http://example.com/ns#","http://id.example/ns/1#"]`),
		},



		{
			Contexts: []jsonld.Context{
				jsonld.Context{
					Prefix: "something",
					Names: []string{"apple","banana","cherry"},
				},
			},
			Expected: []byte(`{"apple":"something:apple","banana":"something:banana","cherry":"something:cherry"}`),
		},



		{
			Contexts: []jsonld.Context{
				jsonld.Context{
					NameSpace: "http://example.com/ns#",
					Prefix: "something",
					Names: []string{"apple","banana","cherry"},
				},
			},
			Expected: []byte(`{"something":"http://example.com/ns#","apple":"something:apple","banana":"something:banana","cherry":"something:cherry"}`),
		},



		{
			Contexts: []jsonld.Context{
				jsonld.Context{
					NameSpace: "http://example.com/ns#",
				},
				jsonld.Context{
					NameSpace: "http://id.example/ns/1#",
				},
				jsonld.Context{
					NameSpace: "http://somewhere.example/namespace#",
					Prefix: "something",
					Names: []string{"apple","banana","cherry"},
				},
			},
			Expected: []byte(`["http://example.com/ns#","http://id.example/ns/1#",{"something":"http://somewhere.example/namespace#","apple":"something:apple","banana":"something:banana","cherry":"something:cherry"}]`),
		},
	}

	for testNumber, test := range tests {

		actual, err := test.Contexts.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual jsonld-marshaled contexts is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			continue
		}
	}
}
