package wrapper

import (
	"fmt"
	"math"
)

var _ Wrapper[complex64, Complex64] = Complex64(0)
var _ Wrapper[complex128, Complex128] = Complex128(0)

// Complex64 is a wrapper type for complex64.
type Complex64 complex64

// Real returns the real part of c.
func (c Complex64) Real() float64 {

	return real(complex128(c))

}

// Imag returns the imaginary part of c.
func (c Complex64) Imag() float64 {

	return imag(complex128(c))

}

// Norm returns the norm of c.
//
// The norm of a complex number is equal to sqrt(c.Real() ^ 2 + c.Imag() ^ 2).
func (c Complex64) Norm() float64 {

	return math.Sqrt(math.Pow(c.Real(), 2) + math.Pow(c.Imag(), 2))

}

// Equal returns true if c and o are equals.
func (c Complex64) Equal(o Complex64) bool {

	return c == o

}

// Compare returns -1 if c is less than o,
// 1 if c is greater than o,
// 0 if c and o are equals.
func (c Complex64) Compare(o Complex64) int {

	if c.Norm() < o.Norm() {

		return -1

	}
	if c.Norm() == o.Norm() {

		return 0

	}
	return 1

}

// Hash returns the hash code of c.
func (c Complex64) Hash() string {

	return fmt.Sprintf("%v", c.Norm())

}

// ToValue returns the wrapped value by c.
func (c Complex64) ToValue() complex64 {

	return complex64(c)

}

// Complex128 is a wrapper type for complex128.
type Complex128 complex128

// Real returns the real part of c.
func (c Complex128) Real() float64 {

	return real(complex128(c))

}

// Imag returns the imaginary part of c.
func (c Complex128) Imag() float64 {

	return imag(complex128(c))

}

// Norm returns the norm of c.
//
// The norm of a complex number is equal to sqrt(c.Real() ^ 2 + c.Imag() ^ 2).
func (c Complex128) Norm() float64 {

	return math.Sqrt(math.Pow(c.Real(), 2) + math.Pow(c.Imag(), 2))

}

// Equal returns true if c and o are equals.
func (c Complex128) Equal(o Complex128) bool {

	return c == o

}

// Compare returns -1 if c is less than o,
// 1 if c is greater than o,
// 0 if c and o are equals.
func (c Complex128) Compare(o Complex128) int {

	if c.Norm() < o.Norm() {

		return -1

	}
	if c.Norm() == o.Norm() {

		return 0

	}
	return 1

}

// Hash returns the hash code of c.
func (c Complex128) Hash() string {

	return fmt.Sprintf("%v", c.Norm())

}

// ToValue returns the wrapped value by c.
func (c Complex128) ToValue() complex128 {

	return complex128(c)

}
