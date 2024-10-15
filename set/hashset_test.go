package set

import (
	"testing"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util/wrapper"
)

func TestNewHashSet(t *testing.T) {

	var set structures.Structure[wrapper.Int] = NewHashSet[wrapper.Int]()

	if set == nil {
		t.Log("set is nil")
		t.Fail()
	}
	if set.Len() != 0 {
		t.Log("length is not 0")
		t.Fail()
	}
}
func TestNewHashSetFromSlice(t *testing.T) {

	var set BaseSet[wrapper.Int] = NewHashSetFromSlice[wrapper.Int]([]wrapper.Int{1, 4, -5, 2, 1})

	if set == nil {
		t.Log("set is nil")
		t.Fail()
	}
	if set.Len() != 4 {
		t.Log("length is not 4")
		t.Fail()
	}
}
func TestContainsHashSet(t *testing.T) {

	var set Set[wrapper.Int] = NewHashSet[wrapper.Int](4, -5, 2, 1)

	if set.Contains(-1) {
		t.Log("found -1 in set")
		t.Fail()
	}
	if !set.Contains(2) {
		t.Log("not found 2 in set")
		t.Fail()
	}
}
func TestAddHashSet(t *testing.T) {

	var set *HashSet[wrapper.Int] = NewHashSet[wrapper.Int](4, -5, 2, 1)

	set.Add(-5)
	if set.Len() != 4 {
		t.Log("length is not 4")
		t.Fail()
	}
	set.Add(12, 56, 2, -4, 4)
	if set.Len() != 7 {
		t.Log("length is not 7")
		t.Fail()
	}
}
func TestRemoveHashSet(t *testing.T) {

	var set *HashSet[wrapper.Int] = NewHashSet[wrapper.Int](4, -5, 2, 1)

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
func TestEqualHashSet(t *testing.T) {

	var set Set[wrapper.Int] = NewHashSet[wrapper.Int](1, 2, 3, 5)
	var setTest Set[test] = NewHashSet[test](test{n1: 1, n2: 2}, test{n1: -2, n2: -4})

	if !set.Equal(NewHashSetFromSlice([]wrapper.Int{1, 2, 3, 5})) {
		t.Log("sets are not equals")
		t.Fail()
	}
	if set.Equal(NewHashSetFromSlice([]wrapper.Int{-1, 2, 3, 5})) {
		t.Log("sets are equals")
		t.Fail()
	}
	if !set.Equal(NewTreeSetFromSlice([]wrapper.Int{2, 1, 3, 5})) {
		t.Log("sets are not equals")
		t.Fail()
	}
	if set.Equal(NewTreeSetFromSlice([]wrapper.Int{-1, 2, 3, 5})) {
		t.Log("sets are equals")
		t.Fail()
	}
	if !setTest.Equal(NewHashSet[test](test{n1: 2, n2: 1}, test{n1: 0, n2: -6})) {
		t.Log("sets are not equals")
		t.Fail()
	}
	if setTest.Equal(NewTreeSet[test](test{n1: 1, n2: 1}, test{n1: -2, n2: -4})) {
		t.Log("sets are equals")
		t.Fail()
	}
}

type test struct {
	n1, n2 int
}

func (t test) Compare(o any) int {
	value, ok := o.(test)
	if !ok {
		return -2
	}
	return wrapper.Int(t.n1 + t.n2).Compare(wrapper.Int(value.n1 + value.n2))
}

func (t test) Hash() string {
	return wrapper.Int(t.n1 + t.n2).Hash()
}
