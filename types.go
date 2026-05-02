package jsonld

import (
	gobytes "bytes"

	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-json"
	"github.com/reiver/go-pckstr"
)

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

		return erorr.Wrap(err, "cannot json-unmarshal a empty []byte into a jsonld.Types")
	}

	switch {
	case gobytes.Equal(bytes, jsonArrayEmpty):
		*receiver = NoType()
		return nil
	case gobytes.Equal(bytes, jsonNULL):
		*receiver = NoType()
		return nil
	case '"' == bytes[0]:
		var target string

		e := json.Unmarshal(bytes, &target)
		if nil != e {
			var err error = erorr.Errors{ErrJSONUnmarshalFailure, e}

			return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.Types",
				field.String("value", string(bytes)),
			)
		}

		*receiver = SomeType(target)
		return nil
	case '[' == bytes[0]:
		var target []string

		e := json.Unmarshal(bytes, &target)
		if nil != e {
			var err error = erorr.Errors{ErrJSONUnmarshalFailure, e}

			return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.Types",
				field.String("value", string(bytes)),
			)
		}

		*receiver = SomeTypes(target...)
		return nil
	default:
		var err error = ErrJSONUnmarshalFailure

		return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.Types",
			field.String("value", string(bytes)),
		)
	}
}
