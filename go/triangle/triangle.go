package triangle

type Kind int

const (
	NaT = iota
	Equ // equilateral
	Iso // isosceles
	Sca // scalene
	Deg // degenerate
)

// KindFromSides returns the type of triangle
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	// sort decending to make `a` the longest side
	if b > a {
		a, b = b, a
	}
	if c > b {
		b, c = c, b
	}
	if b > a {
		a, b = b, a
	}

	switch {
	case c == 0:
		k = NaT
	case a > b+c:
		k = NaT
	case a == b+c:
		k = Deg
	case a == b && b == c:
		k = Equ
	case a == b || b == c:
		k = Iso
	default:
		k = Sca
	}
	return k
}
