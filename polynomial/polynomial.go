package polynomial

import (
	"fmt"
	"lukianov_2023/ratfrac"
)

type Polynomial struct {
	coefs []*ratfrac.RatFrac
	Err   error
}

func New(coefs []*ratfrac.RatFrac) *Polynomial {
	return &Polynomial{coefs: coefs}
}

func (p *Polynomial) Copy() (other *Polynomial) {
	copiedCoefs := make([]*ratfrac.RatFrac, len(p.coefs))
	copy(copiedCoefs, p.coefs) // `copy()` makes a deep copy by default
	return &Polynomial{coefs: copiedCoefs}
}

func (p *Polynomial) Neg() (self *Polynomial) {
	for i := range p.coefs {
		p.coefs[i].Neg()
	}
	return p
}

func (p *Polynomial) Add(other *Polynomial) (self *Polynomial) {
	if p.Err != nil {
		return p
	}

	l := len(p.coefs)
	m := len(other.coefs)

	if m > l {
		add := make([]*ratfrac.RatFrac, m-l)
		p.coefs = append(p.coefs, add...)
	}

	for i := 0; i < l; i++ {
		if err := p.coefs[i].Add(other.coefs[i]).Err; err != nil {
			p.Err = fmt.Errorf("Add(), %d, %w", i, err)
			return p
		}
	}

	return p
}
