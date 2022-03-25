package triangle

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	m := map[float64]bool{}
	m[a] = true
	m[b] = true
	m[c] = true

	if len(m) == 1 {
		return Equ
	} else if len(m) == 2 {
		return Iso
	} else if len(m) == 3 {
		return Sca
	}

	var k Kind
	return k
}
