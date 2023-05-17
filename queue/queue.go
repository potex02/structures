// Package queue implements dinamic queues.
package queue

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/potex02/structures"
)

// Queue provides a generic single queue implemented with a slice.
type Queue[T any] struct {
	// contains filtered or unexported fields
	objects []T
}

// New returns a new empty [Queue].
func New[T any]() *Queue[T] {

	return &Queue[T]{}

}

// NewFromElements is a wrapper for NewFromSlice(c).
func NewFromElements[T any](c ...T) *Queue[T] {

	return NewFromSlice(c)

}

// NewFromSlice returns a new [Queue] containing the elements of slice c.
func NewFromSlice[T any](c []T) *Queue[T] {

	return &Queue[T]{c}

}

// Len returns the length of q.
func (q *Queue[T]) Len() int {

	return len(q.objects)

}

// IsEmpty returns a bool which indicate if q is empty or not.
func (q *Queue[T]) IsEmpty() bool {

	return len(q.objects) == 0

}

// Head returns a pointer to the first element of q.
// If q is empty, the method returns nil.
func (q *Queue[T]) Head() *T {

	if q.IsEmpty() {

		return nil

	}
	return &q.objects[0]

}

// Head returns a pointer to the last element of q.
// If q is empty, the method returns nil.
func (q *Queue[T]) Tail() *T {

	if q.IsEmpty() {

		return nil

	}
	return &q.objects[len(q.objects)-1]

}

// ToSLice returns a slice which contains all elements of q.
func (q *Queue[T]) ToSlice() []T {

	slice := make([]T, len(q.objects))
	copy(slice, q.objects)
	return slice

}

// Add adds the element e at the tail of q.
func (q *Queue[T]) Add(e T) {

	q.objects = append(q.objects, e)

}

// Remove removes an element from the head of q and returns the removed element.
// It returns an error is q is empty.
func (q *Queue[T]) Remove() (T, error) {

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

// Clear removes all element from q.
func (q *Queue[T]) Clear() {

	q.objects = []T{}

}

// Equals returns true if q and st are both queues and their elements are equals.
// In any other case, it returns false.
func (q *Queue[T]) Equals(st structures.Structure[T]) bool {

	queue, ok := st.(*Queue[T])
	return ok && reflect.DeepEqual(q.ToSlice(), queue.ToSlice())

}

// String returns a rapresentation of q in the form of a string.
func (q *Queue[T]) String() string {

	if q.IsEmpty() {

		return fmt.Sprintf("Queue[%T][%d, %v]", *new(T), len(q.objects), nil)

	}
	return fmt.Sprintf("Queue[%T][%d, %v]", *new(T), len(q.objects), q.objects[0])

}
