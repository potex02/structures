// package table implements dynamic sets.
package set

import (
	"github.com/potex02/structures"
	"github.com/potex02/structures/util"
)

// Set provides all methods to use a generic dynamic set.
// A set contains all the methods of [structures.Structure].
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
type Set[T util.Comparer] interface {
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
	// Iter returns an [Iterator] which permits to iterate a [Set].
	//
	//	for i := set.Iter(); !i.End(); i = i.Next() {
	//		element := i.Element()
	//		// Code
	//	}
	Iter() Iterator[T]
}
