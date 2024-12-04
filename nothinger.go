package jsonld

// Nothinger interacts with 'omitempty'.
//
// If the field of a struct is of a type that fits Nothinger (and thus has a 'IsNothing()bool' method),
// and that field of the struct is tagged with omitempty', then jsonld.Marshal() will call its IsNothing
// method to determine whether it is empty or not.
//
// For example:
//
//	type MyType struct {
//		something bool
//	}
//	
//	func (receiver MyType) IsNothing bool {
//		return !receiver.something
//	}
//	
//	type MyStruct struct{
//		Apple  string `jsonld:"apple"`
//		Banana MyType `jsonld:"banana,omitempty"` // <----
//		Cherry bool   `jsonld:"cherry"`
//	}
//	
//	// ...
//	
//	var myStruct = MyStruct{
//		Cherry: true,
//	}
//	
//	bytes, err := jsonld.Marshal(myStruct)
//	
//	// bytes == []byte(`{"apple":"","cherry":true}`)
//	//
//	// Notice that 'cherry' was omitted from the result JSON-LD.
type Nothinger interface {
	IsNothing() bool
}
