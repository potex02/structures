package stack

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
)

var _ structures.Structure[int] = NewArrayStack[int]()
var _ Stack[int] = NewArrayStack[int]()

// ArrayStack provides a generic stack implemented through an [list.ArrayList].
//
// It implements the interface [Stack].
type ArrayStack[T any] struct {
	// contains filtered or unexported fields
	objects list.List[T]
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

	return &ArrayStack[T]{objects: list.NewArrayListFromSlice(c)}

}

// Len returns the length of s.
func (s *ArrayStack[T]) Len() int {

	return s.objects.Len()

}

// IsEmpty returns a bool which indicates if s is empty or not.
func (s *ArrayStack[T]) IsEmpty() bool {

	return s.objects.IsEmpty()

}

// Top returns the top element of s.
// If s is empty, the method returns an error.
func (s *ArrayStack[T]) Top() (T, error) {

	result, err := s.objects.Get(s.Len() - 1)
	if err != nil {

		return result, errors.New("Empty stack")

	}
	return result, err

}

// ToSlice returns a slice which contains all elements of s.
func (s *ArrayStack[T]) ToSlice() []T {

	return s.objects.ToSlice()

}

// Push adds the elements e at the top of s.
func (s *ArrayStack[T]) Push(e ...T) {

	s.objects.Add(e...)

}

// Pop removes an element from the top of s and returns the removed element.
// If s is empty, the method returns an error.
func (s *ArrayStack[T]) Pop() (T, error) {

	result, err := s.objects.Remove(s.Len() - 1)
	if err != nil {

		return result, errors.New("Empty stack")

	}
	return result, err

}

// Clear removes all element from s.
func (s *ArrayStack[T]) Clear() {

	s.objects.Clear()

}

// Equal returns true if s and st are both stacks and their elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is a [LinkedStack],
// but the elements of s and the elements of st are equals, this method returns anyway true.
func (s *ArrayStack[T]) Equal(st any) bool {

	stack, ok := st.(Stack[T])
	if ok && s != nil && stack != nil {

		return s.objects.Equal(list.NewArrayListFromStructure[T](stack))

	}
	return false

}

// Compare returns 0 if s and st are equals,
// -1 if s is shorten than st,
// 1 if s is longer than st,
// -2 if st is not a [Stack] or if one between s and st is nil.
//
// If s and st have the same length, the result is the comparison
// between the first different element of the two stacks if T implemets [util.Comparer],
// otherwhise the result is 0.
func (s *ArrayStack[T]) Compare(st any) int {

	stack, ok := st.(Stack[T])
	if ok && s != nil && stack != nil {

		return s.objects.Compare(list.NewArrayListFromStructure[T](stack))

	}
	return -2

}

// Hash returns the hash code of s.
func (s *ArrayStack[T]) Hash() string {

	check := reflect.TypeOf(new(T)).String()
	top, _ := s.Top()
	return fmt.Sprintf("%v%v", check[1:], top)

}

// String returns a rapresentation of s in the form of a string.
func (s *ArrayStack[T]) String() string {

	check := reflect.TypeOf(new(T)).String()
	if s.IsEmpty() {

		return fmt.Sprintf("ArrayStack[%v][%d, ]", check[1:], s.objects.Len())

	}
	element, _ := s.Top()
	return fmt.Sprintf("ArrayStack[%v][%d, %v]", check[1:], s.objects.Len(), element)

}
