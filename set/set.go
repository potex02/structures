// package table implements dynamic sets.
package set

import (
	"github.com/potex02/structures"
	"github.com/potex02/structures/util"
)

// BaseSet is the base interface for both [Set] and [MultiSet].
// A baseset contains all the methods of [structures.Structure].
//
// It provides methods for a generic dynamic table can have unique or duplicate keys.
type BaseSet[T util.Comparer] interface {
	structures.Structure[T]
	// Contains returns if e is present in the set.
	Contains(e T) bool
	// Add adds the elements e at the set.
	Add(e ...T)
	// AddSlice adds the elements of e at the set.
	AddSlice(e []T)
	// Remove removes the element e from the set if it is present.
	// In that case, the method returns true, otherwhise it returns false.
	Remove(e T) bool
	// Each executes fun for all elements of the set.
	//
	// This method should be used to remove elements. Use Iter insted.
	Each(fun func(element T))
	// Stream returns a [Stream] rapresenting the set.
	Stream() *Stream[T]
	// Iter returns an [Iterator] which permits to iterate a [Set].
	//
	//	for i := set.Iter(); !i.End(); i = i.Next() {
	//		element := i.Element()
	//		// Code
	//	}
	Iter() Iterator[T]
	// RangeIter returns a function that allows to iterate a [Set] using the range keyword.
	//
	//	for i := range set.RangeIter() {
	//		// Code
	//	}
	//
	// Unlike [Set.Iter], it doesn't allow to remove elements during the iteration.
	RangeIter() func(yield func(T) bool)
}

// Set provides all methods to use a generic dynamic set.
// A set contains all the methods of [BaseSet].
//
// The check on the equality of the elements is done with the Compare method.
type Set[T util.Comparer] interface {
	util.Copier[Set[T]]
	BaseSet[T]
}

// MultiSet provides all methods to use a generic dynamic set with duplicate elements.
// A multiset contains all the methods of [BaseSet].
//
// The check on the equality of the elements is done with the Compare method.
type MultiSet[T util.Comparer] interface {
	util.Copier[MultiSet[T]]
	BaseSet[T]
	// RemoveAll removes all occurrences of e from the set.
	RemoveAll(e T)
	// Count returns the number of occurrences of e in the set.
	Count(e T) int
	// ToSet returns a [Set] containing the elements of the multiset.
	ToSet() Set[T]
}

const obj uint8 = 0
