package complexnumbers

import "math"

type Number struct {
	R, I float64
}

func (n Number) Real() float64 {
	return n.R
}

func (n Number) Imaginary() float64 {
	return n.I
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.R + n2.R, n1.I + n2.I}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.R - n2.R, n1.I - n2.I}
}

func (n1 Number) Multiply(n2 Number) Number {
	// (a + i * b) * (c + i * d) = (a * c - b * d) + (b * c + a * d) * i
	return Number{n1.R*n2.R - n1.I*n2.I, n1.I*n2.R + n1.R*n2.I}
}

func (n Number) Times(factor float64) Number {
	return Number{n.R * factor, n.I * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	// (a * c + b * d)/(c^2 + d^2) + (b * c - a * d)/(c^2 + d^2) * i
	d := n2.R*n2.R + n2.I*n2.I
	return Number{(n1.R*n2.R + n1.I*n2.I) / d, (n1.I*n2.R - n1.R*n2.I) / d}
}

func (n Number) Conjugate() Number {
	return Number{n.R, -n.I}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.R*n.R + n.I*n.I)
}

func (n Number) Exp() Number {
	// e^(a + i * b) = e^a * e^(i * b)
	// the last term of which is given by Euler's formula
	// e^(i * b) = cos(b) + i * sin(b).
	return Number{math.Cos(n.I), math.Sin(n.I)}.Times(math.Exp(n.R))
}
