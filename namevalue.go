package jsonld

//NameValue represents a name-value pair (which is also somtimes called a key-value pair).
//
// It is used in jsonld.Context.
type NameValue struct {
	Name string
	Value any
}
