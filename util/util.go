// package utils provides interfaces to peform equality, comparion and hashing operations.
package util

// Equaler defines a method to check the equality between two variables.
type Equaler[T any] interface {
	// Equal returns true if the receiver and o are equals.
	Equal(o T) bool
}

// Comparer defines a method used to compare two variable of type T.
type Comparer[T any] interface {
	// Compare returns a int which indicates the order of the elements.
	//
	// If the result is less than zero, the receiver is placed before o.
	// If the result is greater than zero, the receiver is placed after the parameter.
	// If the the result is zero, the elements are ordered randomly.
	Compare(o T) int
}

// Hasher defines a method that provides hashing for the variable.
type Hasher[T any] interface {
	Comparer[T]
	// Hash returns a string that is used to perform the hashing.
	Hash() string
}
