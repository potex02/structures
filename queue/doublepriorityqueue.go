package queue

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
	"github.com/potex02/structures/tree"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ structures.Structure[wrapper.Int] = NewDoublePriorityQueue[wrapper.Int]()
var _ BaseDoubleQueue[wrapper.Int] = NewDoublePriorityQueue[wrapper.Int]()

// DoublePriorityQueue provides a generic double queue which mantains the order of the elements.
// It is implemented through an [tree.BinaryTree].
//
// It implements the interface [BaseDoublePriorityQueue].
type DoublePriorityQueue[T util.Comparer] struct {
	// contains filtered or unexported fields
	objects *tree.BinaryTree[T]
}

// NewDoublePriorityQueue returns a new [DoublePriorityQueue] containing the elements c.
// The head of the queue is the maximum element of c, while the tail is the minimum element.
//
// if no argument is passed, it will be created an empty [DoublePriorityQueue].
func NewDoublePriorityQueue[T util.Comparer](c ...T) *DoublePriorityQueue[T] {
	return NewDoublePriorityQueueFromSlice(c)
}

// NewDoublePriorityQueueFromSlice returns a new [DoublePriorityQueue] containing the elements of slice c.
func NewDoublePriorityQueueFromSlice[T util.Comparer](c []T) *DoublePriorityQueue[T] {
	return &DoublePriorityQueue[T]{objects: tree.NewBinaryTreeFromSlice(c)}
}

// Len returns the length of q.
func (q *DoublePriorityQueue[T]) Len() int {
	return q.objects.Len()
}

// IsEmpty returns a bool which indicates if q is empty or not.
func (q *DoublePriorityQueue[T]) IsEmpty() bool {
	return q.objects.IsEmpty()
}

// Head returns the maximum element of q.
// The method returns false if q is empty.
func (q *DoublePriorityQueue[T]) Head() (T, bool) {
	if q.IsEmpty() {

		var result T

		return result, false
	}
	return q.objects.Root().Max().Element(), true
}

// Tail returns the minimun element element of q.
// The method returns false if q is empty.
func (q *DoublePriorityQueue[T]) Tail() (T, bool) {
	if q.IsEmpty() {

		var result T

		return result, false
	}
	return q.objects.Root().Min().Element(), true
}

// ToSlice returns a slice which contains all elements of q.
func (q *DoublePriorityQueue[T]) ToSlice() []T {
	slice := q.objects.ToSlice()
	sort.SliceStable(slice, func(i, j int) bool {
		return i > j
	})
	return slice
}

// Push adds the elements e at q.
func (q *DoublePriorityQueue[T]) Push(e ...T) {
	q.objects.Add(e...)
}

// PushHead adds the elements e at q.
//
// Since the queue is ordered, it is the same of [DoublePriorityQueue.Push].
func (q *DoublePriorityQueue[T]) PushHead(e ...T) {
	q.Push(e...)
}

// PushHead adds the elements e at q.
//
// Since the queue is ordered, it is the same of [DoublePriorityQueue.Push].
func (q *DoublePriorityQueue[T]) PushTail(e ...T) {
	q.Push(e...)
}

// PopHead removes the maximun element from q and returns the removed element.
// The method returns false if q is empty.
func (q *DoublePriorityQueue[T]) PopHead() (T, bool) {
	if q.IsEmpty() {

		var result T

		return result, false
	}
	result := q.objects.Root().Max().Element()
	q.objects.Remove(result)
	return result, true
}

// PopTail removes the minimum element from q and returns the removed element.
// The method returns false if q is empty.
func (q *DoublePriorityQueue[T]) PopTail() (T, bool) {
	if q.IsEmpty() {

		var result T

		return result, false
	}
	result := q.objects.Root().Min().Element()
	q.objects.Remove(result)
	return result, true
}

// Clear removes all element from q.
func (q *DoublePriorityQueue[T]) Clear() {
	q.objects.Clear()
}

// Equal returns true if q and st are both double queues and their elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is a [DoubleQueue],
// but the elements of q and the elements of st are equals, this method returns anyway true.
func (q *DoublePriorityQueue[T]) Equal(st any) bool {
	queue, ok := st.(BaseDoubleQueue[T])
	if ok && q != nil && queue != nil {
		return list.NewArrayListFromStructure[T](q).Equal(list.NewArrayListFromStructure[T](queue))
	}
	return false
}

// Compare returns 0 if q and st are equals,
// -1 if q is shorten than st,
// 1 if q is longer than st,
// -2 if st is not a [BaseDoubleQueue] or if one between q and st is nil.
//
// If q and st have the same length, the result is the comparison
// between the first different element of the two queues,
// otherwhise the result is 0.
func (q *DoublePriorityQueue[T]) Compare(st any) int {
	queue, ok := st.(BaseDoubleQueue[T])
	if ok && q != nil && queue != nil {
		return q.objects.Compare(list.NewArrayListFromStructure[T](queue))
	}
	return -2
}

// Hash returns the hash code of q.
func (q *DoublePriorityQueue[T]) Hash() uint64 {
	return q.objects.Hash()
}

// String returns a rapresentation of q in the form of a string.
func (q *DoublePriorityQueue[T]) String() string {
	check := reflect.TypeOf(new(T)).String()
	if q.IsEmpty() {
		return fmt.Sprintf("DoublePriorityQueue[%v][%d, ]", check[1:], q.Len())
	}
	head, _ := q.Head()
	tail, _ := q.Tail()
	return fmt.Sprintf("DoublePriorityQueue[%v][%d, %v %v]", check[1:], q.Len(), head, tail)
}
