package strings_test

import (
	"testing"

	"bytes"
	"encoding/json"

	"github.com/reiver/go-jsonld/strings"
)

func TestStrings_MarshalJSON(t *testing.T) {

	tests := []struct{
		Value strings.Strings
		Expected []byte
	}{
		{
			Expected: []byte("null"),
		},
		{
			Value: strings.Nothing(),
			Expected: []byte("null"),
		},



		{
			Value: strings.Somethings(),
			Expected: []byte("[]"),
		},



		{
			Value: strings.Something("once"),
			Expected:        []byte(`"once"`),
		},
		{
			Value: strings.Something("twice"),
			Expected:        []byte(`"twice"`),
		},
		{
			Value: strings.Something("thrice"),
			Expected:        []byte(`"thrice"`),
		},
		{
			Value: strings.Something("fource"),
			Expected:        []byte(`"fource"`),
		},



		{
			Value: strings.Somethings("once","twice"),
			Expected:        []byte(`["once","twice"]`),
		},
		{
			Value: strings.Somethings("once","twice","thrice"),
			Expected:        []byte(`["once","twice","thrice"]`),
		},
		{
			Value: strings.Somethings("once","twice","thrice","fource"),
			Expected:        []byte(`["once","twice","thrice","fource"]`),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("VALUE:\n%#v", test.Value)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual marshaled-json is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual),  actual)
			t.Logf("VALUE:\n%#v", test.Value)
			continue
		}
	}
}
