package list

import (
	"reflect"
	"testing"
)

func TestNewStream(t *testing.T) {

	var list List[int] = NewArrayList[int](1, 2, -10, 20, 30, -2, 12)
	var stream *Stream[int] = NewStream[int](list, reflect.ValueOf(NewArrayList[int]))

	if stream == nil {
		t.Log("s is nil")
		t.Fail()
	}
	if _, ok := stream.Collect().(*ArrayList[int]); !ok {
		t.Log("result is not an ArrayList")
		t.Fail()
	}
	list = NewLinkedList[int](1, 2, -10, 20, 30, -2, 12)
	stream = NewStream[int](list, reflect.ValueOf(NewLinkedList[int]))
	if stream == nil {
		t.Log("s is nil")
		t.Fail()
	}
	if _, ok := stream.Collect().(*LinkedList[int]); !ok {
		t.Log("result is not a LinkedList")
		t.Fail()
	}
}
func TestMap(t *testing.T) {

	var stream *Stream[int] = NewArrayList[int](1, 2, -10, 20, 30, -2, 12).Stream()

	stream.Map(func(index int, element int) int {
		return element*2 + index
	})
	if !stream.Collect().Equal(NewArrayList[int](2, 5, -18, 43, 64, 1, 30)) {
		t.Log("result is", stream.Collect())
		t.Fail()
	}
}
func TestFilter(t *testing.T) {

	var stream *Stream[int] = NewLinkedList[int](1, 2, -10, 20, 30, -2, 12).Stream()

	stream.Filter(func(index int, element int) bool {
		return element > 0
	})
	if !stream.Collect().Equal(NewLinkedList[int](1, 2, 20, 30, 12)) {
		t.Log("result is", stream.Collect())
		t.Fail()
	}
}
func TestFilterMap(t *testing.T) {

	var stream *Stream[int] = NewArrayList[int](1, 2, -10, 20, 30, -2, 12).Stream()

	stream.FilterMap(func(index int, element int) (int, bool) {
		return element + 10, element > 0
	})
	if !stream.Collect().Equal(NewLinkedList[int](11, 12, 30, 40, 22)) {
		t.Log("result is", stream.Collect())
		t.Fail()
	}
}
func TestAny(t *testing.T) {

	var stream *Stream[int] = NewArrayList[int](1, 2, -10, 20, 30, -2, 12).Stream()

	if !stream.Any(func(index int, element int) bool {
		return element < 0
	}) {
		t.Log("result is false")
		t.Fail()
	}
	stream = NewArrayList[int](1, 2, 20, 30, 12).Stream()
	if stream.Any(func(index int, element int) bool {
		return element < 0
	}) {
		t.Log("result is true")
		t.Fail()
	}
}
func TestAll(t *testing.T) {

	var stream *Stream[int] = NewLinkedList[int](1, 2, -10, 20, 30, -2, 12).Stream()

	if stream.All(func(index int, element int) bool {
		return element > 0
	}) {
		t.Log("result is true")
		t.Fail()
	}
	stream = NewLinkedList[int](1, 2, 20, 30, 12).Stream()
	if !stream.All(func(index int, element int) bool {
		return element > 0
	}) {
		t.Log("result is false")
		t.Fail()
	}
}
func TestNone(t *testing.T) {

	var stream *Stream[int] = NewArrayList[int](1, 2, -10, 20, 30, -2, 12).Stream()

	if stream.None(func(index int, element int) bool {
		return element < 0
	}) {
		t.Log("result is true")
		t.Fail()
	}
	stream = NewLinkedList[int](1, 2, 20, 30, 12).Stream()
	if !stream.None(func(index int, element int) bool {
		return element < 0
	}) {
		t.Log("result is false")
		t.Fail()
	}
}
func TestCount(t *testing.T) {

	var stream *Stream[int] = NewLinkedList[int](1, 2, -10, 20, 30, -2, 12).Stream()

	if stream.Count(func(index int, element int) bool {
		return element > 0
	}) != 5 {
		t.Log("result is not 5")
		t.Fail()
	}
}
