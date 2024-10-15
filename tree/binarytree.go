package tree

import (
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ structures.Structure[wrapper.Int] = NewBinaryTree[wrapper.Int]()
var _ Tree[wrapper.Int] = NewBinaryTree[wrapper.Int]()

// BinaryTree provides a generic binary search tree.
//
// the type T of the tree must implement [util.Comparer].
//
// It implements the interface [Tree].
type BinaryTree[T util.Comparer] struct {
	// contains filtered or unexported fields
	root *Node[T]
	len  int
}

// NewBinaryTree returns a new [BinaryTree] containing the elements c.
//
// if no argument is passed, it will be created an empty [BinaryTree].
func NewBinaryTree[T util.Comparer](c ...T) *BinaryTree[T] {
	return NewBinaryTreeFromSlice[T](c)
}

// NewBinaryTreeFromSlice returns a new [BinaryTree] containing the elements of slice c.
func NewBinaryTreeFromSlice[T util.Comparer](c []T) *BinaryTree[T] {
	tree := &BinaryTree[T]{root: nil, len: 0}
	if len(c) != 0 {
		tree.AddSlice(c)
	}
	return tree
}

// Len returns the length of t.
func (t *BinaryTree[T]) Len() int {
	return t.len
}

// IsEmpty returns a bool which indicates if t is empty or not.
func (t *BinaryTree[T]) IsEmpty() bool {
	return t.len == 0
}

// Root returns the root [Mode] of t.
func (t *BinaryTree[T]) Root() *Node[T] {
	return t.root
}

// Contains returns if e is present in t.
func (t *BinaryTree[T]) Contains(e T) bool {
	return t.contains(t.root, e)
}

// ToSlice returns a slice which contains all elements of t.
func (t *BinaryTree[T]) ToSlice() []T {
	slice := make([]T, 0)
	t.Each(t.root, func(i *Node[T]) { slice = append(slice, i.Element()) })
	return slice
}

// Add adds the elements e at t.
func (t *BinaryTree[T]) Add(e ...T) {
	t.AddSlice(e)
}

// AddSlice adds the elements of e at t.
func (t *BinaryTree[T]) AddSlice(e []T) {
	for _, i := range e {
		t.add(i)
	}
}

// Remove removes the element e if present.
// In that case, the method returns true.
func (t *BinaryTree[T]) Remove(e T) bool {
	return t.Any(t.root, func(i *Node[T]) bool {
		if e.Compare(i.Element()) == 0 {
			t.remove(i)
			return true
		}
		return false
	})
}

// Remove removes the first element that satisfies fun, if present.
// In that case, the method returns true.
func (t *BinaryTree[T]) RemoveFunc(e T, fun func(i T, other *Node[T]) bool) bool {
	return t.Any(t.root, func(i *Node[T]) bool {
		if fun(e, i) {
			t.remove(i)
			return true
		}
		return false
	})
}

// Each executes fun for all elements of a subtree.
//
// node is the root node of the subtree,
// fun is the function to be executed.
//
// This method should be used to remove elements. Use Iter insted.
func (t *BinaryTree[T]) Each(node *Node[T], fun func(i *Node[T])) {
	if node == nil {
		return
	}
	t.Each(node.Left(), fun)
	fun(node)
	t.Each(node.Right(), fun)
}

// Map executes fun for all elements of a subtree and returns a [BinaryTree] containing the resulting elements.
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *BinaryTree[T]) Map(node *Node[T], fun func(i *Node[T]) T) *BinaryTree[T] {
	result := NewBinaryTree[T]()
	t.Each(t.root, func(i *Node[T]) {
		result.add(fun(i))
	})
	return result
}

// Filter returns a [BinaryTree] containing the elements of a subtree that satisfy fun.
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *BinaryTree[T]) Filter(node *Node[T], fun func(i *Node[T]) bool) *BinaryTree[T] {
	result := NewBinaryTree[T]()
	t.Each(t.root, func(i *Node[T]) {
		if fun(i) {
			result.add(i.element)
		}
	})
	return result
}

// FilterMap executes fun for all elements of a subtree and returns a [BinaryTree] containing the resulting elements that satisfy fun.
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *BinaryTree[T]) FilterMap(node *Node[T], fun func(i *Node[T]) (T, bool)) *BinaryTree[T] {
	result := NewBinaryTree[T]()
	t.Each(t.root, func(i *Node[T]) {
		if element, ok := fun(i); ok {
			result.add(element)
		}
	})
	return result
}

// Any returns true if at least one element of a subtree satisfies fun.
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *BinaryTree[T]) Any(node *Node[T], fun func(i *Node[T]) bool) bool {
	if node == nil {
		return false
	}
	return t.Any(node.Left(), fun) || fun(node) || t.Any(node.Right(), fun)
}

// / All returns true if all elements of a subtree table satisfy fun.
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *BinaryTree[T]) All(node *Node[T], fun func(i *Node[T]) bool) bool {
	if node == nil {
		return true
	}
	return t.All(node.Left(), fun) && fun(node) && t.All(node.Right(), fun)
}

