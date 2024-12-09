package jsonld_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-jsonld"
)


func TestMarshal(t *testing.T) {

	type FruitBasket struct {
		Apple string                `json:"apple,omitempty"`
		Banana opt.Optional[string] `json:"banana,omitempty"`
		Cherry opt.Optional[int64]  `json:"cherry,omitempty"`
	}

	type TestStruct1 struct{
		one string
		two int

		Three string                    `json:"three,omitempty"`
		Four  opt.Optional[string]      `json:"four,omitempty"`
		Five  opt.Optional[FruitBasket] `json:"five,omitempty"`
	}

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
					Apple string `json:"apple"`
					Banana int
					Cherry bool  `json:"cherry"`
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
					Apple string `json:"apple"`
					Banana int   `json:"banana"`
					Cherry bool  `json:"cherry"`
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
					Apple string `json:"apple"`
					Banana int
					Cherry bool  `json:"cherry"`
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
					Apple string `json:"apple"`
					Banana int
					Cherry bool  `json:"cherry"`
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
						`"Banana":"ex:Banana"`+
						`,`+
						`"apple":"ex:apple"`+
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

					AttributionDomains string `json:"attributionDomains"`
					Emoji              string
					Discoverable       bool   `json:"discoverable"`
					Featured           bool   `json:"featured"`
					FeaturedTags       string `json:"featuredTags"`
					FocalPoint       []any    `json:"focalPoint"`
					Indexable          bool   `json:"indexable"`
					Memorial           bool   `json:"memorial"`
					Suspended          bool   `json:"suspended"`
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
						`"Emoji":"toot:Emoji"`+
						`,`+
						`"attributionDomains":"toot:attributionDomains"`+
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
					ID      string `json:"id"`
					Name    string `json:"name"`
					Summary string `json:"summary"`
					Type    string `json:"type"`
				}{
					ID: "urn:uuid:88c6a753-d6d5-45eb-bc18-9b1089e7b1f8",
					Name: "Joe Blow",
					Summary: "Hello world!",
					Type: "Person",
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"http://joinmastodon.org/ns#"`
					Prefix    jsonld.Prefix    `jsonld:"toot"`

					AttributionDomains string `json:"attributionDomains"`
					Emoji              string
					Discoverable       bool   `json:"discoverable"`
					Featured           bool   `json:"featured"`
					FeaturedTags       string `json:"featuredTags"`
					FocalPoint       []any    `json:"focalPoint"`
					Indexable          bool   `json:"indexable"`
					Memorial           bool   `json:"memorial"`
					Suspended          bool   `json:"suspended"`
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

					AlsoKnownAs []string `json:"alsoKnownAs"`
					Hashtag     []string `json:"Hashtag"`
					MovedTo       string `json:"movedTo"`
				}{
					Hashtag: []string{
						"#fediverse",
					},
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/security/v1"`

					Signature string `json:"signature"`
				}{
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"https://banana.example/ns#"`

					Color string `json:"colour"`
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

							`"Emoji":"toot:Emoji"`+
							`,`+
							`"attributionDomains":"toot:attributionDomains"`+
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

							`"Hashtag":"as:Hashtag"`+
							`,`+
							`"alsoKnownAs":"as:alsoKnownAs"`+
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
					ID      string `json:"id"`
					Name    string `json:"name"`
					Summary string `json:"summary"`
					Type    string `json:"type"`
				}{
					ID: "urn:uuid:88c6a753-d6d5-45eb-bc18-9b1089e7b1f8",
					Name: "Joe Blow",
					Summary: "Hello world!",
					Type: "Person",
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"http://joinmastodon.org/ns#"`
					Prefix    jsonld.Prefix    `jsonld:"toot"`

					AttributionDomains string `json:"attributionDomains,omitempty"`
					Emoji              string `json:",omitempty"`
					Discoverable       bool   `json:"discoverable"`
					Featured           bool   `json:"featured"`
					FeaturedTags       string `json:"featuredTags,omitempty"`
					FocalPoint       []any    `json:"focalPoint,omitempty"`
					Indexable          bool   `json:"indexable"`
					Memorial           bool   `json:"memorial"`
					Suspended          bool   `json:"suspended"`
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

					AlsoKnownAs []string `json:"alsoKnownAs,omitempty"`
					HashTag     []string `json:"Hashtag,omitempty"`
					MovedTo       string `json:"movedTo,omitempty"`
				}{
					HashTag: []string{
						"#fediverse",
					},
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/security/v1"`

					Signature string `json:"signature,omitempty"`
				}{
				},
				struct{
					NameSpace jsonld.NameSpace `jsonld:"https://banana.example/ns#"`

					Color string `json:"colour"`
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

							`"Emoji":"toot:Emoji"`+
							`,`+
							`"attributionDomains":"toot:attributionDomains"`+
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

							`"Hashtag":"as:Hashtag"`+
							`,`+
							`"alsoKnownAs":"as:alsoKnownAs"`+
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
					Apple  string `json:"apple,omitempty"`
					Banana int    `json:"banana"`
					Cherry string `json:"cherry,omitempty"`
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








		// 10
		{
			Values: []any{
				TestStruct1{},
			},
			Expected: []byte(`{}`),
		},









		// 11
		{
			Values: []any{
				struct{
					NameSpace jsonld.NameSpace `jsonld:"http://ns.example/outer"`

					One string `json:"one,omitempty"`
					Two string `json:"two,omitempty"`
					Something struct{
						NameSpace jsonld.NameSpace `jsonld:"http://example.com/inner"`

						Three string `json:"three,omitempty,bare"`
						Four string `json:"four,omitempty,bare"`
					} `json:"something,omitempty"`
				}{
					One: "1",
					Something: struct{
						NameSpace jsonld.NameSpace `jsonld:"http://example.com/inner"`

						Three string `json:"three,omitempty,bare"`
						Four string `json:"four,omitempty,bare"`
					}{
						Three: "3",
					},
				},
			},
			Expected: []byte(`{"@context":["http://ns.example/outer","http://example.com/inner"],"one":"1","something":{"three":3}}`),
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
