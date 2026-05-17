package jsonld_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
)

func TestID_MarshalJSON(t *testing.T) {

	tests := []struct{
		ID       jsonld.ID
		Expected []byte
	}{
		{
			ID: jsonld.NoID(),
			Expected: []byte(`null`),
		},



		{
			ID: jsonld.SomeID(""),
			Expected: []byte(`""`),
		},



		{
			ID: jsonld.SomeID("apple"),
			Expected: []byte(`"apple"`),
		},
		{
			ID: jsonld.SomeID("BANANA"),
			Expected: []byte(`"BANANA"`),
		},
		{
			ID: jsonld.SomeID("Cherry"),
			Expected: []byte(`"Cherry"`),
		},
		{
			ID: jsonld.SomeID("dAtE"),
			Expected: []byte(`"dAtE"`),
		},



		{
			ID: jsonld.SomeID("https://example.com/something"),
			Expected: []byte(`"https://example.com/something"`),
		},
		{
			ID: jsonld.SomeID("HTTPS://example.com/something"),
			Expected: []byte(`"https://example.com/something"`),
		},
		{
			ID: jsonld.SomeID("https://EXAMPLE.COM/something"),
			Expected: []byte(`"https://example.com/something"`),
		},



		{
			ID: jsonld.SomeID("http://درود.com/"),
			Expected: []byte(`"http://درود.com/"`),
		},
		{
			ID: jsonld.SomeID("http://xn--ugbaf6g.com/"),
			Expected: []byte(`"http://درود.com/"`),
		},
		{
			ID: jsonld.SomeID("http://XN--UGBAF6G.COM/"),
			Expected: []byte(`"http://درود.com/"`),
		},



		{
			ID: jsonld.SomeID("http://안녕.kr/세상"),
			Expected: []byte(`"http://안녕.kr/세상"`),
		},
		{
			ID: jsonld.SomeID("http://xn--o70b819a.kr/%EC%84%B8%EC%83%81"),
			Expected: []byte(`"http://안녕.kr/세상"`),
		},
		{
			ID: jsonld.SomeID("http://XN--O70B819A.KR/%EC%84%B8%EC%83%81"),
			Expected: []byte(`"http://안녕.kr/세상"`),
		},



		{
			ID: jsonld.SomeID("finger://joeblow@😈.example"),
			Expected: []byte(`"finger://joeblow@😈.example"`),
		},



		{
			ID: jsonld.SomeID("PuNcH://eXaMpLe.CoM/😈-%F0%9F%98%88.md"),
			Expected: []byte(`"punch://example.com/😈-😈.md"`),
		},



		{
			ID: jsonld.SomeID("http://example.net/apple/BANANA/Cherry/dAtE.php?once=1&TWICE=22&Thrice=333&fOuRcE=4444#aBc123"),
			Expected: []byte(`"http://example.net/apple/BANANA/Cherry/dAtE.php?once=1\u0026TWICE=22\u0026Thrice=333\u0026fOuRcE=4444#aBc123"`),
		},



		// Reserved percent-encoded characters in path are preserved
		{
			ID: jsonld.SomeID("http://example.com/foo%2Fbar"),
			Expected: []byte(`"http://example.com/foo%2Fbar"`),
		},
		{
			ID: jsonld.SomeID("http://example.com/a%3Fb%3Dc"),
			Expected: []byte(`"http://example.com/a%3Fb%3Dc"`),
		},
		{
			ID: jsonld.SomeID("http://example.com/100%25done"),
			Expected: []byte(`"http://example.com/100%25done"`),
		},



		// Relative IRIs (no scheme)
		{
			ID: jsonld.SomeID("//EXAMPLE.COM/path"),
			Expected: []byte(`"//example.com/path"`),
		},
		{
			ID: jsonld.SomeID("//XN--UGBAF6G.COM/page"),
			Expected: []byte(`"//درود.com/page"`),
		},
		{
			ID: jsonld.SomeID("//example.com/%E4%B8%96%E7%95%8C"),
			Expected: []byte(`"//example.com/世界"`),
		},



		// Query strings with percent-encoded unreserved characters
		{
			ID: jsonld.SomeID("http://example.com/page?name=%41lice"),
			Expected: []byte(`"http://example.com/page?name=Alice"`),
		},
		{
			ID: jsonld.SomeID("http://example.com/search?q=%E4%B8%96%E7%95%8C"),
			Expected: []byte(`"http://example.com/search?q=世界"`),
		},
		{
			ID: jsonld.SomeID("http://example.com/search?q=hello%20world\u0026lang=%65%6E"),
			Expected: []byte(`"http://example.com/search?q=hello%20world\u0026lang=en"`),
		},



		// Fragments with percent-encoded characters
		{
			ID: jsonld.SomeID("http://example.com/page#hello%20world"),
			Expected: []byte(`"http://example.com/page#hello%20world"`),
		},
		{
			ID: jsonld.SomeID("http://example.com/page#%E4%B8%96%E7%95%8C"),
			Expected: []byte(`"http://example.com/page#世界"`),
		},
		{
			ID: jsonld.SomeID("http://example.com/page#foo%2Fbar"),
			Expected: []byte(`"http://example.com/page#foo%2Fbar"`),
		},



		// Opaque URIs with fragments
		{
			ID: jsonld.SomeID("mailto:user@example.com#fragment"),
			Expected: []byte(`"mailto:user@example.com#fragment"`),
		},
		{
			ID: jsonld.SomeID("mailto:joe@example.net#sec1"),
			Expected: []byte(`"mailto:joe@example.net#sec1"`),
		},
		{
			ID: jsonld.SomeID("urn:isbn:0451450523#chapter3"),
			Expected: []byte(`"urn:isbn:0451450523#chapter3"`),
		},
		{
			ID: jsonld.SomeID("acct:joeblow@host.example#AbCdEfG"),
			Expected: []byte(`"acct:joeblow@host.example#AbCdEfG"`),
		},
		{
			ID: jsonld.SomeID("acct:joeblow@host.example#hello%20world%20🙂"),
			Expected: []byte(`"acct:joeblow@host.example#hello%20world%20🙂"`),
		},
		{
			ID: jsonld.SomeID("acct:joeblow@host.example#hello%20world%20%F0%9F%99%82"),
			Expected: []byte(`"acct:joeblow@host.example#hello%20world%20🙂"`),
		},

		// Opaque URIs with queries
		{
			ID: jsonld.SomeID("mailto:user@example.com?subject=hello"),
			Expected: []byte(`"mailto:user@example.com?subject=hello"`),
		},
		{
			ID: jsonld.SomeID("mailto:joe@example.net?subject=meeting&body=hi"),
			Expected: []byte(`"mailto:joe@example.net?subject=meeting\u0026body=hi"`),
		},

		// Opaque URIs with queries and fragments
		{
			ID: jsonld.SomeID("mailto:user@example.com?subject=hello#frag"),
			Expected: []byte(`"mailto:user@example.com?subject=hello#frag"`),
		},



		// Blank Node Identifiers with their Blank Node Labels
		{
			ID: jsonld.SomeID("_:abc123"),
			Expected: []byte(`"_:abc123"`),
		},
		{
			ID: jsonld.SomeID("_:ABC123"),
			Expected: []byte(`"_:ABC123"`),
		},
		{
			ID: jsonld.SomeID("_:Hello-World"),
			Expected: []byte(`"_:Hello-World"`),
		},
		{
			ID: jsonld.SomeID("_:b0"),
			Expected: []byte(`"_:b0"`),
		},
		{
			ID: jsonld.SomeID("_:address84"),
			Expected: []byte(`"_:address84"`),
		},
		{
			ID: jsonld.SomeID("_:n1"),
			Expected: []byte(`"_:n1"`),
		},
		{
			ID: jsonld.SomeID("_:ed7ba470-8e54-465e-825c-99712043e01c"),
			Expected: []byte(`"_:ed7ba470-8e54-465e-825c-99712043e01c"`),
		},
		{
			ID: jsonld.SomeID("_:label123"),
			Expected: []byte(`"_:label123"`),
		},
	}

	for testNumber, test := range tests {

		actual, err := test.ID.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual json-marshaled id is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			continue
		}
	}
}
