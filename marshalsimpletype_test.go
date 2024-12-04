package jsonld

import (
	"testing"

	"bytes"
)

func TestMarshalSimpleType(t *testing.T) {
	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Expected: []byte(`null`),
		},
		{
			Value: nil,
			Expected: []byte(`null`),
		},



		{
			Value: false,
			Expected: []byte(`false`),
		},
		{
			Value: true,
			Expected: []byte(`true`),
		},



		{
			Value: int(0),
			Expected: []byte(`0`),
		},
		{
			Value: int8(0),
			Expected: []byte(`0`),
		},
		{
			Value: int16(0),
			Expected: []byte(`0`),
		},
		{
			Value: int32(0),
			Expected: []byte(`0`),
		},
		{
			Value: int64(0),
			Expected: []byte(`0`),
		},



		{
			Value: "Hello world!",
			Expected: []byte(`"Hello world!"`),
		},



		{
			Value: uint(0),
			Expected: []byte(`0`),
		},
		{
			Value: uint8(0),
			Expected: []byte(`0`),
		},
		{
			Value: uint16(0),
			Expected: []byte(`0`),
		},
		{
			Value: uint32(0),
			Expected: []byte(`0`),
		},
		{
			Value: uint64(0),
			Expected: []byte(`0`),
		},









/*
		{
			Value: [4]string{"once","twice","thrice","fource"},
			Expected: false,
		},
		{
			Value: []string{"once","twice","thrice","fource"},
			Expected: false,
		},



		{
			Value: map[string]string{"once":"1","twice":"2","thrice":"3","fource":"4"},
			Expected: false,
		},



		{
			Value: struct{
				Once string
				Twice int
				Thrice string
				Fource int
			}{
				Once:"1",
				Twice:2,
				Thrice:"3",
				Fource:4,
			},
			Expected: false,
		},
*/
	}

	for testNumber, test := range tests {

		actual, err := marshalSimpleType(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual jsonld-marshaled result of a simple-type is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}
	}
}
