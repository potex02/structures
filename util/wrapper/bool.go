package wrapper

import (
	"fmt"
)

var _ Wrapper[bool] = Bool(false)

// Bool is a wrapper type for bool.
type Bool bool

// Equal returns true if b and o are both [Bool] and are equals.
func (b Bool) Equal(o any) bool {

	value, ok := o.(Bool)
	return ok && b == value

}

// Compare returns -1 if b is false and o is true,
// 1 if b is true and o is false,
// 0 if b and o are equals,
// -2 if o is not [Bool].
func (b Bool) Compare(o any) int {

	value, ok := o.(Bool)
	if !ok {

		return -2

	}
	if !b && value {

		return -1

	}
	if b == value {

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
