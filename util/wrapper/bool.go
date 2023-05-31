package wrapper

import (
	"fmt"
)

var _ Wrapper[bool, Bool] = Bool(false)

// Bool is a wrapper type for bool.
type Bool bool

// Equal returns true if b and o are equals.
func (b Bool) Equal(o Bool) bool {

	return b == o

}

// Compare returns -1 if b is false and o is true,
// 1 if b is true and o is false,
// 0 if b and o are equals.
func (b Bool) Compare(o Bool) int {

	if !b && o {

		return -1

	}
	if b == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of b.
func (b Bool) Hash() string {

	return fmt.Sprintf("%v", b)

}

// ToValue returns the wrapped value by b.
func (b Bool) ToValue() bool {

	return bool(b)

}
