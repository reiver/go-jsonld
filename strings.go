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
//	[]
//
//	["hello"]
//
//	{"once","twice"}
//
//	{"once","twice","thrice"}
//
//	{"once","twice","thrice","fource"}
type Strings = strings.Strings

// NoStrings returns a [Strings] with no value.
func NoStrings() Strings {
	return strings.Nothing()
}

// SomeString returns a [Strings] with some value, and in paticular containing a single string.
func SomeString(value string) Strings {
	return strings.Something(value)
}

// SomeString returns a [Strings] with some value, and in paticular containing zero, one, two, three, ..., strings..
func SomeStrings(values ...string) Strings {
	return strings.Somethings(values...)
}
