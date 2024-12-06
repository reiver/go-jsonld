package jsonld

import (
	"github.com/reiver/go-json"
)

func isSimpleType(value any) bool {
	if nil == value {
		return true
	}

	switch value.(type) {
	case json.Marshaler:
		return true
	case
		bool,
		int,int8,int16,int32,int64,
		string,
		uint,uint8,uint16,uint32,uint64:
		return true

	}

	return false
}

func marshalSimpleType(value any) ([]byte, error) {
	if nil == value {
		return []byte(`null`), nil
	}

	return json.Marshal(value)
}
