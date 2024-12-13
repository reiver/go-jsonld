package jsonld_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

func TestPrefix_JSONOmitAlways(t *testing.T) {

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: struct{
				Prefix jsonld.Prefix
			}{},
			Expected: []byte(`{}`),
		},
		{
			Value: struct{
				Prefix jsonld.Prefix
				Apple bool
			}{},
			Expected: []byte(`{"Apple":false}`),
		},
		{
			Value: struct{
				Prefix jsonld.Prefix
				Apple bool
				Banana int `json:"banana"`
			}{},
			Expected: []byte(`{"Apple":false,"banana":0}`),
		},
		{
			Value: struct{
				Prefix jsonld.Prefix
				Apple bool
				Banana int    `json:"banana"`
				Cherry string `json:"cherry,omitempty"`
			}{},
			Expected: []byte(`{"Apple":false,"banana":0}`),
		},
		{
			Value: struct{
				Prefix jsonld.Prefix
				Apple bool
				Banana int    `json:"banana"`
				Cherry string `json:"cherry,omitempty"`
			}{
				Apple: true,
				Banana: -5,
				Cherry: "wow",
			},
			Expected: []byte(`{"Apple":true,"banana":-5,"cherry":"wow"}`),
		},









		{
			Value: struct{
				Prefix jsonld.Prefix `jsonld:"http://example.com/ns"`
			}{},
			Expected: []byte(`{}`),
		},
		{
			Value: struct{
				Prefix jsonld.Prefix `jsonld:"http://example.com/ns"`
				Apple bool
			}{},
			Expected: []byte(`{"Apple":false}`),
		},
		{
			Value: struct{
				Prefix jsonld.Prefix `jsonld:"http://example.com/ns"`
				Apple bool
				Banana int `json:"banana"`
			}{},
			Expected: []byte(`{"Apple":false,"banana":0}`),
		},
		{
			Value: struct{
				Prefix jsonld.Prefix `jsonld:"http://example.com/ns"`
				Apple bool
				Banana int    `json:"banana"`
				Cherry string `json:"cherry,omitempty"`
			}{},
			Expected: []byte(`{"Apple":false,"banana":0}`),
		},
		{
			Value: struct{
				Prefix jsonld.Prefix `jsonld:"http://example.com/ns"`
				Apple bool
				Banana int    `json:"banana"`
				Cherry string `json:"cherry,omitempty"`
			}{
				Apple: true,
				Banana: -5,
				Cherry: "wow",
			},
			Expected: []byte(`{"Apple":true,"banana":-5,"cherry":"wow"}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.Marshal(test.Value) // <--------- note that is json.Marshal() and not jsonld.Marshal()
		if nil != err{
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("EROR: (%T) %s", err, err)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual json-marshaled (and NOT jsonld-marshaled) value is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}
	}
}
