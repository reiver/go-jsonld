package jsonld_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-jsonld"
)

func TestContextOf(t *testing.T) {

	tests := []struct{
		Value any
		Expected jsonld.Context
	}{
		{
			Value: struct{}{},
		},



		{
			Value: struct{
				NameSpace jsonld.NameSpace
			}{},
		},



		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"http://example.com/ns"`
			}{},
			Expected: jsonld.Context{
				NameSpace: "http://example.com/ns",
			},
		},
		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"uuid:02def968-d55f-426f-af0b-8a5fa7439bdd"`
			}{},
			Expected: jsonld.Context{
				NameSpace: "uuid:02def968-d55f-426f-af0b-8a5fa7439bdd",
			},
		},



		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"http://example.com/ns"`
				Prefix    jsonld.Prefix    `jsonld:"ex"`
			}{},
			Expected: jsonld.Context{
				NameSpace: "http://example.com/ns",
				Prefix: "ex",
			},
		},
		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"uuid:02def968-d55f-426f-af0b-8a5fa7439bdd"`
				Prefix    jsonld.Prefix    `jsonld:"xap"`
			}{},
			Expected: jsonld.Context{
				NameSpace: "uuid:02def968-d55f-426f-af0b-8a5fa7439bdd",
				Prefix: "xap",
			},
		},



		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"http://example.com/ns"`

				Apple bool
				Banana int
				Cherry string
			}{},
			Expected: jsonld.Context{
				NameSpace: "http://example.com/ns",
				Names: []string{"Apple","Banana","Cherry"},
			},
		},
		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"uuid:02def968-d55f-426f-af0b-8a5fa7439bdd"`

				Once uint8
				Twice uint16
				Thrice uint32
				Fource uint64
			}{},
			Expected: jsonld.Context{
				NameSpace: "uuid:02def968-d55f-426f-af0b-8a5fa7439bdd",
				Names: []string{"Fource","Once","Thrice","Twice"},
			},
		},



		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"http://example.com/ns"`
				Prefix    jsonld.Prefix    `jsonld:"ex"`

				Apple bool
				Banana int
				Cherry string
			}{},
			Expected: jsonld.Context{
				NameSpace: "http://example.com/ns",
				Prefix: "ex",
				Names: []string{"Apple","Banana","Cherry"},
			},
		},
		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"uuid:02def968-d55f-426f-af0b-8a5fa7439bdd"`
				Prefix    jsonld.Prefix    `jsonld:"xap"`

				Once uint8
				Twice uint16
				Thrice uint32
				Fource uint64
			}{},
			Expected: jsonld.Context{
				NameSpace: "uuid:02def968-d55f-426f-af0b-8a5fa7439bdd",
				Prefix: "xap",
				Names: []string{"Fource","Once","Thrice","Twice"},
			},
		},



		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"http://example.com/ns"`

				Apple  bool   `json:"apple"`
				Banana int
				Cherry string `json:"cherry"`
			}{},
			Expected: jsonld.Context{
				NameSpace: "http://example.com/ns",
				Names: []string{"Banana","apple","cherry"},
			},
		},
		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"uuid:02def968-d55f-426f-af0b-8a5fa7439bdd"`

				Once   uint8
				Twice  uint16 `json:"twice"`
				Thrice uint32
				Fource uint64 `json:"fource"`
			}{},
			Expected: jsonld.Context{
				NameSpace: "uuid:02def968-d55f-426f-af0b-8a5fa7439bdd",
				Names: []string{"Once","Thrice","fource","twice"},
			},
		},



		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"http://example.com/ns"`
				Prefix    jsonld.Prefix    `jsonld:"ex"`

				Apple  bool   `json:"apple"`
				Banana int
				Cherry string `json:"cherry"`
			}{},
			Expected: jsonld.Context{
				NameSpace: "http://example.com/ns",
				Prefix: "ex",
				Names: []string{"Banana","apple","cherry"},
			},
		},
		{
			Value: struct{
				NameSpace jsonld.NameSpace `jsonld:"uuid:02def968-d55f-426f-af0b-8a5fa7439bdd"`
				Prefix    jsonld.Prefix    `jsonld:"xap"`

				Once   uint8
				Twice  uint16 `json:"twice"`
				Thrice uint32
				Fource uint64 `json:"fource"`
			}{},
			Expected: jsonld.Context{
				NameSpace: "uuid:02def968-d55f-426f-af0b-8a5fa7439bdd",
				Prefix: "xap",
				Names: []string{"Once","Thrice","fource","twice"},
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonld.ContextOf(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual 'context' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}
	}
}
