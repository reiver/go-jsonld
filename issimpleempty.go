package jsonld

import (
	"reflect"
)

func isSimpleEmpty(value any) bool {
	switch casted := value.(type) {
	case Emptier:
		return casted.IsEmpty()
	case Nothinger:
		return casted.IsNothing()
	case bool:
		var empty bool
		return empty == casted
	case int:
		var empty int
		return empty == casted
	case int8:
		var empty int8
		return empty == casted
	case int16:
		var empty int16
		return empty == casted
	case int32:
		var empty int32
		return empty == casted
	case int64:
		var empty int64
		return empty == casted
	case string:
		var empty string
		return empty == casted
	case uint:
		var empty uint
		return empty == casted
	case uint8:
		var empty uint8
		return empty == casted
	case uint16:
		var empty uint16
		return empty == casted
	case uint32:
		var empty uint32
		return empty == casted
	case uint64:
		var empty uint64
		return empty == casted
	}

	{
		var reflectedValue reflect.Value = reflect.ValueOf(value)

		var kind reflect.Kind = reflectedValue.Kind()

		switch kind {
		case reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
			if reflectedValue.IsNil() {
				return true
			}
		}

		switch kind {
//@TODO: Should we also handle pointers to these things?
		case reflect.Array, reflect.Map, reflect.Slice:
			if reflectedValue.Len() <= 0 {
				return true
			}
		}
	}

	return false
}
