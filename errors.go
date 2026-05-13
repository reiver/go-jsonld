package jsonld

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrInvalidTypeValue     = erorr.Error("json-ld invalid type value")
	ErrNilReceiver          = erorr.Error("nil receiver")
	ErrJSONUnmarshalFailure = erorr.Error("json-unmarshal failure")
)

const (
	errNilReflectedType = erorr.Error("jsonld: nil reflected-value")
	errNilValue         = erorr.Error("jsonld: nil value")
	errNothing          = erorr.Error("jsonld: nothing")
	errNotStruct        = erorr.Error("jsonld: not struct")
)
