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
				NameValues: []NameValue{
					NameValue{
						Name: "Apple",
						Value: "ONE",
					},
					NameValue{
						Name: "Banana",
						Value: "2",
					},
					NameValue{
						Name: "Cherry",
						Value: "پ",
					},
				},
			},
		},



		{
			Struct: struct{
				Apple string  `jsonld:"apple"`
				Banana string
				Cherry string
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				NameValues: []NameValue{
					NameValue{
						Name: "apple",
						Value: "ONE",
					},
					NameValue{
						Name: "Banana",
						Value: "2",
					},
					NameValue{
						Name: "Cherry",
						Value: "پ",
					},
				},
			},
		},
		{
			Struct: struct{
				Apple string
				Banana string `jsonld:"banana"`
				Cherry string
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				NameValues: []NameValue{
					NameValue{
						Name: "Apple",
						Value: "ONE",
					},
					NameValue{
						Name: "banana",
						Value: "2",
					},
					NameValue{
						Name: "Cherry",
						Value: "پ",
					},
				},
			},
		},
		{
			Struct: struct{
				Apple string
				Banana string
				Cherry string `jsonld:"cherry"`
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				NameValues: []NameValue{
					NameValue{
						Name: "Apple",
						Value: "ONE",
					},
					NameValue{
						Name: "Banana",
						Value: "2",
					},
					NameValue{
						Name: "cherry",
						Value: "پ",
					},
				},
			},
		},



		{
			Struct: struct{
				Apple string
				Banana string `jsonld:"banana"`
				Cherry string `jsonld:"cherry"`
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				NameValues: []NameValue{
					NameValue{
						Name: "Apple",
						Value: "ONE",
					},
					NameValue{
						Name: "banana",
						Value: "2",
					},
					NameValue{
						Name: "cherry",
						Value: "پ",
					},
				},
			},
		},
		{
			Struct: struct{
				Apple string  `jsonld:"apple"`
				Banana string
				Cherry string `jsonld:"cherry"`
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				NameValues: []NameValue{
					NameValue{
						Name: "apple",
						Value: "ONE",
					},
					NameValue{
						Name: "Banana",
						Value: "2",
					},
					NameValue{
						Name: "cherry",
						Value: "پ",
					},
				},
			},
		},
		{
			Struct: struct{
				Apple string  `jsonld:"apple"`
				Banana string `jsonld:"banana"`
				Cherry string
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				NameValues: []NameValue{
					NameValue{
						Name: "apple",
						Value: "ONE",
					},
					NameValue{
						Name: "banana",
						Value: "2",
					},
					NameValue{
						Name: "Cherry",
						Value: "پ",
					},
				},
			},
		},



		{
			Struct: struct{
				Apple string  `jsonld:"apple"`
				Banana string `jsonld:"banana"`
				Cherry string `jsonld:"cherry"`
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				NameValues: []NameValue{
					NameValue{
						Name: "apple",
						Value: "ONE",
					},
					NameValue{
						Name: "banana",
						Value: "2",
					},
					NameValue{
						Name: "cherry",
						Value: "پ",
					},
				},
			},
		},



		{
			Struct: struct{
				NameSpace NameSpace `jsonld:"http://example.com/ns/"`
				Apple string  `jsonld:"apple"`
				Banana string `jsonld:"banana"`
				Cherry string `jsonld:"cherry"`
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				NameSpace: "http://example.com/ns/",
				NameValues: []NameValue{
					NameValue{
						Name: "apple",
						Value: "ONE",
					},
					NameValue{
						Name: "banana",
						Value: "2",
					},
					NameValue{
						Name: "cherry",
						Value: "پ",
					},
				},
			},
		},



		{
			Struct: struct{
				NameSpace NameSpace `jsonld:"http://example.com/ns/"`
				Prefix Prefix `jsonld:"ex"`
				Apple string  `jsonld:"apple"`
				Banana string `jsonld:"banana"`
				Cherry string `jsonld:"cherry"`
			}{
				Apple: "ONE",
				Banana: "2",
				Cherry: "پ",
			},
			Expected: Context{
				NameSpace: "http://example.com/ns/",
				Prefix: "ex",
				NameValues: []NameValue{
					NameValue{
						Name: "apple",
						Value: "ONE",
					},
					NameValue{
						Name: "banana",
						Value: "2",
					},
					NameValue{
						Name: "cherry",
						Value: "پ",
					},
				},
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
