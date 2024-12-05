package jsonld_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Values []any
		Expected []byte
	}{
		// 0
		{
			Expected: []byte(`{}`),
		},



		// 1
		{
			Values: []any{
				struct{
					Apple string
					Banana int
					Cherry bool
				}{
					Apple: "ONE",
					Banana: 2,
					Cherry: true,
				},
			},
			Expected: []byte(
				`{`+
					`"Apple":"ONE"`+
					`,`+
					`"Banana":2`+
					`,`+
					`"Cherry":true`+
				`}`),
		},



		// 2
		{
			Values: []any{
				struct{
					Apple string `jsonld:"apple"`
					Banana int
					Cherry bool  `jsonld:"cherry"`
				}{
					Apple: "ONE",
					Banana: 2,
					Cherry: true,
				},
			},
			Expected: []byte(
				`{`+
					`"apple":"ONE"`+
					`,`+
					`"Banana":2`+
					`,`+
					`"cherry":true`+
				`}`),
		},



		// 3
		{
			Values: []any{
				struct{
					Apple string `jsonld:"apple"`
					Banana int   `jsonld:"banana"`
					Cherry bool  `jsonld:"cherry"`
				}{
					Apple: "ONE",
					Banana: 2,
					Cherry: true,
				},
			},
			Expected: []byte(
				`{`+
					`"apple":"ONE"`+
					`,`+
					`"banana":2`+
					`,`+
					`"cherry":true`+
				`}`),
		},



		// 4
		{
			Values: []any{
				struct{
					NS jsonld.NameSpace `jsonld:"http://example.com/ns#"`
					Apple string `jsonld:"apple"`
					Banana int
					Cherry bool  `jsonld:"cherry"`
				}{
					Apple: "ONE",
					Banana: 2,
					Cherry: true,
				},
			},
			Expected: []byte(
				`{`+
					`"@context":[`+
						`"http://example.com/ns#"`+
					`]`+
					`,`+
					`"apple":"ONE"`+
					`,`+
					`"Banana":2`+
					`,`+
					`"cherry":true`+
				`}`),
		},



		// 5
		{
			Values: []any{
				struct{
					NS jsonld.NameSpace `jsonld:"http://example.com/ns#"`
					P  jsonld.Prefix    `jsonld:"ex"`
					Apple string `jsonld:"apple"`
					Banana int
					Cherry bool  `jsonld:"cherry"`
				}{
					Apple: "ONE",
					Banana: 2,
					Cherry: true,
				},
			},
			Expected: []byte(
				`{`+
					`"@context":{`+
						`"ex":"http://example.com/ns#"`+
						`,`+
						`"apple":"ex:apple"`+
						`,`+
						`"Banana":"ex:Banana"`+
						`,`+
						`"cherry":"ex:cherry"`+
					`}`+
					`,`+
					`"apple":"ONE"`+
					`,`+
					`"Banana":2`+
					`,`+
					`"cherry":true`+
				`}`),
		},



		// 6
		{
			Values: []any{
				struct{
					NameSpace jsonld.NameSpace `jsonld:"http://joinmastodon.org/ns#"`
					Prefix    jsonld.Prefix    `jsonld:"toot"`

					AttributionDomains string `jsonld:"attributionDomains"`
					Emoji              string
					Discoverable       bool   `jsonld:"discoverable"`
					Featured           bool   `jsonld:"featured"`
					FeaturedTags       string `jsonld:"featuredTags"`
					FocalPoint       []any    `jsonld:"focalPoint"`
					Indexable          bool   `jsonld:"indexable"`
					Memorial           bool   `jsonld:"memorial"`
					Suspended          bool   `jsonld:"suspended"`
				}{
					Discoverable: true,
					Featured:     true,
					FeaturedTags: "#once #twice #thrice #fource",
					FocalPoint: []any{12,100},
					Indexable:    true,
					Memorial:     false,
					Suspended:    false,
				},
			},
			Expected: []byte(
				`{`+
					`"@context":{`+
						`"toot":"http://joinmastodon.org/ns#"`+
						`,`+
						`"attributionDomains":"toot:attributionDomains"`+
						`,`+
						`"Emoji":"toot:Emoji"`+
						`,`+
						`"discoverable":"toot:discoverable"`+
						`,`+
						`"featured":"toot:featured"`+
						`,`+
						`"featuredTags":"toot:featuredTags"`+
						`,`+
						`"focalPoint":"toot:focalPoint"`+
						`,`+
						`"indexable":"toot:indexable"`+
						`,`+
						`"memorial":"toot:memorial"`+
						`,`+
						`"suspended":"toot:suspended"`+
					`}`+
					`,`+
					`"attributionDomains":""`+
					`,`+
					`"Emoji":""`+
					`,`+
					`"discoverable":true`+
					`,`+
					`"featured":true`+
					`,`+
					`"featuredTags":"#once #twice #thrice #fource"`+
					`,`+
					`"focalPoint":[12,100]`+
					`,`+
					`"indexable":true`+
					`,`+
					`"memorial":false`+
					`,`+
					`"suspended":false`+
				`}`),
		},



		// 7
		{
			Values: []any{
				struct{
					ID      string `jsonld:"id"`
					Name    string `jsonld:"name"`
					Summary string `jsonld:"summary"`
					Type    string `jsonld:"type"`
				}{
					ID: "urn:uuid:88c6a753-d6d5-45eb-bc18-9b1089e7b1f8",
					Name: "Joe Blow",
					Summary: "Hello world!",
					Type: "Person",
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"http://joinmastodon.org/ns#"`
					Prefix    jsonld.Prefix    `jsonld:"toot"`

					AttributionDomains string `jsonld:"attributionDomains"`
					Emoji              string
					Discoverable       bool   `jsonld:"discoverable"`
					Featured           bool   `jsonld:"featured"`
					FeaturedTags       string `jsonld:"featuredTags"`
					FocalPoint       []any    `jsonld:"focalPoint"`
					Indexable          bool   `jsonld:"indexable"`
					Memorial           bool   `jsonld:"memorial"`
					Suspended          bool   `jsonld:"suspended"`
				}{
					Discoverable: true,
					Featured:     true,
					FeaturedTags: "#once #twice #thrice #fource",
					FocalPoint: []any{12,100},
					Indexable:    true,
					Memorial:     false,
					Suspended:    false,
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"https://www.w3.org/ns/activitystreams"`
					Prefix    jsonld.Prefix    `jsonld:"as"`

					AlsoKnownAs []string `jsonld:"alsoKnownAs"`
					Hashtag     []string `jsonld:"Hashtag"`
					MovedTo       string `jsonld:"movedTo"`
				}{
					Hashtag: []string{
						"#fediverse",
					},
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/security/v1"`

					Signature string `jsonld:"signature"`
				}{
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"https://banana.example/ns#"`

					Color string `jsonld:"colour"`
				}{
					Color: "yellow",
				},
			},
			Expected: []byte(
				`{`+
					`"@context":[`+
						`"https://w3id.org/security/v1"`+
						`,`+
						`"https://banana.example/ns#"`+
						`,`+
						`{`+
							`"toot":"http://joinmastodon.org/ns#"`+
							`,`+
							`"as":"https://www.w3.org/ns/activitystreams"`+

							`,`+

							`"attributionDomains":"toot:attributionDomains"`+
							`,`+
							`"Emoji":"toot:Emoji"`+
							`,`+
							`"discoverable":"toot:discoverable"`+
							`,`+
							`"featured":"toot:featured"`+
							`,`+
							`"featuredTags":"toot:featuredTags"`+
							`,`+
							`"focalPoint":"toot:focalPoint"`+
							`,`+
							`"indexable":"toot:indexable"`+
							`,`+
							`"memorial":"toot:memorial"`+
							`,`+
							`"suspended":"toot:suspended"`+

							`,`+

							`"alsoKnownAs":"as:alsoKnownAs"`+
							`,`+
							`"Hashtag":"as:Hashtag"`+
							`,`+
							`"movedTo":"as:movedTo"`+
						`}`+
					`]`+
					`,`+

					`"id":"urn:uuid:88c6a753-d6d5-45eb-bc18-9b1089e7b1f8"`+
					`,`+
					`"name":"Joe Blow"`+
					`,`+
					`"summary":"Hello world!"`+
					`,`+
					`"type":"Person"`+

					`,`+

					`"attributionDomains":""`+
					`,`+
					`"Emoji":""`+
					`,`+
					`"discoverable":true`+
					`,`+
					`"featured":true`+
					`,`+
					`"featuredTags":"#once #twice #thrice #fource"`+
					`,`+
					`"focalPoint":[12,100]`+
					`,`+
					`"indexable":true`+
					`,`+
					`"memorial":false`+
					`,`+
					`"suspended":false`+

					`,`+

					`"alsoKnownAs":[]`+
					`,`+
					`"Hashtag":["#fediverse"]`+
					`,`+
					`"movedTo":""`+

					`,`+

					`"signature":""`+

					`,`+

					`"colour":"yellow"`+
				`}`),
		},



		// 8
		{
			Values: []any{
				struct{
					ID      string `jsonld:"id"`
					Name    string `jsonld:"name"`
					Summary string `jsonld:"summary"`
					Type    string `jsonld:"type"`
				}{
					ID: "urn:uuid:88c6a753-d6d5-45eb-bc18-9b1089e7b1f8",
					Name: "Joe Blow",
					Summary: "Hello world!",
					Type: "Person",
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"http://joinmastodon.org/ns#"`
					Prefix    jsonld.Prefix    `jsonld:"toot"`

					AttributionDomains string `jsonld:"attributionDomains,omitempty"`
					Emoji              string `jsonld:",omitempty"`
					Discoverable       bool   `jsonld:"discoverable"`
					Featured           bool   `jsonld:"featured"`
					FeaturedTags       string `jsonld:"featuredTags,omitempty"`
					FocalPoint       []any    `jsonld:"focalPoint,omitempty"`
					Indexable          bool   `jsonld:"indexable"`
					Memorial           bool   `jsonld:"memorial"`
					Suspended          bool   `jsonld:"suspended"`
				}{
					Discoverable: true,
					Featured:     true,
					FeaturedTags: "#once #twice #thrice #fource",
					FocalPoint: []any{12,100},
					Indexable:    true,
					Memorial:     false,
					Suspended:    false,
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"https://www.w3.org/ns/activitystreams"`
					Prefix    jsonld.Prefix    `jsonld:"as"`

					AlsoKnownAs []string `jsonld:"alsoKnownAs,omitempty"`
					HashTag     []string `jsonld:"Hashtag,omitempty"`
					MovedTo       string `jsonld:"movedTo,omitempty"`
				}{
					HashTag: []string{
						"#fediverse",
					},
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/security/v1"`

					Signature string `jsonld:"signature,omitempty"`
				}{
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"https://banana.example/ns#"`

					Color string `jsonld:"colour"`
				}{
					Color: "yellow",
				},
			},
			Expected: []byte(
				`{`+
					`"@context":[`+
						`"https://w3id.org/security/v1"`+
						`,`+
						`"https://banana.example/ns#"`+
						`,`+
						`{`+
							`"toot":"http://joinmastodon.org/ns#"`+
							`,`+
							`"as":"https://www.w3.org/ns/activitystreams"`+

							`,`+

							`"attributionDomains":"toot:attributionDomains"`+
							`,`+
							`"Emoji":"toot:Emoji"`+
							`,`+
							`"discoverable":"toot:discoverable"`+
							`,`+
							`"featured":"toot:featured"`+
							`,`+
							`"featuredTags":"toot:featuredTags"`+
							`,`+
							`"focalPoint":"toot:focalPoint"`+
							`,`+
							`"indexable":"toot:indexable"`+
							`,`+
							`"memorial":"toot:memorial"`+
							`,`+
							`"suspended":"toot:suspended"`+

							`,`+

							`"alsoKnownAs":"as:alsoKnownAs"`+
							`,`+
							`"Hashtag":"as:Hashtag"`+
							`,`+
							`"movedTo":"as:movedTo"`+
						`}`+
					`]`+
					`,`+

					`"id":"urn:uuid:88c6a753-d6d5-45eb-bc18-9b1089e7b1f8"`+
					`,`+
					`"name":"Joe Blow"`+
					`,`+
					`"summary":"Hello world!"`+
					`,`+
					`"type":"Person"`+

					`,`+

					`"discoverable":true`+
					`,`+
					`"featured":true`+
					`,`+
					`"featuredTags":"#once #twice #thrice #fource"`+
					`,`+
					`"focalPoint":[12,100]`+
					`,`+
					`"indexable":true`+
					`,`+
					`"memorial":false`+
					`,`+
					`"suspended":false`+

					`,`+

					`"Hashtag":["#fediverse"]`+

					`,`+

					`"colour":"yellow"`+
				`}`),
		},








		// 9
		{
			Values: []any{
				struct{
					Apple  string `jsonld:"apple,omitempty"`
					Banana int    `jsonld:"banana"`
					Cherry string `jsonld:"cherry,omitempty"`
					date   int
				}{
					Apple:"",
					Banana:2,
					Cherry:"",
					date: 4,
				},
			},
			Expected: []byte(`{"banana":2}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonld.Marshal(test.Values...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual JSON-LD marshaled value is not what was expected.", testNumber)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				continue
			}
		}
	}
}
