// Package list implements dynamic lists.
package list

import (
	"github.com/potex02/structures"
	"github.com/potex02/structures/util"
)

// List provides all methods to use a generic dynamic list.
// A list contains all the methods of [structures.Structure].
//
// A list is indexed starting from 0, but negative indexes are supported. Negative indexes start from the end, meaning that -1 corresponds to the last element.
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
type List[T any] interface {
	structures.Structure[T]
	util.Copier[List[T]]
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
	// GetDefault returns the elements at the specifies index.
	// It returns the T zero value if the the index is out of bounds.
	GetDefault(index int) T
	// GetDefaultValue returns the elements at the specifies index.
	// It returns value if the the index is out of bounds.
	GetDefaultValue(index int, value T) T
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
	// Each executes fun for all elements of the list.
	//
	// This method should be used to remove elements. Use Iter insted.
	Each(fun func(index int, element T))
	// Stream returns a [Stream] rapresenting the list.
	Stream() *Stream[T]
	// Sort sorts the elements of the list.
	//
	// This method panics if T does not implement [util.Comparer]
	Sort()
	//  SortFunc sorts the elements of the list as determined by the less function.
	SortFunc(less func(i T, j T) int)
	// Iter returns an [Iterator] which permits to iterate a [List].
	//
	//	for i := list.Iter(); !i.End(); i = i.Next() {
	//		element := i.Element()
	//		index := i.Index()
	//		// Code
	//	}
	Iter() Iterator[T]
	// IterReverse returns an [Iterator] which permits to iterate a [List] in reverse order.
	//
	//	for i := list.IterReverse(); !i.End(); i = i.Prev() {
	//		element := i.Element()
	//		index := i.Index()
	//		// Code
	//	}
	IterReverse() Iterator[T]
}

func rangeCheck[T any](list List[T], index *int) bool {
	if *index < 0 {
		*index += list.Len()
	}
	return *index >= 0 && *index < list.Len()
}
