package jsonld_test

import (
	"testing"

	"encoding/json"
	"reflect"

	"github.com/reiver/go-jsonld"
)

func TestTypes_UnmarshalJSON(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected jsonld.Types
	}{
		{
			Data: []byte(`null`),
			Expected: jsonld.NoType(),
		},
		{
			Data: []byte(`[]`),
			Expected: jsonld.NoType(),
		},



		{
			Data: []byte(`"apple"`),
			Expected: jsonld.SomeType("apple"),
		},
		{
			Data: []byte(`"BANANA"`),
			Expected: jsonld.SomeType("BANANA"),
		},
		{
			Data: []byte(`"Cherry"`),
			Expected: jsonld.SomeType("Cherry"),
		},
		{
			Data: []byte(`"dAtE"`),
			Expected: jsonld.SomeType("dAtE"),
		},



		{
			Data: []byte(`["apple"]`),
			Expected: jsonld.SomeType("apple"),
		},
		{
			Data: []byte(`["apple","BANANA"]`),
			Expected: jsonld.SomeTypes("apple", "BANANA"),
		},
		{
			Data: []byte(`["apple","BANANA","Cherry"]`),
			Expected: jsonld.SomeTypes("apple", "BANANA", "Cherry"),
		},
		{
			Data: []byte(`["apple","BANANA","Cherry","dAtE"]`),
			Expected: jsonld.SomeTypes("apple", "BANANA", "Cherry", "dAtE"),
		},



		{
			Data: []byte(`{"@id":"apple"}`),
			Expected: jsonld.SomeType("apple"),
		},
		{
			Data: []byte(`{"@id":"BANANA"}`),
			Expected: jsonld.SomeType("BANANA"),
		},
		{
			Data: []byte(`{"@id":"Cherry"}`),
			Expected: jsonld.SomeType("Cherry"),
		},
		{
			Data: []byte(`{"@id":"dAtE"}`),
			Expected: jsonld.SomeType("dAtE"),
		},



		{
			Data: []byte(`{"id":"apple"}`),
			Expected: jsonld.SomeType("apple"),
		},
		{
			Data: []byte(`{"id":"BANANA"}`),
			Expected: jsonld.SomeType("BANANA"),
		},
		{
			Data: []byte(`{"id":"Cherry"}`),
			Expected: jsonld.SomeType("Cherry"),
		},
		{
			Data: []byte(`{"id":"dAtE"}`),
			Expected: jsonld.SomeType("dAtE"),
		},



		{
			Data: []byte(`{"@id":"apple","id":"BANANA"}`),
			Expected: jsonld.SomeType("apple"),
		},
		{
			Data: []byte(`{"id":"apple","@id":"BANANA"}`),
			Expected: jsonld.SomeType("BANANA"),
		},
		{
			Data: []byte(`{"@id":"Cherry","id":"dAtE"}`),
			Expected: jsonld.SomeType("Cherry"),
		},
		{
			Data: []byte(`{"id":"Cherry","@id":"dAtE"}`),
			Expected: jsonld.SomeType("dAtE"),
		},
		{
			Data: []byte(`[{"@id":"apple","id":"BANANA"}]`),
			Expected: jsonld.SomeType("apple"),
		},
		{
			Data: []byte(`[{"id":"apple","@id":"BANANA"}]`),
			Expected: jsonld.SomeType("BANANA"),
		},
		{
			Data: []byte(`[{"@id":"apple","id":"BANANA"},{"@id":"Cherry","id":"dAtE"}]`),
			Expected: jsonld.SomeTypes("apple", "Cherry"),
		},
		{
			Data: []byte(`[{"id":"apple","@id":"BANANA"},{"id":"Cherry","@id":"dAtE"}]`),
			Expected: jsonld.SomeTypes("BANANA", "dAtE"),
		},



		{
			Data: []byte(`[{"@id":"apple"}]`),
			Expected: jsonld.SomeType("apple"),
		},
		{
			Data: []byte(`[{"@id":"apple"},{"@id":"BANANA"}]`),
			Expected: jsonld.SomeTypes("apple", "BANANA"),
		},
		{
			Data: []byte(`[{"@id":"apple"},{"@id":"BANANA"},{"@id":"Cherry"}]`),
			Expected: jsonld.SomeTypes("apple", "BANANA", "Cherry"),
		},
		{
			Data: []byte(`[{"@id":"apple"},{"@id":"BANANA"},{"@id":"Cherry"},{"@id":"dAtE"}]`),
			Expected: jsonld.SomeTypes("apple", "BANANA", "Cherry", "dAtE"),
		},



		{
			Data: []byte(`[{"id":"apple"}]`),
			Expected: jsonld.SomeType("apple"),
		},
		{
			Data: []byte(`[{"id":"apple"},{"id":"BANANA"}]`),
			Expected: jsonld.SomeTypes("apple", "BANANA"),
		},
		{
			Data: []byte(`[{"id":"apple"},{"id":"BANANA"},{"id":"Cherry"}]`),
			Expected: jsonld.SomeTypes("apple", "BANANA", "Cherry"),
		},
		{
			Data: []byte(`[{"id":"apple"},{"id":"BANANA"},{"id":"Cherry"},{"id":"dAtE"}]`),
			Expected: jsonld.SomeTypes("apple", "BANANA", "Cherry", "dAtE"),
		},



		{
			Data: []byte(`[{"@id":"apple"},"BANANA",{"id":"Cherry"},"dAtE"]`),
			Expected: jsonld.SomeTypes("apple", "BANANA", "Cherry", "dAtE"),
		},



	}

	for testNumber, test := range tests {

		var actual jsonld.Types

		err := json.Unmarshal(test.Data, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DATA:\n%s", test.Data)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual unmarshaled-json types is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%#v", expected)
			t.Logf("ACTUAL:\n%#v", actual)
			t.Logf("DATA:\n%s", test.Data)
			continue
		}
	}
}

