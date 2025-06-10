package strings_test

import (
	"testing"

	"encoding/json"
	"reflect"

	"github.com/reiver/go-jsonld/strings"
)

func TestStrings_UnmarshalJSON(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected strings.Strings
	}{
		{
			Data: []byte("null"),
		},
		{
			Data: []byte("null"),
			Expected: strings.Nothing(),
		},



		{
			Data: []byte("[]"),
			Expected: strings.Somethings(),
		},



		{
			Data:               []byte(`"once"`),
			Expected: strings.Something("once"),
		},
		{
			Data:               []byte(`"twice"`),
			Expected: strings.Something("twice"),
		},
		{
			Data:               []byte(`"thrice"`),
			Expected: strings.Something("thrice"),
		},
		{
			Data:               []byte(`"fource"`),
			Expected: strings.Something("fource"),
		},



		{
			Data:               []byte(`["once","twice"]`),
			Expected: strings.Somethings("once","twice"),
		},
		{
			Data:               []byte(`["once","twice","thrice"]`),
			Expected: strings.Somethings("once","twice","thrice"),
		},
		{
			Data:               []byte(`["once","twice","thrice","fource"]`),
			Expected: strings.Somethings("once","twice","thrice","fource"),
		},
	}

	for testNumber, test := range tests {

		var actual strings.Strings

		err := json.Unmarshal(test.Data, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("DATA:\n%s", test.Data)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual marshaled-json is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%#v", expected)
			t.Logf("ACTUAL:\n%#v", actual)
			t.Logf("DATA:\n%s", test.Data)
			continue
		}
	}
}
