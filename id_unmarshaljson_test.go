package jsonld_test

import (
	"testing"

	"github.com/reiver/go-jsonld"
)

func TestID_UnmarshalJSON(t *testing.T) {

	tests := []struct{
		JSON     []byte
		Expected jsonld.ID
	}{
		{
			JSON:          []byte(`null`),
			Expected: jsonld.NoID(),
		},



		{
			JSON:           []byte(`""`),
			Expected: jsonld.SomeID(""),
		},



		{
			JSON:           []byte(`"apple"`),
			Expected: jsonld.SomeID("apple"),
		},
		{
			JSON:           []byte(`"BANANA"`),
			Expected: jsonld.SomeID("BANANA"),
		},
		{
			JSON:           []byte(`"Cherry"`),
			Expected: jsonld.SomeID("Cherry"),
		},
		{
			JSON:           []byte(`"dAtE"`),
			Expected: jsonld.SomeID("dAtE"),
		},



		{
			JSON:           []byte(`"https://example.com/something"`),
			Expected: jsonld.SomeID("https://example.com/something"),
		},
		{
			JSON:           []byte(`"HTTPS://example.com/something"`),
			Expected: jsonld.SomeID("https://example.com/something"),
		},
		{
			JSON:           []byte(`"https://EXAMPLE.COM/something"`),
			Expected: jsonld.SomeID("https://example.com/something"),
		},



		{
			JSON:           []byte(`"http://درود.com/"`),
			Expected: jsonld.SomeID("http://درود.com/"),
		},
		{
			JSON:           []byte(`"http://xn--ugbaf6g.com/"`),
			Expected: jsonld.SomeID("http://درود.com/"),
		},
		{
			JSON:           []byte(`"http://XN--UGBAF6G.COM/"`),
			Expected: jsonld.SomeID("http://درود.com/"),
		},



		{
			JSON:           []byte(`"http://안녕.kr/세상"`),
			Expected: jsonld.SomeID("http://안녕.kr/세상"),
		},
		{
			JSON:           []byte(`"http://xn--o70b819a.kr/%EC%84%B8%EC%83%81"`),
			Expected: jsonld.SomeID("http://안녕.kr/세상"),
		},
		{
			JSON:           []byte(`"http://XN--O70B819A.KR/%EC%84%B8%EC%83%81"`),
			Expected: jsonld.SomeID("http://안녕.kr/세상"),
		},



		{
			JSON:           []byte(`"finger://joeblow@😈.example"`),
			Expected: jsonld.SomeID("finger://joeblow@😈.example"),
		},



		{
			JSON:           []byte(`"PuNcH://eXaMpLe.CoM/😈-%F0%9F%98%88.md"`),
			Expected: jsonld.SomeID("punch://example.com/😈-😈.md"),
		},



		{
			JSON:           []byte(`"http://example.net/apple/BANANA/Cherry/dAtE.php?once=1&TWICE=22&Thrice=333&fOuRcE=4444#aBc123"`),
			Expected: jsonld.SomeID("http://example.net/apple/BANANA/Cherry/dAtE.php?once=1&TWICE=22&Thrice=333&fOuRcE=4444#aBc123"),
		},



		// Reserved percent-encoded characters in path are preserved
		{
			JSON:           []byte(`"http://example.com/foo%2Fbar"`),
			Expected: jsonld.SomeID("http://example.com/foo%2Fbar"),
		},
		{
			JSON:           []byte(`"http://example.com/a%3Fb%3Dc"`),
			Expected: jsonld.SomeID("http://example.com/a%3Fb%3Dc"),
		},
		{
			JSON:           []byte(`"http://example.com/100%25done"`),
			Expected: jsonld.SomeID("http://example.com/100%25done"),
		},



		// Relative IRIs (no scheme)
		{
			JSON:           []byte(`"//EXAMPLE.COM/path"`),
			Expected: jsonld.SomeID("//example.com/path"),
		},
		{
			JSON:           []byte(`"//XN--UGBAF6G.COM/page"`),
			Expected: jsonld.SomeID("//درود.com/page"),
		},
		{
			JSON:           []byte(`"//example.com/%E4%B8%96%E7%95%8C"`),
			Expected: jsonld.SomeID("//example.com/世界"),
		},



		// Query strings with percent-encoded unreserved characters
		{
			JSON:           []byte(`"http://example.com/page?name=%41lice"`),
			Expected: jsonld.SomeID("http://example.com/page?name=Alice"),
		},
		{
			JSON:           []byte(`"http://example.com/search?q=%E4%B8%96%E7%95%8C"`),
			Expected: jsonld.SomeID("http://example.com/search?q=世界"),
		},
		{
			JSON:           []byte(`"http://example.com/search?q=hello%20world&lang=%65%6E"`),
			Expected: jsonld.SomeID("http://example.com/search?q=hello%20world&lang=en"),
		},



		// Fragments with percent-encoded characters
		{
			JSON:           []byte(`"http://example.com/page#hello%20world"`),
			Expected: jsonld.SomeID("http://example.com/page#hello%20world"),
		},
		{
			JSON:           []byte(`"http://example.com/page#%E4%B8%96%E7%95%8C"`),
			Expected: jsonld.SomeID("http://example.com/page#世界"),
		},
		{
			JSON:           []byte(`"http://example.com/page#foo%2Fbar"`),
			Expected: jsonld.SomeID("http://example.com/page#foo%2Fbar"),
		},



		// Opaque URIs with fragments
		{
			JSON:           []byte(`"mailto:user@example.com#fragment"`),
			Expected: jsonld.SomeID("mailto:user@example.com#fragment"),
		},
		{
			JSON:           []byte(`"mailto:joe@example.net#sec1"`),
			Expected: jsonld.SomeID("mailto:joe@example.net#sec1"),
		},
		{
			JSON:           []byte(`"urn:isbn:0451450523#chapter3"`),
			Expected: jsonld.SomeID("urn:isbn:0451450523#chapter3"),
		},
		{
			JSON:           []byte(`"acct:joeblow@host.example#AbCdEfG"`),
			Expected: jsonld.SomeID("acct:joeblow@host.example#AbCdEfG"),
		},
		{
			JSON:           []byte(`"acct:joeblow@host.example#hello%20world%20🙂"`),
			Expected: jsonld.SomeID("acct:joeblow@host.example#hello%20world%20🙂"),
		},
		{
			JSON:           []byte(`"acct:joeblow@host.example#hello%20world%20%F0%9F%99%82"`),
			Expected: jsonld.SomeID("acct:joeblow@host.example#hello%20world%20🙂"),
		},

		// Opaque URIs with queries
		{
			JSON:           []byte(`"mailto:user@example.com?subject=hello"`),
			Expected: jsonld.SomeID("mailto:user@example.com?subject=hello"),
		},
		{
			JSON:           []byte(`"mailto:joe@example.net?subject=meeting&body=hi"`),
			Expected: jsonld.SomeID("mailto:joe@example.net?subject=meeting&body=hi"),
		},

		// Opaque URIs with queries and fragments
		{
			JSON:           []byte(`"mailto:user@example.com?subject=hello#frag"`),
			Expected: jsonld.SomeID("mailto:user@example.com?subject=hello#frag"),
		},



		// Blank Node Identifiers with their Blank Node Labels
		{
			JSON:           []byte(`"_:abc123"`),
			Expected: jsonld.SomeID("_:abc123"),
		},
		{
			JSON:           []byte(`"_:ABC123"`),
			Expected: jsonld.SomeID("_:ABC123"),
		},
		{
			JSON:           []byte(`"_:Hello-World"`),
			Expected: jsonld.SomeID("_:Hello-World"),
		},
		{
			JSON:           []byte(`"_:b0"`),
			Expected: jsonld.SomeID("_:b0"),
		},
		{
			JSON:           []byte(`"_:address84"`),
			Expected: jsonld.SomeID("_:address84"),
		},
		{
			JSON:           []byte(`"_:n1"`),
			Expected: jsonld.SomeID("_:n1"),
		},
		{
			JSON:           []byte(`"_:ed7ba470-8e54-465e-825c-99712043e01c"`),
			Expected: jsonld.SomeID("_:ed7ba470-8e54-465e-825c-99712043e01c"),
		},
		{
			JSON:           []byte(`"_:label123"`),
			Expected: jsonld.SomeID("_:label123"),
		},
	}

	for testNumber, test := range tests {

		var actual jsonld.ID

		err := actual.UnmarshalJSON(test.JSON)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual json-unmarshaled id is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}

func TestID_UnmarshalJSON_nilReceiver(t *testing.T) {

	var id *jsonld.ID

	err := id.UnmarshalJSON([]byte(`"hello"`))
	if nil == err {
		t.Errorf("expected an error but did not actually get one.")
		return
	}
}

func TestID_UnmarshalJSON_fail(t *testing.T) {

	tests := []struct{
		JSON []byte
	}{
		{
			JSON: nil,
		},
		{
			JSON: []byte{},
		},



		{
			JSON: []byte(`123`),
		},
		{
			JSON: []byte(`45.67`),
		},
		{
			JSON: []byte(`true`),
		},
		{
			JSON: []byte(`false`),
		},



		{
			JSON: []byte(`{}`),
		},
		{
			JSON: []byte(`{"key":"value"}`),
		},
		{
			JSON: []byte(`[]`),
		},
		{
			JSON: []byte(`["apple","banana"]`),
		},
	}

	for testNumber, test := range tests {

		var actual jsonld.ID

		err := actual.UnmarshalJSON(test.JSON)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one." , testNumber)
			t.Logf("JSON: %s", test.JSON)
			t.Logf("ACTUAL: %#v", actual)
			continue
		}
	}
}
