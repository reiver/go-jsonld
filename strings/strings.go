package strings

import (
	"bytes"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-json"
)

type Strings struct {
	values []string
	something bool
}

var _ json.Marshaler = Strings{}
var _ json.Unmarshaler = &Strings{}

func Something(value string) Strings {
	return Strings{
		something:true,
		values: []string{value},
	}
}

func Somethings(values ...string) Strings {
	return Strings{
		something:true,
		values: values,
	}
}

func Nothing() Strings {
	return Strings{}
}

func (receiver Strings) IsEmpty() bool {
	return len(receiver.values) <= 0
}

func (receiver Strings) IsNothing() bool {
	return !receiver.something
}

func (receiver Strings) Get(index int) (string, bool) {
	var nada string

	if receiver.IsNothing() {
		return nada, false
	}
	if receiver.IsEmpty() {
		return nada, false
	}

	if receiver.Len() <= index {
		return nada, false
	}

	return receiver.values[index], true
}

func (receiver Strings) GetElse(index int, alternative string) string {
	value, found := receiver.Get(index)
	if !found {
		return alternative
	}
	return value
}

func (receiver Strings) Len() int {
	return len(receiver.values)
}

func (receiver Strings) MarshalJSON() ([]byte, error) {

	if receiver.IsNothing() {
		return null, nil
	}

	if receiver.IsEmpty() {
		return empty, nil
	}

	if 1 == len(receiver.values) {
		return json.Marshal(receiver.values[0])
	}

	return json.Marshal(receiver.values)
}

func (receiver *Strings) UnmarshalJSON (data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if len(data) <= 0 {
		return errEmptyData
	}

	var data0 byte = data[0]

	switch {
	case bytes.Equal(null, data):
		*receiver = Nothing()
		return nil
	case bytes.Equal(empty, data):
		*receiver = Somethings()
		return nil
	case '"' == data0:
		var str string
		err := json.Unmarshal(data, &str)
		if nil != err {
			return erorr.Errorf("json: problem JSON-unmarshaling supposed JSON-string: %w", err)
		}
		*receiver = Something(str)
		return nil
	case '[' == data0:
		var strs []string
		err := json.Unmarshal(data, &strs)
		if nil != err {
			return erorr.Errorf("json: problem JSON-unmarshaling supposed JSON-array of string: %w", err)
		}
		*receiver = Somethings(strs...)
		return nil
	default:
		return erorr.Errorf("json: cannot JSON-unmarshal the data into a %T", receiver.values)
	}
}
