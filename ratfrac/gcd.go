package ratfrac

import (
	"math"
	"math/bits"
)

func gcd(x, y uint64) uint64 {
	xn, yn := (x == 0), (y == 0)

	if xn && !yn {
		return y
	}
	if !xn && yn {
		return x
	}
	if xn && yn {
		return math.MaxUint64
	}

	if x == 1 || y == 1 {
		return 1
	}

	xt := bits.TrailingZeros64(x)
	x >>= xt
	yt := bits.TrailingZeros64(y)
	y >>= yt
	var k int
	if xt < yt {
		k = xt
	} else {
		k = yt
	}

	for {
		if x < y {
			x, y = y, x
		}
		x -= y

		if x == 0 {
			return y << k
		}

		x >>= bits.TrailingZeros(uint(x))
	}
}
