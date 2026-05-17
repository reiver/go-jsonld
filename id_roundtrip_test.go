package jsonld_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
)

func TestID_RoundTrip(t *testing.T) {

	tests := []struct{
		JSON []byte
	}{
		{
			JSON: []byte(`null`),
		},



		{
			JSON: []byte(`""`),
		},



		{
			JSON: []byte(`"apple"`),
		},
		{
			JSON: []byte(`"BANANA"`),
		},
		{
			JSON: []byte(`"Cherry"`),
		},
		{
			JSON: []byte(`"dAtE"`),
		},



		{
			JSON: []byte(`"https://example.com/something"`),
		},
		{
			JSON: []byte(`"http://example.net/apple/BANANA/Cherry/dAtE.php?once=1#aBc123"`),
		},



		{
			JSON: []byte(`"http://درود.com/"`),
		},
		{
			JSON: []byte(`"http://안녕.kr/세상"`),
		},
		{
			JSON: []byte(`"finger://joeblow@😈.example"`),
		},
		{
			JSON: []byte(`"punch://example.com/😈-😈.md"`),
		},



		// Reserved percent-encoded characters in path are preserved
		{
			JSON: []byte(`"http://example.com/foo%2Fbar"`),
		},
		{
			JSON: []byte(`"http://example.com/a%3Fb%3Dc"`),
		},
		{
			JSON: []byte(`"http://example.com/100%25done"`),
		},



		// Relative IRIs (no scheme)
		{
			JSON: []byte(`"//example.com/path"`),
		},
		{
			JSON: []byte(`"//درود.com/page"`),
		},
		{
			JSON: []byte(`"//example.com/世界"`),
		},



		// Query strings with normalized unreserved characters
		{
			JSON: []byte(`"http://example.com/page?name=Alice"`),
		},
		{
			JSON: []byte(`"http://example.com/search?q=世界"`),
		},



		// Fragments with percent-encoded characters
		{
			JSON: []byte(`"http://example.com/page#hello%20world"`),
		},
		{
			JSON: []byte(`"http://example.com/page#世界"`),
		},
		{
			JSON: []byte(`"http://example.com/page#foo%2Fbar"`),
		},



		// Opaque URIs with fragments
		{
			JSON: []byte(`"mailto:user@example.com#fragment"`),
		},
		{
			JSON: []byte(`"mailto:joe@example.net#sec1"`),
		},
		{
			JSON: []byte(`"urn:isbn:0451450523#chapter3"`),
		},
		{
			JSON: []byte(`"acct:joeblow@host.example#AbCdEfG"`),
		},
		{
			JSON: []byte(`"acct:joeblow@host.example#hello%20world%20🙂"`),
		},

		// Opaque URIs with queries
		{
			JSON: []byte(`"mailto:user@example.com?subject=hello"`),
		},

		// Opaque URIs with queries and fragments
		{
			JSON: []byte(`"mailto:user@example.com?subject=hello#frag"`),
		},



		// Blank Node Identifiers with their Blank Node Labels
		{
			JSON: []byte(`"_:abc123"`),
		},
		{
			JSON: []byte(`"_:ABC123"`),
		},
		{
			JSON: []byte(`"_:Hello-World"`),
		},
		{
			JSON: []byte(`"_:b0"`),
		},
		{
			JSON: []byte(`"_:address84"`),
		},
		{
			JSON: []byte(`"_:n1"`),
		},
		{
			JSON: []byte(`"_:ed7ba470-8e54-465e-825c-99712043e01c"`),
		},
		{
			JSON: []byte(`"_:label123"`),
		},
	}

	for testNumber, test := range tests {

		var id jsonld.ID

		err := id.UnmarshalJSON(test.JSON)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error from UnmarshalJSON but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON: %s", test.JSON)
			continue
		}

		actual, err := id.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error from MarshalJSON but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON: %s", test.JSON)
			continue
		}

		if !bytes.Equal(test.JSON, actual) {
			t.Errorf("For test #%d, the round-trip (unmarshal then marshal) did not produce the same bytes.", testNumber)
			t.Logf("ORIGINAL:\n%s", test.JSON)
			t.Logf("ROUND-TRIPPED:\n%s", actual)
			continue
		}
	}
}
