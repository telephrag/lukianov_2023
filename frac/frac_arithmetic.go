package frac

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
		if nl == nr { // null is always non-negative
			rf.sign = POS
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

	if rf.sign == other.sign {
		if nl < nr {
			nl, nr = nr, nl
			rf.sign *= -1
		}
		rf.num = nl - nr // overflow is impossible since (nl => nr) is guaranteed
	} else {
		rf.num = rf.add(nl, nr) // -a-b=-(a+b); a-(-b)=a+b
	}

	if nl == nr { // null is always non-negative
		rf.sign = POS
	}

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
