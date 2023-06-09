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

	var set *HashSet[wrapper.Int] = NewHashSetFromSlice[wrapper.Int]([]wrapper.Int{1, 4, -5, 2, 1})

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
