package table

import (
	"github.com/potex02/structures/list"
	"github.com/potex02/structures/tree"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ Iterator[wrapper.Int, int] = NewHashTableIterator[wrapper.Int, int](NewHashTable[wrapper.Int, int]())
var _ Iterator[wrapper.Int, int] = NewTreeTableIterator[wrapper.Int, int](NewTreeTable[wrapper.Int, int]())
var _ Iterator[wrapper.Int, int] = NewMultiHashTableIterator[wrapper.Int, int](NewMultiHashTable[wrapper.Int, int]())
var _ Iterator[wrapper.Int, int] = &endIterator[wrapper.Int, int]{}

// Iterator provides the methods to iterate over a [Table] or a [MultiTable].
type Iterator[K util.Comparer, T any] interface {
	// Elements returns the element of the iterator.
	Element() T
	// Index returns the key of the element the iterator.
	Key() K
	// Remove removes the element from the table and returns the iterator of the next element.
	//
	// The result of this method must be assigned in most cases to himself.
	//
	//	i = i.Remove()
	//
	// An example of the use of the method is the following:
	//
	//	for i := table.Iter(); !i.End(); i = i.Next() {
	//		// Code
	//		if /*some condition*/ {
	//			i = i.Remove()
	//		}
	//		// Code
	//	}
	//
	// The following code, instead, can lead to undefined program behavior:
	//
	//	for i := table.Iter(); !i.End(); i = i.Next() {
	//		// Code
	//		if /*some condition*/ {
	//			i.Remove()
	//		}
	//		// Code
	//	}
	//
	Remove() Iterator[K, T]
	// Next returns the iterator of the next element.
	Next() Iterator[K, T]
	// End checks if the iteration is finished.
	End() bool
}

// HashTableIterator is an iterator of an [HashTable].
type HashTableIterator[K util.Hasher, T any] struct {
	// contains filtered or unexported fields
	table    *HashTable[K, T]
	iterator list.Iterator[*Entry[K, T]]
	keys     list.List[uint64]
	index    int
}

// NewHashTableIterator returns a new [HashTableIterator] associated at the table parameter.
func NewHashTableIterator[K util.Hasher, T any](table *HashTable[K, T]) Iterator[K, T] {
	if table.IsEmpty() {
		return &endIterator[K, T]{}
	}
	keys := list.NewArrayList[uint64]()
	for i := range table.objects {
		keys.Add(i)
	}
	key, _ := keys.Get(0)
	return &HashTableIterator[K, T]{table: table, iterator: table.objects[key].Iter(), keys: keys, index: 0}
}

// Elements returns the element of the iterator.
func (i *HashTableIterator[K, T]) Element() T {
	return i.iterator.Element().Element()
}

// Index returns the key of the element the iterator.
func (i *HashTableIterator[K, T]) Key() K {
	return i.iterator.Element().Key()
}

// Remove removes the element from the table and returns the iterator of the next element.
//
// The result of this method must be assigned in most cases to himself.
//
//	i = i.Remove()
//
// An example of the use of the method is the following:
//
//	for i := table.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i = i.Remove()
//		}
//		// Code
//	}
//
// The following code, instead, can lead to undefined program behavior:
//
//	for i := table.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i.Remove()
//		}
//		// Code
//	}
func (i *HashTableIterator[K, T]) Remove() Iterator[K, T] {
	key := i.iterator.Element().Key()
	i.table.Remove(key)
	return i.Next()
}

// Next returns the iterator of the next element.
func (i *HashTableIterator[K, T]) Next() Iterator[K, T] {
	i.iterator = i.iterator.Next()
	if i.iterator.End() {
		return i.nextKey()
	}
	return i
}

// End checks if the iteration is finished.
func (i *HashTableIterator[K, T]) End() bool {
	return false
}

func (i *HashTableIterator[K, T]) nextKey() Iterator[K, T] {
	i.index++
	key, err := i.keys.Get(i.index)
	if err != nil {
		return &endIterator[K, T]{}
	}
	i.iterator = i.table.objects[key].Iter()
	return i
}

// TreeTableIterator is an iterator of a [TreeTable].
type TreeTableIterator[K util.Comparer, T any] struct {
	// contains filtered or unexported fields
	table    *TreeTable[K, T]
	iterator tree.Iterator[*Entry[K, T]]
}

// NewTreeTableIterator returns a new [TreeTableIterator] associated at the table parameter.
func NewTreeTableIterator[K util.Comparer, T any](table *TreeTable[K, T]) Iterator[K, T] {
	if table.IsEmpty() {
		return &endIterator[K, T]{}
	}
	return &TreeTableIterator[K, T]{table: table, iterator: tree.NewTreeIterator[*Entry[K, T]](table.objects)}
}

// Elements returns the element of the iterator.
func (i *TreeTableIterator[K, T]) Element() T {
	return i.iterator.Element().Element()
}

// Index returns the key of the element the iterator.
func (i *TreeTableIterator[K, T]) Key() K {
	return i.iterator.Element().Key()
}

// Remove removes the element from the table and returns the iterator of the next element.
//
// The result of this method must be assigned in most cases to himself.
//
//	i = i.Remove()
//
// An example of the use of the method is the following:
//
//	for i := table.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i = i.Remove()
//		}
//		// Code
//	}
//
// The following code, instead, can lead to undefined program behavior:
//
//	for i := table.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i.Remove()
//		}
//		// Code
//	}
func (i *TreeTableIterator[K, T]) Remove() Iterator[K, T] {
	i.iterator = i.iterator.Remove()
	if i.iterator.End() {
		return &endIterator[K, T]{}
	}
	return i
}

// Next returns the iterator of the next element.
func (i *TreeTableIterator[K, T]) Next() Iterator[K, T] {
	i.iterator = i.iterator.Next()
	if i.iterator.End() {
		return &endIterator[K, T]{}
	}
	return i
}

// End checks if the iteration is finished.
func (i *TreeTableIterator[K, T]) End() bool {
	return false
}

// MultiHashTableIterator is an iterator of a [MultiHashTable].
type MultiHashTableIterator[K util.Hasher, T any] struct {
	// contains filtered or unexported fields
	table    *MultiHashTable[K, T]
	iterator list.Iterator[*Entry[K, T]]
	keys     list.List[uint64]
	index    int
}

// NewMultiHashTableIterator returns a new [MultiHashTableIterator] associated at the table parameter.
func NewMultiHashTableIterator[K util.Hasher, T any](table *MultiHashTable[K, T]) Iterator[K, T] {
	if table.IsEmpty() {
		return &endIterator[K, T]{}
	}
	keys := list.NewArrayList[uint64]()
	for i := range table.objects {
		keys.Add(i)
	}
	key, _ := keys.Get(0)
	return &MultiHashTableIterator[K, T]{table: table, iterator: table.objects[key].Iter(), keys: keys, index: 0}
}

// Elements returns the element of the iterator.
func (i *MultiHashTableIterator[K, T]) Element() T {
	return i.iterator.Element().Element()
}

// Index returns the key of the element the iterator.
func (i *MultiHashTableIterator[K, T]) Key() K {
	return i.iterator.Element().Key()
}

// Remove removes the element from the table and returns the iterator of the next element.
//
// The result of this method must be assigned in most cases to himself.
//
//	i = i.Remove()
//
// An example of the use of the method is the following:
//
//	for i := table.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i = i.Remove()
//		}
//		// Code
//	}
//
// The following code, instead, can lead to undefined program behavior:
//
//	for i := table.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i.Remove()
//		}
//		// Code
//	}
func (i *MultiHashTableIterator[K, T]) Remove() Iterator[K, T] {
	key := i.iterator.Element().Key()
	element := i.iterator.Element().Element()
	i.table.Remove(key, element)
	return i.Next()
}

// Next returns the iterator of the next element.
func (i *MultiHashTableIterator[K, T]) Next() Iterator[K, T] {
	i.iterator = i.iterator.Next()
	if i.iterator.End() {
		return i.nextKey()
	}
	return i
}

// End checks if the iteration is finished.
func (i *MultiHashTableIterator[K, T]) End() bool {
	return false
}

func (i *MultiHashTableIterator[K, T]) nextKey() Iterator[K, T] {
	i.index++
	key, err := i.keys.Get(i.index)
	if err != nil {
		return &endIterator[K, T]{}
	}
	i.iterator = i.table.objects[key].Iter()
	return i
}

// MultiTreeTableIterator is an iterator of a [MultiTreeTable].
type MultiTreeTableIterator[K util.Comparer, T any] struct {
	// contains filtered or unexported fields
	table    *MultiTreeTable[K, T]
	iterator tree.Iterator[*Entry[K, T]]
}

// NewMultiTreeTableIterator returns a new [MultiTreeTableIterator] associated at the table parameter.
func NewMultiTreeTableIterator[K util.Comparer, T any](table *MultiTreeTable[K, T]) Iterator[K, T] {
	if table.IsEmpty() {
		return &endIterator[K, T]{}
	}
	return &MultiTreeTableIterator[K, T]{table: table, iterator: tree.NewTreeIterator[*Entry[K, T]](table.objects)}
}

// Elements returns the element of the iterator.
func (i *MultiTreeTableIterator[K, T]) Element() T {
	return i.iterator.Element().Element()
}

// Index returns the key of the element the iterator.
func (i *MultiTreeTableIterator[K, T]) Key() K {
	return i.iterator.Element().Key()
}

// Remove removes the element from the table and returns the iterator of the next element.
//
// The result of this method must be assigned in most cases to himself.
//
//	i = i.Remove()
//
// An example of the use of the method is the following:
//
//	for i := table.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i = i.Remove()
//		}
//		// Code
//	}
//
// The following code, instead, can lead to undefined program behavior:
//
//	for i := table.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i.Remove()
//		}
//		// Code
//	}
func (i *MultiTreeTableIterator[K, T]) Remove() Iterator[K, T] {
	i.iterator = i.iterator.Remove()
	if i.iterator.End() {
		return &endIterator[K, T]{}
	}
	return i
}

// Next returns the iterator of the next element.
func (i *MultiTreeTableIterator[K, T]) Next() Iterator[K, T] {
	i.iterator = i.iterator.Next()
	if i.iterator.End() {
		return &endIterator[K, T]{}
	}
	return i
}

// End checks if the iteration is finished.
func (i *MultiTreeTableIterator[K, T]) End() bool {
	return false
}

type endIterator[K util.Comparer, T any] struct{}

func (i *endIterator[K, T]) Element() T {
	return *new(T)
}

func (i *endIterator[K, T]) Key() K {
	return *new(K)
}

func (i *endIterator[K, T]) Remove() Iterator[K, T] {
	return i
}

func (i *endIterator[K, T]) Next() Iterator[K, T] {
	return i
}

func (i *endIterator[K, T]) End() bool {
	return true
}
