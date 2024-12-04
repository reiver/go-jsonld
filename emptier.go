package jsonld

// Emptier interacts with 'omitempty'.
//
// If the field of a struct is of a type that fits Emptier (and thus has a 'IsEmpty()bool' method),
// and that field of the struct is tagged with omitempty', then jsonld.Marshal() will call its IsEmpty
// method to determine whether it is empty or not.
//
// For example:
//
//	type MyType struct {
//		notEmpty bool
//	}
//	
//	func (receiver MyType) IsEmpty bool {
//		return !receiver.notEmpty
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
type Emptier interface {
	IsEmpty() bool
}
