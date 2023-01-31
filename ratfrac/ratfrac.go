package ratfrac

import (
	"fmt"
)

var (
	SIGN_NEG = -1
	SIGN_POS = 1
)

type RatFrac struct {
	sign int
	num  uint64
	den  uint64

	Err error
}

func New(num, den uint64, sign int) *RatFrac {
	rf := &RatFrac{}
	rf.num = num
	rf.den = den
	if den == 0 {
		rf.Err = ErrNullDiv
	}
	rf.sign = sign
	return rf
}

func (rf *RatFrac) Copy() (other *RatFrac) {
	return &RatFrac{
		sign: rf.sign,
		num:  rf.num,
		den:  rf.den,
		Err:  rf.Err,
	}
}

func (rf *RatFrac) Simplify() (self *RatFrac) {
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

func (rf *RatFrac) Equals(other *RatFrac) bool {
	if rf.sign != other.sign {
		return false
	}

	n := rf.num >= other.num
	d := rf.den >= other.den

	if rf.num == 0 && other.num == 0 {
		return true
	}

	if n && d {
		// if fractions are proportional to one another
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

func (rf *RatFrac) Sign() int {
	return rf.sign
}

func (rf *RatFrac) Neg() (self *RatFrac) {
	rf.sign *= -1
	return rf
}

func (rf *RatFrac) Add(other *RatFrac) (self *RatFrac) {
	if rf.Err != nil {
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

func (rf *RatFrac) Sub(other *RatFrac) (self *RatFrac) {
	if rf.Err != nil {
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

func (rf *RatFrac) Mul(other *RatFrac) (self *RatFrac) {
	if rf.Err != nil {
		return rf
	}

	if rf.sign == other.sign {
		rf.sign = SIGN_POS
	} else {
		rf.sign = SIGN_NEG
	}

	rf.num = rf.mul(rf.num, other.num)
	rf.den = rf.mul(rf.den, other.den)

	return rf
}

func (rf *RatFrac) Div(other *RatFrac) (self *RatFrac) {
	if rf.Err != nil {
		return rf
	}

	if other.num == 0 {
		rf.Err = ErrNullDiv
		return rf
	}

	if rf.sign == other.sign {
		rf.sign = SIGN_POS
	} else {
		rf.sign = SIGN_NEG
	}

	rf.num = rf.mul(rf.num, other.den)
	rf.den = rf.mul(rf.den, other.num)

	return rf
}

func (rf *RatFrac) String() string {
	format := "%d/%d"

	if rf.sign == SIGN_NEG {
		format = "-" + format
	}

	if rf.Err != nil {
		format += " " + fmt.Sprintf("[%s]", rf.Err)
	}

	return fmt.Sprintf(format, rf.num, rf.den)
}
