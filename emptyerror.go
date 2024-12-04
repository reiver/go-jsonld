package jsonld

// Empty is an jsonld.EmptyError error.
//
// It exists as a convenience, so that one does not have to create their own type that fits jsonld.EmptyError.
//
// One would return this from a custom .MarshalJSONLD() method to trigger 'omitempty', if it exists.
var Empty error

func init() {
	Empty  = internalEmptyError{}
}

// If a custom .MarshalJSONLD() method returns an error that fits jsonld.EmptyError, and the struct has 'omitempty', then that field in the struct is omitted.
type EmptyError interface {
	error
	EmptyError()
}

type internalEmptyError struct{}

var _ error = internalEmptyError{}
var _ EmptyError = internalEmptyError{}

func (internalEmptyError) Error() string {
	return "empty"
}

func (internalEmptyError) EmptyError() {
	// nothing here
}
