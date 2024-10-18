package queue

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
)

func TestNewQueue(t *testing.T) {

	var queue structures.Structure[float32] = NewQueue[float32]()

	if queue == nil {
		t.Log("queue is nil")
		t.Fail()
	}
	if queue.Len() != 0 {
		t.Log("length is not 0")
		t.Fail()
	}
}
func TestQueueFromSlice(t *testing.T) {

	var queue *Queue[float32] = NewQueueFromSlice([]float32{1.3, -2.5, 3.0, -4.0})

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
func TestHeadTailQueue(t *testing.T) {

	var queue *Queue[float32] = NewQueue[float32]()

	if _, ok := queue.Head(); ok {
		t.Log("the queue is not empty")
		t.Fail()
	}
	if _, ok := queue.Tail(); ok {
		t.Log("the queue is not empty")
		t.Fail()
	}
	queue = NewQueueFromSlice([]float32{1.3, -2.5, 3.0, -4.0})
	head, ok := queue.Head()
	if head != 1.3 {
		t.Log("head is", head)
		t.Fail()
	}
	if !ok {
		t.Log("the queue is empty")
		t.Fail()
	}
	tail, ok := queue.Tail()
	if tail != -4.0 {
		t.Log("tail is", tail)
		t.Fail()
	}
	if !ok {
		t.Log("the queue is empty")
		t.Fail()
	}
}
func TestPushQueue(t *testing.T) {

	var queue *Queue[float32] = NewQueue[float32]()

	queue.Push(1, 3)
	if !reflect.DeepEqual(queue.ToSlice(), []float32{1, 3}) {
		t.Log("queue is", queue.ToSlice())
		t.Fail()
	}
	queue.Push(-3)
	if e, _ := queue.Tail(); e != -3 {
		t.Log("queue tail is", e)
		t.Fail()
	}
}
func TestPopQueue(t *testing.T) {

	var queue BaseQueue[float32] = NewQueue[float32](1.3, -2.5)

	e, ok := queue.Pop()
	if !ok {
		t.Log("the queue is empty")
		t.Fail()
	}
	if e != 1.3 {
		t.Log("e is", e)
		t.Fail()
	}
	if queue.Len() != 1 {
		t.Log("length is not 1")
		t.Fail()
	}
	e, ok = queue.Pop()
	if !ok {
		t.Log("the queue is empty")
		t.Fail()
	}
	if e != -2.5 {
		t.Log("e is", e)
		t.Fail()
	}
	if !queue.IsEmpty() {
		t.Log("the queue is not empty")
		t.Fail()
	}
	if _, ok := queue.Pop(); ok {
		t.Log("the queue is not empty")
		t.Fail()
	}
}
func TestEqualQueue(t *testing.T) {

	var queue *Queue[float32] = NewQueue[float32](1.3, -2.5)

	if !queue.Equal(NewQueue[float32](1.3, -2.5)) {
		t.Log("queues are not equals")
		t.Fail()
	}
	if queue.Equal(NewQueue[float32](1.3, -2.5, -1)) {
		t.Log("queues are equals")
		t.Fail()
	}
}
