package tree

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
)

func TestNewNAryTree(t *testing.T) {

	var tree structures.Structure[int] = NewNAryTree[int](3)

	if tree == nil {
		t.Log("tree is nil")
		t.Fail()
	}
	if tree.Len() != 0 {
		t.Log("length is not 0")
		t.Fail()
	}
}
func TestNewNAryTreeFromSlice(t *testing.T) {

	var tree *NAryTree[int] = NewNAryTreeFromSlice[int](3, []int{1, 8, -3, 5})

	if tree == nil {
		t.Log("tree is nil")
		t.Fail()
	}
	if tree.Len() != 4 {
		t.Log("length is not 4")
		t.Fail()
	}
	if !reflect.DeepEqual(tree.ToSlice(), []int{1, 8, -3, 5}) {
		t.Log("tree is", tree)
		t.Fail()
	}
}
func TestContainsNAryTree(t *testing.T) {

	var tree Tree[int] = NewNAryTree[int](2, 1, 8, -3, 5)

	if tree.Contains(-1) {
		t.Log("found -1 in tree")
		t.Fail()
	}
	if !tree.Contains(1) {
		t.Log("not found 1 in tree")
		t.Fail()
	}
}
func TestToSliceNAryTree(t *testing.T) {

	var tree *NAryTree[int] = NewNAryTree[int](2, -3, 1, 5, 8)

	if slice := tree.ToSlice(); !reflect.DeepEqual(slice, []int{-3, 1, 8, 5}) {
		t.Log("slice is", slice)
		t.Fail()
	}
}
func TestAddNAryTree(t *testing.T) {

	var tree *NAryTree[int] = NewNAryTree[int](3, -3, 1, 5, 8)

	tree.Add(-2, 6, 1)
	if slice := tree.ToSlice(); !reflect.DeepEqual(slice, []int{-3, 1, -2, 6, 1, 5, 8}) {
		t.Log("slice is", slice)
		t.Fail()
	}
}
func TestRemoveNAryTree(t *testing.T) {

	var tree *NAryTree[float32] = NewNAryTree[float32](3, 12.5, 7, -7.6, 3.4, 9, 0.9, 50, -120)

	if ok := tree.Remove(9); !ok {
		t.Log("not found 9 in tree")
	}
	if ok := tree.Remove(-9); ok {
		t.Log("found 9 in tree")
	}
	if slice := tree.ToSlice(); !reflect.DeepEqual(slice, []float32{12.5, 7, -120, 0.9, 50, -7.6, 3.4}) {
		t.Log("slice is", slice)
		t.Fail()
	}
}
func TestIterNAryTree(t *testing.T) {

	var tree *NAryTree[float32] = NewNAryTree[float32](3, 3, 12.5, 7, -7.6, 3.4, 9, 0.9, 50, -120)

	slice := tree.ToSlice()
	j := 0
	for i := tree.Iter(); !i.End() && j != tree.Len(); i = i.Next() {
		if reflect.DeepEqual(i, slice[j]) {
			t.Log("element is", i.Element())
			t.Fail()
		}
		j++
	}
}
func TestEqualNAryTree(t *testing.T) {

	var tree *NAryTree[int] = NewNAryTree[int](3, -3, 1, 5, 8)

	if !tree.Equal(NewNAryTree[int](3, 8, 5, -3, 1)) {
		t.Log("trees are not equals")
		t.Fail()
	}
	if tree.Equal(NewNAryTree[int](2, 8)) {
		t.Log("trees are equals")
		t.Fail()
	}
	if tree.Equal(NewNAryTree[int](3, 12, 5, -3, 1)) {
		t.Log("trees are equals")
		t.Fail()
	}
}
func TestCompareNAryTree(t *testing.T) {

	var tree *NAryTree[int] = NewNAryTree[int](3, -3, 1, 5, 8)

	if tree.Compare(NewNAryTree[int](3, 8, 5, -3, 1)) != 0 {
		t.Log("compare is not 0")
		t.Fail()
	}
	if tree.Compare(NewNAryTree[int](8)) != 1 {
		t.Log("compare is not -1")
		t.Fail()
	}
	if tree.Compare(NewNAryTree[int](12, 5, -3, 1, -1)) != -1 {
		t.Log("compare is not -1")
		t.Fail()
	}
}
