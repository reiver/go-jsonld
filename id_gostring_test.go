package jsonld_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
)

func TestID_GoString(t *testing.T) {

	tests := []struct{
		ID jsonld.ID
		Expected string
	}{
		{
			ID:        jsonld.NoID(),
			Expected: "jsonld.NoID()",
		},



		{
			ID:        jsonld.SomeID(""),
			Expected: `jsonld.SomeID("")`,
		},



		{
			ID:        jsonld.SomeID("apple"),
			Expected: `jsonld.SomeID("apple")`,
		},
		{
			ID:        jsonld.SomeID("BANANA"),
			Expected: `jsonld.SomeID("BANANA")`,
		},
		{
			ID:        jsonld.SomeID("Cherry"),
			Expected: `jsonld.SomeID("Cherry")`,
		},
		{
			ID:        jsonld.SomeID("dAtE"),
			Expected: `jsonld.SomeID("dAtE")`,
		},



		{
			ID:        jsonld.SomeID("https://example.com/something"),
			Expected: `jsonld.SomeID("https://example.com/something")`,
		},
		{
			ID:        jsonld.SomeID("HTTPS://example.com/something"),
			Expected: `jsonld.SomeID("https://example.com/something")`,
		},
		{
			ID:        jsonld.SomeID("https://EXAMPLE.COM/something"),
			Expected: `jsonld.SomeID("https://example.com/something")`,
		},



		{
			ID:        jsonld.SomeID("http://درود.com/"),
			Expected: `jsonld.SomeID("http://درود.com/")`,
		},
		{
			ID:        jsonld.SomeID("http://xn--ugbaf6g.com/"),
			Expected: `jsonld.SomeID("http://درود.com/")`,
		},
		{
			ID:        jsonld.SomeID("http://XN--UGBAF6G.COM/"),
			Expected: `jsonld.SomeID("http://درود.com/")`,
		},



		{
			ID:        jsonld.SomeID("http://안녕.kr/세상"),
			Expected: `jsonld.SomeID("http://안녕.kr/세상")`,
		},
		{
			ID:        jsonld.SomeID("http://xn--o70b819a.kr/%EC%84%B8%EC%83%81"),
			Expected: `jsonld.SomeID("http://안녕.kr/세상")`,
		},
		{
			ID:        jsonld.SomeID("http://XN--O70B819A.KR/%EC%84%B8%EC%83%81"),
			Expected: `jsonld.SomeID("http://안녕.kr/세상")`,
		},



		{
			ID:        jsonld.SomeID("finger://joeblow@😈.example"),
			Expected: `jsonld.SomeID("finger://joeblow@😈.example")`,
		},



		{
			ID:        jsonld.SomeID("PuNcH://eXaMpLe.CoM/😈-%F0%9F%98%88.md"),
			Expected: `jsonld.SomeID("punch://example.com/😈-😈.md")`,
		},



		{
			ID:        jsonld.SomeID("http://example.net/apple/BANANA/Cherry/dAtE.php?once=1&TWICE=22&Thrice=333&fOuRcE=4444#aBc123"),
			Expected: `jsonld.SomeID("http://example.net/apple/BANANA/Cherry/dAtE.php?once=1&TWICE=22&Thrice=333&fOuRcE=4444#aBc123")`,
		},



		// Reserved percent-encoded characters in path are preserved
		{
			ID:        jsonld.SomeID("http://example.com/foo%2Fbar"),
			Expected: `jsonld.SomeID("http://example.com/foo%2Fbar")`,
		},
		{
			ID:        jsonld.SomeID("http://example.com/a%3Fb%3Dc"),
			Expected: `jsonld.SomeID("http://example.com/a%3Fb%3Dc")`,
		},
		{
			ID:        jsonld.SomeID("http://example.com/100%25done"),
			Expected: `jsonld.SomeID("http://example.com/100%25done")`,
		},



		// Relative IRIs (no scheme)
		{
			ID:        jsonld.SomeID("//EXAMPLE.COM/path"),
			Expected: `jsonld.SomeID("//example.com/path")`,
		},
		{
			ID:        jsonld.SomeID("//XN--UGBAF6G.COM/page"),
			Expected: `jsonld.SomeID("//درود.com/page")`,
		},
		{
			ID:        jsonld.SomeID("//example.com/%E4%B8%96%E7%95%8C"),
			Expected: `jsonld.SomeID("//example.com/世界")`,
		},



		// Query strings with percent-encoded unreserved characters
		{
			ID:        jsonld.SomeID("http://example.com/page?name=%41lice"),
			Expected: `jsonld.SomeID("http://example.com/page?name=Alice")`,
		},
		{
			ID:        jsonld.SomeID("http://example.com/search?q=%E4%B8%96%E7%95%8C"),
			Expected: `jsonld.SomeID("http://example.com/search?q=世界")`,
		},
		{
			ID:        jsonld.SomeID("http://example.com/search?q=hello%20world&lang=%65%6E"),
			Expected: `jsonld.SomeID("http://example.com/search?q=hello%20world&lang=en")`,
		},



		// Fragments with percent-encoded characters
		{
			ID:        jsonld.SomeID("http://example.com/page#hello%20world"),
			Expected: `jsonld.SomeID("http://example.com/page#hello%20world")`,
		},
		{
			ID:        jsonld.SomeID("http://example.com/page#%E4%B8%96%E7%95%8C"),
			Expected: `jsonld.SomeID("http://example.com/page#世界")`,
		},
		{
			ID:        jsonld.SomeID("http://example.com/page#foo%2Fbar"),
			Expected: `jsonld.SomeID("http://example.com/page#foo%2Fbar")`,
		},



		// Opaque URIs with fragments
		{
			ID:        jsonld.SomeID("mailto:user@example.com#fragment"),
			Expected: `jsonld.SomeID("mailto:user@example.com#fragment")`,
		},
		{
			ID:        jsonld.SomeID("mailto:joe@example.net#sec1"),
			Expected: `jsonld.SomeID("mailto:joe@example.net#sec1")`,
		},
		{
			ID:        jsonld.SomeID("urn:isbn:0451450523#chapter3"),
			Expected: `jsonld.SomeID("urn:isbn:0451450523#chapter3")`,
		},
		{
			ID:        jsonld.SomeID("acct:joeblow@host.example#AbCdEfG"),
			Expected: `jsonld.SomeID("acct:joeblow@host.example#AbCdEfG")`,
		},
		{
			ID:        jsonld.SomeID("acct:joeblow@host.example#hello%20world%20🙂"),
			Expected: `jsonld.SomeID("acct:joeblow@host.example#hello%20world%20🙂")`,
		},
		{
			ID:        jsonld.SomeID("acct:joeblow@host.example#hello%20world%20%F0%9F%99%82"),
			Expected: `jsonld.SomeID("acct:joeblow@host.example#hello%20world%20🙂")`,
		},

		// Opaque URIs with queries
		{
			ID:        jsonld.SomeID("mailto:user@example.com?subject=hello"),
			Expected: `jsonld.SomeID("mailto:user@example.com?subject=hello")`,
		},
		{
			ID:        jsonld.SomeID("mailto:joe@example.net?subject=meeting&body=hi"),
			Expected: `jsonld.SomeID("mailto:joe@example.net?subject=meeting&body=hi")`,
		},

		// Opaque URIs with queries and fragments
		{
			ID:        jsonld.SomeID("mailto:user@example.com?subject=hello#frag"),
			Expected: `jsonld.SomeID("mailto:user@example.com?subject=hello#frag")`,
		},



		// Blank Node Identifiers with their Blank Node Labels
		{
			ID:        jsonld.SomeID("_:abc123"),
			Expected: `jsonld.SomeID("_:abc123")`,
		},
		{
			ID:        jsonld.SomeID("_:ABC123"),
			Expected: `jsonld.SomeID("_:ABC123")`,
		},
		{
			ID:        jsonld.SomeID("_:Hello-World"),
			Expected: `jsonld.SomeID("_:Hello-World")`,
		},
		{
			ID:        jsonld.SomeID("_:b0"),
			Expected: `jsonld.SomeID("_:b0")`,
		},
		{
			ID:        jsonld.SomeID("_:address84"),
			Expected: `jsonld.SomeID("_:address84")`,
		},
		{
			ID:        jsonld.SomeID("_:n1"),
			Expected: `jsonld.SomeID("_:n1")`,
		},
		{
			ID:        jsonld.SomeID("_:ed7ba470-8e54-465e-825c-99712043e01c"),
			Expected: `jsonld.SomeID("_:ed7ba470-8e54-465e-825c-99712043e01c")`,
		},
		{
			ID:        jsonld.SomeID("_:label123"),
			Expected: `jsonld.SomeID("_:label123")`,
		},
	}

	for testNumber, test := range tests {

		actual := test.ID.GoString()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual go-string is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			continue
		}
	}
}
