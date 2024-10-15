package table

import (
	"reflect"
	"testing"

	"github.com/potex02/structures/util/wrapper"
)

func TestNewStream(t *testing.T) {

	var table Table[wrapper.Int, int] = NewHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, -2, 30, 20, -10, 2, 1},
	)
	var stream *Stream[wrapper.Int, int] = NewStream[wrapper.Int, int](table, reflect.ValueOf(NewHashTable[wrapper.Int, int]))

	if stream == nil {
		t.Log("s is nil")
		t.Fail()
	}
	if _, ok := stream.Collect().(*HashTable[wrapper.Int, int]); !ok {
		t.Log("result is not an HashTable")
		t.Fail()
	}
	table = NewTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, -2, 30, 20, -10, 2, 1},
	)
	stream = NewStream[wrapper.Int, int](table, reflect.ValueOf(NewTreeTable[wrapper.Int, int]))
	if stream == nil {
		t.Log("s is nil")
		t.Fail()
	}
	if _, ok := stream.Collect().(*TreeTable[wrapper.Int, int]); !ok {
		t.Log("result is not a TreeTable")
		t.Fail()
	}
}
func TestMap(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	stream.Map(func(key wrapper.Int, element int) int {
		return element + key.ToValue()
	})
	if !stream.Collect().Equal(NewHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{13, 0, 20, 40, 20, 0, 13},
	)) {
		t.Log("result is", stream.Collect())
		t.Fail()
	}
}
func TestFilter(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	stream.Filter(func(key wrapper.Int, element int) bool {
		return element > 0
	})
	if !stream.Collect().Equal(NewTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, -10, 20, -2, 12},
		[]int{12, 30, 20, 2, 1},
	)) {
		t.Log("result is", stream.Collect())
		t.Fail()
	}
}
func TestFilterMap(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	stream.FilterMap(func(key wrapper.Int, element int) (int, bool) {
		return element + 10, key > 0
	})
	if !stream.Collect().Equal(NewTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, 20, 30, 12},
		[]int{22, 8, 30, 0, 11},
	)) {
		t.Log("result is", stream.Collect())
		t.Fail()
	}
}
func TestAny(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	if !stream.Any(func(key wrapper.Int, element int) bool {
		return element < 0
	}) {
		t.Log("result is false")
		t.Fail()
	}
	stream = NewHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, 2, 30, 20, 10, 2, 1},
	).Stream()
	if stream.Any(func(key wrapper.Int, element int) bool {
		return element < 0
	}) {
		t.Log("result is true")
		t.Fail()
	}
}
func TestAll(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	if stream.All(func(key wrapper.Int, element int) bool {
		return element > 0
	}) {
		t.Log("result is true")
		t.Fail()
	}
	stream = NewTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, 2, 30, 20, 10, 2, 1},
	).Stream()
	if !stream.All(func(key wrapper.Int, element int) bool {
		return element > 0
	}) {
		t.Log("result is false")
		t.Fail()
	}
}
func TestNone(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	if stream.None(func(key wrapper.Int, element int) bool {
		return element < 0
	}) {
		t.Log("result is true")
		t.Fail()
	}
	stream = NewTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, 2, 30, 20, 10, 2, 1},
	).Stream()
	if !stream.None(func(key wrapper.Int, element int) bool {
		return element < 0
	}) {
		t.Log("result is false")
		t.Fail()
	}
}
func TestCount(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	if stream.Count(func(key wrapper.Int, element int) bool {
		return key < 0
	}) != 2 {
		t.Log("result is not 2")
		t.Fail()
	}
}
func TestNewMultiStream(t *testing.T) {

	var table MultiTable[wrapper.Int, int] = NewMultiHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, 1, -10},
		[]int{12, -2, 30, 20, -10, 2, 1},
	)
	var stream *Stream[wrapper.Int, int] = NewStream[wrapper.Int, int](table, reflect.ValueOf(NewMultiHashTable[wrapper.Int, int]))

	if stream == nil {
		t.Log("s is nil")
		t.Fail()
	}
	if _, ok := stream.CollectMulti().(*MultiHashTable[wrapper.Int, int]); !ok {
		t.Log("result is not a MultiHashTable")
		t.Fail()
	}
	table = NewMultiTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, 1, -10},
		[]int{12, -2, 30, 20, -10, 2, 1},
	)
	stream = NewStream[wrapper.Int, int](table, reflect.ValueOf(NewMultiTreeTable[wrapper.Int, int]))
	if stream == nil {
		t.Log("s is nil")
		t.Fail()
	}
	if _, ok := stream.CollectMulti().(*MultiTreeTable[wrapper.Int, int]); !ok {
		t.Log("result is not a MultiTreeTable")
		t.Fail()
	}
}
func TestMapMulti(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewMultiHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, 1, -10},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	stream.Map(func(key wrapper.Int, element int) int {
		return element + key.ToValue()
	})
	if !stream.CollectMulti().Equal(NewMultiHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, 1, -10},
		[]int{13, 0, 20, 40, 20, 3, -9},
	)) {
		t.Log("result is", stream.CollectMulti())
		t.Fail()

	}

}
func TestFilterMulti(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewMultiTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, 1, -10},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	stream.Filter(func(key wrapper.Int, element int) bool {
		return element > 0
	})
	if !stream.CollectMulti().Equal(NewMultiTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, -10, 20, 1, -10},
		[]int{12, 30, 20, 2, 1},
	)) {
		t.Log("result is", stream.CollectMulti())
		t.Fail()
	}
}
func TestFilterMapMulti(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewMultiHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, 1, -10},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	stream.FilterMap(func(key wrapper.Int, element int) (int, bool) {
		return element + 10, key > 0
	})
	if !stream.CollectMulti().Equal(NewMultiTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, 20, 30, 1},
		[]int{22, 8, 30, 0, 12},
	)) {
		t.Log("result is", stream.CollectMulti())
		t.Fail()
	}
}
func TestAnyMulti(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewMultiHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, 1, -10},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	if !stream.Any(func(key wrapper.Int, element int) bool {
		return element < 0
	}) {
		t.Log("result is false")
		t.Fail()
	}
	stream = NewMultiHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, 1, -10},
		[]int{12, 2, 30, 20, 10, 2, 1},
	).Stream()
	if stream.Any(func(key wrapper.Int, element int) bool {
		return element < 0
	}) {
		t.Log("result is true")
		t.Fail()
	}
}
func TestAllMutli(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewMultiTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, 1, -10},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	if stream.All(func(key wrapper.Int, element int) bool {
		return element > 0
	}) {
		t.Log("result is true")
		t.Fail()
	}
	stream = NewMultiTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, 1, -10},
		[]int{12, 2, 30, 20, 10, 2, 1},
	).Stream()
	if !stream.All(func(key wrapper.Int, element int) bool {
		return element > 0
	}) {
		t.Log("result is false")
		t.Fail()
	}
}
func TestNoneMulti(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewMultiHashTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	if stream.None(func(key wrapper.Int, element int) bool {
		return element < 0
	}) {
		t.Log("result is true")
		t.Fail()
	}
	stream = NewMultiTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -2, 12},
		[]int{12, 2, 30, 20, 10, 2, 1},
	).Stream()
	if !stream.None(func(key wrapper.Int, element int) bool {
		return element < 0
	}) {
		t.Log("result is false")
		t.Fail()
	}
}
func TestCountMulti(t *testing.T) {

	var stream *Stream[wrapper.Int, int] = NewMultiTreeTableFromSlice[wrapper.Int, int](
		[]wrapper.Int{1, 2, -10, 20, 30, -1, -10},
		[]int{12, -2, 30, 20, -10, 2, 1},
	).Stream()

	if stream.Count(func(key wrapper.Int, element int) bool {
		return key < 0
	}) != 3 {
		t.Log("result is not 2")
		t.Fail()
	}
}
