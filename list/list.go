// Package list implements dynamic lists.
package list

import (
	"sort"

	"github.com/potex02/structures"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// List provides all methods to use a generic dynamic list.
// A list contains all the methods of [structures.Structure].
//
// A list is indexed starting from 0.
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
	// Iter returns a chan which permits to iterate a [List] with the range keyword.
	//
	//	for i := range l.Iter() {
	//		// code
	//	}
	//
	// This method can only be used to iterate a [List] if the index is not needed.
	// if you need to iterate a [List] with the index there are two options:
	//
	//	for i := 0; i < list.Len(); i++ {
	//		element, err := list.Get(i)
	//		// Code
	//	}
	//
	//	j := 0
	//	for i := range l.Iter() {
	//		// code
	//		j++
	//	}
	Iter() chan T
	// IterReverse returns a chan which permits to iterate a [List] in reverse order with the range keyword.
	//
	//	for i := range l.IterReverse() {
	//		// code
	//	}
	//
	// This method can only be used to iterate a [List] if the index is not needed.
	// if you need to iterate a [List] in reverse order with the index there are two options:
	//
	//	for i := list.Len() - 1; i >= 0; i-- {
	//		element, err := list.Get(i)
	//		// Code
	//	}
	//
	//	j := l.Len() -1
	//	for i := range l.Iter() {
	//		// code
	//		j--
	//	}
	IterReverse() chan T
	// Copy returns a copy of the list.
	Copy() List[T]
}

// Comparator defines a method used to sort a [List] though the Sort method
type Comparator[T any] interface {
	// Compare returns a bool which indicates how the elements have to be sorted.
	//
	// If the result is false, the receiver is placed before o, otherwhise it is placed after the parameter.
	Compare(o T) bool
}

// Sort returns a [List] which contains all elements of l that have been sorted.
//
// This function can be used only with types which implements the [Comparator] interface.
func Sort[T Comparator[T]](l List[T]) List[T] {

	slice := l.ToSlice()
	sort.Slice(slice, func(i, j int) bool {

		return slice[i].Compare(slice[j])

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
