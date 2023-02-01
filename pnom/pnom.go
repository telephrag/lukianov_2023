package pnom

import (
	"fmt"
	"lukianov_2023/frac"
)

type Polynomial struct {
	coefs []*frac.Frac
	Err   error
}

func New(coefs []*frac.Frac) *Polynomial {
	p := &Polynomial{
		coefs: make([]*frac.Frac, len(coefs)),
	}
	copy(p.coefs, coefs) // `copy()` makes a deep copy by default
	return p
}

func (p *Polynomial) Copy() (other *Polynomial) {
	copiedCoefs := make([]*frac.Frac, len(p.coefs))
	copy(copiedCoefs, p.coefs) // `copy()` makes a deep copy by default
	return &Polynomial{coefs: copiedCoefs}
}

func (p *Polynomial) Neg() (self *Polynomial) {
	for i := range p.coefs {
		p.coefs[i].Neg()
	}
	return p
}

func (p *Polynomial) Len() int {
	tail := 0
	for i := len(p.coefs) - 1; i >= 0; i-- {
		if p.coefs[i].Equals(frac.NULL) {
			tail++
		} else {
			break
		}
	}
	return len(p.coefs) - tail
}

func (p *Polynomial) Equals(other *Polynomial) bool {
	if p.Len() != other.Len() {
		return false
	}

	for i := range p.coefs {
		if !p.coefs[i].Equals(other.coefs[i]) {
			return false
		}
	}
	return true
}

func (p *Polynomial) Add(other *Polynomial) (self *Polynomial) {
	if p.Err != nil {
		return p
	}

	l := len(p.coefs)
	m := len(other.coefs)

	if m > l {
		add := make([]*frac.Frac, m-l)
		for i := range add {
			add[i] = frac.NULL.Copy()
		}
		p.coefs = append(p.coefs, add...)
		l = m
	}

	for i := 0; i < l; i++ {
		if err := p.coefs[i].Add(other.coefs[i]).Err; err != nil {
			p.Err = fmt.Errorf("Add(), %d, %w", i, err)
			return p
		}
	}

	return p
}
