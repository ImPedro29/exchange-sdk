package common

import "errors"

var (
	ErrNotSupported = errors.New("not supported")
	ErrReturnedLen0 = errors.New("returned length 0")
)
