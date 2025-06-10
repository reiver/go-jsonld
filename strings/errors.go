package strings

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilReceiver = erorr.Error("json: nil recevier")
	errEmptyData   = erorr.Error("json: empty data")
)
