// package utils provides interfaces to peform equality, comparion and hashing operations.
package util

import "reflect"

// Equaler defines a method to check the equality between two variables.
type Equaler interface {
	// Equal returns true if the receiver and o are equals.
	Equal(o any) bool
}

// Comparer defines a method used to compare two variable of type T.
type Comparer interface {
	// Compare returns a int which indicates the order of the elements.
	//
	// If the result is less than zero, the receiver is placed before o.
	// If the result is greater than zero, the receiver is placed after the parameter.
	// If the the result is zero, the elements are ordered randomly.
	Compare(o any) int
}

// Hasher defines a method that provides hashing for the variable.
type Hasher interface {
	Comparer
	// Hash returns a string that is used to perform the hashing.
	Hash() string
}

// Copier defines a method that permits to make copies of a variable.
type Copier[T any] interface {
	// Copy returns a copy of the receiver.
	Copy() T
}

// EqualFunction generate a function that can be used to check the equality of e and other.
//
// if T implements [Equaler], the resulting function use the Equal method,
// otherwhise it use [reflect.DeepEqual].
func EqualFunction(e any) func(other any) bool {

	if element, ok := interface{}(e).(Equaler); ok {

		return func(other any) bool {
			return element.Equal(other)
		}

	}
	return func(other any) bool {

		return reflect.DeepEqual(e, other)

	}

}

// Copy returns a copy of e.
//
// if T implements [Copier], the method returns e.Copy(),
// otherwhise it returns e.
func Copy[T any](e T) T {

	if element, ok := interface{}(e).(Copier[T]); ok {

		return element.Copy()

	}
	return e

}
