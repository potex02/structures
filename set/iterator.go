package set

import (
	"github.com/potex02/structures/table"
	"github.com/potex02/structures/tree"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ Iterator[wrapper.Int] = NewHashSetIterator[wrapper.Int](NewHashSet[wrapper.Int]())
var _ Iterator[wrapper.Int] = NewTreeSetIterator[wrapper.Int](NewTreeSet[wrapper.Int]())
var _ Iterator[wrapper.Int] = &endIterator[wrapper.Int]{}

// Iterator provides the methods to iterate over a [Set] or a [MultiSet].
type Iterator[T util.Comparer] interface {
	// Elements returns the element of the iterator.
	Element() T
	// Remove removes the element from the set and returns the iterator of the next element.
	//
	// The result of this method must be assigned in most cases to himself.
	//
	//	i = i.Remove()
	//
	// An example of the use of the method is the following:
	//
	//	for i := set.Iter(); !i.End(); i = i.Next() {
	//		// Code
	//		if /*some condition*/ {
	//			i = i.Remove()
	//		}
	//		// Code
	//	}
	//
	// The following code, instead, can lead to undefined program behavior:
	//
	//	for i := set.Iter(); !i.End(); i = i.Next() {
	//		// Code
	//		if /*some condition*/ {
	//			i.Remove()
	//		}
	//		// Code
	//	}
	//
	Remove() Iterator[T]
	// Next returns the iterator of the next element.
	Next() Iterator[T]
	// End checks if the iteration is finished.
	End() bool
}

// HashSetIterator is an iterator of a [HashSet] or [MultiHashSet].
type HashSetIterator[T util.Hasher] struct {
	// contains filtered or unexported fields
	iterator table.Iterator[T, uint8]
}

// NewHashSetIterator returns a new [HashSetIterator] for a [HashSet] associated at the set parameter.
func NewHashSetIterator[T util.Hasher](set *HashSet[T]) Iterator[T] {
	if set.IsEmpty() {
		return &endIterator[T]{}
	}
	return &HashSetIterator[T]{iterator: table.NewHashTableIterator(set.objects.(*table.HashTable[T, uint8]))}
}

// NewMultiHashSetIterator returns a new [HashSetIterator] for a [MultiHashSet] associated at the set parameter.
func NewMultiHashSetIterator[T util.Hasher](set *MultiHashSet[T]) Iterator[T] {
	if set.IsEmpty() {
		return &endIterator[T]{}
	}
	return &HashSetIterator[T]{iterator: table.NewMultiHashTableIterator(set.objects.(*table.MultiHashTable[T, uint8]))}
}

// Elements returns the element of the iterator.
func (i *HashSetIterator[T]) Element() T {
	return i.iterator.Key()
}

// Remove removes the element from the set and returns the iterator of the next element.
//
// The result of this method must be assigned in most cases to himself.
//
//	i = i.Remove()
//
// An example of the use of the method is the following:
//
//	for i := set.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i = i.Remove()
//		}
//		// Code
//	}
//
// The following code, instead, can lead to undefined program behavior:
//
//	for i := set.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i.Remove()
//		}
//		// Code
//	}
func (i *HashSetIterator[T]) Remove() Iterator[T] {
	i.iterator = i.iterator.Remove()
	if i.iterator.End() {
		return &endIterator[T]{}
	}
	return i
}

// Next returns the iterator of the next element.
func (i *HashSetIterator[T]) Next() Iterator[T] {
	i.iterator = i.iterator.Next()
	if i.iterator.End() {
		return &endIterator[T]{}
	}
	return i
}

// End checks if the iteration is finished.
func (i *HashSetIterator[T]) End() bool {
	return false
}

// TreeSetIterator is an iterator of a [TreeSet] or a [MultiTreeSet].
type TreeSetIterator[T util.Comparer] struct {
	// contains filtered or unexported fields
	iterator tree.Iterator[T]
}

// NewTreeSetIterator returns a new [TreeSetIterator] for a [TreeSet] associated at the set parameter.
func NewTreeSetIterator[T util.Comparer](set *TreeSet[T]) Iterator[T] {
	if set.IsEmpty() {
		return &endIterator[T]{}
	}
	return &TreeSetIterator[T]{iterator: tree.NewTreeIterator[T](set.objects)}
}

// NewTreeSetIterator returns a new [TreeSetIterator] for a [MultiTreeSet] associated at the set parameter.
func NewMultiTreeSetIterator[T util.Comparer](set *MultiTreeSet[T]) Iterator[T] {
	if set.IsEmpty() {
		return &endIterator[T]{}
	}
	return &TreeSetIterator[T]{iterator: tree.NewTreeIterator[T](set.objects)}
}

// Elements returns the element of the iterator.
func (i *TreeSetIterator[T]) Element() T {
	return i.iterator.Element()
}

// Remove removes the element from the set and returns the iterator of the next element.
//
// The result of this method must be assigned in most cases to himself.
//
//	i = i.Remove()
//
// An example of the use of the method is the following:
//
//	for i := set.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i = i.Remove()
//		}
//		// Code
//	}
//
// The following code, instead, can lead to undefined program behavior:
//
//	for i := set.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i.Remove()
//		}
//		// Code
//	}
func (i *TreeSetIterator[T]) Remove() Iterator[T] {
	i.iterator = i.iterator.Remove()
	if i.iterator.End() {
		return &endIterator[T]{}
	}
	return i
}

// Next returns the iterator of the next element.
func (i *TreeSetIterator[T]) Next() Iterator[T] {
	i.iterator = i.iterator.Next()
	if i.iterator.End() {
		return &endIterator[T]{}
	}
	return i
}

// End checks if the iteration is finished.
func (i *TreeSetIterator[T]) End() bool {
	return false
}

type endIterator[T util.Comparer] struct{}

func (i *endIterator[T]) Element() T {
	return *new(T)
}

func (i *endIterator[T]) Remove() Iterator[T] {
	return i
}

func (i *endIterator[T]) Next() Iterator[T] {
	return i
}

func (i *endIterator[T]) End() bool {
	return true
}
