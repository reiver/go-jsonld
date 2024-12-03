package jsonld

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilValue         = erorr.Error("jsonld: nil value")
	errNilReflectedType = erorr.Error("jsonld: nil reflected-value")
	errNotStruct        = erorr.Error("jsonld: not struct")
)
