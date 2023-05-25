package queue

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/potex02/structures"
)

// DoubleArrayQueue provides a generic double queue implemented with a slice.
//
// It implements the interface [DoubleQueue].
type DoubleArrayQueue[T any] struct {
	// contains filtered or unexported fields
	objects []T
}

// NewDoubleArrayQueue returns a new [DoubleArrayQueue] containing the elements c.
// The head of the queue is the first element of c, while the tail is the last element.
//
// if no argument is passed, it will be created an empty [DoubleArrayQueue].
func NewDoubleArrayQueue[T any](c ...T) *DoubleArrayQueue[T] {

	return NewDoubleArrayQueueFromSlice(c)

}

// NewDoubleArrayQueueFromSlice returns a new [DoubleArrayQueue] containing the elements of slice c.
func NewDoubleArrayQueueFromSlice[T any](c []T) *DoubleArrayQueue[T] {

	return &DoubleArrayQueue[T]{objects: c}

}

// Len returns the length of q.
func (q *DoubleArrayQueue[T]) Len() int {

	return len(q.objects)

}

// IsEmpty returns a bool which indicate if q is empty or not.
func (q *DoubleArrayQueue[T]) IsEmpty() bool {

	return len(q.objects) == 0

}

// Head returns the head element of q.
// If q is empty, the method returns an error.
func (q *DoubleArrayQueue[T]) Head() (T, error) {

	if q.IsEmpty() {

		var result T

		return result, errors.New("Empty queue")

	}
	return q.objects[0], nil

}

// Tail returns the tail element element of q.
// If q is empty, the method returns an error.
func (q *DoubleArrayQueue[T]) Tail() (T, error) {

	if q.IsEmpty() {

		var result T

		return result, errors.New("Empty queue")

	}
	return q.objects[len(q.objects)-1], nil

}

// ToSLice returns a slice which contains all elements of q.
func (q *DoubleArrayQueue[T]) ToSlice() []T {

	slice := make([]T, len(q.objects))
	copy(slice, q.objects)
	return slice

}

// PushHead adds the elements e at the head of q.
func (q *DoubleArrayQueue[T]) PushHead(e ...T) {

	q.objects = append(e, q.objects...)

}

// PushTail adds the elements e at the tail of q.
func (q *DoubleArrayQueue[T]) PushTail(e ...T) {

	q.objects = append(q.objects, e...)

}

// PopHead removes an element from the head of q and returns the removed element.
// If q is empty, the method returns an error.
func (q *DoubleArrayQueue[T]) PopHead() (T, error) {

	var result T

	if q.IsEmpty() {

		return result, errors.New("Empty queue")

	}
	result = q.objects[0]
	if len(q.objects) > 1 {

		q.objects = q.objects[1:]

	} else {

		q.Clear()

	}
	return result, nil

}

// PopTail removes an element from the tail of q and returns the removed element.
// If q is empty, the method returns an error.
func (q *DoubleArrayQueue[T]) PopTail() (T, error) {

	var result T

	if q.IsEmpty() {

		return result, errors.New("Empty queue")

	}
	result = q.objects[len(q.objects)-1]
	if len(q.objects) > 1 {

		q.objects = q.objects[:len(q.objects)-1]

	} else {

		q.Clear()

	}
	return result, nil

}

// Clear removes all element from q.
func (q *DoubleArrayQueue[T]) Clear() {

	q.objects = []T{}

}

// Equals returns true if q and st are both double queues and their elements are equals.
// In any other case, it returns false.
//
// Equals does not take into account the effective type of st. This means that if st is a [DoubleLinkedQueue],
// but the elements of q and the elements of st are equals, this method returns anyway true.
func (q *DoubleArrayQueue[T]) Equals(st structures.Structure[T]) bool {

	queue, ok := st.(DoubleQueue[T])
	return ok && reflect.DeepEqual(q.ToSlice(), queue.ToSlice())

}

// String returns a rapresentation of q in the form of a string.
func (q *DoubleArrayQueue[T]) String() string {

	if q.IsEmpty() {

		return fmt.Sprintf("DoubleArrayQueue[%T][%d, ]", *new(T), len(q.objects))

	}
	return fmt.Sprintf("DoubleArrayQueue[%T][%d, %v %v]", *new(T), len(q.objects), q.objects[0], q.objects[len(q.objects)-1])

}
