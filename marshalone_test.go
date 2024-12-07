package jsonld

import (
	"testing"

	"bytes"
)

func TestMarshalOne(t *testing.T) {

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
				Apple string `json:"apple"`
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
				Apple string `json:"apple"`
				Banana int
			}{
				Apple:"1",
				Banana:2,
			},
			Expected: []byte(`{"apple":"1","Banana":2}`),
		},
		{
			Value: struct{
				Apple string `json:"apple"`
				Banana int   `json:"banana"`
			}{
				Apple:"1",
				Banana:2,
			},
			Expected: []byte(`{"apple":"1","banana":2}`),
		},



		{
			Value: struct{
				Apple  string `json:"apple"`
				Banana int    `json:"banana"`
				Cherry string `json:"cherry"`
			}{
				Apple:"1",
				Banana:2,
				Cherry:"three",
			},
			Expected: []byte(`{"apple":"1","banana":2,"cherry":"three"}`),
		},



		{
			Value: struct{
				Apple  string `json:"apple,omitempty"`
				Banana int    `json:"banana"`
				Cherry string `json:"cherry,omitempty"`
			}{
				Apple:"1",
				Banana:2,
				Cherry:"three",
			},
			Expected: []byte(`{"apple":"1","banana":2,"cherry":"three"}`),
		},
		{
			Value: struct{
				Apple  string `json:"apple,omitempty"`
				Banana int    `json:"banana"`
				Cherry string `json:"cherry,omitempty"`
			}{
				Apple:"",
				Banana:2,
				Cherry:"three",
			},
			Expected: []byte(`{"banana":2,"cherry":"three"}`),
		},
		{
			Value: struct{
				Apple  string `json:"apple,omitempty"`
				Banana int    `json:"banana"`
				Cherry string `json:"cherry,omitempty"`
			}{
				Apple:"1",
				Banana:2,
				Cherry:"",
			},
			Expected: []byte(`{"apple":"1","banana":2}`),
		},
		{
			Value: struct{
				Apple  string `json:"apple,omitempty"`
				Banana int    `json:"banana"`
				Cherry string `json:"cherry,omitempty"`
			}{
				Apple:"",
				Banana:2,
				Cherry:"",
			},
			Expected: []byte(`{"banana":2}`),
		},









		{
			Value: map[string]string(nil),
			Expected: []byte(`{}`),
		},
		{
			Value: map[string]string{},
			Expected: []byte(`{}`),
		},



		{
			Value: map[string]string{
				"apple":"1",
			},
			Expected: []byte(`{"apple":"1"}`),
		},
		{
			Value: map[string]string{
				"apple":"1",
				"banana":"2",
			},
			Expected: []byte(`{"apple":"1","banana":"2"}`),
		},
		{
			Value: map[string]string{
				"apple":"1",
				"banana":"2",
				"cherry":"3",
			},
			Expected: []byte(`{"apple":"1","banana":"2","cherry":"3"}`),
		},









		{
			Value: []string(nil),
			Expected: []byte(`[]`),
		},
		{
			Value: []string{},
			Expected: []byte(`[]`),
		},
		{
			Value: [0]string{},
			Expected: []byte(`[]`),
		},



		{
			Value: []string{
				"apple",
			},
			Expected: []byte(`["apple"]`),
		},
		{
			Value: []string{
				"apple",
				"banana",
			},
			Expected: []byte(`["apple","banana"]`),
		},
		{
			Value: []string{
				"apple",
				"banana",
				"cherry",
			},
			Expected: []byte(`["apple","banana","cherry"]`),
		},



		{
			Value: [1]string{
				"apple",
			},
			Expected: []byte(`["apple"]`),
		},
		{
			Value: [2]string{
				"apple",
				"banana",
			},
			Expected: []byte(`["apple","banana"]`),
		},
		{
			Value: [3]string{
				"apple",
				"banana",
				"cherry",
			},
			Expected: []byte(`["apple","banana","cherry"]`),
		},



		{
			Value: []any{
				"ONE",
				2,
				"3",
			},
			Expected: []byte(`["ONE",2,"3"]`),
		},



		{
			Value: [3]any{
				"ONE",
				2,
				"3",
			},
			Expected: []byte(`["ONE",2,"3"]`),
		},








		{
			Value: struct{
				Apple  string `json:"apple,omitempty"`
				Banana int    `json:"banana"`
				Cherry string `json:"cherry,omitempty"`
				date   int

				List []string `json:"list,omitempty"`
				Thing  string `json:"thing"`
			}{
				Apple:"",
				Banana:2,
				Cherry:"",
				date: 4,
				List: []string{"#one","#two","#three","#four"},
				Thing: "wow",
			},
			Expected: []byte(`{"banana":2,"list":["#one","#two","#three","#four"],"thing":"wow"}`),
		},
	}

	for testNumber, test := range tests {
		actual, err := marshalOne(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual jsonld-marshaled (one) is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d) %q", len(expected), expected)
			t.Logf("ACTUAL:   (%d) %q", len(actual), actual)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}
	}
}
