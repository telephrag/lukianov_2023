package pnom

import (
	"fmt"
	"lukianov_2023/frac"
	"testing"
)

func must(t *testing.T, location string, should, is *Polynomial) {
	if !is.Equals(should) {
		t.Fatalf("%-16s expected: %-8s got: %-8s\n",
			fmt.Sprintf("[%s]", location),
			should,
			is,
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

	must(t, "sample", p0.Add(p1), pShould)

	tailedAddShould := New([]*frac.Frac{
		frac.New(1, 2, frac.POS),
		frac.New(2, 7, frac.NEG),
		frac.New(0, 1, frac.POS),
		frac.New(7, 4, frac.POS),
		frac.NULL,
		frac.New(1, 9, frac.NEG),
		frac.NULL,
	})
	tailedAdd := p1.Add(New([]*frac.Frac{
		frac.NULL,
		frac.NULL,
		frac.NULL,
		frac.NULL,
		frac.NULL,
		frac.New(1, 9, frac.NEG),
		frac.NULL,
	}))
	must(t, "tailedAdd", tailedAddShould, tailedAdd)
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
		frac.NULL,
		frac.New(1, 9, frac.NEG),
		frac.NULL,
	})
	if p0.Equals(p1) {
		t.Fatalf("%-16s expected \"%s\" to be unequal to \"%s\"\n",
			"[equalsWithTail]", p0, p1,
		)
	}
}
