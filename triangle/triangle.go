// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// The different kinds of triangles
type Kind int

// All the different kinds of triangles represented by the Kind type, which is an int
const (
	NaT = iota
	Equ
	Iso
	Sca
)

// KindFromSides returns what kind of triangle is formed by the three sides
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 || a+b < c || b+c < a || a+c < b {
		return NaT
	}

	if a == b && b == c {
		return Equ
	} else if a == b || b == c || c == a {
		return Iso
	}

	return Sca
}
