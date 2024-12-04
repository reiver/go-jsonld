package jsonld

import (
	"testing"

	"reflect"
)

func TestParseStructField(t *testing.T) {

	tests := []struct{
		Value any
		ExpectedName string
		ExpectedOmitEmpty bool
	}{
		{
			Value: struct{
				First bool
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First bool `jsonld:"first"`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First bool `jsonld:",omitempty"`
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: true,
		},
		{
			Value: struct{
				First bool `jsonld:"first,omitempty"`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: true,
		},



		{
			Value: struct{
				First string
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:"first"`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:",omitempty"`
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: true,
		},
		{
			Value: struct{
				First string `jsonld:"first,omitempty"`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: true,
		},



		{
			Value: struct{
				First string
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:" first"`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:",omitempty"`
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: true,
		},
		{
			Value: struct{
				First string `jsonld:" first,omitempty"`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: true,
		},



		{
			Value: struct{
				First string
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:"first "`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:",omitempty"`
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: true,
		},
		{
			Value: struct{
				First string `jsonld:"first ,omitempty"`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: true,
		},



		{
			Value: struct{
				First string
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:"  first "`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:",omitempty"`
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: true,
		},
		{
			Value: struct{
				First string `jsonld:"  first ,omitempty"`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: true,
		},



		{
			Value: struct{
				First string
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:"  first "`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:", omitempty"`
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: true,
		},
		{
			Value: struct{
				First string `jsonld:"  first , omitempty"`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: true,
		},



		{
			Value: struct{
				First string
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:"  first "`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:",omitempty "`
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: true,
		},
		{
			Value: struct{
				First string `jsonld:"  first ,omitempty "`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: true,
		},



		{
			Value: struct{
				First string
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:"  first "`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: false,
		},
		{
			Value: struct{
				First string `jsonld:", omitempty  "`
			}{},
			ExpectedName: "First",
			ExpectedOmitEmpty: true,
		},
		{
			Value: struct{
				First string `jsonld:"  first , omitempty  "`
			}{},
			ExpectedName: "first",
			ExpectedOmitEmpty: true,
		},
	}

	for testNumber, test := range tests {

		var reflectedType reflect.Type = reflect.TypeOf(test.Value)

		{
			var kind reflect.Kind = reflectedType.Kind()

			if reflect.Struct != kind {
				t.Errorf("For test #%d, expected test-value to be a 'struct' but actually wasn't.", testNumber)
				t.Logf("KIND: %s", kind)
				t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
				continue
			}
		}

		var structField reflect.StructField = reflectedType.Field(0)

		actualName, actualOmitEmpty := parseStructField(structField)

		{
			expected := test.ExpectedName
			actual   :=        actualName

			if expected != actual {
				t.Errorf("For test #%d, the actual 'name' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedOmitEmpty
			actual   :=        actualOmitEmpty

			if expected != actual {
				t.Errorf("For test #%d, the actual 'name' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
				continue
			}
		}
	}
}
