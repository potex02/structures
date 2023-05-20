// Package queue implements dinamic queues.
package queue

import "github.com/potex02/structures"

//Queue provides all methods to use a generic queue.
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
