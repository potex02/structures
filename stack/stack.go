// Package stack implements dynamic stacks.
package stack

import (
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
)

var _ structures.Structure[int] = NewStack[int]()

// Stack provides a generic LIFO structure implemented through an [list.ArrayList].
// An stack contains all the methods of [structures.Structure].
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
type Stack[T any] struct {
	// contains filtered or unexported fields
	objects list.List[T]
}

// NewStack returns a new [Stack] containing the elements c.
// The top of the stack is the last element of c.
//
// if no argument is passed, it will be created an empty [Stack].
func NewStack[T any](c ...T) *Stack[T] {
	return NewStackFromSlice(c)
}

// NewStackFromSlice returns a new [Stack] containing the elements of slice c.
// The top of the stack is the last element of c
func NewStackFromSlice[T any](c []T) *Stack[T] {
	return &Stack[T]{objects: list.NewArrayListFromSlice(c)}
}

// Len returns the length of s.
func (s *Stack[T]) Len() int {
	return s.objects.Len()
}

// IsEmpty returns a bool which indicates if s is empty or not.
func (s *Stack[T]) IsEmpty() bool {
	return s.objects.IsEmpty()
}

// Top returns the top element of s.
// The method returns false if s is empty.
func (s *Stack[T]) Top() (T, bool) {
	result, err := s.objects.Get(s.Len() - 1)
	if err != nil {
		return result, false
	}
	return result, true
}

// ToSlice returns a slice which contains all elements of s.
func (s *Stack[T]) ToSlice() []T {
	return s.objects.ToSlice()
}

// Push adds the elements e at the top of s.
func (s *Stack[T]) Push(e ...T) {
	s.objects.Add(e...)
}

// Pop removes an element from the top of s and returns the removed element.
// The method returns false if s is empty.
func (s *Stack[T]) Pop() (T, bool) {
	result, err := s.objects.Remove(s.Len() - 1)
	if err != nil {
		return result, false
	}
	return result, true
}

// Clear removes all element from s.
func (s *Stack[T]) Clear() {
	s.objects.Clear()
}

// Equal returns true if s and st are both stacks and their elements are equals.
// In any other case, it returns false.
//
// but the elements of s and the elements of st are equals, this method returns anyway true.
func (s *Stack[T]) Equal(st any) bool {
	stack, ok := st.(*Stack[T])
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
func (s *Stack[T]) Compare(st any) int {
	stack, ok := st.(*Stack[T])
	if ok && s != nil && stack != nil {
		return s.objects.Compare(list.NewArrayListFromStructure[T](stack))
	}
	return -2
}

// Hash returns the hash code of s.
func (s *Stack[T]) Hash() uint64 {
	return s.objects.Hash()
}

// String returns a rapresentation of s in the form of a string.
func (s *Stack[T]) String() string {
	check := reflect.TypeOf(new(T)).String()
	if s.IsEmpty() {
		return fmt.Sprintf("Stack[%v][%d, ]", check[1:], s.objects.Len())
	}
	element, _ := s.Top()
	return fmt.Sprintf("Stack[%v][%d, %v]", check[1:], s.objects.Len(), element)
}
