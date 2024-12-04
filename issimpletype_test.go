package jsonld

import (
	"testing"
)

func TestIsSimpleType(t *testing.T) {
	tests := []struct{
		Value any
		Expected bool
	}{
		{
			Expected: true,
		},
		{
			Value: nil,
			Expected: true,
		},



		{
			Value: false,
			Expected: true,
		},
		{
			Value: true,
			Expected: true,
		},



		{
			Value: int(0),
			Expected: true,
		},
		{
			Value: int8(0),
			Expected: true,
		},
		{
			Value: int16(0),
			Expected: true,
		},
		{
			Value: int32(0),
			Expected: true,
		},
		{
			Value: int64(0),
			Expected: true,
		},



		{
			Value: "Hello world",
			Expected: true,
		},



		{
			Value: uint(0),
			Expected: true,
		},
		{
			Value: uint8(0),
			Expected: true,
		},
		{
			Value: uint16(0),
			Expected: true,
		},
		{
			Value: uint32(0),
			Expected: true,
		},
		{
			Value: uint64(0),
			Expected: true,
		},









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
	}

	for testNumber, test := range tests {

		actual := isSimpleType(test.Value)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual is-simple-type result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}
	}
}