// None returns true if none of the elements of a subtree table satisfies fun.
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *BinaryTree[T]) None(node *Node[T], fun func(i *Node[T]) bool) bool {
	if node == nil {
		return true
	}
	return t.None(node.Left(), fun) && !fun(node) && t.None(node.Right(), fun)
}

// Count returns the number of elements of a subtree that satisfy fun..
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *BinaryTree[T]) Count(node *Node[T], fun func(i *Node[T]) bool) int {
	result := 0
	if node == nil {
		return result
	}
	if fun(node) {
		result = 1
	}
	return t.Count(node.Left(), fun) + result + t.Count(node.Right(), fun)
}

// Clear removes all element from t.
func (t *BinaryTree[T]) Clear() {
	t.root = nil
	t.len = 0
}

// Iter returns an [Iterator] which permits to iterate a [BinaryTree].
//
//	for i := t.Iter(); !i.End(); i = i.Next() {
//		element := i.Element()
//		// Code
//	}
func (t *BinaryTree[T]) Iter() Iterator[T] {
	return NewTreeIterator[T](t)
}

// Equal returns true if t and st are both [BinaryTree] and their elements are equals.
// In any other case, it returns false.
func (t *BinaryTree[T]) Equal(st any) bool {
	tree, ok := st.(*BinaryTree[T])
	if ok && t != nil && tree != nil {
		if t.Len() != tree.Len() {
			return false
		}
		others := tree.ToSlice()
		j := 0
		return t.All(t.root, func(i *Node[T]) bool {
			j++
			return util.EqualFunction(i.element)(others[j-1])
		})
	}
	return false
}

// Compare returns -1 if t is shorten than st,
// 1 if t is longer than st,
// -2 if st is not a [BinaryTree] or if one between t and st is nil.
//
// If t and st have the same length, the result is the comparison
// between the first different element of the two trees.
// If they are all equals, the result is 0.
func (t *BinaryTree[T]) Compare(st any) int {
	tree, ok := st.(*BinaryTree[T])
	if ok && t != nil && tree != nil {
		if t.Len() < tree.Len() {
			return -1
		}
		if t.Len() > tree.Len() {
			return 1
		}
		others := tree.ToSlice()
		j := 0
		result := 0
		t.All(t.root, func(i *Node[T]) bool {
			j++
			result = i.Element().Compare(others[j-1])
			return result == 0
		})
		return result
	}
	return -2
}

// Hash returns the hash code of t.
func (t *BinaryTree[T]) Hash() string {
	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("%v%v", check[1:], t.Len())
}

// String returns a rapresentation of t in the form of a string.
func (t *BinaryTree[T]) String() string {
	check := reflect.TypeOf(new(T)).String()
	objects := make([]T, 0)
	t.Each(t.root, func(i *Node[T]) { objects = append(objects, i.Element()) })
	return fmt.Sprintf("BinaryTree[%v]%v", check[1:], objects)
}

func (t *BinaryTree[T]) contains(node *Node[T], e T) bool {
	if node == nil {
		return false
	}
	check := e.Compare(node.Element())
	if check == 0 {
		return true
	}
	if check < 0 {
		return t.contains(node.Left(), e)
	}
	return t.contains(node.Right(), e)
}

func (t *BinaryTree[T]) add(e T) {
	if t.root == nil {
		t.root = NewNode[T](e, nil, nil, nil)
		t.len++
		return
	}
	t.checkNext(t.root, e)
}

func (t *BinaryTree[T]) addLeft(parent *Node[T], e T) {
	if parent.Left() == nil {
		node := NewNode[T](e, parent, nil, nil)
		parent.SetLeft(node)
		t.len++
		return
	}
	t.checkNext(parent.Left(), e)
}

func (t *BinaryTree[T]) addRight(parent *Node[T], e T) {
	if parent.Right() == nil {
		node := NewNode[T](e, parent, nil, nil)
		parent.SetRight(node)
		t.len++
		return
	}
	t.checkNext(parent.Right(), e)
}

func (t *BinaryTree[T]) checkNext(parent *Node[T], e T) {
	compare := e.Compare(parent.Element())
	if compare < 0 {
		t.addLeft(parent, e)
		return
	}
	t.addRight(parent, e)
}

func (t *BinaryTree[T]) remove(node *Node[T]) {
	t.len--
	if node.Right() == nil {
		if node == t.root {
			t.root = node.Left()
		} else {
			if node == node.Parent().Left() {
				node.Parent().SetLeft(node.Left())
			} else {
				node.Parent().SetRight(node.Left())
			}
		}
		if node.Left() != nil {
			node.Left().SetParent(node.Parent())
		}
		return
	}
	min := node.Right().Min()
	node.SetElement(min.Element())
	if min == min.Parent().Left() {
		min.Parent().SetLeft(min.Right())
		if min.Right() != nil {
			min.Right().SetParent(min.Parent())
		}
	} else {
		min.Parent().SetRight(min.Right())
		if min.Right() != nil {
			min.Right().SetParent(min.Parent())
		}
	}
	min.SetParent(nil)
}
