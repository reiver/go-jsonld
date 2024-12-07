package jsonld

import (
	"testing"

	"reflect"
)

func TestContext_forNames(t *testing.T) {

	tests := []struct{
		Context Context
		Expected []struct{
			Prefix string
			Name string
		}
	}{
		{
			Context: Context{},
		},
		{
			Context: Context{
				NameSpace: "http://example.com/ns#",
			},
		},
		{
			Context: Context{
				Prefix: "nom",
				NameSpace: "http://example.com/ns#",
			},
		},
		{
			Context: Context{
				NameSpace: "http://example.com/ns#",
				Names: []string{"once","twice","thrice","fource"},
			},
		},
		{
			Context: Context{
				Names: []string{"once","twice","thrice","fource"},
			},
		},


		{
			Context: Context{
				Prefix: "nom",
				NameSpace: "http://example.com/ns#",
				Names: []string{"once","twice","thrice","fource"},
			},
			Expected: []struct{
				Prefix string
				Name string
			}{
				{
					Prefix:"nom",
					Name: "once",
				},
				{
					Prefix:"nom",
					Name: "twice",
				},
				{
					Prefix:"nom",
					Name: "thrice",
				},
				{
					Prefix:"nom",
					Name: "fource",
				},
			},
		},
		{
			Context: Context{
				Prefix: "nom",
				Names: []string{"once","twice","thrice","fource"},
			},
			Expected: []struct{
				Prefix string
				Name string
			}{
				{
					Prefix:"nom",
					Name: "once",
				},
				{
					Prefix:"nom",
					Name: "twice",
				},
				{
					Prefix:"nom",
					Name: "thrice",
				},
				{
					Prefix:"nom",
					Name: "fource",
				},
			},
		},
	}

	for testNumber, test := range tests {

		var actual []struct{
			Prefix string
			Name string
		}

		test.Context.forNames(func(prefix string, name string){
			actual = append(actual, struct{
				Prefix string
				Name string
			}{
				Prefix: prefix,
				Name: name,
			})
		})

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual list of prefix-name is not what wad expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				for index, datum := range expected {
					t.Logf("\t[%d] ⁃ prefix = %q , name = %q", index, datum.Prefix, datum.Name)
				}
				t.Logf("ACTUAL:   %#v", actual)
				for index, datum := range actual {
					t.Logf("\t[%d] ⁃ prefix = %q , name = %q", index, datum.Prefix, datum.Name)
				}
				t.Logf("CONTEXT: %#v", test.Context)
				continue
			}
		}
	}
}
