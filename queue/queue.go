// Package queue implements single and double dynamic queues.
package queue

import "github.com/potex02/structures"

// Queue provides all methods to use a generic queue.
// A queue contains all the methods of [structures.Structure].
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
type Queue[T any] interface {
	structures.Structure[T]
	// Head returns the head element of the queue.
	// The method returns false if the queue is empty.
	Head() (T, bool)
	// Tail returns the tail element of the queue.
	// The method returns false if the queue is empty.
	Tail() (T, bool)
	// Push adds the elements e at the tail of the queue.
	Push(e ...T)
	// Pop removes an element from the head of the queue and returns the removed element.
	// The method returns false if the queue is empty.
	Pop() (T, bool)
}

// DoubleQueue provides all methods to use a generic double queue.
// A double queue contains all the methods of [structures.Structure].
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
type DoubleQueue[T any] interface {
	structures.Structure[T]
	// Head returns the head element of the queue.
	// The method returns false if the queue is empty.
	Head() (T, bool)
	// Tail returns the tail element of the queue.
	// The method returns false if the queue is empty.
	Tail() (T, bool)
	// PushHead adds the elements e at the head of the queue.
	PushHead(e ...T)
	// PushTail adds the elements e at the tail of the queue.
	PushTail(e ...T)
	// PopHead removes an element from the head of the queue and returns the removed element.
	// The method returns false if the queue is empty.
	PopHead() (T, bool)
	// PopTail removes an element from the tail of the queue and returns the removed element.
	// The method returns false if the queue is empty.
	PopTail() (T, bool)
}
