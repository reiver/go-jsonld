package jsonld

import (
	"testing"

	"reflect"
)

func TestContextOfStruct(t *testing.T) {

	tests := []struct{
		Struct any
		Expected Context
	}{
		{
			Struct: struct{}{},
		},



		{
			Struct: struct{
				Apple string
				Banana string
				Cherry string
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				Names: []string{
					"Apple",
					"Banana",
					"Cherry",
				},
			},
		},



		{
			Struct: struct{
				Apple string  `json:"apple"`
				Banana string
				Cherry string
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				Names: []string{
					"apple",
					"Banana",
					"Cherry",
				},
			},
		},
		{
			Struct: struct{
				Apple string
				Banana string `json:"banana"`
				Cherry string
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				Names: []string{
					"Apple",
					"banana",
					"Cherry",
				},
			},
		},
		{
			Struct: struct{
				Apple string
				Banana string
				Cherry string `json:"cherry"`
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				Names: []string{
					"Apple",
					"Banana",
					"cherry",
				},
			},
		},



		{
			Struct: struct{
				Apple string
				Banana string `json:"banana"`
				Cherry string `json:"cherry"`
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				Names: []string{
					"Apple",
					"banana",
					"cherry",
				},
			},
		},
		{
			Struct: struct{
				Apple string  `json:"apple"`
				Banana string
				Cherry string `json:"cherry"`
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				Names: []string{
					"apple",
					"Banana",
					"cherry",
				},
			},
		},
		{
			Struct: struct{
				Apple string  `json:"apple"`
				Banana string `json:"banana"`
				Cherry string
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				Names: []string{
					"apple",
					"banana",
					"Cherry",
				},
			},
		},



		{
			Struct: struct{
				Apple string  `json:"apple"`
				Banana string `json:"banana"`
				Cherry string `json:"cherry"`
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				Names: []string{
					"apple",
					"banana",
					"cherry",
				},
			},
		},



		{
			Struct: struct{
				NameSpace NameSpace `jsonld:"http://example.com/ns/"`
				Apple string  `json:"apple"`
				Banana string `json:"banana"`
				Cherry string `json:"cherry"`
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				NameSpace: "http://example.com/ns/",
				Names: []string{
					"apple",
					"banana",
					"cherry",
				},
			},
		},



		{
			Struct: struct{
				NameSpace NameSpace `jsonld:"http://example.com/ns/"`
				Prefix Prefix `jsonld:"ex"`
				Apple string  `json:"apple"`
				Banana string `json:"banana"`
				Cherry string `json:"cherry"`
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				NameSpace: "http://example.com/ns/",
				Prefix: "ex",
				Names: []string{
					"apple",
					"banana",
					"cherry",
				},
			},
		},


		{
			Struct: struct{
				NS NameSpace `jsonld:"http://example.com/ns#"`
				P  Prefix    `jsonld:"ex"`
				Apple string `json:"apple"`
				Banana int
				Cherry bool  `json:"cherry"`
			}{
				Apple: "ONE",
				Banana: 2,
				Cherry: true,
			},
			Expected: Context{
				NameSpace: "http://example.com/ns#",
				Prefix: "ex",
				Names: []string{"apple","Banana","cherry"},
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := contextOfStruct(test.Struct)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("STRUCT: %#v", test.Struct)
			continue
		}

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual context of-the-struct is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("STRUCT: %#v", test.Struct)
				continue
			}
		}
	}
}
