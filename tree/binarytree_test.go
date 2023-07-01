package tree

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util/wrapper"
)

func TestNewBinaryTree(t *testing.T) {

	var tree structures.Structure[wrapper.Int] = NewBinaryTree[wrapper.Int]()

	if tree == nil {

		t.Log("tree is nil")
		t.Fail()

	}
	if tree.Len() != 0 {

		t.Log("length is not 0")
		t.Fail()

	}

}
func TestNewBinaryTreeFromSlice(t *testing.T) {

	var tree *BinaryTree[wrapper.Int] = NewBinaryTreeFromSlice[wrapper.Int]([]wrapper.Int{1, 8, -3, 5})

	if tree == nil {

		t.Log("tree is nil")
		t.Fail()

	}
	if tree.Len() != 4 {

		t.Log("length is not 4")
		t.Fail()

	}
	if !reflect.DeepEqual(tree.ToSlice(), []wrapper.Int{-3, 1, 5, 8}) {

		t.Log("tree is", tree)
		t.Fail()

	}

}
func TestContainsBinaryTree(t *testing.T) {

	var tree Tree[wrapper.Int] = NewBinaryTree[wrapper.Int](1, 8, -3, 5)

	if tree.Contains(-1) {

		t.Log("found -1 in tree")
		t.Fail()

	}
	if !tree.Contains(1) {

		t.Log("not found 1 in tree")
		t.Fail()

	}

}
func TestToSliceBinaryTree(t *testing.T) {

	var tree *BinaryTree[wrapper.Int] = NewBinaryTree[wrapper.Int](-3, 1, 5, 8)

	if slice := tree.ToSlice(); !reflect.DeepEqual(slice, []wrapper.Int{-3, 1, 5, 8}) {

		t.Log("slice is", slice)
		t.Fail()

	}

}
func TestAddBinaryTree(t *testing.T) {

	var tree *BinaryTree[wrapper.Int] = NewBinaryTree[wrapper.Int](-3, 1, 5, 8)

	tree.Add(-2, 6, 1)
	if slice := tree.ToSlice(); !reflect.DeepEqual(slice, []wrapper.Int{-3, -2, 1, 1, 5, 6, 8}) {

		t.Log("slice is", slice)
		t.Fail()

	}

}
func TestRemoveBinaryTree(t *testing.T) {

	var tree *BinaryTree[wrapper.Float32] = NewBinaryTree[wrapper.Float32](12.5, 7, -7.6, 3.4, 9, 0.9, 50, -120)

	if ok := tree.Remove(9); !ok {

		t.Log("not found 9 in tree")

	}
	if ok := tree.Remove(-9); ok {

		t.Log("found 9 in tree")

	}
	if slice := tree.ToSlice(); !reflect.DeepEqual(slice, []wrapper.Float32{-120, -7.6, 0.9, 3.4, 7, 12.5, 50}) {

		t.Log("slice is", slice)
		t.Fail()

	}

}
func TestIterBinaryTree(t *testing.T) {

	var tree *BinaryTree[wrapper.Float32] = NewBinaryTree[wrapper.Float32](12.5, 7, -7.6, 3.4, 9, 0.9, 50, -120)

	slice := tree.ToSlice()
	j := 0
	for i := tree.Iter(); !i.End() && j != tree.Len(); i = i.Next() {

		if !i.Element().Equal(slice[j]) {

			t.Log("element is", i.Element())
			t.Fail()

		}
		j++

	}

}
func TestEqualBinaryTree(t *testing.T) {

	var tree *BinaryTree[wrapper.Int] = NewBinaryTree[wrapper.Int](-3, 1, 5, 8)

	if !tree.Equal(NewBinaryTree[wrapper.Int](8, 5, -3, 1)) {

		t.Log("trees are not equals")
		t.Fail()

	}
	if tree.Equal(NewBinaryTree[wrapper.Int](8)) {

		t.Log("trees are equals")
		t.Fail()

	}
	if tree.Equal(NewBinaryTree[wrapper.Int](12, 5, -3, 1)) {

		t.Log("trees are equals")
		t.Fail()

	}

}
func TestCompareBinaryTree(t *testing.T) {

	var tree *BinaryTree[wrapper.Int] = NewBinaryTree[wrapper.Int](-3, 1, 5, 8)

	if tree.Compare(NewBinaryTree[wrapper.Int](8, 5, -3, 1)) != 0 {

		t.Log("compare is not 0")
		t.Fail()

	}
	if tree.Compare(NewBinaryTree[wrapper.Int](8)) != 1 {

		t.Log("compare is not -1")
		t.Fail()

	}
	if tree.Compare(NewBinaryTree[wrapper.Int](12, 5, -3, 1)) != -1 {

		t.Log("compare is not -1")
		t.Fail()

	}

}
