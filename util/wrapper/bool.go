package wrapper

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
func (b Bool) Hash() uint64 {
	if b {
		return 1
	}
	return 0
}

// Copy returns a copy of b.
func (b Bool) Copy() Wrapper[bool] {
	return b
}

// ToValue returns the wrapped value by b.
func (b Bool) ToValue() bool {
	return bool(b)
}
