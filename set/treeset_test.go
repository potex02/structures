package set

import (
	"testing"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util/wrapper"
)

func TestNewTreeSet(t *testing.T) {

	var set structures.Structure[wrapper.Int] = NewTreeSet[wrapper.Int]()

	if set == nil {

		t.Log("set is nil")
		t.Fail()

	}
	if set.Len() != 0 {

		t.Log("length is not 0")
		t.Fail()

	}

}
func TestNewTreeSetFromSlice(t *testing.T) {

	var set *TreeSet[wrapper.Int] = NewTreeSetFromSlice[wrapper.Int]([]wrapper.Int{1, 4, -5, 2, 1})

	if set == nil {

		t.Log("set is nil")
		t.Fail()

	}
	if set.Len() != 4 {

		t.Log("length is not 4")
		t.Fail()

	}

}
func TestContainsTreeSet(t *testing.T) {

	var set Set[wrapper.Int] = NewTreeSet[wrapper.Int](4, -5, 2, 1)

	if set.Contains(-1) {

		t.Log("found -1 in set")
		t.Fail()

	}
	if !set.Contains(2) {

		t.Log("not found 2 in set")
		t.Fail()

	}

}
func TestAddTreeSet(t *testing.T) {

	var set *TreeSet[wrapper.Int] = NewTreeSet[wrapper.Int](4, -5, 2, 1)

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
func TestRemoveTreeSet(t *testing.T) {

	var set *TreeSet[wrapper.Int] = NewTreeSet[wrapper.Int](4, -5, 2, 1)

	if ok := set.Remove(4); !ok {

		t.Log("not found 4 is set")
		t.Fail()

	}
	if ok := set.Remove(4); ok {

		t.Log("found 4 is set")
		t.Fail()

	}
	if ok := set.Remove(10); ok {

		t.Log("found 10 is set")
		t.Fail()

	}

}
func TestEqualTreeSet(t *testing.T) {

	var list Set[wrapper.Int] = NewTreeSet[wrapper.Int](1, 2, 3, 5)
	var listTest Set[test] = NewTreeSet[test](test{n1: 1, n2: 2}, test{n1: -2, n2: -4})

	if !list.Equal(NewTreeSetFromSlice([]wrapper.Int{1, 2, 3, 5})) {

		t.Log("sets are not equals")
		t.Fail()

	}
	if list.Equal(NewTreeSetFromSlice([]wrapper.Int{-1, 2, 3, 5})) {

		t.Log("sets are equals")
		t.Fail()

	}
	if !list.Equal(NewHashSetFromSlice([]wrapper.Int{2, 1, 3, 5})) {

		t.Log("sets are not equals")
		t.Fail()

	}
	if list.Equal(NewHashSetFromSlice([]wrapper.Int{-1, 2, 3, 5})) {

		t.Log("sets are equals")
		t.Fail()

	}
	if !listTest.Equal(NewTreeSet[test](test{n1: 2, n2: 1}, test{n1: 0, n2: -6})) {

		t.Log("sets are not equals")
		t.Fail()

	}
	if listTest.Equal(NewHashSet[test](test{n1: 1, n2: 1}, test{n1: -2, n2: -4})) {

		t.Log("sets are equals")
		t.Fail()

	}

}
