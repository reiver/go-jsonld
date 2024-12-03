package jsonld_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
)

func TestMarshalContexts(t *testing.T) {

	tests := []struct{
		Contexts []jsonld.Context
		Expected []byte
	}{
		{
			Expected: []byte(`{}`),
		},



		{
			Contexts: []jsonld.Context{
				jsonld.Context{
					Names: []string{
						"apple",
						"banana",
						"cherry",
					},
				},
			},
			Expected: []byte(`{}`),
		},
		{
			Contexts: []jsonld.Context{
				jsonld.Context{
					Prefix: "ex",
					Names: []string{
						"apple",
						"banana",
						"cherry",
					},
				},
			},
			Expected: []byte(`{}`),
		},
		{
			Contexts: []jsonld.Context{
				jsonld.Context{
					NameSpace: "http://example.com/ns/",
					Prefix: "ex",
					Names: []string{
						"apple",
						"banana",
						"cherry",
					},
				},
			},
			Expected: []byte(
				`{`+
					`"ex":"http://example.com/ns/"`+
					`,`+
					`"apple":"ex:apple"`+
					`,`+
					`"banana":"ex:banana"`+
					`,`+
					`"cherry":"ex:cherry"`+
				`}`,
			),
		},
		{
			Contexts: []jsonld.Context{
				jsonld.Context{
					NameSpace: "http://example.com/ns/",
					Names: []string{
						"apple",
						"banana",
						"cherry",
					},
				},
			},
			Expected: []byte(
				`[`+
					`"http://example.com/ns/"`+
				`]`,
			),
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonld.MarshalContexts(test.Contexts...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual jsonld-marshaled value is not what was expected." , testNumber)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				continue
			}
		}
	}
}