func TestTypes_UnmarshalJSON_fail(t *testing.T) {

	tests := []struct{
		Data []byte
	}{
		{
			Data: nil,
		},
		{
			Data: []byte(``),
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
			Data: []byte(`[{}]`),
		},



		{
			Data: []byte(`{"ONCE":"apple","TWICE":"BANANA","THRICE":"Cherry","FOURCE":"dAtE"}`),
		},
		{
			Data: []byte(`[{"ONCE":"apple","TWICE":"BANANA","THRICE":"Cherry","FOURCE":"dAtE"}]`),
		},
		{
			Data: []byte(`[{"ONCE":"apple"},{"TWICE":"BANANA"},{"THRICE":"Cherry"},{"FOURCE":"dAtE"}]`),
		},



		{
			Data: []byte(`[{"@id":"apple"},"BANANA",{"id":"Cherry"},"dAtE", null]`),
		},
		{
			Data: []byte(`[{"@id":"apple"},"BANANA",{"id":"Cherry"},"dAtE", false]`),
		},
		{
			Data: []byte(`[{"@id":"apple"},"BANANA",{"id":"Cherry"},"dAtE", true]`),
		},
		{
			Data: []byte(`[{"@id":"apple"},"BANANA",{"id":"Cherry"},"dAtE", -1]`),
		},
		{
			Data: []byte(`[{"@id":"apple"},"BANANA",{"id":"Cherry"},"dAtE", 0]`),
		},
		{
			Data: []byte(`[{"@id":"apple"},"BANANA",{"id":"Cherry"},"dAtE", 1]`),
		},
		{
			Data: []byte(`[{"@id":"apple"},"BANANA",{"id":"Cherry"},"dAtE", {}]`),
		},
		{
			Data: []byte(`[{"@id":"apple"},"BANANA",{"id":"Cherry"},"dAtE", []]`),
		},



		{
			Data: []byte(`["apple", {}]`),
		},
		{
			Data: []byte(`[{}, "apple"]`),
		},
		{
			Data: []byte(`[{}, "apple", {}]`),
		},
		{
			Data: []byte(`[{}, "apple", {}, "BANANA", {}]`),
		},
		{
			Data: []byte(`[{"@id":"apple"}, {}, {"id":"BANANA"}, {}, "Cherry"]`),
		},
		{
			Data: []byte(`[{}, {"@id":"apple"}, {}, {"id":"BANANA"}, {}, "Cherry", {}, "dAtE", {}]`),
		},
	}

	for testNumber, test := range tests {

		var actual jsonld.Types

		err := json.Unmarshal(test.Data, &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("DATA:\n%s", test.Data)
			t.Logf("ACTUAL:\n%#v", actual)
			continue
		}
	}
}
