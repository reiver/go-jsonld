package jsonld

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilFunc          = erorr.Error("jsonld: nil func")
	errNilReflectedType = erorr.Error("jsonld: nil reflected-value")
	errNilValue         = erorr.Error("jsonld: nil value")
	errNotStruct        = erorr.Error("jsonld: not struct")
)
