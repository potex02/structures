// Package queue implements single and double dinamic queues.
package queue

import "github.com/potex02/structures"

// Queue provides all methods to use a generic queue.
// A queue contains all the methods of [structures.Structure].
type Queue[T any] interface {
	structures.Structure[T]
	// Head returns the head element of the queue.
	// If the queue is empty, the method returns an error.
	Head() (T, error)
	// Tail returns the tail element of the queue.
	// If the queue is empty, the method returns an error.
	Tail() (T, error)
	// Push adds the elements e at the tail of the queue.
	Push(e ...T)
	// Pop removes an element from the head of the queue and returns the removed element.
	// If the queue is empty, the method returns an error.
	Pop() (T, error)
}

// DoubleQueue provides all methods to use a generic double queue.
// A double queue contains all the methods of [structures.Structure].
type DoubleQueue[T any] interface {
	structures.Structure[T]
	// Head returns the head element of the queue.
	// If the queue is empty, the method returns an error.
	Head() (T, error)
	// Tail returns the tail element of the queue.
	// If the queue is empty, the method returns an error.
	Tail() (T, error)
	// PushHead adds the elements e at the head of the queue.
	PushHead(e ...T)
	// PushTail adds the elements e at the tail of the queue.
	PushTail(e ...T)
	// PopHead removes an element from the head of the queue and returns the removed element.
	// If the queue is empty, the method returns an error.
	PopHead() (T, error)
	// PopTail removes an element from the tail of the queue and returns the removed element.
	// If the queue is empty, the method returns an error.
	PopTail() (T, error)
}
