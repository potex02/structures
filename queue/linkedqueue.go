package queue

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/potex02/structures"
)

// ArrayQueue provides a generic single queue implemented.
// The queue is implemented through a series of linked [structures.Entry].
//
// It implements the interface [Queue].
type LinkedQueue[T any] struct {
	head *structures.Entry[T]
	tail *structures.Entry[T]
	len  int
}

// NewLinkedQueue returns a new [LinkedQueue] containing the elements c.
// The head of the queue is the first element of c, while the tail is the last element.
//
// if no argument is passed, it will be created an empty [LinkedQueue].
func NewLinkedQueue[T any](c ...T) *LinkedQueue[T] {

	return NewLinkedQueueFromSlice(c)

}

// NewLinkedQueueFromSlice returns a new [LinkedQueue] containing the elements of slice c.
func NewLinkedQueueFromSlice[T any](c []T) *LinkedQueue[T] {

	queue := &LinkedQueue[T]{head: nil, tail: nil, len: 0}
	if len(c) != 0 {

		queue.Push(c...)

	}
	return queue

}

// Len returns the length of q.
func (q *LinkedQueue[T]) Len() int {

	return q.len

}

// IsEmpty returns a bool which indicate if q is empty or not.
func (q *LinkedQueue[T]) IsEmpty() bool {

	return q.len == 0

}

// Head returns the head element of q.
// If q is empty, the method returns an error.
func (q *LinkedQueue[T]) Head() (T, error) {

	if q.IsEmpty() {

		var result T

		return result, errors.New("Empty queue")

	}
	return q.head.Element(), nil

}

// Tail returns the tail element element of q.
// If q is empty, the method returns an error.
func (q *LinkedQueue[T]) Tail() (T, error) {

	if q.IsEmpty() {

		var result T

		return result, errors.New("Empty queue")

	}
	return q.tail.Element(), nil

}

// ToSLice returns a slice which contains all elements of q.
func (q *LinkedQueue[T]) ToSlice() []T {

	slice := make([]T, q.len)
	j := 0
	for i := q.head; i != nil; i = i.Next() {

		slice[j] = i.Element()
		j++

	}
	return slice

}

// Push adds the elements e at the tail of q.
func (q *LinkedQueue[T]) Push(e ...T) {

	if len(e) == 0 {

		return

	}
	elements := make([]T, len(e))
	for i := 0; i != len(e); i++ {

		elements[i] = e[len(e)-i-1]

	}
	first, last := structures.NewEntrySliceSingle(elements)
	if q.len == 0 {

		q.head = first
		q.tail = last
		q.len = len(e)
		return

	}
	q.tail.SetNext(first)
	q.tail = last
	q.len += len(e)

}

// Pop removes an element from the head of q and returns the removed element.
// If q is empty, the method returns an error.
func (q *LinkedQueue[T]) Pop() (T, error) {

	var result T

	if q.IsEmpty() {

		return result, errors.New("Empty queue")

	}
	result = q.head.Element()
	if q.len > 1 {

		q.head = q.head.Next()
		q.len--

	} else {

		q.Clear()

	}
	return result, nil

}

// Clear removes all element from q.
func (q *LinkedQueue[T]) Clear() {

	q.head = nil
	q.tail = nil
	q.len = 0

}

// Equal returns true if q and st are both queues and their elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is an [ArrayQueue],
// but the elements of q and the elements of st are equals, this method returns anyway true.
func (q *LinkedQueue[T]) Equal(st structures.Structure[T]) bool {

	queue, ok := st.(Queue[T])
	return ok && reflect.DeepEqual(q.ToSlice(), queue.ToSlice())

}

// String returns a rapresentation of q in the form of a string.
func (q *LinkedQueue[T]) String() string {

	if q.IsEmpty() {

		return fmt.Sprintf("LinkedQueue[%T][%d, ]", *new(T), q.len)

	}
	return fmt.Sprintf("LinkedQueue[%T][%d, %v %v]", *new(T), q.len, q.head.Element(), q.tail.Element())

}
