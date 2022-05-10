package common

import "errors"

var (
	ErrNotSupported     = errors.New("not supported")
	ErrReturnedLen0     = errors.New("returned length 0")
	ErrIdReturnedWrong  = errors.New("connection id returned dont match with sent")
	ErrConnectionClosed = errors.New("connection is already closed")
	ErrEmptyData        = errors.New("data to buy is empty")
)
