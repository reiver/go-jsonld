package jsonld

import (
	"testing"

	"bytes"
)

func TestMarshalOneMap(t *testing.T) {

	tests := []struct{
		Value any
		Expected []byte
	}{
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
	}

	for testNumber, test := range tests {
		actual, err := marshalOneMap(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual jsonld-marshaled map is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d) %q", len(expected), expected)
			t.Logf("ACTUAL:   (%d) %q", len(actual), actual)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}
	}
}
