package ratfrac

import (
	"math/bits"
)

// Methods in this file merely implicitly set internal `Err` value in case
// of an error while wrapping functions from `bits` package.
// In case of an error, return value equals to the first argument.
// This is done to keep original value inside `rf` the same
// instead of corrupting it with invalid value.

func (rf *RatFrac) add(x, y uint64) uint64 {
	res, carry := bits.Add64(x, y, 0)
	if carry != 0 {
		rf.Err = ErrOverflow
		return x
	}
	return res
}

func (rf *RatFrac) mul(x, y uint64) uint64 {
	hi, lo := bits.Mul64(x, y)
	if hi != 0 {
		rf.Err = ErrOverflow
		return x
	}
	return lo
}
