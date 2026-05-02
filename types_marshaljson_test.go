package jsonld_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
)

func TestTypes_MarshalJSON(t *testing.T) {

	tests := []struct{
		Types jsonld.Types
		Expected []byte
	}{
		{
			Types: jsonld.NoType(),
			Expected: []byte(`null`),
		},



		{
			Types: jsonld.SomeType("apple"),
			Expected: []byte(`"apple"`),
		},
		{
			Types: jsonld.SomeType("BANANA"),
			Expected: []byte(`"BANANA"`),
		},
		{
			Types: jsonld.SomeType("Cherry"),
			Expected: []byte(`"Cherry"`),
		},
		{
			Types: jsonld.SomeType("dAtE"),
			Expected: []byte(`"dAtE"`),
		},



		{
			Types: jsonld.SomeTypes("apple", "BANANA"),
			Expected: []byte(`["apple","BANANA"]`),
		},
		{
			Types: jsonld.SomeTypes("apple", "BANANA", "Cherry"),
			Expected: []byte(`["apple","BANANA","Cherry"]`),
		},
		{
			Types: jsonld.SomeTypes("apple", "BANANA", "Cherry", "dAtE"),
			Expected: []byte(`["apple","BANANA","Cherry","dAtE"]`),
		},
	}

	for testNumber, test := range tests {

		actual, err := test.Types.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual json-marshaled types is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			continue
		}
	}
}
