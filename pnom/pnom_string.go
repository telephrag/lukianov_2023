package pnom

import "lukianov_2023/frac"

var superscripts = [10]string{
	"\u2070", "\u00b9", "\u00b2", "\u00b3", "\u2074", "\u2075", "\u2076", "\u2077", "\u2078", "\u2079",
}

func getXToPow(n int) string {
	var output string

	nn := n != 0

	for n > 1 {
		output = superscripts[n%10] + output
		n = n / 10
	}

	if nn {
		output = "x" + output
	}

	return output
}

func (p *Pnom) String() string {
	var res string

	plusPrev := ""
	for i, c := range p.coefs {
		if c.Equals(frac.NULL()) {
			continue
		}

		res = c.String() + getXToPow(i) + plusPrev + res
		if c.Sign() == frac.POS {
			plusPrev = "+"
		} else {
			plusPrev = ""
		}
	}

	return res
}
