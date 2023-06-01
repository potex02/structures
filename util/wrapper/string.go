package wrapper

var _ Wrapper[string] = String("")

// String is a wrapper type for string.
type String string

// Equal returns true if s and o are both [String] and are equals.
func (s String) Equal(o any) bool {

	value, ok := o.(String)
	return ok && s == value

}

// Compare returns -1 if s is less than o,
// 1 if s is greater than o,
// 0 if s and o are equals,
// -2 if o is not [String].
//
// The comparison is made in lexicographical order.
func (s String) Compare(o any) int {

	value, ok := o.(String)
	if !ok {

		return -2

	}
	if s < value {

		return -1

	}
	if s == value {

		return 0

	}
	return 1

}

// Hash returns the hash code of s.
func (s String) Hash() string {

	return string(s)

}

// ToValue returns the wrapped value by s.
func (s String) ToValue() string {

	return string(s)

}
