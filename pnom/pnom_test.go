package pnom

import (
	"fmt"
	"lukianov_2023/frac"
	"math/big"
	"testing"
)

func must(t *testing.T, location string, should, is *Pnom) {
	if !is.Equals(should) {
		t.Fatalf("%-16s expected: %-8s got: %-8s\n",
			fmt.Sprintf("[%s]", location),
			should,
			is,
		)
	}
}

func TestRemoveTrailingNulls(t *testing.T) {
	p := New([]*frac.Frac{
		frac.New(1, 1, frac.NEG),
		frac.NULL(),
		frac.New(1, 1, frac.POS),
		frac.NULL(),
		frac.NULL(),
		frac.NULL(),
		frac.NULL(),
	}).RemoveTrailingNulls()

	if p.len() != 3 {
		t.Fatalf("expected p.Len() to be 3; got: %d\n", p.len())
	}
}

func TestEquals(t *testing.T) {
	p0 := New([]*frac.Frac{
		frac.New(1, 2, frac.POS),
		frac.New(2, 7, frac.NEG),
		frac.New(0, 1, frac.POS),
		frac.New(7, 4, frac.POS),
	})
	p1 := New([]*frac.Frac{
		frac.New(1, 2, frac.POS),
		frac.New(2, 7, frac.NEG),
		frac.New(0, 1, frac.POS),
		frac.New(7, 4, frac.POS),
	})
	if !p0.Equals(p1) {
		t.Fatalf("%-16s expected \"%s\" to be equal to \"%s\"\n",
			"[equals]", p0, p1,
		)
	}

	p1 = New([]*frac.Frac{
		frac.New(1, 2, frac.POS),
		frac.New(2, 7, frac.NEG),
		frac.New(0, 1, frac.POS),
		frac.New(7, 4, frac.POS),
		frac.NULL(),
		frac.New(1, 9, frac.NEG),
		frac.NULL(),
	})
	if p0.Equals(p1) {
		t.Fatalf("%-16s expected \"%s\" to be unequal to \"%s\"\n",
			"[equalsWithTail]", p0, p1,
		)
	}
}

func TestAdd(t *testing.T) {
	coefs0 := []*frac.Frac{
		frac.New(1, 2, frac.POS),
		frac.New(2, 7, frac.NEG),
		frac.New(0, 1, frac.POS),
		frac.New(7, 4, frac.POS),
	}
	p0 := New(coefs0)

	coefs1 := []*frac.Frac{
		frac.New(1, 2, frac.POS),
		frac.New(2, 7, frac.NEG),
		frac.New(0, 1, frac.POS),
		frac.New(7, 4, frac.POS),
	}
	p1 := New(coefs1)

	coefsShould := []*frac.Frac{
		frac.New(1, 1, frac.POS),
		frac.New(4, 7, frac.NEG),
		frac.New(0, 1, frac.POS),
		frac.New(7, 2, frac.POS),
	}
	pShould := New(coefsShould)

	must(t, "simpleAdd", p0.Add(p1), pShould)

	tailedAddShould := New([]*frac.Frac{
		frac.New(1, 2, frac.POS),
		frac.New(2, 7, frac.NEG),
		frac.New(0, 1, frac.POS),
		frac.New(7, 4, frac.POS),
		frac.NULL(),
		frac.New(1, 9, frac.NEG),
		frac.NULL(),
	})
	tailedAdd := p1.Add(New([]*frac.Frac{
		frac.NULL(),
		frac.NULL(),
		frac.NULL(),
		frac.NULL(),
		frac.NULL(),
		frac.New(1, 9, frac.NEG),
		frac.NULL(),
	}))
	must(t, "tailedAdd", tailedAddShould, tailedAdd)
}

func TestSub(t *testing.T) {
	coefsToNull := []*frac.Frac{
		frac.New(1, 2, frac.POS),
		frac.New(2, 7, frac.NEG),
		frac.New(0, 1, frac.POS),
		frac.New(7, 4, frac.POS),
	}
	toNull := New(coefsToNull).Sub(
		New(coefsToNull),
	)
	toNullShould := New([]*frac.Frac{})
	must(t, "toNull", toNullShould, toNull)

	complexLeft := New([]*frac.Frac{
		frac.New(1, 2, frac.POS),
		frac.New(2, 7, frac.NEG),
		frac.New(0, 1, frac.POS),
		frac.New(7, 4, frac.POS),
	})
	complexRigth := New([]*frac.Frac{
		frac.New(1, 2, frac.NEG),
		frac.New(0, 7, frac.POS),
		frac.New(3, 1, frac.NEG),
		frac.New(7, 4, frac.POS),
		frac.New(0, 1, frac.POS),
		frac.NULL(),
		frac.New(1, 1, frac.NEG),
		frac.NULL(),
		frac.NULL(),
	})
	complexShould := New([]*frac.Frac{
		frac.New(1, 1, frac.POS),
		frac.New(2, 7, frac.NEG),
		frac.New(3, 1, frac.POS),
		frac.NULL(),
		frac.New(0, 28, frac.POS),
		frac.NULL(),
		frac.New(1, 1, frac.POS),
		frac.NULL(),
	})
	must(t, "complex", complexShould, complexLeft.Sub(complexRigth))
}

func TestMul(t *testing.T) {
	sampleLeft := New([]*frac.Frac{
		frac.New(2, 1, frac.NEG),
		frac.New(3, 7, frac.NEG),
		frac.New(1, 2, frac.POS),
		frac.New(1, 1, frac.POS),
		frac.NULL(),
	})
	sampleRight := New([]*frac.Frac{
		frac.New(1, 1, frac.NEG),
		frac.New(3, 1, frac.POS),
		frac.New(2, 1, frac.NEG),
		frac.New(1, 1, frac.NEG),
		frac.NULL(),
		frac.NULL(),
		frac.NULL(),
	})
	sampleShould := New([]*frac.Frac{
		frac.New(2, 1, frac.POS),
		frac.New(39, 7, frac.NEG),
		frac.New(31, 14, frac.POS),
		frac.New(47, 14, frac.POS),
		frac.New(17, 7, frac.POS),
		frac.New(5, 2, frac.NEG),
		frac.New(1, 1, frac.NEG),
	})
	must(t, "sample", sampleShould, sampleLeft.Mul(sampleRight))
}

func TestMulByK(t *testing.T) {
	sample := New([]*frac.Frac{
		frac.NULL(),
		frac.New(1, 2, frac.POS),
		frac.New(3, 8, frac.NEG),
		frac.NULL(),
	})
	sampleShould := New([]*frac.Frac{
		frac.NULL(),
		frac.New(2, 2, frac.POS),
		frac.New(3, 4, frac.NEG),
	})
	must(t, "sample", sampleShould, sample.MullByK(frac.New(2, 1, frac.POS)))
}

func TestCalc(t *testing.T) {
	res, err := New([]*frac.Frac{
		frac.New(1, 2, frac.POS),
		frac.New(7, 1, frac.NEG),
		frac.New(2, 4, frac.POS),
		frac.New(9, 9, frac.POS),
	}).Calc(1.6)

	if err != nil {
		t.Fatalf("expected no error; got: %s\n", err)
	}

	rb := big.NewFloat(res).SetPrec(3)
	if rb.Cmp(big.NewFloat(-5.324).SetPrec(3)) != 0 {
		t.Fatalf("expected result to be: %f; got: %f\n", rb, -5.324)
	}
}
