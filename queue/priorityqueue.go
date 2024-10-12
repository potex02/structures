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

var _ structures.Structure[wrapper.Int] = NewPriorityQueue[wrapper.Int]()
var _ Queue[wrapper.Int] = NewPriorityQueue[wrapper.Int]()

// PriorityQueue provides a generic priority queue which mantains the order of the elements.
// It is implemented through an [tree.BinaryTree].
//
// It implements the interface [Queue].
type PriorityQueue[T util.Comparer] struct {
	// contains filtered or unexported fields
	objects *tree.BinaryTree[T]
}

// NewPriorityQueue returns a new [PriorityQueue] containing the elements c.
//
// if no argument is passed, it will be created an empty [PriorityQueue].
func NewPriorityQueue[T util.Comparer](c ...T) *PriorityQueue[T] {
	return NewPriorityQueueFromSlice(c)
}

// NewPriorityQueueFromSlice returns a new [PriorityQueue] containing the elements of slice c.
func NewPriorityQueueFromSlice[T util.Comparer](c []T) *PriorityQueue[T] {
	return &PriorityQueue[T]{objects: tree.NewBinaryTreeFromSlice(c)}
}

// Len returns the length of q.
func (q *PriorityQueue[T]) Len() int {
	return q.objects.Len()
}

// IsEmpty returns a bool which indicates if q is empty or not.
func (q *PriorityQueue[T]) IsEmpty() bool {
	return q.objects.IsEmpty()
}

// Head returns the max element of q.
// The method returns false if q is empty.
func (q *PriorityQueue[T]) Head() (T, bool) {
	if q.IsEmpty() {

		var result T

		return result, false
	}
	return q.objects.Root().Max().Element(), true
}

// Tail returns the min element element of q.
// The method returns false if q is empty.
func (q *PriorityQueue[T]) Tail() (T, bool) {
	if q.IsEmpty() {

		var result T

		return result, false
	}
	return q.objects.Root().Min().Element(), true
}

// ToSlice returns a slice which contains all elements of q.
func (q *PriorityQueue[T]) ToSlice() []T {
	slice := q.objects.ToSlice()
	sort.SliceStable(slice, func(i, j int) bool {
		return i > j
	})
	return slice
}

// Push adds the elements e at q.
func (q *PriorityQueue[T]) Push(e ...T) {
	q.objects.Add(e...)
}

// Pop removes the max element from q and returns the removed element.
// The method returns false if q is empty.
func (q *PriorityQueue[T]) Pop() (T, bool) {
	if q.IsEmpty() {

		var result T

		return result, false
	}
	result := q.objects.Root().Max().Element()
	q.objects.Remove(result)
	return result, true
}

// Clear removes all element from q.
func (q *PriorityQueue[T]) Clear() {
	q.objects.Clear()
}

// Equal returns true if q and st are both queues and their elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is an [ArrayQueue] or a [LinkedQueue],
// but the elements of q and the elements of st are equals, this method returns anyway true.
func (q *PriorityQueue[T]) Equal(st any) bool {
	queue, ok := st.(Queue[T])
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
func (q *PriorityQueue[T]) Compare(st any) int {
	queue, ok := st.(Queue[T])
	if ok && q != nil && queue != nil {
		return list.NewArrayListFromStructure[T](q).Compare(list.NewArrayListFromStructure[T](queue))
	}
	return -2
}

// Hash returns the hash code of q.
func (q *PriorityQueue[T]) Hash() string {
	check := reflect.TypeOf(new(T)).String()
	head, _ := q.Head()
	tail, _ := q.Tail()
	return fmt.Sprintf("%v%v%v", check[1:], head, tail)
}

// String returns a rapresentation of q in the form of a string.
func (q *PriorityQueue[T]) String() string {
	check := reflect.TypeOf(new(T)).String()
	if q.IsEmpty() {
		return fmt.Sprintf("PriorityQueue[%v][%d, ]", check[1:], q.objects.Len())
	}
	head, _ := q.Head()
	tail, _ := q.Tail()
	return fmt.Sprintf("PriorityQueue[%v][%d, %v %v]", check[1:], q.objects.Len(), head, tail)
}
