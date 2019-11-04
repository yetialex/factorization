package evaluate

import "errors"

var (
	ErrConvert          = errors.New("convert error")
	ErrWrongNumber      = errors.New("number must be >= 2")
	ErrRequestCancelled = errors.New("request cancelled")
)
