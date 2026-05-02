package jsonld_test

import (
	"testing"

	"encoding/json"

	"github.com/reiver/go-jsonld"
)

func TestDirection_UnmarshalJSON(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected jsonld.Direction
	}{
		{
			Data: []byte(`null`),
			Expected: jsonld.DirectionNull(),
		},
		{
			Data: []byte(`"ltr"`),
			Expected: jsonld.DirectionLTR(),
		},
		{
			Data: []byte(`"rtl"`),
			Expected: jsonld.DirectionRTL(),
		},
	}

	for testNumber, test := range tests {

		var actual jsonld.Direction

		err := json.Unmarshal(test.Data, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DATA:\n%s", test.Data)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual unmarshaled-json direction is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%#v", expected)
			t.Logf("ACTUAL:\n%#v", actual)
			t.Logf("DATA:\n%s", test.Data)
			continue
		}
	}
}

func TestDirection_UnmarshalJSON_fail(t *testing.T) {

	tests := []struct{
		Data []byte
	}{
		{
			Data: nil,
		},



		{
			Data: []byte(`""`),
		},
		{
			Data: []byte(`"LTR"`),
		},
		{
			Data: []byte(`"RTL"`),
		},
		{
			Data: []byte(`"Ltr"`),
		},
		{
			Data: []byte(`"Rtl"`),
		},



		{
			Data: []byte(`"something"`),
		},
		{
			Data: []byte(`123`),
		},
		{
			Data: []byte(`true`),
		},
		{
			Data: []byte(`{}`),
		},
		{
			Data: []byte(`[]`),
		},
	}

	for testNumber, test := range tests {

		var actual jsonld.Direction

		err := json.Unmarshal(test.Data, &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("DATA:\n%s", test.Data)
			t.Logf("ACTUAL:\n%#v", actual)
			continue
		}
	}
}
