// package tree implements dynamic trees.
package tree

import "github.com/potex02/structures"

// Tree provides all methods to use a generic dynamic tree.
// A tree contains all the methods of [structures.Structure].
//
// A tree is implemented through the [Node] type.
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
//
type Tree[T any] interface {
	structures.Structure[T]
	// Root returns the root [Node] of the tree.
	Root() *Node[T]
	// Contains returns if e is present in the tree.
	Contains(e T) bool
	// Add adds the elements e at the tree.
	Add(e ...T)
	// AddSlice adds the elements of e at the tree.
	AddSlice(e []T)
	// Remove removes the element e if present.
	// In that case, the method returns true.
	Remove(e T) bool
	// Each executes fun for all elements of the subtree.
	//
	// node is the root node of the subtree,
	// fun is the function to be executed.
	//
	// This method should be used to remove elements. Use Iter insted.
	Each(node *Node[T], fun func(i *Node[T]))
	// Iter returns an [Iterator] which permits to iterate a [Tree].
	//
	//	for i := t.Iter(); !i.End(); i = i.Next() {
	//		element := i.Element()
	//		// Code
	//	}
	Iter() Iterator[T]
}
