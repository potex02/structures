package stack

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
)

var _ structures.Structure[int] = NewLinkedStack[int]()
var _ Stack[int] = NewLinkedStack[int]()

// ArrayStack provides a generic linked stack.
// The stack is implemented through a series of linked [structures.Entry].
//
// It implements the interface [Stack].
type LinkedStack[T any] struct {
	// contains filtered or unexported fields
	top *structures.Entry[T]
	len int
}

// NewLinkedStack returns a new [LinkedStack] containing the elements c.
// The top of the stack is the last element of c.
//
// if no argument is passed, it will be created an empty [LinkedStack].
func NewLinkedStack[T any](c ...T) *LinkedStack[T] {

	return NewLinkedStackFromSlice(c)

}

// NewLinkedStackFromSlice returns a new [LinkedStack] containing the elements of slice c.
// The top of the stack is the last element of c
func NewLinkedStackFromSlice[T any](c []T) *LinkedStack[T] {

	stack := &LinkedStack[T]{top: nil, len: 0}
	if len(c) != 0 {

		stack.Push(c...)

	}
	return stack

}

// Len returns the length of s.
func (s *LinkedStack[T]) Len() int {

	return s.len

}

// IsEmpty returns a bool which indicates if s is empty or not.
func (s *LinkedStack[T]) IsEmpty() bool {

	return s.len == 0

}

// Top returns the top element of s.
// If s is empty, the method returns an error.
func (s *LinkedStack[T]) Top() (T, error) {

	if s.IsEmpty() {

		var result T

		return result, errors.New("Empty stack")

	}
	return s.top.Element(), nil

}

// ToSLice returns a slice which contains all elements of s.
func (s *LinkedStack[T]) ToSlice() []T {

	slice := make([]T, s.len)
	j := 0
	for i := s.top; i != nil; i = i.Next() {

		slice[s.len-1-j] = i.Element()
		j++

	}
	return slice

}

// Push adds the elements e at the top of s.
func (s *LinkedStack[T]) Push(e ...T) {

	if len(e) == 0 {

		return

	}
	first, last := structures.NewEntrySliceSingle(e)
	last.SetNext(s.top)
	s.top = first
	s.len += len(e)

}

// Pop removes an element from the top of s and returns the removed element.
// If s is empty, the method returns an error.
func (s *LinkedStack[T]) Pop() (T, error) {

	var result T

	if s.IsEmpty() {

		return result, errors.New("Empty stack")

	}
	result = s.top.Element()
	if s.len > 1 {

		s.top = s.top.Next()
		s.len--

	} else {

		s.Clear()

	}
	return result, nil

}

// Clear removes all element from s.
func (s *LinkedStack[T]) Clear() {

	s.top = nil
	s.len = 0

}

// Equal returns true if s and st are both stacks and their elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is an [ArrayStack],
// but the elements of s and the elements of st are equals, this method returns anyway true.
func (s *LinkedStack[T]) Equal(st any) bool {

	stack, ok := st.(Stack[T])
	if ok && s != nil && stack != nil {

		return list.NewArrayListFromStructure[T](s).Equal(list.NewArrayListFromStructure[T](stack))

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
func (s *LinkedStack[T]) Compare(st any) int {

	stack, ok := st.(Stack[T])
	if ok && s != nil && stack != nil {

		return list.NewArrayListFromStructure[T](s).Compare(list.NewArrayListFromStructure[T](stack))

	}
	return -2

}

// Hash returns the hash code of s.
func (s *LinkedStack[T]) Hash() string {

	check := reflect.TypeOf(new(T)).String()
	top, _ := s.Top()
	return fmt.Sprintf("%v%v", check[1:], top)

}

// String returns a rapresentation of s in the form of a string.
func (s *LinkedStack[T]) String() string {

	check := reflect.TypeOf(new(T)).String()
	if s.IsEmpty() {

		return fmt.Sprintf("LinkedStack[%T][%d, ]", check[1:], s.len)

	}
	return fmt.Sprintf("LinkedStack[%T][%d, %v]", check[1:], s.len, s.top.Element())

}
