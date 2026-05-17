package jsonld

import (
	gobytes "bytes"
	"fmt"
	gourl "net/url"
	"strings"

	"encoding/hex"

	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-blanknode"
	"github.com/reiver/go-json"
	"github.com/reiver/go-opt"
	"golang.org/x/net/idna"
)

// ID is used as the value of the JSON-LD "@id" construct.
//
// ID is an optional-type (also sometimes called an option-type or maybe-type).
//
// For example:
//
//	type Node struct {
//		ID jsonld.ID `json:"@id"`
//	}
//
// See also:
//
//	• [ID.GoString]
//	• [ID.IsEmpty]
//	• [ID.MarshalJSON]
//	• [ID.UnmarshalJSON]
//	• [NoID]
//	• [SomeID]
type ID struct {
	optional opt.Optional[string]
}

// NoID returns a "nothing" [ID].
//
// "nothing" in the sense of optional-type (also sometimes called an option-type or maybe-type).
//
// For example:
//
//	id := jsonld.NoID()
//
// Although, more typically, it would more likely be used in code similar to:
//
//	type Node struct {
//		ID jsonld.ID `json:"@id"`
//	}
//	
//	var node Node = Node{
//		ID: jsonld.NoID(),
//	}
//
// See also:
//
//	• [ID]
//	• [SomeID]
func NoID() ID {
	return ID{}
}

// SomeID returns a "something" [ID].
//
// "something" in the sense of optional-type (also sometimes called an option-type or maybe-type).
//
// For example:
//
//	id := jsonld.SomeID("http://example.com/banana")
//
// Although, more typically, it would more likely be used in code similar to:
//
//	type Node struct {
//		ID jsonld.ID `json:"@id"`
//	}
//	
//	var node Node = Node{
//		ID: jsonld.SomeID("http://example.com/banana"),
//	}
//
// Note that SomeID will normalze the value stored in [ID].
// So, for example, this:
//
//	id := jsonld.SomeID("HTTPS://EXAMPLE.COM/banana")
//
// Would get stored as:
//
//	"https://example.com/banana"
//
// (Also, note that Blank Node Identifiers and their Blank Node Labels will be left as with respect to normalization. I.e., they won't be changed.)
//
// See also:
//
//	• [ID]
//	• [NoID]
func SomeID(id string) ID {
	if !strings.Contains(id, ":") && !strings.HasPrefix(id, "//") {
		return ID{
			optional: opt.Something(id),
		}
	}
	if blanknode.HasIdentifierPrefixString(id) {
		return ID{
			optional: opt.Something(id),
		}
	}

	url, err := gourl.Parse(id)
	if nil != err {
		return ID{
			optional: opt.Something(id),
		}
	}

	url.Scheme = strings.ToLower(url.Scheme)
	url.Host   = strings.ToLower(url.Host)

	{
		host, err := idna.ToUnicode(url.Host)
		if nil == err {
			url.Host = host
		}
	}

	{
		var buffer [256]byte
		var p []byte = buffer[0:0]

		if "" != url.Scheme {
			p = append(p, url.Scheme...)
			p = append(p, ':')
		}

		switch {
		case "" != url.Opaque:
			p = append(p, url.Opaque...)

			if "" != url.RawQuery {
				p = append(p, '?')
				p = append(p, percentDecodeUnreserved(url.RawQuery)...)
			}

			if "" != url.Fragment {
				p = append(p, '#')
				p = append(p, percentDecodeUnreserved(url.EscapedFragment())...)
			}
		default:
			p = append(p, "//"...)

			if nil != url.User {
				p = append(p, url.User.String()...)
				p = append(p, '@')
			}

			p = append(p, url.Host...)

			p = append(p, percentDecodeUnreserved(url.EscapedPath())...)

			if "" != url.RawQuery {
				p = append(p, '?')
				p = append(p, percentDecodeUnreserved(url.RawQuery)...)
			}

			if "" != url.Fragment {
				p = append(p, '#')
				p = append(p, percentDecodeUnreserved(url.EscapedFragment())...)
			}
		}

		{
			// This comparison is a Go idiom.
			// The Go compiler optimizes these comparisons:
			//
			//	string(byteSlice) == existingString
			//
			// The Go compler compares the bytes directly and avoids allocating a temporary string
			if string(p) == id {
				return ID{
					optional: opt.Something(id),
				}
			}

			return ID{
				optional: opt.Something(string(p)),
			}
		}
	}
}

// Get is used to get the identifier inside of [ID] if there is anything inside of it.
//
// If there is "nothing" inside of the [ID], then it will return false for the second return value.
//
// If there is "something" inside of the [ID], then it wll return true for the second return value, and the first return value will be a string with the value of what is inside of it.
//
// For example:
//
//	value, found := id.Get()
//
// See also:
//
//	• [ID]
func (receiver ID) Get() (string, bool) {
	return receiver.optional.Get()
}

