package jsonld

import (
	"testing"

	"bytes"
)

func TestMarshalOneStruct(t *testing.T) {

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: struct{}{},
			Expected: []byte(`{}`),
		},



		{
			Value: struct{
				Apple string
			}{
				Apple:"1",
			},
			Expected: []byte(`{"Apple":"1"}`),
		},
		{
			Value: struct{
				Apple string `jsonld:"apple"`
			}{
				Apple:"1",
			},
			Expected: []byte(`{"apple":"1"}`),
		},



		{
			Value: struct{
				Apple string
				Banana int
			}{
				Apple:"1",
				Banana:2,
			},
			Expected: []byte(`{"Apple":"1","Banana":2}`),
		},
		{
			Value: struct{
				Apple string `jsonld:"apple"`
				Banana int
			}{
				Apple:"1",
				Banana:2,
			},
			Expected: []byte(`{"apple":"1","Banana":2}`),
		},
		{
			Value: struct{
				Apple string `jsonld:"apple"`
				Banana int   `jsonld:"banana"`
			}{
				Apple:"1",
				Banana:2,
			},
			Expected: []byte(`{"apple":"1","banana":2}`),
		},



		{
			Value: struct{
				Apple  string `jsonld:"apple"`
				Banana int    `jsonld:"banana"`
				Cherry string `jsonld:"cherry"`
			}{
				Apple:"1",
				Banana:2,
				Cherry:"three",
			},
			Expected: []byte(`{"apple":"1","banana":2,"cherry":"three"}`),
		},



		{
			Value: struct{
				Apple  string `jsonld:"apple,omitempty"`
				Banana int    `jsonld:"banana"`
				Cherry string `jsonld:"cherry,omitempty"`
			}{
				Apple:"1",
				Banana:2,
				Cherry:"three",
			},
			Expected: []byte(`{"apple":"1","banana":2,"cherry":"three"}`),
		},
		{
			Value: struct{
				Apple  string `jsonld:"apple,omitempty"`
				Banana int    `jsonld:"banana"`
				Cherry string `jsonld:"cherry,omitempty"`
			}{
				Apple:"",
				Banana:2,
				Cherry:"three",
			},
			Expected: []byte(`{"banana":2,"cherry":"three"}`),
		},
		{
			Value: struct{
				Apple  string `jsonld:"apple,omitempty"`
				Banana int    `jsonld:"banana"`
				Cherry string `jsonld:"cherry,omitempty"`
			}{
				Apple:"1",
				Banana:2,
				Cherry:"",
			},
			Expected: []byte(`{"apple":"1","banana":2}`),
		},
		{
			Value: struct{
				Apple  string `jsonld:"apple,omitempty"`
				Banana int    `jsonld:"banana"`
				Cherry string `jsonld:"cherry,omitempty"`
			}{
				Apple:"",
				Banana:2,
				Cherry:"",
			},
			Expected: []byte(`{"banana":2}`),
		},



		{
			Value: struct{
				Apple  string `jsonld:"apple,omitempty"`
				Banana int    `jsonld:"banana"`
				Cherry string `jsonld:"cherry,omitempty"`
				date   int 
			}{
				Apple:"",
				Banana:2,
				Cherry:"",
				date: 4,
			},
			Expected: []byte(`{"banana":2}`),
		},
	}

	for testNumber, test := range tests {
		actual, err := marshalOneStruct(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual jsonld-marshaled struct is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d) %q", len(expected), expected)
			t.Logf("ACTUAL:   (%d) %q", len(actual), actual)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}
	}
}
