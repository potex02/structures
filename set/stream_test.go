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
	if !stream.CollectBase().Equal(NewHashSet[wrapper.Int](2, 4, -20, 40, 60, -4, 24)) {
		t.Log("result is", stream.CollectBase())
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
func TestUnion(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewTreeSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12).Stream()

	stream.Union(NewTreeSet[wrapper.Int](-10, -20, 100, -50))
	if !stream.Collect().Equal(NewTreeSet[wrapper.Int](-10, 20, 1, 2, 30, -2, -20, 100, 12, -50)) {
		t.Log("result is", stream.Collect())
		t.Fail()
	}
}
func TestIntersection(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewHashSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12).Stream()

	stream.Intersection(NewTreeSet[wrapper.Int](-10, -20, -2, 100, 30, -50))
	if !stream.Collect().Equal(NewHashSet[wrapper.Int](-10, 30, -2)) {
		t.Log("result is", stream.Collect())
		t.Fail()
	}
}
func TestDifference(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewTreeSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12).Stream()

	stream.Difference(NewMultiTreeSet[wrapper.Int](-10, -20, -2, 100, 30, -50))
	if !stream.Collect().Equal(NewHashSet[wrapper.Int](1, 2, 20, 12)) {
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
func TestNewMultiStream(t *testing.T) {

	var set MultiSet[wrapper.Int] = NewMultiHashSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12)
	var stream *Stream[wrapper.Int] = NewStream[wrapper.Int](set, reflect.ValueOf(NewMultiHashSet[wrapper.Int]))

	if stream == nil {
		t.Log("s is nil")
		t.Fail()
	}
	if _, ok := stream.CollectMulti().(*MultiHashSet[wrapper.Int]); !ok {
		t.Log("result is not an MultiHashSet")
		t.Fail()
	}
	set = NewMultiTreeSet[wrapper.Int](1, 2, -10, 20, 30, -2, 12)
	stream = NewStream[wrapper.Int](set, reflect.ValueOf(NewMultiTreeSet[wrapper.Int]))
	if stream == nil {
		t.Log("s is nil")
		t.Fail()
	}
	if _, ok := stream.CollectMulti().(*MultiTreeSet[wrapper.Int]); !ok {
		t.Log("result is not a TreeSet")
		t.Fail()
	}
}
func TestMapMulti(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewMultiHashSet[wrapper.Int](1, 2, -10, 20, 2, -2, 12).Stream()

	stream.Map(func(element wrapper.Int) wrapper.Int {
		return element * 2
	})
	if !stream.CollectBase().Equal(NewMultiHashSet[wrapper.Int](2, 4, -20, 40, 4, -4, 24)) {
		t.Log("result is", stream.CollectBase())
		t.Fail()
	}
}
func TestFilterMulti(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewMultiTreeSet[wrapper.Int](1, 2, -10, 20, 2, -2, 12).Stream()

	stream.Filter(func(element wrapper.Int) bool {
		return element > 0
	})
	if !stream.CollectMulti().Equal(NewMultiTreeSet[wrapper.Int](1, 2, 20, 2, 12)) {
		t.Log("result is", stream.CollectMulti())
		t.Fail()
	}
}
func TestFilterMapMulti(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewMultiHashSet[wrapper.Int](1, 2, -10, 20, 2, -2, 12).Stream()

	stream.FilterMap(func(element wrapper.Int) (wrapper.Int, bool) {
		return element + 10, element > 0
	})
	if !stream.CollectMulti().Equal(NewMultiTreeSet[wrapper.Int](11, 12, 30, 12, 22)) {
		t.Log("result is", stream.CollectMulti())
		t.Fail()
	}
}
func TestUnionMulti(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewMultiTreeSet[wrapper.Int](1, 2, -10, 20, 2, -2, 12).Stream()

	stream.Union(NewMultiTreeSet[wrapper.Int](-10, -20, 100, -50))
	if !stream.CollectMulti().Equal(NewMultiTreeSet[wrapper.Int](1, 2, -10, 20, 2, -2, 12, -10, -20, 100, -50)) {
		t.Log("result is", stream.CollectMulti())
		t.Fail()
	}
}
func TestIntersectionMulti(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewMultiHashSet[wrapper.Int](1, 2, -10, 20, 2, -2, -2).Stream()

	stream.Intersection(NewTreeSet[wrapper.Int](-10, -20, -2, 100, 30, -50))
	if !stream.CollectMulti().Equal(NewMultiHashSet[wrapper.Int](-10, -2, -2)) {
		t.Log("result is", stream.CollectMulti())
		t.Fail()
	}
}
func TestDifferenceMulti(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewMultiTreeSet[wrapper.Int](1, 2, -10, 20, 2, -2, 12).Stream()

	stream.Difference(NewTreeSet[wrapper.Int](-10, -20, -2, 100, 30, -50))
	if !stream.CollectMulti().Equal(NewMultiHashSet[wrapper.Int](1, 2, 20, 12, 2)) {
		t.Log("result is", stream.CollectMulti())
		t.Fail()
	}
}
func TestAnyMulti(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewMultiHashSet[wrapper.Int](1, 2, -10, 20, 2, -2, 12).Stream()

	if !stream.Any(func(element wrapper.Int) bool {
		return element < 0
	}) {
		t.Log("result is false")
		t.Fail()
	}
	stream = NewMultiHashSet[wrapper.Int](1, 2, 20, 30, 12, 30).Stream()
	if stream.Any(func(element wrapper.Int) bool {
		return element < 0
	}) {
		t.Log("result is true")
		t.Fail()
	}
}
func TestAllMulti(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewMultiTreeSet[wrapper.Int](1, 2, -10, 20, 2, -2, 12).Stream()

	if stream.All(func(element wrapper.Int) bool {
		return element > 0
	}) {
		t.Log("result is true")
		t.Fail()
	}
	stream = NewMultiTreeSet[wrapper.Int](1, 2, 20, 30, 12, 30).Stream()
	if !stream.All(func(element wrapper.Int) bool {
		return element > 0
	}) {
		t.Log("result is false")
		t.Fail()
	}
}
func TestNoneMulti(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewMultiHashSet[wrapper.Int](1, 2, -10, 20, 2, -2, 12).Stream()

	if stream.None(func(element wrapper.Int) bool {
		return element < 0
	}) {
		t.Log("result is true")
		t.Fail()
	}
	stream = NewMultiTreeSet[wrapper.Int](1, 2, 20, 30, 12, 30).Stream()
	if !stream.None(func(element wrapper.Int) bool {
		return element < 0
	}) {
		t.Log("result is false")
		t.Fail()
	}
}
func TestCountMulti(t *testing.T) {

	var stream *Stream[wrapper.Int] = NewMultiTreeSet[wrapper.Int](1, 2, -10, 20, 2, -2, 12).Stream()

	if stream.Count(func(element wrapper.Int) bool {
		return element > 0
	}) != 5 {
		t.Log("result is not 5")
		t.Fail()
	}
}
