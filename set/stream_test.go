package set

import (
	"reflect"
	"testing"

	"github.com/potex02/structures/util/wrapper"
)

func TestNewStream(t *testing.T) {

	var set Set[wrapper.Int] = NewHashSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12)
	var stream *Stream[wrapper.Int] = NewStream[wrapper.Int](set, reflect.ValueOf(NewHashSet[wrapper.Int]))

	if stream == nil {

		t.Log("s is nil")
		t.Fail()

	}
	if _, ok := stream.Collect().(*HashSet[wrapper.Int]); !ok {

		t.Log("result is not an HashSet")
		t.Fail()

	}
	set = NewTreeSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12)
	stream = NewStream[wrapper.Int](set, reflect.ValueOf(NewTreeSet[wrapper.Int]))
	if stream == nil {

		t.Log("s is nil")
		t.Fail()

	}
	if _, ok := stream.Collect().(*TreeSet[wrapper.Int]); !ok {

		t.Log("result is not a TreeSet")
		t.Fail()

	}

}
func TestMap(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewHashSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12).Stream()

	stream.Map(func(element wrapper.Int) wrapper.Int {
		return element * 2
	})
	if !stream.Collect().Equal(NewHashSet[wrapper.Int](2, 4, -20, 40, 60, -4, 24)) {

		t.Log("result is", stream.Collect())
		t.Fail()

	}

}
func TestFilter(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewTreeSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12).Stream()

	stream.Filter(func(element wrapper.Int) bool {
		return element > 0
	})
	if !stream.Collect().Equal(NewTreeSet[wrapper.Int](1, 2, 20, 30, 12)) {

		t.Log("result is", stream.Collect())
		t.Fail()

	}

}
func TestFilterMap(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewHashSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12).Stream()

	stream.FilterMap(func(element wrapper.Int) (wrapper.Int, bool) {
		return element + 10, element > 0
	})
	if !stream.Collect().Equal(NewTreeSet[wrapper.Int](11, 12, 30, 40, 22)) {

		t.Log("result is", stream.Collect())
		t.Fail()

	}

}
func TestAny(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewHashSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12).Stream()

	if !stream.Any(func(element wrapper.Int) bool {
		return element < 0
	}) {

		t.Log("result is false")
		t.Fail()

	}
	stream = NewHashSet[wrapper.Int](1, 2, 20, 30, 12).Stream()
	if stream.Any(func(element wrapper.Int) bool {
		return element < 0
	}) {

		t.Log("result is true")
		t.Fail()

	}

}
func TestAll(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewTreeSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12).Stream()

	if stream.All(func(element wrapper.Int) bool {
		return element > 0
	}) {

		t.Log("result is true")
		t.Fail()

	}
	stream = NewTreeSet[wrapper.Int](1, 2, 20, 30, 12).Stream()
	if !stream.All(func(element wrapper.Int) bool {
		return element > 0
	}) {

		t.Log("result is false")
		t.Fail()

	}

}
func TestNone(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewHashSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12).Stream()

	if stream.None(func(element wrapper.Int) bool {
		return element < 0
	}) {

		t.Log("result is true")
		t.Fail()

	}
	stream = NewTreeSet[wrapper.Int](1, 2, 20, 30, 12).Stream()
	if !stream.None(func(element wrapper.Int) bool {
		return element < 0
	}) {

		t.Log("result is false")
		t.Fail()

	}

}
func TestCount(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewTreeSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12).Stream()

	if stream.Count(func(element wrapper.Int) bool {
		return element > 0
	}) != 5 {

		t.Log("result is not 5")
		t.Fail()

	}

}
