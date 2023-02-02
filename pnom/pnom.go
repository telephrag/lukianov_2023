package pnom

import (
	"lukianov_2023/frac"
)

type Pnom struct {
	coefs []*frac.Frac
	Err   error
}

func New(coefs []*frac.Frac) *Pnom {
	p := &Pnom{
		coefs: make([]*frac.Frac, len(coefs)),
	}
	copy(p.coefs, coefs) // `copy()` makes a deep copy by default
	return p
}

func (p *Pnom) Copy() (other *Pnom) {
	copiedCoefs := make([]*frac.Frac, len(p.coefs))
	copy(copiedCoefs, p.coefs) // `copy()` makes a deep copy by default
	return &Pnom{coefs: copiedCoefs}
}

func (p *Pnom) Len() int {
	tail := 0
	for i := len(p.coefs) - 1; i >= 0; i-- {
		if p.coefs[i].Equals(frac.NULL()) {
			tail++
		} else {
			break
		}
	}
	return len(p.coefs) - tail
}

func (p *Pnom) Equals(other *Pnom) bool {
	if p.Len() != other.Len() {
		return false
	}

	if p.Len() == 0 {
		return true
	}

	for i := range p.coefs {
		if !p.coefs[i].Equals(other.coefs[i]) {
			return false
		}
	}
	return true
}

func (p *Pnom) At(index int) *frac.Frac {
	if index < len(p.coefs) {
		return p.coefs[index]
	}
	return frac.NULL().Copy()
}
