package frac

import (
	"fmt"
)

var (
	NEG  = -1
	POS  = 1
	NULL = New(0, 1, POS)
)

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
	if den == 0 {
		rf.Err = ErrNullDiv
	}
	rf.sign = sign
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

func (rf *Frac) Neg() (self *Frac) {
	rf.sign *= -1
	return rf
}

func (rf *Frac) Add(other *Frac) (self *Frac) {
	if rf.Err != nil || other.Err != nil {
		return
	}

	// prep
	gcd := gcd(rf.den, other.den)
	left := rf.den/gcd + rf.den%gcd
	right := other.den/gcd + other.den%gcd

	// calculationg denominator in advance
	rf.den = rf.mul(left, right)
	rf.den = rf.mul(rf.den, gcd)

	nl := rf.mul(rf.num, right)
	nr := rf.mul(other.num, left)

	if rf.sign == other.sign {
		rf.num = rf.add(nl, nr)
	} else {
		if nl < nr {
			nl, nr = nr, nl
			rf.sign *= -1
		}
		rf.num = nl - nr // overflow is impossible since (nl => nr) is guaranteed
	}

	return rf
}

func (rf *Frac) Sub(other *Frac) (self *Frac) {
	if rf.Err != nil || other.Err != nil {
		return rf
	}

	// prep
	gcd := gcd(rf.den, other.den)
	left := rf.den / gcd
	right := other.den / gcd

	// calculationg denominator in advance
	rf.den = rf.mul(left, right)
	rf.den = rf.mul(rf.den, gcd)

	nl := rf.mul(rf.num, right)
	nr := rf.mul(other.num, left)

	if nl < nr {
		nl, nr = nr, nl
		rf.sign *= -1
	}
	rf.num = nl - nr // overflow is impossible since (nl => nr) is guaranteed

	return rf
}

func (rf *Frac) Mul(other *Frac) (self *Frac) {
	if rf.Err != nil || other.Err != nil {
		return rf
	}

	if rf.sign == other.sign {
		rf.sign = POS
	} else {
		rf.sign = NEG
	}

	rf.num = rf.mul(rf.num, other.num)
	rf.den = rf.mul(rf.den, other.den)

	return rf
}

func (rf *Frac) Div(other *Frac) (self *Frac) {
	if rf.Err != nil || other.Err != nil {
		return rf
	}

	if other.num == 0 {
		rf.Err = ErrNullDiv
		return rf
	}

	if rf.sign == other.sign {
		rf.sign = POS
	} else {
		rf.sign = NEG
	}

	rf.num = rf.mul(rf.num, other.den)
	rf.den = rf.mul(rf.den, other.num)

	return rf
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
