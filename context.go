package jsonld

// Context represents a JSON-LD @context.
//
// Example:
//
//	context, err := jsonld.ContextOf(something)
type Context struct {
	NameSpace string
	Prefix string
	NameValues []NameValue
}
