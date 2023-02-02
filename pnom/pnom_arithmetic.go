package pnom

import (
	"fmt"
	"lukianov_2023/frac"
)

func (p *Pnom) Neg() (self *Pnom) {
	for i := range p.coefs {
		p.coefs[i].Neg()
	}
	return p
}

func (p *Pnom) Add(other *Pnom) (self *Pnom) {
	if p.Err != nil || other.Err != nil {
		return p
	}

	l := len(p.coefs)
	m := len(other.coefs)

	if m > l {
		add := make([]*frac.Frac, m-l)
		for i := range add {
			add[i] = frac.NULL().Copy()
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

	// remove "tail" of null elements
	tail := 0
	for i := p.Len() - 1; i >= 0; i-- {
		if p.coefs[i].Equals(frac.NULL()) {
			tail++
		} else {
			break
		}
	}
	p.coefs = p.coefs[:p.Len()-tail]

	return p
}

func (p *Pnom) Sub(other *Pnom) (self *Pnom) {
	if p.Err != nil || other.Err != nil {
		return p
	}

	l := len(p.coefs)
	m := len(other.coefs)

	if m > l {
		add := make([]*frac.Frac, m-l)
		for i := range add {
			add[i] = frac.NULL().Copy()
		}
		p.coefs = append(p.coefs, add...)
		l = m
	}

	for i := 0; i < l; i++ {
		if err := p.coefs[i].Sub(other.coefs[i]).Err; err != nil {
			p.Err = fmt.Errorf("Sub(), %d, %w", i, err)
			return p
		}
	}

	// remove "tail" of null elements
	tail := 0
	for i := p.Len() - 1; i >= 0; i-- {
		if p.coefs[i].Equals(frac.NULL()) {
			tail++
		} else {
			break
		}
	}
	p.coefs = p.coefs[:p.Len()-tail]

	return p
}
