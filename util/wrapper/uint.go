package wrapper

import "fmt"

var _ Wrapper[uint] = Uint(0)
var _ Wrapper[uint8] = Uint8(0)
var _ Wrapper[uint16] = Uint16(0)
var _ Wrapper[uint32] = Uint32(0)
var _ Wrapper[uint64] = Uint64(0)
var _ Wrapper[uintptr] = UintPtr(0)
var _ Wrapper[byte] = Byte('a')

// Uint is a wrapper type for uint.
type Uint uint

// Equal returns true if u and o are both [Uint] and are equals.
func (u Uint) Equal(o any) bool {
	value, ok := o.(Uint)
	return ok && u == value
}

// Compare returns -1 if u is less than o,
// 1 if u is greater than o,
// 0 if u and o are equals,
// -2 if o is not [Uint].
func (u Uint) Compare(o any) int {
	value, ok := o.(Uint)
	if !ok {
		return -2
	}
	if u < value {
		return -1
	}
	if u == value {
		return 0
	}
	return 1
}

// Hash returns the hash code of u.
func (u Uint) Hash() string {
	return fmt.Sprintf("%v", u)
}

// Copy returns a copy of u.
func (u Uint) Copy() Wrapper[uint] {
	return u
}

// ToValue returns the wrapped value by u.
func (u Uint) ToValue() uint {
	return uint(u)
}

// Uint8 is a wrapper type for uint8.
type Uint8 uint8

// Equal returns true if u and o are both [Uint8] and are equals.
func (u Uint8) Equal(o any) bool {
	value, ok := o.(Uint8)
	return ok && u == value
}

// Compare returns -1 if u is less than o,
// 1 if u is greater than o,
// 0 if u and o are equals,
// -2 if o is not [Uint8].
func (u Uint8) Compare(o any) int {
	value, ok := o.(Uint8)
	if !ok {
		return -2
	}
	if u < value {
		return -1
	}
	if u == value {
		return 0
	}
	return 1
}

// Hash returns the hash code of u.
func (u Uint8) Hash() string {
	return fmt.Sprintf("%v", u)
}

// Copy returns a copy of u.
func (u Uint8) Copy() Wrapper[uint8] {
	return u
}

// ToValue returns the wrapped value by u.
func (u Uint8) ToValue() uint8 {
	return uint8(u)
}

// Uint16 is a wrapper type for uint16.
type Uint16 uint16

// Equal returns true if u and o are both [Uint16] and are equals.
func (u Uint16) Equal(o any) bool {
	value, ok := o.(Uint16)
	return ok && u == value
}

// Compare returns -1 if u is less than o,
// 1 if u is greater than o,
// 0 if u and o are equals,
// -2 if o is not [Uint16].
func (u Uint16) Compare(o any) int {
	value, ok := o.(Uint16)
	if !ok {
		return -2
	}
	if u < value {
		return -1
	}
	if u == value {
		return 0
	}
	return 1
}

// Hash returns the hash code of u.
func (u Uint16) Hash() string {
	return fmt.Sprintf("%v", u)
}

// Copy returns a copy of u.
func (u Uint16) Copy() Wrapper[uint16] {
	return u
}

// ToValue returns the wrapped value by u.
func (u Uint16) ToValue() uint16 {
	return uint16(u)
}

// Uint32 is a wrapper type for uint32.
type Uint32 uint32

// Equal returns true if u and o are both [Uint32] and are equals.
func (u Uint32) Equal(o any) bool {
	value, ok := o.(Uint32)
	return ok && u == value
}

// Compare returns -1 if u is less than o,
// 1 if u is greater than o,
// 0 if u and o are equals,
// -2 if o is not [Uint32].
func (u Uint32) Compare(o any) int {
	value, ok := o.(Uint32)
	if !ok {
		return -2
	}
	if u < value {
		return -1
	}
	if u == value {
		return 0
	}
	return 1
}

// Hash returns the hash code of u.
func (u Uint32) Hash() string {
	return fmt.Sprintf("%v", u)
}

// Copy returns a copy of u.
func (u Uint32) Copy() Wrapper[uint32] {
	return u
}

// ToValue returns the wrapped value by u.
func (u Uint32) ToValue() uint32 {
	return uint32(u)
}

// Uint64 is a wrapper type for uint64.
type Uint64 uint64

// Equal returns true if u and o are both [Uint64] and are equals.
func (u Uint64) Equal(o any) bool {
	value, ok := o.(Uint64)
	return ok && u == value
}

// Compare returns -1 if u is less than o,
// 1 if u is greater than o,
// 0 if u and o are equals,
// -2 if o is not [Uint64].
func (u Uint64) Compare(o any) int {
	value, ok := o.(Uint64)
	if !ok {
		return -2
	}
	if u < value {
		return -1
	}
	if u == value {
		return 0
	}
	return 1
}

// Hash returns the hash code of u.
func (u Uint64) Hash() string {
	return fmt.Sprintf("%v", u)
}

// Copy returns a copy of u.
func (u Uint64) Copy() Wrapper[uint64] {
	return u
}

// ToValue returns the wrapped value by u.
func (u Uint64) ToValue() uint64 {
	return uint64(u)
}

// UintPtr is a wrapper type for uintptr.
type UintPtr uintptr

// Equal returns true if u and o are both [UintPtr] and are equals.
func (u UintPtr) Equal(o any) bool {
	value, ok := o.(UintPtr)
	return ok && u == value
}

// Compare returns -1 if u is less than o,
// 1 if u is greater than o,
// 0 if u and o are equals,
// -2 if o is not [UintPtr].
func (u UintPtr) Compare(o any) int {
	value, ok := o.(UintPtr)
	if !ok {
		return -2
	}
	if u < value {
		return -1
	}
	if u == value {
		return 0
	}
	return 1
}

// Hash returns the hash code of u.
func (u UintPtr) Hash() string {
	return fmt.Sprintf("%v", u)
}

// Copy returns a copy of u.
func (u UintPtr) Copy() Wrapper[uintptr] {
	return u
}

// ToValue returns the wrapped value by u.
func (u UintPtr) ToValue() uintptr {
	return uintptr(u)
}

// Byte is a wrapper type for byte.
type Byte = Uint8
