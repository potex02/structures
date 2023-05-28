package queue

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
)

// DoubleArrayQueue provides a generic double queue implemented through an [list.ArrayList].
//
// It implements the interface [DoubleQueue].
type DoubleArrayQueue[T any] struct {
	// contains filtered or unexported fields
	objects list.List[T]
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

	return &DoubleArrayQueue[T]{objects: list.NewArrayListFromSlice(c)}

}

// Len returns the length of q.
func (q *DoubleArrayQueue[T]) Len() int {

	return q.objects.Len()

}

// IsEmpty returns a bool which indicate if q is empty or not.
func (q *DoubleArrayQueue[T]) IsEmpty() bool {

	return q.objects.IsEmpty()

}

// Head returns the head element of q.
// If q is empty, the method returns an error.
func (q *DoubleArrayQueue[T]) Head() (T, error) {

	result, err := q.objects.Get(0)
	if err != nil {

		return result, errors.New("Empty queue")

	}
	return result, err

}

// Tail returns the tail element element of q.
// If q is empty, the method returns an error.
func (q *DoubleArrayQueue[T]) Tail() (T, error) {

	result, err := q.objects.Get(q.Len() - 1)
	if err != nil {

		return result, errors.New("Empty queue")

	}
	return result, err

}

// ToSLice returns a slice which contains all elements of q.
func (q *DoubleArrayQueue[T]) ToSlice() []T {

	return q.objects.ToSlice()

}

// PushHead adds the elements e at the head of q.
func (q *DoubleArrayQueue[T]) PushHead(e ...T) {

	elements := make([]T, len(e))
	for i := 0; i != len(e); i++ {

		elements[i] = e[len(e)-i-1]

	}
	q.objects.AddAtIndex(0, elements...)

}

// PushTail adds the elements e at the tail of q.
func (q *DoubleArrayQueue[T]) PushTail(e ...T) {

	q.objects.Add(e...)

}

// PopHead removes an element from the head of q and returns the removed element.
// If q is empty, the method returns an error.
func (q *DoubleArrayQueue[T]) PopHead() (T, error) {

	result, err := q.objects.Remove(0)
	if err != nil {

		return result, errors.New("Empty queue")

	}
	return result, err

}

// PopTail removes an element from the tail of q and returns the removed element.
// If q is empty, the method returns an error.
func (q *DoubleArrayQueue[T]) PopTail() (T, error) {

	result, err := q.objects.Remove(q.Len() - 1)
	if err != nil {

		return result, errors.New("Empty queue")

	}
	return result, err

}

// Clear removes all element from q.
func (q *DoubleArrayQueue[T]) Clear() {

	q.objects.Clear()

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

		return fmt.Sprintf("DoubleArrayQueue[%T][%d, ]", *new(T), q.objects.Len())

	}
	head, _ := q.Head()
	tail, _ := q.Tail()
	return fmt.Sprintf("DoubleArrayQueue[%T][%d, %v %v]", *new(T), q.objects.Len(), head, tail)

}
