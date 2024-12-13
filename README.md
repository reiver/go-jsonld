# go-jsonld

Package **jsonld** provides JSON-LD encoders and decoders, for the Go programming language.

And in particular, handles the way the Fediverse, ActivityPub, and ActivityStreams uses JSON-LD.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-jsonld

[![GoDoc](https://godoc.org/github.com/reiver/go-jsonld?status.svg)](https://godoc.org/github.com/reiver/go-jsonld)

## Example

Here is a simple example with only one namespace:

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
