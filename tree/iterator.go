package tree

import (
	"github.com/potex02/structures/util/wrapper"
)

var _ Iterator[wrapper.Int] = NewTreeIterator[wrapper.Int](NewBinaryTree[wrapper.Int]())
var _ Iterator[int] = NewTreeIterator[int](NewNAryTree[int](3))
var _ Iterator[wrapper.Int] = &endIterator[wrapper.Int]{}

// Iterator provides the methods to iterate over a [Tree].
type Iterator[T any] interface {
	// Elements returns the element of the iterator.
	Element() T
	// Remove removes the element from the tree and returns the iterator of the next element.
	//
	// The result of this method must be assigned in most cases to himself.
	//
	//	i = i.Remove()
	//
	// An example of the use of the method is the following:
	//
	//	for i := tree.Iter(); !i.End(); i = i.Next() {
	//		// Code
	//		if /*some condition*/ {
	//			i = i.Remove()
	//		}
	//		// Code
	//	}
	//
	// The following code, instead, can lead to undefined program behavior:
	//
	//	for i := tree.Iter(); !i.End(); i = i.Next() {
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

// TreeIterator is an iterator of a [Tree].
type TreeIterator[T any] struct {
	// contains filtered or unexported fields
	tree    Tree[T]
	element T
	next    *TreeIterator[T]
}

// NewTreeIterator returns a new [TreeIterator] associated at the tree parameter.
func NewTreeIterator[T any](tree Tree[T]) Iterator[T] {

	if tree.IsEmpty() {

		return &endIterator[T]{}

	}
	iterator := &TreeIterator[T]{tree: tree, element: tree.Root().Min().Element()}
	current := iterator
	first := true
	tree.Each(tree.Root(), func(i *Node[T]) {
		if !first {
			current.setNext(&TreeIterator[T]{tree: tree, element: i.Element()})
			current = current.next
		}
		first = false
	})
	return iterator

}

// Elements returns the element of the iterator.
func (i *TreeIterator[T]) Element() T {

	return i.element

}

// Remove removes the element from the tree and returns the iterator of the next element.
//
// The result of this method must be assigned in most cases to himself.
//
//	i = i.Remove()
//
// An example of the use of the method is the following:
//
//	for i := tree.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i = i.Remove()
//		}
//		// Code
//	}
//
// The following code, instead, can lead to undefined program behavior:
//
//	for i := tree.Iter(); !i.End(); i = i.Next() {
//		// Code
//		if /*some condition*/ {
//			i.Remove()
//		}
//		// Code
//	}
func (i *TreeIterator[T]) Remove() Iterator[T] {

	next := i.Next()
	i.tree.Remove(i.element)
	return next

}

// Next returns the iterator of the next element.
func (i *TreeIterator[T]) Next() Iterator[T] {

	if i.next == nil {

		return &endIterator[T]{}

	}
	return i.next

}

// End checks if the iteration is finished.
func (i *TreeIterator[T]) End() bool {

	return false

}

func (i *TreeIterator[T]) setNext(next *TreeIterator[T]) {

	i.next = next

}

type endIterator[T any] struct{}

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
