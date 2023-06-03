package queue

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
)

var _ structures.Structure[int] = NewArrayQueue[int]()
var _ Queue[int] = NewArrayQueue[int]()

// ArrayQueue provides a generic single queue implemented through an [list.ArrayList].
//
// It implements the interface [Queue].
type ArrayQueue[T any] struct {
	// contains filtered or unexported fields
	objects list.List[T]
}

// NewArrayQueue returns a new [ArrayQueue] containing the elements c.
// The head of the queue is the first element of c, while the tail is the last element.
//
// if no argument is passed, it will be created an empty [ArrayQueue].
func NewArrayQueue[T any](c ...T) *ArrayQueue[T] {

	return NewArrayQueueFromSlice(c)

}

// NewArrayQueueFromSlice returns a new [ArrayQueue] containing the elements of slice c.
func NewArrayQueueFromSlice[T any](c []T) *ArrayQueue[T] {

	return &ArrayQueue[T]{objects: list.NewArrayListFromSlice(c)}

}

// Len returns the length of q.
func (q *ArrayQueue[T]) Len() int {

	return q.objects.Len()

}

// IsEmpty returns a bool which indicate if q is empty or not.
func (q *ArrayQueue[T]) IsEmpty() bool {

	return q.objects.IsEmpty()

}

// Head returns the head element of q.
// If q is empty, the method returns an error.
func (q *ArrayQueue[T]) Head() (T, error) {

	result, err := q.objects.Get(0)
	if err != nil {

		return result, errors.New("Empty queue")

	}
	return result, err

}

// Tail returns the tail element element of q.
// If q is empty, the method returns an error.
func (q *ArrayQueue[T]) Tail() (T, error) {

	result, err := q.objects.Get(q.Len() - 1)
	if err != nil {

		return result, errors.New("Empty queue")

	}
	return result, err

}

// ToSLice returns a slice which contains all elements of q.
func (q *ArrayQueue[T]) ToSlice() []T {

	return q.objects.ToSlice()

}

// Push adds the elements e at the tail of q.
func (q *ArrayQueue[T]) Push(e ...T) {

	q.objects.Add(e...)

}

// Pop removes an element from the head of q and returns the removed element.
// If q is empty, the method returns an error.
func (q *ArrayQueue[T]) Pop() (T, error) {

	result, err := q.objects.Remove(0)
	if err != nil {

		return result, errors.New("Empty queue")

	}
	return result, err

}

// Clear removes all element from q.
func (q *ArrayQueue[T]) Clear() {

	q.objects.Clear()

}

// Equal returns true if q and st are both queues and their elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is a [LinkedQueue],
// but the elements of q and the elements of st are equals, this method returns anyway true.
func (q *ArrayQueue[T]) Equal(st any) bool {

	queue, ok := st.(Queue[T])
	if ok && q != nil && queue != nil {

		return q.objects.Equal(list.NewArrayListFromStructure[T](queue))

	}
	return false

}

// Compare returns 0 if q and st are equals,
// -1 if q is shorten than st,
// 1 if q is longer than st,
// -2 if st is not a [Queue] or if one between q and st is nil.
//
// If q and st have the same length, the result is the comparison
// between the first different element of the two queues if T implemets [util.Comparer],
// otherwhise the result is 0.
func (q *ArrayQueue[T]) Compare(st any) int {

	queue, ok := st.(Queue[T])
	if ok && q != nil && queue != nil {

		return q.objects.Compare(list.NewArrayListFromStructure[T](queue))

	}
	return -2

}

// Hash returns the hash code of q.
func (q *ArrayQueue[T]) Hash() string {

	check := reflect.TypeOf(new(T)).String()
	head, _ := q.Head()
	tail, _ := q.Tail()
	return fmt.Sprintf("%v%v%v", check[1:], head, tail)

}

// String returns a rapresentation of q in the form of a string.
func (q *ArrayQueue[T]) String() string {

	check := reflect.TypeOf(new(T)).String()
	if q.IsEmpty() {

		return fmt.Sprintf("ArrayQueue[%T][%d, ]", check[1:], q.objects.Len())

	}
	head, _ := q.Head()
	tail, _ := q.Tail()
	return fmt.Sprintf("ArrayQueue[%T][%d, %v %v]", check[1:], q.objects.Len(), head, tail)

}
