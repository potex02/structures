package queue

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
)

func TestNewLinkedQueue(t *testing.T) {

	var queue structures.Structure[float32] = NewLinkedQueue[float32]()

	if queue == nil {

		t.Log("queue is nil")
		t.Fail()

	}
	if queue.Len() != 0 {

		t.Log("length is not 0")
		t.Fail()

	}

}
func TestLinkedQueueFromSlice(t *testing.T) {

	var queue *LinkedQueue[float32] = NewLinkedQueueFromSlice([]float32{1.3, -2.5, 3.0, -4.0})

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
func TestHeadTailLinkedQueue(t *testing.T) {

	var queue *LinkedQueue[float32] = NewLinkedQueue[float32]()

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
	queue = NewLinkedQueueFromSlice([]float32{1.3, -2.5, 3.0, -4.0})
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
func TestRemoveLinkedQueue(t *testing.T) {

	var queue Queue[float32] = NewLinkedQueue[float32](1.3, -2.5)

	e, err := queue.Pop()
	if err != nil {

		t.Log("err is", err)
		t.Fail()

	}
	if e != 1.3 {

		t.Log("e is", e)
		t.Fail()

	}
	if queue.Len() != 1 {

		t.Log("Size is not 1")
		t.Fail()

	}
	e, err = queue.Pop()
	if err != nil {

		t.Log("err is", err)
		t.Fail()

	}
	if e != -2.5 {

		t.Log("e is", e)
		t.Fail()

	}
	if !queue.IsEmpty() {

		t.Log("queue not empty")
		t.Fail()

	}
	_, err = queue.Pop()
	if err == nil {

		t.Log("err is nil")
		t.Fail()

	}

}
func TestEqualsLinkedQueue(t *testing.T) {

	var queue *LinkedQueue[float32] = NewLinkedQueue[float32](1.3, -2.5)

	if !queue.Equals(NewArrayQueue[float32](1.3, -2.5)) {

		t.Log("queues are not equals")
		t.Fail()

	}
	if queue.Equals(NewLinkedQueue[float32](1.3, -2.5, -1)) {

		t.Log("queues are equals")
		t.Fail()

	}

}
