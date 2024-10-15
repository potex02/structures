package wrapper

import "strings"

var _ Wrapper[string] = String("")

// String is a wrapper type for string.
//
// a String can be indexed with the [] as a normal string.
type String string

// Len returns the numbers of characters of s.
func (s String) Len() int {
	return len(s)
}

// IsEmpty returns a bool which indicates if s is empty or not.
func (s String) IsEmpty() bool {
	return len(s) == 0
}

// Clone returns a copy of s.
func (s String) Clone() String {
	return String(strings.Clone(s.ToValue()))
}

// Contains returns if s contains the str substring.
func (s String) Contains(str String) bool {
	return strings.Contains(s.ToValue(), str.ToValue())
}

// HasPrefix tests whether the string s begins with prefix.
func (s String) HasPrefix(prefix String) bool {
	return strings.HasPrefix(s.ToValue(), prefix.ToValue())
}

// HasSuffix tests whether the string s ends with prefix.
func (s String) HasSuffix(suffix String) bool {
	return strings.HasSuffix(s.ToValue(), suffix.ToValue())
}

// IndexOf returns the first position of the str substring in s.
// If str is not present, the result is -1.
func (s String) IndexOf(str String) int {
	return strings.Index(s.ToValue(), str.ToValue())
}

// IndexOf returns the first position of the c Byte in s.
// If c is not present, the result is -1.
func (s String) IndexByteOf(c Byte) int {
	return strings.IndexByte(s.ToValue(), c.ToValue())
}

// LastIndexOf returns the last position of the str substring in s.
// If str is not present, the result is -1.
func (s String) LastIndexOf(str String) int {
	return strings.LastIndex(s.ToValue(), str.ToValue())
}

// LastIndexOf returns the last position of the c Byte in s.
// If c is not present, the result is -1.
func (s String) LastIndexByteOf(c Byte) int {
	return strings.LastIndexByte(s.ToValue(), c.ToValue())
}

// Split slices s into all substrings separated by sep
// and returns a slice of the substrings between those separators.
func (s String) Split(sep String) []String {
	slice := strings.Split(s.ToValue(), sep.ToValue())
	result := make([]String, len(slice))
	for i, j := range slice {
		result[i] = String(j)
	}
	return result
}

// ToLowerCase returns a copy of s with all characters in lower case.
func (s String) ToLowerCase() String {
	return String(strings.ToLower(s.ToValue()))
}

// ToUpperCase returns a copy of s with all characters in upper case.
func (s String) ToUpperCase() String {
	return String(strings.ToUpper(s.ToValue()))
}

// Trim returns a copy of s with all leading and trailing characters contained in cutset removed.
func (s String) Trim(cutset String) String {
	return String(strings.Trim(s.ToValue(), cutset.ToValue()))
}

// TrimLeft returns a copy of s with all leading characters contained in cutset removed.
func (s String) TrimLeft(cutset String) String {
	return String(strings.TrimLeft(s.ToValue(), cutset.ToValue()))
}

// TrimRight returns a copy of s with all trailing characters contained in cutset removed.
func (s String) TrimRight(cutset String) String {
	return String(strings.TrimRight(s.ToValue(), cutset.ToValue()))
}

// TrimSpace returns a copy of s with all leading and trailing empty spaces removed.
func (s String) TrimSpace() String {
	return String(strings.TrimSpace(s.ToValue()))
}

// Equal returns true if s and o are both [String] and are equals.
func (s String) Equal(o any) bool {
	value, ok := o.(String)
	return ok && s == value
}

// EqualFold returns true if s and o are both [String] and are equals.
//
// This method make the comparison in a case insensitive way.
func (s String) EqualFold(o any) bool {
	value, ok := o.(String)
	return ok && strings.EqualFold(s.ToValue(), value.ToValue())
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

// Copy returns a copy of s.
func (s String) Copy() Wrapper[string] {
	return s
}

// ToValue returns the wrapped value by s.
func (s String) ToValue() string {
	return string(s)
}
