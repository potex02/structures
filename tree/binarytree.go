// package tree implements dynamic trees.
package tree

import (
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ structures.Structure[wrapper.Int] = NewBinaryTree[wrapper.Int]()

// BinaryTree provides a generic binary tree.
//
// the type T of the tree must implement [util.Comparer].
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
type BinaryTree[T util.Comparer] struct {
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

// Root returns the root node of t.
func (t *BinaryTree[T]) Root() *Node[T] {

	return t.root

}

// Contains returns if e is present in t.
func (t *BinaryTree[T]) Contains(e T) bool {

	result := false
	t.each(t.root, func(i *Node[T]) {
		if e.Compare(i.Element()) == 0 {
			result = true
		}
	})
	return result

}

// ToSLice returns a slice which contains all elements of t.
func (t *BinaryTree[T]) ToSlice() []T {

	slice := make([]T, 0)
	t.each(t.root, func(i *Node[T]) { slice = append(slice, i.Element()) })
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

// Remove removes the element if present.
// In that case, the method returns true.
func (t *BinaryTree[T]) Remove(e T) bool {

	result := false
	t.each(t.root, func(i *Node[T]) {
		if e.Compare(i.Element()) == 0 && !result {
			t.remove(i)
			result = true
		}
	})
	return result

}

// Clear removes all element from t.
func (t *BinaryTree[T]) Clear() {

	t.root = nil
	t.len = 0

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
		return t.all(t.root, func(i *Node[T]) bool {

			element, ok := interface{}(i).(util.Equaler)
			j++
			if ok {

				return element.Equal(others[j-1])

			}
			return reflect.DeepEqual(i.element, others[j-1])

		})

	}
	return false

}

// Compare returns -1 if t is shorten than st,
// 1 if t is longer than st,
// -2 if st is not a [BinaryTree] or if one between t and st is nil.
//
// If t and st have the same length, the result is the comparison
// between the first different element of the two tree.
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
		t.all(t.root, func(i *Node[T]) bool {

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
	t.each(t.root, func(i *Node[T]) { objects = append(objects, i.Element()) })
	return fmt.Sprintf("BinaryTree[%v]%v", check[1:], objects)

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

	if node.Right() == nil {

		if node == t.root {

			t.root = node.Left()

		} else {

			if node == node.Parent().Left() {

				node.Parent().SetLeft(nil)

			} else {

				node.Parent().SetRight(nil)

			}

		}
		if node.Left() != nil {

			node.Left().SetParent(node.Parent())

		}
		t.len--
		return

	}
	min := node.Right().Min()
	node.SetElement(min.element)
	if min == min.Parent().Left() {

		min.Parent().SetLeft(nil)

	} else {

		min.Parent().SetRight(nil)

	}
	min.SetParent(nil)
	t.len--

}

func (t *BinaryTree[T]) each(node *Node[T], fun func(i *Node[T])) {

	if node == nil {

		return

	}
	t.each(node.Left(), fun)
	fun(node)
	t.each(node.Right(), fun)

}

func (t *BinaryTree[T]) all(node *Node[T], fun func(i *Node[T]) bool) bool {

	if node == nil {

		return true

	}
	return t.all(node.Left(), fun) && fun(node) && t.all(node.Right(), fun)

}
