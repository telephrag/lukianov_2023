package ratfrac

import (
	"fmt"
	"testing"
)

func must(t *testing.T, location string, should, is *RatFrac) {
	if !is.Equals(should) {
		t.Fatalf("%-16s expected: %-8s got: %-8s\n",
			fmt.Sprintf("[%s]", location),
			should,
			is,
		)
	}
}

func TestSimplify(t *testing.T) {

	simplifyPrime := New(273, 487, SIGN_POS).Simplify() // denominator is prime, will do nothing
	must(t, "simplifyPrime", simplifyPrime, simplifyPrime)

	simplify := New(288, 486, SIGN_POS).Simplify()
	simplifyShould := New(16, 27, SIGN_POS)
	must(t, "simplify", simplifyShould, simplify)
}

func TestAdd(t *testing.T) {

	prime := New(273, 487, SIGN_POS)

	add := prime.Copy().Add(New(273, 487, SIGN_NEG)) // 273/487 - 273/487
	addShould := New(0, 1, SIGN_POS)
	must(t, "add", addShould, add)

	improper := add.Copy().
		Add(New(200, 487, SIGN_POS)).
		Add(New(300, 487, SIGN_POS)) // 0 + 500/487
	improperShould := New(500, 487, SIGN_POS)
	must(t, "improper", improperShould, improper)

	whole := improper.Copy().Add(New(13, 487, SIGN_NEG)) // 500/487 - 13/487 = 1
	wholeShould := New(1, 1, SIGN_POS)
	must(t, "whole", wholeShould, whole)
}

func TestSub(t *testing.T) {
	prime := New(273, 487, SIGN_POS)

	toNull := prime.Copy().Sub(New(273, 487, SIGN_POS)) // 273/487 - 273/487
	toNullShould := New(0, 1, SIGN_POS)
	must(t, "toNull", toNullShould, toNull)

	minusNull := New(500, 487, SIGN_POS).Sub(toNull) // 500/487 - 0
	minusNullShould := New(500, 487, SIGN_POS)
	must(t, "minusNull", minusNullShould, minusNull)

	whole := minusNull.Copy().Sub(New(13, 487, SIGN_POS)) // 500/487 - 13/487 = 1
	wholeShould := New(1, 1, SIGN_POS)
	must(t, "whole", wholeShould, whole)
}

func TestMul(t *testing.T) {
	some := New(16, 27, SIGN_POS).Mul(New(12, 5, SIGN_POS)) // 16/27 * 12/5 = 64/45
	someShould := New(64, 45, SIGN_POS)
	must(t, "some", someShould, some)

	signChange := some.Mul(New(1, 8, SIGN_NEG)) // 64/45 * -1/8 = -8/45
	signChangeShould := New(8, 45, SIGN_NEG)
	must(t, "signChange", signChangeShould, signChange)
}

func TestDiv(t *testing.T) {
	some := New(16, 27, SIGN_POS).Div(New(12, 5, SIGN_POS)) // 16/27 / 12/5 = 20/81
	someShould := New(20, 81, SIGN_POS)
	must(t, "some", someShould, some)

	signChange := some.Div(New(1, 8, SIGN_NEG)) // 20/81 / -1/8 = -160/81
	signChangeShould := New(160, 81, SIGN_NEG)
	must(t, "signChange", signChangeShould, signChange)
}
