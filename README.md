# go-jsonld

Package **jsonld** provides JSON-LD encoders and decoders, for the Go programming language.

And in particular, handles the way the Fediverse, ActivityPub, and ActivityStreams uses JSON-LD.

And import thing to understand is — you use separate Go `struct`s to represent each JSON-LD namespace.
Both for _marshaling_ and _unmarshaing_.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-jsonld

[![GoDoc](https://godoc.org/github.com/reiver/go-jsonld?status.svg)](https://godoc.org/github.com/reiver/go-jsonld)

## Example

Here is a simple example with only one JSON-LD namespace:

```golang
import "github.com/reiver/go-jsonld"

// ...

// Note that "jsonld" struct-tags are used on the first 2 fields,
// and "json" struct-tags are used on the rest of the fields.
type MyStruct struct {
	NameSpace jsonld.NameSpace `jsonld:"http://example.com/ns"`
	Prefix    jsonld.Prefix    `jsonld:"ex"`

	Once   bool   `json:"once"`
	Twice  int    `json:"twice"`
	Thrice string `json:"thrice,omitempty"`
	Fource uint   `json:"fource"`
}

// ...

var value MyStruct // = ...

bytes, err := jsonld.Marshal(value)
```

Here is a more typical example with multiple JSON-LD namespaces.

```golang
import "github.com/reiver/go-jsonld"

// ...

type Person struct {
	NameSpace jsonld.NameSpace `jsonld:"http://ns.example/person"`
	Prefix    jsonld.Prefix    `jsonld:"person"`

	GivenName         string `json:"given-name,omitempty"`
	AdditionalNames []string `json:"additional-names,omitempty"`
	FamilyName        string `json:"family-name,omitempty"`
}

type Programmer struct {
	NameSpace jsonld.NameSpace `jsonld:"http://example.com/programmer"`
	Prefix    jsonld.Prefix    `jsonld:"programmer"`

	ProgrammingLanguage string `json:"programming-language,omitempty"`
}

// ...

var person Person // = ...
var programmer Programmer // = ...

bytes, err := jsonld.Marshal(person, programmer)
```

## Import

To import package **jsonld** use `import` code like the follownig:
```
import "github.com/reiver/go-jsonld"
```

## Installation

To install package **jsonld** do the following:
```
GOPROXY=direct go get github.com/reiver/go-jsonld
```

## Author

Package **jsonld** was written by [Charles Iliya Krempeaux](http://reiver.link)
