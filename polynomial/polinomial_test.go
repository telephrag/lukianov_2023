package polynomial

import (
	"lukianov_2023/ratfrac"
	"testing"
)

func TestSample(t *testing.T) {
	coefs := []*ratfrac.RatFrac{
		ratfrac.New(1, 2, ratfrac.SIGN_POS),
		ratfrac.New(2, 7, ratfrac.SIGN_NEG),
		ratfrac.New(0, 1, ratfrac.SIGN_POS),
		ratfrac.New(7, 4, ratfrac.SIGN_POS),
	}

	p := New(coefs)
	t.Log(p)
}
