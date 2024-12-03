package jsonld

// NameSpace is used to specify a default JSON-LD name-space for a struct that can be marshaled to JSON-LD.
//
// For example:
//
//	type MyType struct {
//		NameSpace jsonld.NameSpace `jsonld:"http://example.com/ns#"`
//		Prefix    jsonld.Prefix    `jsonld:"ex"`
//		
//		GivenNames string `jsonld:"givenNames"`
//		FamilyName string `jsonld:"familyName"`
//	}
//
// And, another example:
//
//	type MyOtherType struct {
//		JSONLDNameSpace jsonld.NameSpace `jsonld:"http://fruitbasket.example/ns/"`
//		JSONLDPrefix    jsonld.Prefix    `jsonld:"fb"`
//		
//		NumApples   uint `jsonld:"NumApples"`
//		NumBananas  uint `jsonld:"NumBananas"`
//		NumCherries uint `jsonld:"NumCherries"`
//	}
//
// One thing to notice is that — it does not matter what the name of the field of type jsonld.NameSpace is.
// In the first example it was "NameSpace". In the second example it was "JSONLDNameSpace".
// You coud name it anything you want (so long it is a valid Go field-name).
// The name isn't the important part — the important part is the tag.
type NameSpace struct{}
