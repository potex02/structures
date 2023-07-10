// Package stack implements dynamic stacks.
package stack

import "github.com/potex02/structures"

// Stack provides all methods to use a generic stack.
// A stack contains all the methods of [structures.Structure].
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
type Stack[T any] interface {
	structures.Structure[T]
	// Top returns the top element of the stack.
	// The method returns false if the stack is empty.
	Top() (T, bool)
	// Push adds the elements e at the top of the stack.
	Push(e ...T)
	// Pop removes an element from the top of the stack and returns the removed element.
	// The method returns false if the stack is empty.
	Pop() (T, bool)
}
