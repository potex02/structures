package set

import (
	"testing"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util/wrapper"
)

func TestNewMultiHashSet(t *testing.T) {

	var set structures.Structure[wrapper.Int] = NewMultiHashSet[wrapper.Int]()

	if set == nil {
		t.Log("set is nil")
		t.Fail()
	}
	if set.Len() != 0 {
		t.Log("length is not 0")
		t.Fail()
	}

}
func TestNewMultiHashSetFromSlice(t *testing.T) {

	var set BaseSet[wrapper.Int] = NewMultiHashSetFromSlice[wrapper.Int]([]wrapper.Int{1, 4, -5, 1})

	if set == nil {
		t.Log("set is nil")
		t.Fail()
	}
	if set.Len() != 4 {
		t.Log("length is not 4")
		t.Fail()
	}

}
func TestContainsMultiHashSet(t *testing.T) {

	var set MultiSet[wrapper.Int] = NewMultiHashSet[wrapper.Int](4, -5, 2, 1, 4)

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
func TestAddMultiHashSet(t *testing.T) {

	var set *MultiHashSet[wrapper.Int] = NewMultiHashSet[wrapper.Int](4, -5, 2, 1, 4)

	set.Add(12, 56, 2, -4, 4)
	if set.Len() != 10 {
		t.Log("length is not 10")
		t.Fail()
	}

}
func TestRemoveMultiHashSet(t *testing.T) {

	var set *MultiHashSet[wrapper.Int] = NewMultiHashSet[wrapper.Int](4, -5, 2, 1, 4)

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
func TestRemoveAllMultiHashSet(t *testing.T) {

	var set *MultiHashSet[wrapper.Int] = NewMultiHashSet[wrapper.Int](4, -5, 2, 1, 4, 10, -5, -5, 3)

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
func TestCountMultiHashSet(t *testing.T) {

	var set *MultiHashSet[wrapper.Int] = NewMultiHashSet[wrapper.Int](4, -5, 2, 1, 4)

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
func TestToSetMultiHashSet(t *testing.T) {

	var set *MultiHashSet[wrapper.Int] = NewMultiHashSet[wrapper.Int](4, -5, 2, 1, 4)

	if !set.ToSet().Equal(NewHashSet[wrapper.Int](-5, 2, 1, 4)) {
		t.Log("set is", set.ToSet())
	}
}
func TestEqualMultiHashSet(t *testing.T) {

	var set MultiSet[wrapper.Int] = NewMultiHashSet[wrapper.Int](1, 2, 3, 5, 2)
	var setTest MultiSet[test] = NewMultiHashSet[test](test{n1: 1, n2: 2}, test{n1: -2, n2: -4}, test{n1: -2, n2: -4})

	if !set.Equal(NewMultiHashSetFromSlice([]wrapper.Int{1, 2, 3, 2, 5})) {
		t.Log("sets are not equals")
		t.Fail()
	}
	if set.Equal(NewMultiHashSetFromSlice([]wrapper.Int{-1, 2, 3, 5, 5})) {
		t.Log("sets are equals")
		t.Fail()
	}
	if set.Equal(NewMultiTreeSetFromSlice([]wrapper.Int{-1, 2, 3, 5})) {
		t.Log("sets are equals")
		t.Fail()
	}
	if !setTest.Equal(NewMultiHashSet[test](test{n1: 2, n2: 1}, test{n1: 0, n2: -6}, test{n1: -10, n2: 4})) {
		t.Log("sets are not equals")
		t.Fail()
	}
	if setTest.Equal(NewMultiTreeSet[test](test{n1: 1, n2: 1}, test{n1: -2, n2: -4}, test{n1: 0, n2: -6}, test{n1: 10, n2: 4})) {
		t.Log("sets are equals")
		t.Fail()
	}
}
