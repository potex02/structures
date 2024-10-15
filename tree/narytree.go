package tree

import (
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ structures.Structure[int] = NewNAryTree[int](3)
var _ Tree[int] = NewNAryTree[int](3)

// NAryTree provides a generic complete n-ary tree implemented through the left-child right-sibiling rapresentation.
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
//
// It implements the interface [Tree].
type NAryTree[T any] struct {
	// contains filtered or unexported fields
	n    int
	root *Node[T]
	last *Node[T]
	len  int
}

// NewNAryTree returns a new [NAryTree] containing the elements c.
//
// n is the max number of children for a node.
//
// if no extra argument is passed, it will be created an empty [NAryTree].
//
// This function panics if n is less than 2.
func NewNAryTree[T any](n int, c ...T) *NAryTree[T] {
	return NewNAryTreeFromSlice(n, c)
}

// NewNAryTreeFromSlice returns a new [NAryTree] containing the elements of slice c.
//
// This function panics if n is less than 2.
func NewNAryTreeFromSlice[T any](n int, c []T) *NAryTree[T] {
	if n < 2 {
		panic(fmt.Sprintf("Cannot create a %v-ary tree", n))
	}
	tree := &NAryTree[T]{n: n, root: nil, last: nil, len: 0}
	if len(c) != 0 {
		tree.AddSlice(c)
	}
	return tree
}

// Len returns the length of t.
func (t *NAryTree[T]) Len() int {
	return t.len
}

// IsEmpty returns a bool which indicates if t is empty or not.
func (t *NAryTree[T]) IsEmpty() bool {
	return t.len == 0
}

// N returns the max number of children for a node.
func (t *NAryTree[T]) N() int {
	return t.n
}

// Root returns the root [Node] of t.
func (t *NAryTree[T]) Root() *Node[T] {
	return t.root
}

// Contains returns if e is present in t.
func (t *NAryTree[T]) Contains(e T) bool {
	fun := func(i *Node[T]) bool {
		return reflect.DeepEqual(e, i.Element())
	}
	if value, ok := interface{}(e).(util.Equaler); ok {
		fun = func(i *Node[T]) bool {
			return value.Equal(i.Element())
		}
	}
	return t.Any(t.root, func(i *Node[T]) bool {
		return fun(i)
	})
}

// ToSlice returns a slice which contains all elements of t.
func (t *NAryTree[T]) ToSlice() []T {
	slice := make([]T, 0)
	t.Each(t.root, func(i *Node[T]) { slice = append(slice, i.Element()) })
	return slice
}

// Add adds the elements e at t.
func (t *NAryTree[T]) Add(e ...T) {
	t.AddSlice(e)
}

// AddSlice adds the elements of e at t.
func (t *NAryTree[T]) AddSlice(e []T) {
	for _, i := range e {
		t.add(i)
	}
}

// Remove removes the element e if present.
// In that case, the method returns true.
func (t *NAryTree[T]) Remove(e T) bool {
	fun := func(i *Node[T]) bool {
		return reflect.DeepEqual(e, i.Element())
	}
	if value, ok := interface{}(e).(util.Equaler); ok {
		fun = func(i *Node[T]) bool {
			return value.Equal(i.Element())
		}
	}
	return t.Any(t.root, func(i *Node[T]) bool {
		if fun(i) {
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
func (t *NAryTree[T]) Each(node *Node[T], fun func(i *Node[T])) {
	if node == nil {
		return
	}
	fun(node)
	if node.Left() == nil {
		return
	}
	child := node.Left()
	for child != nil {
		t.Each(child, fun)
		child = child.Right()
	}
}

// Map executes fun for all elements of a subtree and returns a [NAryTree] containing the resulting elements.
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *NAryTree[T]) Map(node *Node[T], fun func(i *Node[T]) T) *NAryTree[T] {
	result := NewNAryTree[T](t.n)
	t.Each(t.root, func(i *Node[T]) {
		result.add(fun(i))
	})
	return result
}

// Filter returns a [NAryTree] containing the elements of a subtree that satisfy fun.
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *NAryTree[T]) Filter(node *Node[T], fun func(i *Node[T]) bool) *NAryTree[T] {
	result := NewNAryTree[T](t.n)
	t.Each(t.root, func(i *Node[T]) {
		if fun(i) {
			result.add(i.element)
		}
	})
	return result
}

// FilterMap executes fun for all elements of a subtree and returns a [NAryTree] containing the resulting elements that satisfy fun.
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *NAryTree[T]) FilterMap(node *Node[T], fun func(i *Node[T]) (T, bool)) *NAryTree[T] {
	result := NewNAryTree[T](t.n)
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
func (t *NAryTree[T]) Any(node *Node[T], fun func(i *Node[T]) bool) bool {
	if node == nil {
		return false
	}
	return t.Any(node.Left(), fun) || fun(node) || t.Any(node.Right(), fun)
}

// / All returns true if all elements of a subtree table satisfy fun.
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *NAryTree[T]) All(node *Node[T], fun func(i *Node[T]) bool) bool {
	if node == nil {
		return true
	}
	return t.All(node.Left(), fun) && fun(node) && t.All(node.Right(), fun)
}

// None returns true if none of the elements of a subtree table satisfies fun.
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *NAryTree[T]) None(node *Node[T], fun func(i *Node[T]) bool) bool {
	if node == nil {
		return true
	}
	return t.None(node.Left(), fun) && !fun(node) && t.None(node.Right(), fun)
}

// Count returns the number of elements of a subtree that satisfy fun.
//
// node is the root node of the subtree,
// fun is the function to be executed.
func (t *NAryTree[T]) Count(node *Node[T], fun func(i *Node[T]) bool) int {
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
func (t *NAryTree[T]) Clear() {
	t.root = nil
	t.len = 0
}

// Iter returns an [Iterator] which permits to iterate a [NAryTree].
//
//	for i := t.Iter(); !i.End(); i = i.Next() {
//		element := i.Element()
//		// Code
//	}
func (t *NAryTree[T]) Iter() Iterator[T] {
	return NewTreeIterator[T](t)
}

// Equal returns true if t and st are both [NAryTree] and their elements are equals.
// In any other case, it returns false.
func (t *NAryTree[T]) Equal(st any) bool {
	tree, ok := st.(*NAryTree[T])
	if ok && t != nil && tree != nil {
		if t.Len() != tree.Len() {
			return false
		}
		return t.All(t.root, func(i *Node[T]) bool {
			return tree.Contains(i.Element())
		})
	}
	return false
}

// Compare returns 0 if t and st have the same length,
// -1 if t is shorten than st,
// 1 if t is longer than st,
// -2 if st is not a [NAryTree] or if one between t and st is nil.
//
// If t and st have the same length, the result is the comparison
// between the max number of children of the two trees.
func (t *NAryTree[T]) Compare(st any) int {
	tree, ok := st.(*NAryTree[T])
	if ok && t != nil && tree != nil {
		if t.Len() < tree.Len() {
			return -1
		}
		if t.Len() > tree.Len() {
			return 1
		}
		return wrapper.Int(t.n).Compare(wrapper.Int(tree.N()))
	}
	return -2
}

// Hash returns the hash code of t.
func (t *NAryTree[T]) Hash() string {
	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("%v%v", check[1:], t.Len())
}

// String returns a rapresentation of t in the form of a string.
func (t *NAryTree[T]) String() string {
	check := reflect.TypeOf(new(T)).String()
	objects := make([]T, 0)
	t.Each(t.root, func(i *Node[T]) { objects = append(objects, i.Element()) })
	return fmt.Sprintf("%v-AryTree[%v]%v", t.n, check[1:], objects)
}

func (t *NAryTree[T]) add(e T) {
	t.len++
	if t.root == nil {
		t.root = NewNode[T](e, nil, nil, nil)
		t.last = t.root
		return
	}
	if !t.isParentComplete() {
		t.last.SetRight(NewNode[T](e, t.last.parent, nil, nil))
		t.last = t.last.Right()
		return
	} else {
		node := t.nextParent()
		node.SetLeft(NewNode[T](e, node, nil, nil))
		t.last = node.Left()
	}
}

func (t *NAryTree[T]) remove(node *Node[T]) {
	node.SetElement(t.last.Element())
	t.removeLast()
	t.len--
}

func (t *NAryTree[T]) isParentComplete() bool {
	if t.last.Parent() == nil {
		return true
	}
	node := t.last.parent.Left()
	if node == nil {
		return false
	}
	children := 1
	for node.Right() != nil {
		node = node.Right()
		children++
	}
	return children == t.N()
}

func (t *NAryTree[T]) nextParent() *Node[T] {
	if t.last == t.root {
		return t.root
	}
	node := t.last.Parent().Right()
	if node != nil {
		return node
	}
	node = t.root
	for node.Left() != nil {
		node = node.Left()
	}
	return node
}

func (t *NAryTree[T]) removeLast() {
	if t.last == t.root {
		t.Clear()
		return
	}
	node := t.last.Parent()
	t.last.SetParent(nil)
	if t.last == node.Left() {
		node.SetLeft(nil)
		t.last = node
		return
	}
	node = node.Left()
	for t.last != node.Right() {
		node = node.Right()
	}
	node.SetRight(nil)
	t.last = node
}
