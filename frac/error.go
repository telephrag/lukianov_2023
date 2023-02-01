package frac

import (
	"errors"
)

var (
	ErrOverflow = errors.New("overflow")
	ErrNullDiv  = errors.New("division by null")
)
