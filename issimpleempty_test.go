package jsonld

import (
	"testing"
)

func TestIsSimpleEmpty(t *testing.T) {

	tests := []struct{
		Value any
		Expected bool
	}{
		// bool
		{
			Value: true,
			Expected: false,
		},
		{
			Value: false,
			Expected: true,
		},



		// int
		{
			Value: int(-123),
			Expected: false,
		},
		{
			Value: int(-1),
			Expected: false,
		},
		{
			Value: int(0),
			Expected: true,
		},
		{
			Value: int(1),
			Expected: false,
		},
		{
			Value: int(123),
			Expected: false,
		},



		// int8
		{
			Value: int8(-123),
			Expected: false,
		},
		{
			Value: int8(-1),
			Expected: false,
		},
		{
			Value: int8(0),
			Expected: true,
		},
		{
			Value: int8(1),
			Expected: false,
		},
		{
			Value: int8(123),
			Expected: false,
		},



		// int16
		{
			Value: int16(-123),
			Expected: false,
		},
		{
			Value: int16(-1),
			Expected: false,
		},
		{
			Value: int16(0),
			Expected: true,
		},
		{
			Value: int16(1),
			Expected: false,
		},
		{
			Value: int16(123),
			Expected: false,
		},



		// int32
		{
			Value: int32(-123),
			Expected: false,
		},
		{
			Value: int32(-1),
			Expected: false,
		},
		{
			Value: int32(0),
			Expected: true,
		},
		{
			Value: int32(1),
			Expected: false,
		},
		{
			Value: int32(123),
			Expected: false,
		},



		// int64
		{
			Value: int64(-123),
			Expected: false,
		},
		{
			Value: int64(-1),
			Expected: false,
		},
		{
			Value: int64(0),
			Expected: true,
		},
		{
			Value: int64(1),
			Expected: false,
		},
		{
			Value: int64(123),
			Expected: false,
		},



		// string
		{
			Value: "",
			Expected: true,
		},
		{
			Value: "apple",
			Expected: false,
		},
		{
			Value: "banana",
			Expected: false,
		},
		{
			Value: "cherry",
			Expected: false,
		},



		// uint
		{
			Value: int(0),
			Expected: true,
		},
		{
			Value: int(1),
			Expected: false,
		},
		{
			Value: int(123),
			Expected: false,
		},



		// uint8
		{
			Value: int8(0),
			Expected: true,
		},
		{
			Value: int8(1),
			Expected: false,
		},
		{
			Value: int8(123),
			Expected: false,
		},



		// uint16
		{
			Value: int16(0),
			Expected: true,
		},
		{
			Value: int16(1),
			Expected: false,
		},
		{
			Value: int16(123),
			Expected: false,
		},



		// uint32
		{
			Value: int32(0),
			Expected: true,
		},
		{
			Value: int32(1),
			Expected: false,
		},
		{
			Value: int32(123),
			Expected: false,
		},



		// uint64
		{
			Value: int64(0),
			Expected: true,
		},
		{
			Value: int64(1),
			Expected: false,
		},
		{
			Value: int64(123),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := isSimpleEmpty(test.Value)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result is not what was expected." , testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}
	}
}
