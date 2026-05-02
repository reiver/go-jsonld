package jsonld

import (
	gobytes "bytes"

	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
	"github.com/reiver/go-nul"
)

// Direction is used as the value of the JSON-LD "@direction" construct.
type Direction struct {
	value nul.Nullable[string]
}

func DirectionNull() Direction {
	return Direction{
		value: nul.Null[string](),
	}
}

func DirectionLTR() Direction {
	return Direction{
		value: nul.Something("ltr"),
	}
}

func DirectionRTL() Direction {
	return Direction{
		value: nul.Something("rtl"),
	}
}

func (receiver Direction) MarshalJSON() ([]byte, error) {
	switch receiver {
	case DirectionNull():
		return jsonNULL, nil
	case DirectionLTR():
		return jsonStringLTR, nil
	case DirectionRTL():
		return jsonStringRTL, nil
	default:
		return nil, errNothing
	}
}

func (receiver *Direction) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrNilReceiver
	}

	switch {
	case gobytes.Equal(bytes, jsonNULL):
		*receiver = DirectionNull()
		return nil
	case gobytes.Equal(bytes, jsonStringLTR):
		*receiver = DirectionLTR()
		return nil
	case gobytes.Equal(bytes, jsonStringRTL):
		*receiver = DirectionRTL()
		return nil
	case nil == bytes:
		var err error = ErrJSONUnmarshalFailure

		return erorr.Wrap(err, "cannot json-unmarshal a nil []byte into a jsonld.Direction")
	default:
		var err error = ErrJSONUnmarshalFailure

		return erorr.Wrap(err, "cannot json-unmarshal into a jsonld.Direction",
			field.String("value", string(bytes)),
		)
	}
}
