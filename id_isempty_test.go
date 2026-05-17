package jsonld_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
)

func TestID_IsEmpty(t *testing.T) {

	tests := []struct{
		ID       jsonld.ID
		Expected bool
	}{
		{
			ID:       jsonld.NoID(),
			Expected: true,
		},



		{
			ID:       jsonld.SomeID(""),
			Expected: false,
		},



		{
			ID:       jsonld.SomeID("apple"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("BANANA"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("Cherry"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("dAtE"),
			Expected: false,
		},



		{
			ID:       jsonld.SomeID("https://example.com/something"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("HTTPS://example.com/something"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("https://EXAMPLE.COM/something"),
			Expected: false,
		},



		{
			ID:       jsonld.SomeID("http://درود.com/"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("http://xn--ugbaf6g.com/"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("http://XN--UGBAF6G.COM/"),
			Expected: false,
		},



		{
			ID:       jsonld.SomeID("http://안녕.kr/세상"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("http://xn--o70b819a.kr/%EC%84%B8%EC%83%81"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("http://XN--O70B819A.KR/%EC%84%B8%EC%83%81"),
			Expected: false,
		},



		{
			ID:       jsonld.SomeID("finger://joeblow@😈.example"),
			Expected: false,
		},



		{
			ID:       jsonld.SomeID("PuNcH://eXaMpLe.CoM/😈-%F0%9F%98%88.md"),
			Expected: false,
		},



		{
			ID:       jsonld.SomeID("http://example.net/apple/BANANA/Cherry/dAtE.php?once=1&TWICE=22&Thrice=333&fOuRcE=4444#aBc123"),
			Expected: false,
		},



		// Reserved percent-encoded characters in path are preserved
		{
			ID:       jsonld.SomeID("http://example.com/foo%2Fbar"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("http://example.com/a%3Fb%3Dc"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("http://example.com/100%25done"),
			Expected: false,
		},



		// Relative IRIs (no scheme)
		{
			ID:       jsonld.SomeID("//EXAMPLE.COM/path"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("//XN--UGBAF6G.COM/page"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("//example.com/%E4%B8%96%E7%95%8C"),
			Expected: false,
		},



		// Query strings with percent-encoded unreserved characters
		{
			ID:       jsonld.SomeID("http://example.com/page?name=%41lice"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("http://example.com/search?q=%E4%B8%96%E7%95%8C"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("http://example.com/search?q=hello%20world&lang=%65%6E"),
			Expected: false,
		},



		// Fragments with percent-encoded characters
		{
			ID:       jsonld.SomeID("http://example.com/page#hello%20world"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("http://example.com/page#%E4%B8%96%E7%95%8C"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("http://example.com/page#foo%2Fbar"),
			Expected: false,
		},



		// Opaque URIs with fragments
		{
			ID:       jsonld.SomeID("mailto:user@example.com#fragment"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("mailto:joe@example.net#sec1"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("urn:isbn:0451450523#chapter3"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("acct:joeblow@host.example#AbCdEfG"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("acct:joeblow@host.example#hello%20world%20🙂"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("acct:joeblow@host.example#hello%20world%20%F0%9F%99%82"),
			Expected: false,
		},

		// Opaque URIs with queries
		{
			ID:       jsonld.SomeID("mailto:user@example.com?subject=hello"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("mailto:joe@example.net?subject=meeting&body=hi"),
			Expected: false,
		},

		// Opaque URIs with queries and fragments
		{
			ID:       jsonld.SomeID("mailto:user@example.com?subject=hello#frag"),
			Expected: false,
		},



		// Blank Node Identifiers with their Blank Node Labels
		{
			ID:       jsonld.SomeID("_:abc123"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("_:ABC123"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("_:Hello-World"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("_:b0"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("_:address84"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("_:n1"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("_:ed7ba470-8e54-465e-825c-99712043e01c"),
			Expected: false,
		},
		{
			ID:       jsonld.SomeID("_:label123"),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := test.ID.IsEmpty()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual is-empty is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}
