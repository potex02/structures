package wrapper

import (
	"math"
	"math/bits"
)

var _ Wrapper[float32] = Float32(0)
var _ Wrapper[float64] = Float64(0)

// Float32 is a wrapper type for float32.
type Float32 float32

// Equal returns true if f and o are both [Float32] and are equals.
func (f Float32) Equal(o any) bool {
	value, ok := o.(Float32)
	return ok && f == value
}

// Compare returns -1 if f is less than o,
// 1 if f is greater than o,
// 0 if f and o are equals,
// -2 if o is not [Float32].
func (f Float32) Compare(o any) int {
	value, ok := o.(Float32)
	if !ok {
		return -2
	}
	if f < value {
		return -1
	}
	if f == value {
		return 0
	}
	return 1
}

// Hash returns the hash code of f.
func (f Float32) Hash() uint64 {
	if math.IsNaN(float64(f)) {
		return 0
	}
	if math.IsInf(float64(f), 1) {
		return 2e9
	}
	if math.IsInf(float64(f), -1) {
		return uint64(bits.Reverse(2e9))
	}
	intPart, fracPart := math.Modf(float64(f))
	hash := uint64(int(intPart))
	if fracPart != 0 {
		hash ^= uint64(fracPart * 1e9)
	}
	return hash
}

// Copy returns a copy of f.
func (f Float32) Copy() Wrapper[float32] {
	return f
}

// ToValue returns the wrapped value by f.
func (f Float32) ToValue() float32 {
	return float32(f)
}

// Float64 is a wrapper type for float64.
type Float64 float64

// Equal returns true if f and o are both [Float32] and are equals.
func (f Float64) Equal(o any) bool {
	value, ok := o.(Float64)
	return ok && f == value
}

// Compare returns -1 if f is less than o,
// 1 if f is greater than o,
// 0 if f and o are equals,
// -2 if o is not [Float64].
func (f Float64) Compare(o any) int {
	value, ok := o.(Float64)
	if !ok {
		return -2
	}
	if f < value {
		return -1
	}
	if f == value {
		return 0
	}
	return 1
}

// Hash returns the hash code of f.
func (f Float64) Hash() uint64 {
	if math.IsNaN(float64(f)) {
		return 0
	}
	if math.IsInf(float64(f), 1) {
		return 2e9
	}
	if math.IsInf(float64(f), -1) {
		return uint64(bits.Reverse(2e9))
	}
	intPart, fracPart := math.Modf(float64(f))
	hash := uint64(int(intPart))
	if fracPart != 0 {
		hash ^= uint64(fracPart * 1e9)
	}
	return hash
}

// Copy returns a copy of f.
func (f Float64) Copy() Wrapper[float64] {
	return f
}

// ToValue returns the wrapped value by f.
func (f Float64) ToValue() float64 {
	return float64(f)
}
