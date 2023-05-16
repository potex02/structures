// Package structures implements the most common used data structures.
package structures

// Structure defines commons methods for all data structures.
//
// A Structure is a generic and can be used with any type T.
type Structure[T any] interface {
	// Len returns the numbers of elements in the structure.
	Len() int
	// IsEmpty returns a bool which indicate if the structure is empty or not.
	IsEmpty() bool
	// ToSLice returns a slice which contains all elements of the structure.
	ToSlice() []T
	// Clear removes all element from the structure.
	Clear()
	// Equals returns true if the structure and st are the same type of structure and their elements are equals.
	// In any other case, it returns false.
	Equals(s Structure[T]) bool
	// String returns a rapresentation of the structure in the form of a string.
	String() string
}
