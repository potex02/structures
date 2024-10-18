package queue

import (
	"fmt"
	"hash/fnv"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
)

var _ structures.Structure[int] = NewQueue[int]()
var _ BaseQueue[int] = NewQueue[int]()

// Queue provides a generic FIFO structure.
// The queue is implemented through a series of linked [structures.Entry].
//
// It implements the interface [BaseQueue].
type Queue[T any] struct {
	head *structures.Entry[T]
	tail *structures.Entry[T]
	len  int
}

// NewQueue returns a new [Queue] containing the elements c.
// The head of the queue is the first element of c, while the tail is the last element.
//
// if no argument is passed, it will be created an empty [Queue].
func NewQueue[T any](c ...T) *Queue[T] {
	return NewQueueFromSlice(c)
}

// NewQueueFromSlice returns a new [Queue] containing the elements of slice c.
func NewQueueFromSlice[T any](c []T) *Queue[T] {
	queue := &Queue[T]{head: nil, tail: nil, len: 0}
	if len(c) != 0 {
		queue.Push(c...)
	}
	return queue
}

// Len returns the length of q.
func (q *Queue[T]) Len() int {
	return q.len
}

// IsEmpty returns a bool which indicates if q is empty or not.
func (q *Queue[T]) IsEmpty() bool {
	return q.len == 0
}

// Head returns the head element of q.
// The method returns false if q is empty.
func (q *Queue[T]) Head() (T, bool) {
	if q.IsEmpty() {

		var result T

		return result, false
	}
	return q.head.Element(), true
}

// Tail returns the tail element element of q.
// The method returns false if q is empty.
func (q *Queue[T]) Tail() (T, bool) {
	if q.IsEmpty() {

		var result T

		return result, false
	}
	return q.tail.Element(), true
}

// ToSlice returns a slice which contains all elements of q.
func (q *Queue[T]) ToSlice() []T {
	slice := make([]T, q.len)
	j := 0
	for i := q.head; i != nil; i = i.Next() {
		slice[j] = i.Element()
		j++
	}
	return slice
}

// Push adds the elements e at the tail of q.
func (q *Queue[T]) Push(e ...T) {
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
// The method returns false if q is empty.
func (q *Queue[T]) Pop() (T, bool) {

	var result T

	if q.IsEmpty() {
		return result, false
	}
	result = q.head.Element()
	if q.len > 1 {
		q.head = q.head.Next()
		q.len--
	} else {
		q.Clear()
	}
	return result, true
}

// Clear removes all element from q.
func (q *Queue[T]) Clear() {
	q.head = nil
	q.tail = nil
	q.len = 0
}

// Equal returns true if q and st are both queues and their elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is an [ArrayQueue] or a [PriorityQueue],
// but the elements of q and the elements of st are equals, this method returns anyway true.
func (q *Queue[T]) Equal(st any) bool {
	queue, ok := st.(BaseQueue[T])
	if ok && q != nil && queue != nil {
		return list.NewArrayListFromStructure[T](q).Equal(list.NewArrayListFromStructure[T](queue))
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
func (q *Queue[T]) Compare(st any) int {
	queue, ok := st.(BaseQueue[T])
	if ok && q != nil && queue != nil {
		return list.NewArrayListFromStructure[T](q).Compare(list.NewArrayListFromStructure[T](queue))
	}
	return -2
}

// Hash returns the hash code of q.
func (q *Queue[T]) Hash() uint64 {
	h := fnv.New64()
	for i := q.head; i != nil; i = i.Next() {
		h.Write([]byte(fmt.Sprintf("%v", i.Hash())))
	}
	return h.Sum64()
}

// String returns a rapresentation of q in the form of a string.
func (q *Queue[T]) String() string {
	check := reflect.TypeOf(new(T)).String()
	if q.IsEmpty() {
		return fmt.Sprintf("Queue[%v][%d, ]", check[1:], q.len)
	}
	return fmt.Sprintf("Queue[%v][%d, %v %v]", check[1:], q.len, q.head.Element(), q.tail.Element())
}
