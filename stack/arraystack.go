package stack

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/potex02/structures"
)

// ArrayStack provides a generic stack implemented with a slice.
//
// It implements the interface [Stack].
type ArrayStack[T any] struct {
	// contains filtered or unexported fields
	objects []T
}

// NewArrayStack returns a new [ArrayStack] containing the elements c.
// The top of the stack is the last element of c.
//
// if no argument is passed, it will be created an empty [ArrayStack].
func NewArrayStack[T any](c ...T) *ArrayStack[T] {

	return NewArrayStackFromSlice(c)

}

// NewArrayStackFromSlice returns a new [ArrayStack] containing the elements of slice c.
// The top of the stack is the last element of c
func NewArrayStackFromSlice[T any](c []T) *ArrayStack[T] {

	return &ArrayStack[T]{objects: c}

}

// Len returns the length of s.
func (s *ArrayStack[T]) Len() int {

	return len(s.objects)

}

// IsEmpty returns a bool which indicate if s is empty or not.
func (s *ArrayStack[T]) IsEmpty() bool {

	return len(s.objects) == 0

}

// Top returns the top element of s.
// If s is empty, the method returns an error.
func (s *ArrayStack[T]) Top() (T, error) {

	if s.IsEmpty() {

		var result T

		return result, errors.New("Empty stack")

	}
	return s.objects[len(s.objects)-1], nil

}

// ToSLice returns a slice which contains all elements of s.
func (s *ArrayStack[T]) ToSlice() []T {

	slice := make([]T, len(s.objects))
	copy(slice, s.objects)
	return slice

}

// Push adds the elements e at the top of s.
func (s *ArrayStack[T]) Push(e ...T) {

	s.objects = append(s.objects, e...)

}

// Pop removes an element from the top of s and returns the removed element.
// If s is empty, the method returns an error.
func (s *ArrayStack[T]) Pop() (T, error) {

	var result T

	if s.IsEmpty() {

		return result, errors.New("Empty stack")

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
func (s *ArrayStack[T]) Clear() {

	s.objects = []T{}

}

// Equals returns true if s and st are both stacks and their elements are equals.
// In any other case, it returns false.
//
// Equals does not take into account the effective type of st. This means that if st is a [LinkedStack],
// but the elements of s and the elements of st are equals, this method returns anyway true.
func (s *ArrayStack[T]) Equals(st structures.Structure[T]) bool {

	stack, ok := st.(Stack[T])
	return ok && reflect.DeepEqual(s.ToSlice(), stack.ToSlice())

}

// String returns a rapresentation of s in the form of a string.
func (s *ArrayStack[T]) String() string {

	if s.IsEmpty() {

		return fmt.Sprintf("ArrayStack[%T][%d, ]", *new(T), len(s.objects))

	}
	return fmt.Sprintf("ArrayStack[%T][%d, %v]", *new(T), len(s.objects), s.objects[len(s.objects)-1])

}
