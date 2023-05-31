package wrapper

import "fmt"

var _ Wrapper[uint, Uint] = Uint(0)
var _ Wrapper[uint8, Uint8] = Uint8(0)
var _ Wrapper[uint16, Uint16] = Uint16(0)
var _ Wrapper[uint32, Uint32] = Uint32(0)
var _ Wrapper[uint64, Uint64] = Uint64(0)
var _ Wrapper[uintptr, UintPtr] = UintPtr(0)
var _ Wrapper[byte, Byte] = Byte('a')

// Uint is a wrapper type for uint.
type Uint uint

// Equal returns true if u and o are equals.
func (u Uint) Equal(o Uint) bool {

	return u == o

}

// Compare returns -1 if u is less than o,
// 1 if u is greater than o,
// 0 if u and o are equals.
func (u Uint) Compare(o Uint) int {

	if u < o {

		return -1

	}
	if u == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of u.
func (u Uint) Hash() string {

	return fmt.Sprintf("%v", u)

}

// ToValue returns the wrapped value by u.
func (u Uint) ToValue() uint {

	return uint(u)

}

// Uint8 is a wrapper type for uint8.
type Uint8 uint8

// Equal returns true if u and o are equals.
func (u Uint8) Equal(o Uint8) bool {

	return u == o

}

// Compare returns -1 if u is less than o,
// 1 if u is greater than o,
// 0 if u and o are equals.
func (u Uint8) Compare(o Uint8) int {

	if u < o {

		return -1

	}
	if u == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of u.
func (u Uint8) Hash() string {

	return fmt.Sprintf("%v", u)

}

// ToValue returns the wrapped value by u.
func (u Uint8) ToValue() uint8 {

	return uint8(u)

}

// Uint16 is a wrapper type for uint16.
type Uint16 uint16

// Equal returns true if u and o are equals.
func (u Uint16) Equal(o Uint16) bool {

	return u == o

}

// Compare returns -1 if u is less than o,
// 1 if u is greater than o,
// 0 if u and o are equals.
func (u Uint16) Compare(o Uint16) int {

	if u < o {

		return -1

	}
	if u == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of u.
func (u Uint16) Hash() string {

	return fmt.Sprintf("%v", u)

}

// ToValue returns the wrapped value by u.
func (u Uint16) ToValue() uint16 {

	return uint16(u)

}

// Uint32 is a wrapper type for uint32.
type Uint32 uint32

// Equal returns true if u and o are equals.
func (u Uint32) Equal(o Uint32) bool {

	return u == o

}

// Compare returns -1 if u is less than o,
// 1 if u is greater than o,
// 0 if u and o are equals.
func (u Uint32) Compare(o Uint32) int {

	if u < o {

		return -1

	}
	if u == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of u.
func (u Uint32) Hash() string {

	return fmt.Sprintf("%v", u)

}

// ToValue returns the wrapped value by u.
func (u Uint32) ToValue() uint32 {

	return uint32(u)

}

// Uint64 is a wrapper type for uint64.
type Uint64 uint64

// Equal returns true if u and o are equals.
func (u Uint64) Equal(o Uint64) bool {

	return u == o

}

// Compare returns -1 if u is less than o,
// 1 if u is greater than o,
// 0 if u and o are equals.
func (u Uint64) Compare(o Uint64) int {

	if u < o {

		return -1

	}
	if u == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of u.
func (u Uint64) Hash() string {

	return fmt.Sprintf("%v", u)

}

// ToValue returns the wrapped value by u.
func (u Uint64) ToValue() uint64 {

	return uint64(u)

}

// UintPtr is a wrapper type for uintptr.
type UintPtr uintptr

// Equal returns true if u and o are equals.
func (u UintPtr) Equal(o UintPtr) bool {

	return u == o

}

// Compare returns -1 if u is less than o,
// 1 if u is greater than o,
// 0 if u and o are equals.
func (u UintPtr) Compare(o UintPtr) int {

	if u < o {

		return -1

	}
	if u == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of u.
func (u UintPtr) Hash() string {

	return fmt.Sprintf("%v", u)

}

// ToValue returns the wrapped value by u.
func (u UintPtr) ToValue() uintptr {

	return uintptr(u)

}

// Byte is a wrapper type for byte.
type Byte = Uint8
