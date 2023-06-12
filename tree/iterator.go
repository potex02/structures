package tree

import (
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ Iterator[wrapper.Int] = NewBinaryTreeIterator[wrapper.Int](NewBinaryTree[wrapper.Int]())
var _ Iterator[wrapper.Int] = &endIterator[wrapper.Int]{}

// Iterator provides the methods to iterate over a [Tree].
type Iterator[T util.Comparer] interface {
	// Elements returns the element of the iterator.
	Element() T
	// Remove removes the element from the tree and returns the iterator of the next element.
	//
	// The result of this method must be assigned in most cases to himself.
	//
	// i = i.Remove()
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

// HashTableIterator is an iterator of an [BinaryTree].
type BinaryTreeIterator[T util.Comparer] struct {
	// contains filtered or unexported fields
	tree *BinaryTree[T]
	node *Node[T]
	next *BinaryTreeIterator[T]
}

// NewBinaryTreeIterator returns a new [BinaryTreeIterator] associated at the tree parameter.
func NewBinaryTreeIterator[T util.Comparer](tree *BinaryTree[T]) Iterator[T] {

	if tree.IsEmpty() {

		return &endIterator[T]{}

	}
	node := tree.Root().Min()
	iterator := &BinaryTreeIterator[T]{tree: tree, node: node}
	current := iterator
	tree.Each(tree.Root(), func(i *Node[T]) {
		if i != node {
			current.setNext(&BinaryTreeIterator[T]{tree: tree, node: i})
			current = current.next
		}
	})
	return iterator

}

// Elements returns the element of the iterator.
func (i *BinaryTreeIterator[T]) Element() T {

	return i.node.Element()

}

// Remove removes the element from the tree and returns the iterator of the next element.
//
// The result of this method must be assigned in most cases to himself.
//
// i = i.Remove()
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
func (i *BinaryTreeIterator[T]) Remove() Iterator[T] {

	next := i.Next()
	i.tree.Remove(i.node.element)
	return next

}

// Next returns the iterator of the next element.
func (i *BinaryTreeIterator[T]) Next() Iterator[T] {

	if i.next == nil {

		return &endIterator[T]{}

	}
	return i.next

}

func (i *BinaryTreeIterator[T]) setNext(next *BinaryTreeIterator[T]) {

	i.next = next

}

// End checks if the iteration is finished.
func (i *BinaryTreeIterator[T]) End() bool {

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
