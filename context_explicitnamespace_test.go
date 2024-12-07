package jsonld

import (
	"testing"
)

func TestContext_explicitNameSpace(t *testing.T) {

	tests := []struct{
		Context Context
		ExpectedPrefix string
		ExpectedNameSpace string
		ExpectedFound bool
	}{
		{
			Context: Context{},
			ExpectedPrefix: "",
			ExpectedNameSpace: "",
			ExpectedFound: false,
		},
		{
			Context: Context{
				NameSpace: "http://example.com/ns#",
			},
			ExpectedPrefix: "",
			ExpectedNameSpace: "",
			ExpectedFound: false,
		},


		{
			Context: Context{
				Prefix: "nom",
				NameSpace: "http://example.com/ns#",
			},
			ExpectedPrefix: "nom",
			ExpectedNameSpace: "http://example.com/ns#",
			ExpectedFound: true,
		},
		{
			Context: Context{
				Prefix: "label",
				NameSpace: "tag:joeblow@example.com,2024-12-06:ns",
			},
			ExpectedPrefix: "label",
			ExpectedNameSpace: "tag:joeblow@example.com,2024-12-06:ns",
			ExpectedFound: true,
		},
	}

	for testNumber, test := range tests {

		actualPrefix, actualNameSpace, actualFound := test.Context.explicitNameSpace()

		{
			expected := test.ExpectedFound
			actual   :=        actualFound

			if expected != actual {
				t.Errorf("For test #%d, the actual 'found' is not what wad expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("CONTEXT: %#v", test.Context)
				continue
			}
		}

		{
			expected := test.ExpectedNameSpace
			actual   :=        actualNameSpace

			if expected != actual {
				t.Errorf("For test #%d, the actual 'name-space' is not what wad expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("CONTEXT: %#v", test.Context)
				continue
			}
		}

		{
			expected := test.ExpectedPrefix
			actual   :=        actualPrefix

			if expected != actual {
				t.Errorf("For test #%d, the actual 'prefix' is not what wad expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("CONTEXT: %#v", test.Context)
				continue
			}
		}
	}
}
