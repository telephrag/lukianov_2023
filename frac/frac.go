package frac

import (
	"fmt"
)

var (
	NEG  = -1
	POS  = 1
	null = New(0, 1, POS)
)

func NULL() *Frac {
	return null.Copy()
}

type Frac struct {
	sign int
	num  uint64
	den  uint64

	Err error
}

func New(num, den uint64, sign int) *Frac {
	rf := &Frac{}

	rf.num = num
	rf.den = den
	rf.sign = sign

	if num == 0 && sign == NEG {
		rf.Err = ErrNullIsNonNegative
		return rf
	}
	if den == 0 {
		rf.Err = ErrNullDiv
		return rf
	}

	return rf
}

func (rf *Frac) Copy() (other *Frac) {
	return &Frac{
		sign: rf.sign,
		num:  rf.num,
		den:  rf.den,
		Err:  rf.Err,
	}
}

func (rf *Frac) Simplify() (self *Frac) {
	if rf.Err != nil {
		return rf
	}

	if rf.num == 0 { // den = 1 makes future ops easier cause gcd calculation becomes simpler
		rf.den = 1
		return rf
	}

	gcd := gcd(rf.num, rf.den)
	rf.num /= gcd
	rf.den /= gcd
	return rf
}

func (rf *Frac) Equals(other *Frac) bool {
	if rf.sign != other.sign {
		return false
	}

	// null check cause, we use division later
	rn, on := rf.num == 0, other.num == 0
	if rn && on {
		return true
	} else if rn || on {
		return false
	}

	n := rf.num >= other.num
	d := rf.den >= other.den
	if n && d {
		// fraction are equal if proportional to one another
		if rf.num/other.num+rf.num%other.num == rf.den/other.den+rf.den%other.den {
			return true
		}
	} else if !n && !d {
		if other.num/rf.num+other.num%rf.num == other.den/rf.den+other.den%rf.den {
			return true
		}
	}

	return false
}

func (rf *Frac) Sign() int {
	return rf.sign
}

func (rf *Frac) String() string {
	format := "%d/%d"

	if rf.sign == NEG {
		format = "-" + format
	}

	if rf.Err != nil {
		format += " " + fmt.Sprintf("[%s]", rf.Err)
	}

	return fmt.Sprintf(format, rf.num, rf.den)
}
