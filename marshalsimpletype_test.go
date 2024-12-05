package jsonld

import (
	"testing"

	"bytes"
)

type TestMarshalSimpleType_type1 struct {}
func (receiver TestMarshalSimpleType_type1) MarshalJSON() ([]byte, error) {
	return []byte(`"it worked!"`), nil
}

type TestMarshalSimpleType_type2 struct {}
func (receiver TestMarshalSimpleType_type2) MarshalJSON() ([]byte, error) {
	return []byte(`["ONCE","TWICE","THRICE","FOURCE"]`), nil
}

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









		{
			Value: TestMarshalSimpleType_type1{},
			Expected: []byte(`"it worked!"`),
		},
		{
			Value: TestMarshalSimpleType_type2{},
			Expected: []byte(`["ONCE","TWICE","THRICE","FOURCE"]`),
		},
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
