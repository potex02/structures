package queue

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/potex02/structures"
)

// DoubleLinkedQueue provides a generic double queue.
// The queue is implemented through a series of linked [structures.Entry].
//
// It implements the interface [DoubleQueue].
type DoubleLinkedQueue[T any] struct {
	// contains filtered or unexported fields
	head *structures.Entry[T]
	tail *structures.Entry[T]
	len  int
}

// NewDoubleLinkedQueue returns a new [DoubleLinkedQueue] containing the elements c.
// The head of the queue is the first element of c, while the tail is the last element.
//
// if no argument is passed, it will be created an empty [DoubleLinkedQueue].
func NewDoubleLinkedQueue[T any](c ...T) *DoubleLinkedQueue[T] {

	return NewDoubleLinkedQueueFromSlice(c)

}

// NewDoubleLinkedQueueFromSlice returns a new [DoubleLinkedQueue] containing the elements of slice c.
func NewDoubleLinkedQueueFromSlice[T any](c []T) *DoubleLinkedQueue[T] {

	queue := &DoubleLinkedQueue[T]{head: nil, tail: nil, len: 0}
	if len(c) != 0 {

		queue.PushTail(c...)

	}
	return queue

}

// Len returns the length of q.
func (q *DoubleLinkedQueue[T]) Len() int {

	return q.len

}

// IsEmpty returns a bool which indicate if q is empty or not.
func (q *DoubleLinkedQueue[T]) IsEmpty() bool {

	return q.len == 0

}

// Head returns the head element of q.
// If q is empty, the method returns an error.
func (q *DoubleLinkedQueue[T]) Head() (T, error) {

	if q.IsEmpty() {

		var result T

		return result, errors.New("Empty queue")

	}
	return q.head.Element(), nil

}

// Tail returns the tail element element of q.
// If q is empty, the method returns an error.
func (q *DoubleLinkedQueue[T]) Tail() (T, error) {

	if q.IsEmpty() {

		var result T

		return result, errors.New("Empty queue")

	}
	return q.tail.Element(), nil

}

// ToSLice returns a slice which contains all elements of q.
func (q *DoubleLinkedQueue[T]) ToSlice() []T {

	slice := make([]T, q.len)
	j := 0
	for i := q.head; i != nil; i = i.Next() {

		slice[j] = i.Element()
		j++

	}
	return slice

}

// PushHead adds the elements e at the head of q.
func (q *DoubleLinkedQueue[T]) PushHead(e ...T) {

	if len(e) == 0 {

		return

	}
	elements := make([]T, len(e))
	for i := 0; i != len(e); i++ {

		elements[i] = e[len(e)-i-1]

	}
	first, last := structures.NewEntrySlice(elements)
	if q.len == 0 {

		q.head = first
		q.tail = last
		q.len = len(e)
		return

	}
	last.SetNext(q.head)
	q.head.SetPrev(last)
	q.head = first
	q.len += len(e)

}

// PushTail adds the elements e at the tail of q.
func (q *DoubleLinkedQueue[T]) PushTail(e ...T) {

	if len(e) == 0 {

		return

	}
	first, last := structures.NewEntrySlice(e)
	if q.len == 0 {

		q.head = first
		q.tail = last
		q.len = len(e)
		return

	}
	first.SetPrev(q.tail)
	q.tail.SetNext(first)
	q.tail = last
	q.len += len(e)

}

// PopHead removes an element from the head of q and returns the removed element.
// If q is empty, the method returns an error.
func (q *DoubleLinkedQueue[T]) PopHead() (T, error) {

	var result T

	if q.IsEmpty() {

		return result, errors.New("Empty queue")

	}
	result = q.head.Element()
	if q.len > 1 {

		q.head = q.head.Next()
		q.head.SetPrev(nil)
		q.len--

	} else {

		q.Clear()

	}
	return result, nil

}

// PopTail removes an element from the tail of q and returns the removed element.
// If q is empty, the method returns an error.
func (q *DoubleLinkedQueue[T]) PopTail() (T, error) {

	var result T

	if q.IsEmpty() {

		return result, errors.New("Empty queue")

	}
	result = q.tail.Element()
	if q.len > 1 {

		q.tail = q.tail.Prev()
		q.tail.SetNext(nil)
		q.len--

	} else {

		q.Clear()

	}
	return result, nil

}

// Clear removes all element from q.
func (q *DoubleLinkedQueue[T]) Clear() {

	q.head = nil
	q.tail = nil
	q.len--

}

// Equals returns true if q and st are both double queues and their elements are equals.
// In any other case, it returns false.
//
// Equals does not take into account the effective type of st. This means that if st is a [DoubleArrayQueue],
// but the elements of q and the elements of st are equals, this method returns anyway true.
func (q *DoubleLinkedQueue[T]) Equals(st structures.Structure[T]) bool {

	queue, ok := st.(DoubleQueue[T])
	return ok && reflect.DeepEqual(q.ToSlice(), queue.ToSlice())

}

// String returns a rapresentation of q in the form of a string.
func (q *DoubleLinkedQueue[T]) String() string {

	if q.IsEmpty() {

		return fmt.Sprintf("DoubleLinkedQueue[%T][%d, ]", *new(T), q.len)

	}
	return fmt.Sprintf("DoubleLinkedQueue[%T][%d, %v %v]", *new(T), q.len, q.head.Element(), q.tail.Element())

}
