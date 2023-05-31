package wrapper

var _ Wrapper[string, String] = String("")

// String is a wrapper type for string.
type String string

// Equal returns true if s and o are equals.
func (s String) Equal(o String) bool {

	return s == o

}

// Compare returns -1 if s is less than o,
// 1 if s is greater than o,
// 0 if s and o are equals.
//
// The comparison is made in lexicographical order.
func (s String) Compare(o String) int {

	if s < o {

		return -1

	}
	if s == o {

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
