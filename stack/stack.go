// Package stack implements dinamic stacks.
package stack

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/potex02/structures"
)

// Stack provides a generic single stack implemented with a slice.
type Stack[T any] struct {
	objects []T
}

// New returns a new empty [Stack].
func New[T any]() *Stack[T] {

	return &Stack[T]{}

}

// NewFromElements is a wrapper for NewFromSlice(c).
func NewFromElements[T any](c ...T) *Stack[T] {

	return &Stack[T]{c}

}

// NewFromSlice returns a new [Stack] containing the elements of slice c.
func NewFromSlice[T any](c []T) *Stack[T] {

	return &Stack[T]{c}

}

// Len returns the length of s.
func (s *Stack[T]) Len() int {

	return len(s.objects)

}

// IsEmpty returns a bool wich indicate if s is empty or not
func (s *Stack[T]) IsEmpty() bool {

	return len(s.objects) == 0

}

// Head returns a pointer to the top of s.
// If s is empty, the method returns nil.
func (s *Stack[T]) Top() *T {

	if s.IsEmpty() {

		return nil

	}
	return &s.objects[len(s.objects)-1]

}

// ToSLice returns a slice wich contains all elements of s.
func (s *Stack[T]) ToSlice() []T {

	slice := make([]T, len(s.objects))
	copy(slice, s.objects)
	return slice

}

// Add adds the element e at the top of s.
func (s *Stack[T]) Add(e T) {

	s.objects = append(s.objects, e)

}

// Remove removes an element from the top of s and returns the removed element.
// It returns an error is q is empty.
func (s *Stack[T]) Remove() (T, error) {

	var result T

	if s.IsEmpty() {

		return result, errors.New("Stack empty")

	}
	result = s.objects[len(s.objects)-1]
	if len(s.objects) > 1 {

		s.objects = s.objects[:len(s.objects)-1]

	} else {

		s.Clear()

	}
	return result, nil

}

// Clear removes all element from s.
func (s *Stack[T]) Clear() {

	s.objects = []T{}

}

// Equals returns true if s and st are both stacks and their elements are equals.
// In any other case, it returns false.
func (s *Stack[T]) Equals(st structures.Structure[T]) bool {

	stack, ok := st.(*Stack[T])
	return ok && reflect.DeepEqual(s.ToSlice(), stack.ToSlice())

}

// String returns a rapresentation of s in the form of a string.
func (s *Stack[T]) String() string {

	if s.IsEmpty() {

		return fmt.Sprintf("Stack[%T][%d, %v]", *new(T), len(s.objects), nil)

	}
	return fmt.Sprintf("Stack[%T][%d, %v]", *new(T), len(s.objects), s.objects[len(s.objects)-1])

}
