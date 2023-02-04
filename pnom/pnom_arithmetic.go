package pnom

import (
	"fmt"
	"lukianov_2023/frac"
)

func (p *Pnom) RemoveTrailingNulls() (self *Pnom) {
	tail := 0
	for i := p.len() - 1; i >= 0; i-- {
		if p.coefs[i].Equals(frac.NULL()) {
			tail++
		} else {
			break
		}
	}
	p.coefs = p.coefs[:p.len()-tail]
	return p
}

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

	l := p.RemoveTrailingNulls().len()
	m := other.RemoveTrailingNulls().len()

	if m > l {
		add := make([]*frac.Frac, m-l)
		for i := range add {
			add[i] = frac.NULL()
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

func (p *Pnom) Sub(other *Pnom) (self *Pnom) {
	if p.Err != nil || other.Err != nil {
		return p
	}

	l := p.RemoveTrailingNulls().len()
	m := other.RemoveTrailingNulls().len()

	if m > l {
		add := make([]*frac.Frac, m-l)
		for i := range add {
			add[i] = frac.NULL()
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

	return p
}

func (p *Pnom) Mul(other *Pnom) (self *Pnom) {
	if p.Err != nil || other.Err != nil {
		return p
	}

	// length of `p` will be max power of `p` + max power of `other`
	res := make([]*frac.Frac,
		other.RemoveTrailingNulls().len()+p.RemoveTrailingNulls().len(),
	)
	for i := range res {
		res[i] = frac.NULL()
	}

	for i := range other.coefs {
		for j := range p.coefs {
			res[i+j].Add(
				other.coefs[i].Copy().Mul(p.coefs[j]),
			)
		}
	}

	p.coefs = res

	return p
}
