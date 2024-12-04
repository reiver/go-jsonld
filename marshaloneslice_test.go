package jsonld

import (
	"testing"

	"bytes"
)

func TestMarshalOneSlice(t *testing.T) {

	tests := []struct{
		Value any
		Expected []byte
	}{
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
	}

	for testNumber, test := range tests {
		actual, err := marshalOneSlice(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual jsonld-marshaled slice/array is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d) %q", len(expected), expected)
			t.Logf("ACTUAL:   (%d) %q", len(actual), actual)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}
	}
}
