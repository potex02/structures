package queue

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
)

func TestNewDoubleQueue(t *testing.T) {

	var queue structures.Structure[float32] = NewDoubleQueue[float32]()

	if queue == nil {
		t.Log("queue is nil")
		t.Fail()
	}
	if queue.Len() != 0 {
		t.Log("length is not 0")
		t.Fail()
	}
}
func TestNewDoubleQueueFromSlice(t *testing.T) {

	var queue *DoubleQueue[float32] = NewDoubleQueueFromSlice([]float32{1.3, -2.5, 3.0, -4.0})

	if queue == nil {
		t.Log("queue is nil")
		t.Fail()
	}
	if queue.Len() != 4 {
		t.Log("length is not 4")
		t.Fail()
	}
	if !reflect.DeepEqual(queue.ToSlice(), []float32{1.3, -2.5, 3.0, -4.0}) {
		t.Log("queue objects are", queue.ToSlice())
		t.Fail()
	}
}
func TestHeadTailDoubleQueue(t *testing.T) {

	var queue *DoubleQueue[float32] = NewDoubleQueue[float32]()

	if _, ok := queue.Head(); ok {
		t.Log("the queue is not empty")
		t.Fail()
	}
	if _, ok := queue.Tail(); ok {
		t.Log("the queue is not empty")
		t.Fail()
	}
	queue = NewDoubleQueueFromSlice([]float32{1.3, -2.5, 3.0, -4.0})
	head, ok := queue.Head()
	if head != 1.3 {
		t.Log("Head is", head)
		t.Fail()
	}
	if !ok {
		t.Log("the queue is empty")
		t.Fail()
	}
	tail, ok := queue.Tail()
	if tail != -4.0 {
		t.Log("Tail is", tail)
		t.Fail()
	}
	if !ok {
		t.Log("the queue is empty")
		t.Fail()
	}
}
func TestPushDoubleQueue(t *testing.T) {

	var queue *DoubleQueue[float32] = NewDoubleQueue[float32]()

	queue.PushHead(1, 3)
	if !reflect.DeepEqual(queue.ToSlice(), []float32{3, 1}) {
		t.Log("queue is", queue.ToSlice())
		t.Fail()
	}
	queue.PushHead(-3)
	if e, _ := queue.Head(); e != -3 {
		t.Log("queue head is", queue.ToSlice())
		t.Fail()
	}
	queue.PushTail(-1.5)
	if !reflect.DeepEqual(queue.ToSlice(), []float32{-3, 3, 1, -1.5}) {
		t.Log("queue is", queue.ToSlice())
		t.Fail()
	}
	queue.PushTail(2, 12)
	if e, _ := queue.Tail(); e != 12 {
		t.Log("queue tail is", e)
		t.Fail()
	}
}
func TestPopDoubleLinkedyQueue(t *testing.T) {

	var queue BaseDoubleQueue[float32] = NewDoubleQueue[float32](1.3, 3, -2.5)

	e, ok := queue.PopHead()
	if !ok {
		t.Log("the queue is empty")
		t.Fail()
	}
	if e != 1.3 {
		t.Log("e is", e)
		t.Fail()
	}
	if queue.Len() != 2 {
		t.Log("length is not 2")
		t.Fail()
	}
	e, ok = queue.PopTail()
	if !ok {
		t.Log("the queue is empty")
		t.Fail()
	}
	if e != -2.5 {
		t.Log("e is", e)
		t.Fail()
	}
	if queue.Len() != 1 {
		t.Log("length is not 1")
		t.Fail()
	}
	e, ok = queue.PopTail()
	if !ok {
		t.Log("the queue is empty")
		t.Fail()
	}
	if e != 3 {
		t.Log("e is", e)
		t.Fail()
	}
	if !queue.IsEmpty() {
		t.Log("the queue is not empty")
		t.Fail()
	}
	queue = NewDoubleQueue[float32]()
	if _, ok := queue.PopHead(); ok {
		t.Log("the queue is empty")
		t.Fail()
	}
	if _, ok := queue.PopTail(); ok {
		t.Log("the queue is empty")
		t.Fail()
	}
}
func TestEqualDoubleQueue(t *testing.T) {

	var queue *DoubleQueue[float32] = NewDoubleQueue[float32](1.3, -2.5)

	if !queue.Equal(NewDoubleQueue[float32](1.3, -2.5)) {
		t.Log("queues are not equals")
		t.Fail()
	}
	if queue.Equal(NewDoubleQueue[float32](1.3, -2.5, -1)) {
		t.Log("queues are equals")
		t.Fail()
	}
}
