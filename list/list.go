// Package list implements dynamic lists.
package list

import (
	"sort"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// List provides all methods to use a generic dynamic list.
// A list contains all the methods of [structures.Structure].
//
// A list is indexed starting from 0.
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
type List[T any] interface {
	structures.Structure[T]
	// Contains returns if e is present in the list.
	Contains(e T) bool
	// IndexOf returns the first position of e in the list.
	// If e is not present, the result is -1.
	IndexOf(e T) int
	// LastIndexOf returns the last position of e in the list.
	// If e is not present, the result is -1.
	LastIndexOf(e T) int
	// Get returns the elements at the specifies index.
	// It returns an error if the the index is out of bounds.
	Get(index int) (T, error)
	// Set sets the value of element at the specified index and returns the overwritten value.
	// It returns an error if the the index is out of bounds.
	Set(index int, e T) (T, error)
	// Add adds the elements e at the end if the list.
	Add(e ...T)
	// AddAtIndex adds the elements e at the specified index.
	// It returns an error if the the index is out of bounds.
	AddAtIndex(index int, e ...T) error
	// AddSlice adds the elements of e at the end of the list.
	AddSlice(e []T)
	// AddSliceAtIndex adds the elements of e at the specified index.
	// It returns an error if the the index is out of bounds.
	AddSliceAtIndex(index int, e []T) error
	// Remove removes the element at specified index and return the removed value.
	// It returns an error if the the index is out of bounds.
	Remove(index int) (T, error)
	// RemoveElement removes the element e from the list if it is presentt.
	// In that case, the method returns true, otherwhise it returns false.
	RemoveElement(e T) bool
	// Iter returns an [Iterator] which permits to iterate a [List].
	//
	//	for i := list.Iter(); !i.End(); i = i.Next() {
	//		element := i.Element()
	//		index := i.Index()
	//		// Code
	//	}
	Iter() Iterator[T]
	// Iter returns an [Iterator] which permits to iterate a [List] in reverse order.
	//
	//	for i := list.IterReverse(); !i.End(); i = i.Prev() {
	//		element := i.Element()
	//		index := i.Index()
	//		// Code
	//	}
	IterReverse() Iterator[T]
	// Copy returns a copy of the list.
	Copy() List[T]
}

// Sort returns a [List] which contains all elements of l that have been sorted.
//
// This function can be used only with types which implements the [util.Comparer] interface.
func Sort[T util.Comparer](l List[T]) List[T] {

	slice := l.ToSlice()
	sort.Slice(slice, func(i, j int) bool {

		return slice[i].Compare(slice[j]) < 0

	})
	switch l.(type) {

	case *ArrayList[T]:
		return NewArrayListFromSlice(slice)
	case *LinkedList[T]:
		return NewLinkedListFromSlice(slice)

	}
	return nil

}

// SortOrdered returns a [List] which contains all elements of l that have been sorted.
//
// This function can be used only with types which implements the [constraints.Ordered] interface.
func SortOrdered[T constraints.Ordered](l List[T]) List[T] {

	slice := l.ToSlice()
	slices.Sort(slice)
	switch l.(type) {

	case *ArrayList[T]:
		return NewArrayListFromSlice(slice)
	case *LinkedList[T]:
		return NewLinkedListFromSlice(slice)

	}
	return nil

}

// SortCustom returns a [List] which contains all elements of l that have been sorted.
//
// This function can be used to with any list and require a function to make the sorting.
// If the result of comparator is false, the the element in position i is placed before that in position j,
// otherwhise the opposite happens.
func SortCustom[T any](l List[T], comparator func(i T, j T) bool) List[T] {

	slice := l.ToSlice()
	slices.SortFunc(slice, comparator)
	switch l.(type) {

	case *ArrayList[T]:
		return NewArrayListFromSlice(slice)
	case *LinkedList[T]:
		return NewLinkedListFromSlice(slice)

	}
	return nil

}
func rangeCheck[T any](list List[T], index int) bool {

	return index >= 0 && index < list.Len()

}
