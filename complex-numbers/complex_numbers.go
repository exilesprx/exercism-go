// Package complexnumbers implements a simple complex number type and associated methods.
package complexnumbers

import "math"

type Number struct {
	a float64
	b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
		a: n1.Real() + n2.Real(),
		b: n1.Imaginary() + n2.Imaginary(),
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		a: n1.Real() - n2.Real(),
		b: n1.Imaginary() - n2.Imaginary(),
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		a: (n1.Real() * n2.Real()) - (n1.Imaginary() * n2.Imaginary()),
		b: (n1.Imaginary() * n2.Real()) + (n1.Real() * n2.Imaginary()),
	}
}

func (n Number) Times(factor float64) Number {
	return Number{
		a: n.Real() * factor,
		b: n.Imaginary() * factor,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	factor := n2.Real()*n2.Real() + n2.Imaginary()*n2.Imaginary()
	return Number{
		a: (n1.Real()*n2.Real() + n1.Imaginary()*n2.Imaginary()) / factor,
		b: (n1.Imaginary()*n2.Real() - n1.Real()*n2.Imaginary()) / factor,
	}
}

func (n Number) Conjugate() Number {
	return Number{
		a: n.Real(),
		b: n.Imaginary() * -1.0,
	}
}

func (n Number) Abs() float64 {
	return math.Sqrt((n.Real() * n.Real()) + (n.Imaginary() * n.Imaginary()))
}

func (n Number) Exp() Number {
	factor := math.Exp(n.Real())
	return Number{
		a: factor * math.Cos(n.Imaginary()),
		b: factor * math.Sin(n.Imaginary()),
	}
}
