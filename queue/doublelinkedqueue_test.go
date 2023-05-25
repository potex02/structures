package queue

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
)

func TestNewDoubleLinkedQueue(t *testing.T) {

	var queue structures.Structure[float32] = NewDoubleLinkedQueue[float32]()

	if queue == nil {

		t.Log("queue is nil")
		t.Fail()

	}
	if queue.Len() != 0 {

		t.Log("length is not 0")
		t.Fail()

	}

}
func TestNewDoubleLinkedQueueFromSlice(t *testing.T) {

	var queue *DoubleLinkedQueue[float32] = NewDoubleLinkedQueueFromSlice([]float32{1.3, -2.5, 3.0, -4.0})

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

func TestHeadTailDoubleLinkedQueue(t *testing.T) {

	var queue *DoubleLinkedQueue[float32] = NewDoubleLinkedQueue[float32]()

	_, err := queue.Head()
	if err == nil {

		t.Log("the queue is not empty")
		t.Fail()

	}
	_, err = queue.Tail()
	if err == nil {

		t.Log("the queue is not empty")
		t.Fail()

	}
	queue = NewDoubleLinkedQueueFromSlice([]float32{1.3, -2.5, 3.0, -4.0})
	head, err := queue.Head()
	if head != 1.3 {

		t.Log("Head is", head)
		t.Fail()

	}
	if err != nil {

		t.Log("err is", err)
		t.Fail()

	}
	tail, err := queue.Tail()
	if tail != -4.0 {

		t.Log("Tail is", tail)
		t.Fail()

	}
	if err != nil {

		t.Log("err is", err)
		t.Fail()

	}

}

func TestPushDoubleLinkedQueue(t *testing.T) {

	var queue *DoubleLinkedQueue[float32] = NewDoubleLinkedQueue[float32]()

	queue.PushHead(1, 3)
	if !reflect.DeepEqual(queue.ToSlice(), []float32{3, 1}) {

		t.Log("queue is", queue.ToSlice())
		t.Fail()

	}
	queue.PushHead(-3)
	e, _ := queue.Head()
	if e != -3 {

		t.Log("queue head is", queue.ToSlice())
		t.Fail()

	}
	queue.PushTail(-1.5)
	if !reflect.DeepEqual(queue.ToSlice(), []float32{-3, 3, 1, -1.5}) {

		t.Log("queue is", queue.ToSlice())
		t.Fail()

	}
	queue.PushTail(2, 12)
	e, _ = queue.Tail()
	if e != 12 {

		t.Log("queue tail is", e)
		t.Fail()

	}

}

func TestPopDoubleLinkedyQueue(t *testing.T) {

	var queue DoubleQueue[float32] = NewDoubleLinkedQueue[float32](1.3, 3, -2.5)

	e, err := queue.PopHead()
	if err != nil {

		t.Log("err is", err)
		t.Fail()

	}
	if e != 1.3 {

		t.Log("e is", e)
		t.Fail()

	}
	if queue.Len() != 2 {

		t.Log("Size is not 2")
		t.Fail()

	}
	e, err = queue.PopTail()
	if err != nil {

		t.Log("err is", err)
		t.Fail()

	}
	if e != -2.5 {

		t.Log("e is", e)
		t.Fail()

	}
	if queue.Len() != 1 {

		t.Log("Size is not 1")
		t.Fail()

	}
	e, err = queue.PopTail()
	if err != nil {

		t.Log("err is", err)
		t.Fail()

	}
	if e != 3 {

		t.Log("e is", e)
		t.Fail()

	}
	if !queue.IsEmpty() {

		t.Log("queue not empty")
		t.Fail()

	}
	queue = NewDoubleLinkedQueue[float32]()
	_, err = queue.PopHead()
	if err == nil {

		t.Log("err is nil")
		t.Fail()

	}
	_, err = queue.PopTail()
	if err == nil {

		t.Log("err is nil")
		t.Fail()

	}

}

func TestEqualsDoubleLinkedQueue(t *testing.T) {

	var queue *DoubleLinkedQueue[float32] = NewDoubleLinkedQueue[float32](1.3, -2.5)

	if !queue.Equals(NewDoubleArrayQueue[float32](1.3, -2.5)) {

		t.Log("queues are not equals")
		t.Fail()

	}
	if queue.Equals(NewDoubleLinkedQueue[float32](1.3, -2.5, -1)) {

		t.Log("queues are equals")
		t.Fail()

	}

}
