// Package stack implements dynamic stacks.
package stack

import "github.com/potex02/structures"

// Stack provides all methods to use a generic stack.
// A stack contains all the methods of [structures.Structure].
type Stack[T any] interface {
	structures.Structure[T]
	// Top returns the top element of the stack.
	// If the stack is empty, the method returns an error.
	Top() (T, error)
	// Push adds the elements e at the top of the stack.
	Push(e ...T)
	// Pop removes an element from the top of the stack and returns the removed element.
	// If the stack is empty, the method returns an error.
	Pop() (T, error)
}
