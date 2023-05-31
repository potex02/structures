package wrapper

import (
	"fmt"
	"math"
)

var _ Wrapper[float32, Float32] = Float32(0)
var _ Wrapper[float64, Float64] = Float64(0)

// Float32 is a wrapper type for float32.
type Float32 float32

// Equal returns true if f and o are equals.
func (f Float32) Equal(o Float32) bool {

	return f == o

}

// Compare returns -1 if f is less than o,
// 1 if f is greater than o,
// 0 if f and o are equals.
func (f Float32) Compare(o Float32) int {

	if f < o {

		return -1

	}
	if f == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of f.
func (f Float32) Hash() string {

	return fmt.Sprintf("%v", math.Abs(float64(f)))

}

// ToValue returns the wrapped value by f.
func (f Float32) ToValue() float32 {

	return float32(f)

}

// Float64 is a wrapper type for float64.
type Float64 float64

// Equal returns true if f and o are equals.
func (f Float64) Equal(o Float64) bool {

	return f == o

}

// Compare returns -1 if f is less than o,
// 1 if f is greater than o,
// 0 if f and o are equals.
func (f Float64) Compare(o Float64) int {

	if f < o {

		return -1

	}
	if f == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of f.
func (f Float64) Hash() string {

	return fmt.Sprintf("%v", math.Abs(float64(f)))

}

// ToValue returns the wrapped value by f.
func (f Float64) ToValue() float64 {

	return float64(f)

}
