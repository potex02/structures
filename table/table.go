// package table implements dynamic tables.
package table

import (
	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
	"github.com/potex02/structures/util"
)

// BaseTable is the base interface for both [Table] and [MultiTable].
// A basetable contains all the methods of [structures.Structure].
//
// It provides methods for a generic dynamic table can have unique or duplicate keys.
type BaseTable[K util.Comparer, T any] interface {
	structures.Structure[T]
	// ContainsKey returns true if the key is present in the table.
	ContainsKey(key K) bool
	// ContainsElement returns true if the element e is present in the table.
	ContainsElement(e T) bool
	// Keys returns a [list.List] which contains all keys of the table.
	Keys() list.List[K]
	// Elements returns a [list.List] which contains all elements of the table.
	Elements() list.List[T]
	// PutSlice adds the elements of e at the table.
	// It panics if key and e have different lengths.
	PutSlice(key []K, e []T)
	// Each executes fun for all elements of the table.
	//
	// This method should be used to remove elements. Use Iter insted.
	Each(fun func(key K, element T))
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
}

// Table provides all methods to use a generic dynamic table.
// A table contains all the methods of [BaseTable].
//
// The check on the equality of the keys is done with the Compare method.
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
type Table[K util.Comparer, T any] interface {
	BaseTable[K, T]
	util.Copier[Table[K, T]]
	// Get returns the element associated at the key.
	// The method returns false if the key is not found.
	Get(key K) (T, bool)
	// Put set the element e at the key and returns the overwritten value, if present.
	// If the element is not present, the method returns false.
	Put(key K, e T) (T, bool)
	// Remove removes the key from the table and returns the value associated at the key.
	// It returns false if the the key does not exists.
	Remove(key K) (T, bool)
}

// MultiTable provides all methods to use a generic dynamic table with duplicate keys.
// A multitable contains all the methods of [BaseTable].
//
// The check on the equality of the keys is done with the Compare method.
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
type MultiTable[K util.Comparer, T any] interface {
	BaseTable[K, T]
	util.Copier[MultiTable[K, T]]
	// Contains returns true if the key is present in the table associated with the element e.
	Contains(key K, e T) bool
	// Get returns a slice cotaining the elements associated at the key.
	Get(key K) []T
	// Put add the elements of e at the key.
	Put(key K, e ...T)
	// Replace replace all elements associated at the key with e and returns the slice of overwritten values.
	Replace(key K, e ...T) []T
	// ReplaceSlice replace all elements associated at the key with e and returns the slice of overwritten values.
	ReplaceSlice(key []K, e []T) []T
	// Remove removes the key associated at e from the table.
	// It returns false if the the entry does not exists.
	Remove(key K, e T) bool
	// RemoveKey remove all elements associated at the key and returns the slice of removed values.
	RemoveKey(key K) []T
}
