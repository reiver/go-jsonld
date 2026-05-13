package jsonld

import (
	gobytes "bytes"

	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-json"
	"github.com/reiver/go-opt"
	"github.com/reiver/go-pckstr"
)

// internalType represents a single (non-array) type.
//
// I.e.:
//
//	"@type":"BANANA"
//
//	"@type":{"@id":"BANANA"}
//
// It is used by [Types].
type internalType struct {
	optional opt.Optional[string]
}

func (receiver internalType) Get() (string, bool) {
	return receiver.optional.Get()
}

func (receiver *internalType) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrNilReceiver
	}

	if nil == bytes {
		var err error = ErrJSONUnmarshalFailure

		return erorr.Wrap(err, "cannot json-unmarshal a nil []byte into a jsonld.Types")
	}
	if len(bytes) <= 0 {
		var err error = ErrJSONUnmarshalFailure

		return erorr.Wrap(err, "cannot json-unmarshal an empty []byte into a jsonld.Types")
	}

	switch {
	case '"' == bytes[0]:
		var target string

		e := json.Unmarshal(bytes, &target)
		if nil != e {
			var err error = erorr.Errors{ErrJSONUnmarshalFailure, e}

			return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.Types",
				field.String("value", string(bytes)),
			)
		}

		receiver.optional = opt.Something(target)
		return nil
	case '{' == bytes[0]:
		var target struct {
			AtID string `json:"@id,omitempty"`
			ID   string `json:"id,omitempty"`
		}

		e := json.Unmarshal(bytes, &target)
		if nil != e {
			var err error = erorr.Errors{ErrJSONUnmarshalFailure, e}

			return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.Types",
				field.String("value", string(bytes)),
			)
		}

		var id string = target.AtID
		if "" == id {
			id = target.ID
		}
		if "" == id {
			var err error = erorr.Errors{ErrJSONUnmarshalFailure, ErrInvalidTypeValue}

			return erorr.Wrap(err, "cannot json-unmarshal JSON object with \"@id\" or \"id\" into jsonld.Types",
				field.String("value", string(bytes)),
			)
		}

		receiver.optional = opt.Something(id)
		return nil
	default:
		var err error = ErrJSONUnmarshalFailure

		return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.Types",
			field.String("value", string(bytes)),
		)
	}
}

// Types is used as the value of the JSON-LD "@type" construct.
type Types struct {
	values pckstr.PackedStrings
}

func NoType() Types {
	return Types{}
}

func SomeType(t string) Types {
	return SomeTypes(t)
}

func SomeTypes(tt ...string) Types {
	if len(tt) <= 0 {
		return NoType()
	}
	return Types{
		values: pckstr.SomeStrings(tt...),
	}
}

func (receiver Types) IsEmpty() bool {
	if NoType() == receiver {
		return true
	}
	if receiver.values.LenZero() {
		return true
	}

	return false
}

func (receiver Types) MarshalJSON() ([]byte, error) {
	strings := receiver.values.Strings()

	switch len(strings) {
	case 0:
		return jsonNULL, nil
	case 1:
		return json.MarshalString(strings[0]), nil
	default:
		return json.MarshalSlice(strings)
	}
}

func (receiver Types) Strings() []string {
	return receiver.values.Strings()
}

func (receiver *Types) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrNilReceiver
	}

	if nil == bytes {
		var err error = ErrJSONUnmarshalFailure

		return erorr.Wrap(err, "cannot json-unmarshal a nil []byte into a jsonld.Types")
	}
	if len(bytes) <= 0 {
		var err error = ErrJSONUnmarshalFailure

		return erorr.Wrap(err, "cannot json-unmarshal an empty []byte into a jsonld.Types")
	}

	switch {
	case gobytes.Equal(bytes, jsonArrayEmpty):
		*receiver = NoType()
		return nil
	case gobytes.Equal(bytes, jsonNULL):
		*receiver = NoType()
		return nil
	case '"' == bytes[0] || '{' == bytes[0]:
		var target internalType

		e := json.Unmarshal(bytes, &target)
		if nil != e {
			var err error = erorr.Errors{ErrJSONUnmarshalFailure, e}

			return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.Types",
				field.String("value", string(bytes)),
			)
		}

		t, found := target.Get()
		if !found {
			var err error = erorr.Errors{ErrJSONUnmarshalFailure, ErrInvalidTypeValue}

			return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.Types",
				field.String("value", string(bytes)),
			)
		}

		*receiver = SomeType(t)
		return nil
	case '[' == bytes[0]:
		var target []internalType

		e := json.Unmarshal(bytes, &target)
		if nil != e {
			var err error = erorr.Errors{ErrJSONUnmarshalFailure, e}

			return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.Types",
				field.String("value", string(bytes)),
			)
		}

		var types []string
		for _, t := range target {
			value, found := t.Get()
			if !found {
				var err error = erorr.Errors{ErrJSONUnmarshalFailure, ErrInvalidTypeValue}

				return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.Types",
					field.String("value", string(bytes)),
				)
			}
			types = append(types, value)
		}

		*receiver = SomeTypes(types...)
		return nil
	default:
		var err error = ErrJSONUnmarshalFailure

		return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.Types",
			field.String("value", string(bytes)),
		)
	}
}
