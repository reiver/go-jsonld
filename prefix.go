package jsonld

// Prefix is used to specify a default JSON-LD prefix for a struct that can be marshaled to JSON-LD.
//
// For example:
//
//	type MyType struct {
//		NameSpace jsonld.NameSpace `jsonld:"http://example.com/ns#"`
//		Prefix    jsonld.Prefix    `jsonld:"ex"`
//		
//		GivenNames string `json:"givenNames"`
//		FamilyName string `json:"familyName"`
//	}
//
// This would result in a JSON-LD @context similar to:
//
//	"@context":{"ex":"http://example.com/ns#", ...}
//
// And, another example:
//
//	type MyOtherType struct {
//		JSONLDNameSpace jsonld.NameSpace `jsonld:"http://fruitbasket.example/ns/"`
//		JSONLDPrefix    jsonld.Prefix    `jsonld:"fb"`
//		
//		NumApples   uint `json:"NumApples"`
//		NumBananas  uint `json:"NumBananas"`
//		NumCherries uint `json:"NumCherries"`
//	}
//
//
// This would result in a JSON-LD @context similar to:
//
//	"@context":{"fb":"http://fruitbasket.example/ns/", ...}
//
// One thing to notice is that — it does not matter what the name of the field of type jsonld.Prefix is.
// In the first example it was "NameSpace". In the second example it was "JSONLDNameSpace".
// You coud name it anything you want (so long it is a valid Go field-name).
// The name isn't the important part — the important part is the tag.
type Prefix struct{}

// This makes it so the JSON package used omits this, when used as a field on a struct or a value in a map, from the resulting JSON.
func (Prefix) JSONOmitAlways() {}