// GoString returns an [ID] as a string that represents Go code.
//
// Typically, this would be used implicitly with fmt.Printf, fmt.Sprintf, fmt.Fprintf, etc.
// For example:
//
//	type Node struct {
//		ID jsonld.ID `json:"@id"`
//	}
//	
//	var node Node = Node{
//		ID: jsonld.SomeID("http://example.com/banana"),
//	}
//	
//	fmt.Printf("@id = %#v\n", node.ID)
//
// (Note the "%#v" verb used in the example code.)
//
// See also:
//
//	• [ID]
func (receiver ID) GoString() string {
	value, found := receiver.optional.Get()
	if !found {
		return "jsonld.NoID()"
	}

	return fmt.Sprintf("jsonld.SomeID(%q)", value)
}

// IsEmpty returns whether the [ID] is empty or not.
//
// [ID] is an optional-type (also sometimes called an option-type or maybe-type).
// An [ID] is empty if it is "nothing", and it not empty when it is "something".
// "nothing" and "something" are in the sense of optional-type.
//
// For example:
//
//	if id.IsEmpty() {
//		// ...
//	}
//
// See also:
//
//	• [ID]
//	• [NoID]
//	• [SomeID]
func (receiver ID) IsEmpty() bool {
	_, found := receiver.optional.Get()
	return !found
}

// MarshalJSON json-marshals an [ID] into raw JSON as []byte.
//
// Typically, this would be used implicitly with json.Marshal.
// For example:
//
//	type Node struct {
//		ID jsonld.ID `json:"@id"`
//	}
//	
//	var node Node = Node {
//		ID: jsonld.SomeID("http://example.com/banana"),
//	}
//	
//	data, err := json.Marshal(node)
//
// See also:
//
//	• [ID]
//	• [ID.UnmarshalJSON]
func (receiver ID) MarshalJSON() ([]byte, error) {
	value, found := receiver.optional.Get()
	if !found {
		return jsonNULL, nil
	}

	return json.MarshalString(value), nil
}

// UnmarshalJSON json-unmarshals raw JSON as []byte into an [ID].
//
// Typically, this would be used implicitly with json.Unmarshal.
// For example:
//
//	type Node struct {
//		ID jsonld.ID `json:"@id"`
//	}
//	
//	var node Node
//	
//	err := json.Unmarshal(data, &node)
//
// See also:
//
//	• [ID]
//	• [ID.UnmarshalJSON]
func (receiver *ID) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrNilReceiver
	}

	if nil == bytes {
		var err error = ErrJSONUnmarshalFailure

		return erorr.Wrap(err, "cannot json-unmarshal a nil []byte into a jsonld.ID")
	}
	if len(bytes) <= 0 {
		var err error = ErrJSONUnmarshalFailure

		return erorr.Wrap(err, "cannot json-unmarshal an empty []byte into a jsonld.ID")
	}

	switch {
	case gobytes.Equal(jsonNULL, bytes):
		*receiver = NoID()
		return nil
	case '"' == bytes[0]:
		var target string

		e := json.Unmarshal(bytes, &target)
		if nil != e {
			var err error = erorr.Errors{ErrJSONUnmarshalFailure, e}

			return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.ID",
				field.String("value", string(bytes)),
			)
		}

		*receiver = SomeID(target)
		return nil
	default:
		var err error = ErrJSONUnmarshalFailure

		return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.ID",
			field.String("value", string(bytes)),
		)
	}
}

// percentDecodeUnreserved percent-decodes / url-decodes return the selectively decoded form of a string.
// (The parameter is assumed to be a path from a URL/URI/IRI.)
//
// percentDecodeUnreserved selectively decodes percent-encoded / url-encoded sequences in a URL path.
// It decodes only unreserved characters.
// I.e., anything with a byte value greater-than or equal-to 0x80 for UTF-8 multibyte sequences, and unreserved ASCII: digits ('0'-'9') letters ('A'-'Z', 'a'-'z'), '-', '.', '_', '~'.
//
// Reserved characters such as %2F ('/'), %3F ('?'), %23 ('#'), %25 ('%') are left encoded.
func percentDecodeUnreserved(s string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	loop: for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '%' && i+2 < len(s):
			var pair [2]byte
			pair[0] = s[i+1]
			pair[1] = s[i+2]

			var oneByteBuffer [1]byte
			var oneByte []byte = oneByteBuffer[:]

			_, err := hex.Decode(oneByte, pair[:])
			if nil != err {
				p = append(p, s[i])
				continue loop
			}

			var b byte = oneByte[0]

			switch {
			case isUnreserved(b):
				p = append(p, b)
				i += len(pair)
			default:
				p = append(p, s[i], s[i+1], s[i+2])
				i += len(pair)
			}
		default:
			p = append(p, s[i])
		}
	}

	return string(p)
}

func isUnreserved(b byte) bool {
	switch {
	case 'A' <= b && b <= 'Z':
		return true
	case 'a' <= b && b <= 'z':
		return true
	case '0' <= b && b <= '9':
		return true
	case '-' == b, '.' == b, '_' == b, '~' == b:
		return true
	case 0x80 <= b:
		return true
	default:
		return false
	}
}
