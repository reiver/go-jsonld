package jsonld

import (
	"github.com/reiver/go-jsonld/strings"
)

// Strings is a type that can be zero, one, or many strings.
//
// In JSON, all of these would be valid values for this type:
//
//	null
//
//	"hello"
//
//	["hello"]
//
//	{"once","twice"}
//
//	{"once","twice","thrice"}
//
//	{"once","twice","thrice","fource"}
type Strings = strings.Strings
