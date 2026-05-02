package jsonld_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
)

func TestDirection_MarshalJSON(t *testing.T) {

	tests := []struct{
		Direction jsonld.Direction
		Expected []byte
	}{
		{
			Direction: jsonld.DirectionNull(),
			Expected: []byte(`null`),
		},
		{
			Direction: jsonld.DirectionLTR(),
			Expected: []byte(`"ltr"`),
		},
		{
			Direction: jsonld.DirectionRTL(),
			Expected: []byte(`"rtl"`),
		},
	}

	for testNumber, test := range tests {

		actual, err := test.Direction.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual json-marshaled direction is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			continue
		}
	}
}

func TestDirection_MarshalJSON_fail(t *testing.T) {

	var zero jsonld.Direction

	_, err := zero.MarshalJSON()
	if nil == err {
		t.Errorf("Expected an error for zero-value Direction but did not get one.")
		return
	}
}
