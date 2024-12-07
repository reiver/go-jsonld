package jsonld

// NameSpace is used to specify a default JSON-LD name-space for a struct that can be marshaled to JSON-LD.
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
// This would result in a JSON-LD @context similar to:
//
//	"@context":{"fb":"http://fruitbasket.example/ns/", ...}
//
// And also, for example:
//
//	type YetAnotherType struct {
//		NS jsonld.NameSpace `jsonld:"http://count.example/ns/"`
//		
//		Once   uint `json:"once"`
//		Twice  uint `json:"twice"`
//		Thrice uint `json:"thrice"`
//		Fource uint `json:"fource"`
//	}
//
// This would result in a JSON-LD @context similar to:
//
//	"@context":["http://count.example/ns/", ...]
//
// One thing to notice is that — it does not matter what the name of the field of type jsonld.NameSpace is.
// In the first example it was "NameSpace". In the second example it was "JSONLDNameSpace". In the third example is was "NS".
// You coud name it anything you want (so long it is a valid Go field-name).
// The name isn't the important part — the important part is the tag.
type NameSpace struct{}

// This makes it so the JSON package used omits this, when used as a field on a struct or a value in a map, from the resulting JSON.
func (NameSpace) JSONOmitAlways() {}
