// package table implements dynamic tables.
package table

import (
	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
	"github.com/potex02/structures/util"
)

// Table provides all methods to use a generic dynamic table.
// A table contains all the methods of [structures.Structure].
//
// The check on the equality of the keys is done with the Compare method.
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
type Table[K util.Comparer, T any] interface {
	structures.Structure[T]
	// ContainsKey returns true if the key is present in the table.
	ContainsKey(key K) bool
	// ContainsElement returns true if the element e is present in the table.
	ContainsElement(e T) bool
	// Keys returns a [list.List] which contains all keys of the table.
	Keys() list.List[K]
	// Elements returns a [list.List] which contains all elements of the table.
	Elements() list.List[T]
	// Get returns the element associated at the key.
	// The method returns false if the key is not found.
	Get(key K) (T, bool)
	// Put set the element e at the key and returns the overwritten value, if present.
	// If the element is not present, the method returns false.
	Put(key K, e T) (T, bool)
	// PutSlice adds the elements of e at the table if not present.
	// it panics if key and e have different lengths.
	PutSlice(key []K, e []T)
	// Remove removes the key from the table and returns the value associated at the key.
	// It returns false if the the key does not exists.
	Remove(key K) (T, bool)
	// Each executes fun for all elements of the table.
	Each(fun func(Key K, element T))
	// Stream returns a [Stream] rapresenting the table.
	Stream() *Stream[K, T]
	// Iter returns an [Iterator] which permits to iterate a [Table].
	//
	//	for i := table.Iter(); !i.End(); i = i.Next() {
	//		key := i.Key()
	//		element := i.Element()
	//		// Code
	//	}
	Iter() Iterator[K, T]
	// Copy returns a table containing a copy of the elements of the table.
	Copy() Table[K, T]
}
