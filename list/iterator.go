package list

import (
	"github.com/potex02/structures"
)

var _ Iterator[int] = NewArrayListIterator[int](NewArrayList[int]())
var _ Iterator[int] = NewLinkedListIterator[int](NewLinkedList[int]())
var _ Iterator[int] = &endIterator[int]{}

// Iterator provides the methods to iterate over a [List].
type Iterator[T any] interface {
	// Elements returns the element of the iterator.
	Element() T
	// Index returns the index of the element the iterator.
	Index() int
	// Remove removes the element from the list and returns the iterator of the next element.
	//
	// The result of this method must be assigned in most cases to himself.
	//
	//	i = i.Remove()
	//
	// An example of the use of the method is the following:
	//
	//	for i := list.Iter(); !i.End(); i = i.Next() {
	//		// Code
	//		if /*some condition*/ {
	//			i = i.Remove()
	//		}
	//		// Code
	//	}
	//
	// The following code, instead, can lead to undefined program behavior:
	//
	//	for i := list.Iter(); !i.End(); i = i.Next() {
	//		// Code
	//		if /*some condition*/ {
	//			i.Remove()
	//		}
	//		// Code
	//	}
	//
	Remove() Iterator[T]
	// Prev returns the iterator of the previous element.
	Prev() Iterator[T]
	// Next returns the iterator of the next element.
	Next() Iterator[T]
	// End checks if the iteration is finished.
	End() bool
}

// ArrayListIterator is an iterator of an [ArrayList].
type ArrayListIterator[T any] struct {
	// contains filtered or unexported fields
	list    *ArrayList[T]
	element T
	index   int
}

// NewArrayListIterator returns a new [ArrayListIterator] associated at the list parameter.
func NewArrayListIterator[T any](list *ArrayList[T]) Iterator[T] {
	element, err := list.Get(0)
	if err != nil {
		return &endIterator[T]{}
	}
	return &ArrayListIterator[T]{list: list, element: element, index: 0}
}

// NewArrayListReverseIterator returns a new reverse [ArrayListIterator] associated at the list parameter.
func NewArrayListReverseIterator[T any](list *ArrayList[T]) Iterator[T] {
	element, err := list.Get(list.Len() - 1)
	if err != nil {
		return &endIterator[T]{}
	}
	return &ArrayListIterator[T]{list: list, element: element, index: list.Len() - 1}
}

// Elements returns the element of i.
func (i *ArrayListIterator[T]) Element() T {
	return i.element
}

// Index returns the index of the element of i.
func (i *ArrayListIterator[T]) Index() int {
	return i.index
}

// Remove removes the element from the list and returns the iterator of the next element.
//
// The result of this method must be assigned in most cases to himself.
//
//	i = i.Remove()
//
// An example of the use of the method is the following:
//
//	for i := list.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i = i.Remove()
//		}
//		// Code
//	}
//
// The following code, instead, can lead to undefined program behavior:
//
//	for i := list.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i.Remove()
//		}
//		// Code
//	}
func (i *ArrayListIterator[T]) Remove() Iterator[T] {
	i.list.Remove(i.index)
	i.index--
	return i.Next()
}

// Prev returns the iterator of the previous element.
func (i *ArrayListIterator[T]) Prev() Iterator[T] {
	i.index--
	element, err := i.list.Get(i.index)
	if err != nil {
		return &endIterator[T]{}
	}
	i.element = element
	return i
}

// Next returns the iterator of the next element.
func (i *ArrayListIterator[T]) Next() Iterator[T] {
	i.index++
	element, err := i.list.Get(i.index)
	if err != nil {
		return &endIterator[T]{}
	}
	i.element = element
	return i
}

// End checks if the iteration is finished.
func (i *ArrayListIterator[T]) End() bool {
	return false
}

// LinkedListIterator is an iterator of a [LinkedList].
type LinkedListIterator[T any] struct {
	// contains filtered or unexported fields
	list  *LinkedList[T]
	entry *structures.Entry[T]
	index int
}

// NewLinkedListIterator returns a new [LinkedListIterator] associated at the list parameter.
func NewLinkedListIterator[T any](list *LinkedList[T]) Iterator[T] {
	if list.IsEmpty() {
		return &endIterator[T]{}
	}
	return &LinkedListIterator[T]{list: list, entry: list.root, index: 0}
}

// NewLinkedListReverseIterator returns a new reverse [LinkedListIterator] associated at the list parameter.
func NewLinkedListReverseIterator[T any](list *LinkedList[T]) Iterator[T] {
	if list.IsEmpty() {
		return &endIterator[T]{}
	}
	return &LinkedListIterator[T]{list: list, entry: list.tail, index: list.Len() - 1}
}

// Elements returns the element of i.
func (i *LinkedListIterator[T]) Element() T {
	return i.entry.Element()
}

// Index returns the index of the element of i.
func (i *LinkedListIterator[T]) Index() int {
	return i.index
}

// Remove removes the element from the list and returns the iterator of the next element.
//
// The result of this method must be assigned in most cases to himself.
//
//	i = i.Remove()
//
// An example of the use of the method is the following:
//
//	for i := list.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i = i.Remove()
//		}
//		// Code
//	}
//
// The following code, instead, can lead to undefined program behavior:
//
//	for i := list.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i.Remove()
//		}
//		// Code
//	}
func (i *LinkedListIterator[T]) Remove() Iterator[T] {
	i.index--
	i.list.removeEntry(i.entry)
	return i.Next()
}

// Prev returns the iterator of the previous element.
func (i *LinkedListIterator[T]) Prev() Iterator[T] {
	if i.entry.Prev() == nil {
		return &endIterator[T]{}
	}
	i.index--
	i.entry = i.entry.Prev()
	return i
}

// Next returns the iterator of the next element.
func (i *LinkedListIterator[T]) Next() Iterator[T] {
	if i.entry.Next() == nil {
		return &endIterator[T]{}
	}
	i.index++
	i.entry = i.entry.Next()
	return i
}

// End checks if the iteration is finished.
func (i *LinkedListIterator[T]) End() bool {
	return false
}

type endIterator[T any] struct{}

func (i *endIterator[T]) Element() T {
	return *new(T)
}

func (i *endIterator[T]) Index() int {
	return -1
}

func (i *endIterator[T]) Remove() Iterator[T] {
	return i
}

func (i *endIterator[T]) Prev() Iterator[T] {
	return i
}

func (i *endIterator[T]) Next() Iterator[T] {
	return i
}

func (i *endIterator[T]) End() bool {
	return true
}
