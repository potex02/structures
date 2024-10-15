package set

import (
	"testing"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util/wrapper"
)

func TestNewMultiTreeSet(t *testing.T) {

	var set structures.Structure[wrapper.Int] = NewMultiTreeSet[wrapper.Int]()

	if set == nil {
		t.Log("set is nil")
		t.Fail()
	}
	if set.Len() != 0 {
		t.Log("length is not 0")
		t.Fail()
	}
}
func TestNewMultiTreeSetFromSlice(t *testing.T) {

	var set BaseSet[wrapper.Int] = NewMultiTreeSetFromSlice[wrapper.Int]([]wrapper.Int{1, 4, -5, 1})

	if set == nil {
		t.Log("set is nil")
		t.Fail()
	}
	if set.Len() != 4 {
		t.Log("length is not 4")
		t.Fail()
	}
}
func TestContainsMultiTreeSet(t *testing.T) {

	var set MultiSet[wrapper.Int] = NewMultiTreeSet[wrapper.Int](4, -5, 2, 1, 4)

	if set.Contains(-1) {
		t.Log("found -1 in set")
		t.Fail()
	}
	if !set.Contains(2) {
		t.Log("not found 2 in set")
		t.Fail()
	}
	if !set.Contains(4) {
		t.Log("not found 4 in set")
		t.Fail()
	}
}
func TestAddMultiTreeSet(t *testing.T) {

	var set *MultiTreeSet[wrapper.Int] = NewMultiTreeSet[wrapper.Int](4, -5, 2, 1, 4)

	set.Add(12, 56, 2, -4, 4)
	if set.Len() != 10 {
		t.Log("length is not 10")
		t.Fail()
	}
}
func TestRemoveMultiTreeSet(t *testing.T) {

	var set *MultiTreeSet[wrapper.Int] = NewMultiTreeSet[wrapper.Int](4, -5, 2, 1, 4)

	if !set.Remove(4) {
		t.Log("not found 4 is set")
		t.Fail()
	}
	if !set.Remove(4) {
		t.Log("not found 4 is set")
		t.Fail()
	}
	if set.Remove(4) {
		t.Log("found 4 is set")
		t.Fail()
	}
	if set.Remove(10) {
		t.Log("found 10 is set")
		t.Fail()
	}
}
func TestRemoveAllMultiTreeSet(t *testing.T) {

	var set *MultiTreeSet[wrapper.Int] = NewMultiTreeSet[wrapper.Int](4, -5, 2, 1, 4, 10, -5, 1, -5)

	set.RemoveAll(-5)
	if set.Len() != 6 {
		t.Log("length is not 6")
		t.Fail()
	}
	set.RemoveAll(10)
	if set.Len() != 5 {
		t.Log("length is not 5")
		t.Fail()
	}
	set.RemoveAll(0)
	if set.Len() != 5 {
		t.Log("length is not 5")
		t.Fail()
	}
}
func TestCountMultiTreeSet(t *testing.T) {

	var set *MultiTreeSet[wrapper.Int] = NewMultiTreeSet[wrapper.Int](4, -5, 2, 1, 4)

	if set.Count(4) != 2 {
		t.Log("not found 2 occurrences")
		t.Fail()
	}
	if set.Count(-5) != 1 {
		t.Log("not found 1 occurrence")
		t.Fail()
	}
	if set.Count(10) != 0 {
		t.Log("not found 0 occurrences")
		t.Fail()
	}
}
func TestToSetMultiTreeSet(t *testing.T) {

	var set *MultiTreeSet[wrapper.Int] = NewMultiTreeSet[wrapper.Int](4, -5, 2, 1, 4)

	if !set.ToSet().Equal(NewTreeSet[wrapper.Int](-5, 2, 1, 4)) {
		t.Log("set is", set.ToSet())
	}
}
func TestEqualMultiTreeSet(t *testing.T) {

	var set MultiSet[wrapper.Int] = NewMultiTreeSet[wrapper.Int](1, 2, 3, 5, 2)
	var setTest MultiSet[test] = NewMultiTreeSet[test](test{n1: 1, n2: 2}, test{n1: -2, n2: -4}, test{n1: -2, n2: -4})

	if !set.Equal(NewMultiTreeSetFromSlice([]wrapper.Int{1, 2, 3, 2, 5})) {
		t.Log("sets are not equals")
		t.Fail()
	}
	if set.Equal(NewMultiTreeSetFromSlice([]wrapper.Int{-1, 2, 3, 5, 5})) {
		t.Log("sets are equals")
		t.Fail()
	}
	if set.Equal(NewMultiHashSetFromSlice([]wrapper.Int{-1, 2, 3, 5})) {
		t.Log("sets are equals")
		t.Fail()
	}
	if !setTest.Equal(NewMultiTreeSet[test](test{n1: 2, n2: 1}, test{n1: 0, n2: -6}, test{n1: -10, n2: 4})) {
		t.Log("sets are not equals")
		t.Fail()
	}
	if setTest.Equal(NewMultiHashSet[test](test{n1: 1, n2: 1}, test{n1: -2, n2: -4}, test{n1: 0, n2: -6}, test{n1: 10, n2: 4})) {
		t.Log("sets are equals")
		t.Fail()
	}
}
