package pnom

import "lukianov_2023/frac"

var superscripts = [10]string{
	"\u2070", "\u00b9", "\u00b2", "\u00b3", "\u2074", "\u2075", "\u2076", "\u2077", "\u2078", "\u2079",
}

func getSuperscripts(n int) string { // returns member's order as unicode symbols
	var output string = ""
	for n > 1 {
		output = superscripts[n%10] + output
		n = n / 10
	}
	return output
}

func (p *Polynomial) String() string {

	var res string
	for i, c := range p.coefs[1:] {
		var plus string
		if c.Equals(frac.NULL) {
			continue
		}

		if c.Sign() == frac.POS {
			plus = "+"
		}
		res = plus + c.String() + "x" + getSuperscripts(i+1) + res
	}

	if !p.coefs[0].Equals(frac.NULL) {
		var plus string
		if p.coefs[0].Sign() == frac.POS {
			plus = "+"
		}
		res = res + plus + p.coefs[0].String()
	}

	if p.coefs[len(p.coefs)-1].Sign() == frac.POS {
		return res[1:]
	}

	return res
}
