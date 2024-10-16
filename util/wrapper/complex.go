package wrapper

import (
	"math"

	"github.com/potex02/structures/util"
)

var _ Wrapper[complex64] = Complex64(0)
var _ Wrapper[complex128] = Complex128(0)

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

// Equal returns true if c and o are both [Complex64] and are equals.
func (c Complex64) Equal(o any) bool {
	value, ok := o.(Complex64)
	return ok && c == value
}

// Compare returns -1 if c is less than o,
// 1 if c is greater than o,
// 0 if c and o are equals,
// -2 if o is not [Complex64].
func (c Complex64) Compare(o any) int {
	value, ok := o.(Complex64)
	if !ok {
		return -2
	}
	if c.Norm() < value.Norm() {
		return -1
	}
	if c.Norm() == value.Norm() {
		return 0
	}
	return 1
}

// Hash returns the hash code of c.
func (c Complex64) Hash() uint64 {
	return Float32(c.Real()).Hash() + util.Prime*Float32(c.Imag()).Hash()
}

// Copy returns a copy of c.
func (c Complex64) Copy() Wrapper[complex64] {
	return c
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

// Equal returns true if c and o are both [Complex128] and are equals.
func (c Complex128) Equal(o any) bool {
	value, ok := o.(Complex128)
	return ok && c == value
}

// Compare returns -1 if c is less than o,
// 1 if c is greater than o,
// 0 if c and o are equals,
// -2 if o is not [Complex128].
func (c Complex128) Compare(o any) int {
	value, ok := o.(Complex128)
	if !ok {
		return -2
	}
	if c.Norm() < value.Norm() {
		return -1
	}
	if c.Norm() == value.Norm() {
		return 0
	}
	return 1
}

// Hash returns the hash code of c.
func (c Complex128) Hash() uint64 {
	return Float64(c.Real()).Hash() + util.Prime*Float64(c.Imag()).Hash()
}

// Copy returns a copy of c.
func (c Complex128) Copy() Wrapper[complex128] {
	return c
}

// ToValue returns the wrapped value by c.
func (c Complex128) ToValue() complex128 {
	return complex128(c)
}
