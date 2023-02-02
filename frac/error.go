package frac

import (
	"errors"
)

var (
	ErrOverflow          = errors.New("overflow")
	ErrNullDiv           = errors.New("division by null")
	ErrNullIsNonNegative = errors.New("null must be non negative")
)
