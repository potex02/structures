package queue

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
)

func TestNewDoubleArrayQueue(t *testing.T) {

	var queue structures.Structure[float32] = NewDoubleArrayQueue[float32]()

	if queue == nil {

		t.Log("queue is nil")
		t.Fail()

	}
	if queue.Len() != 0 {

		t.Log("length is not 0")
		t.Fail()

	}

}

func TestNewDoubleArrayQueueFromSlice(t *testing.T) {

	var queue *DoubleArrayQueue[float32] = NewDoubleArrayQueueFromSlice([]float32{1.3, -2.5, 3.0, -4.0})

	if queue == nil {

		t.Log("queue is nil")
		t.Fail()

	}
	if queue.Len() != 4 {

		t.Log("length is not 4")
		t.Fail()

	}
	if !reflect.DeepEqual(queue.ToSlice(), []float32{1.3, -2.5, 3.0, -4.0}) {

		t.Log("queue objects are", queue.objects)
		t.Fail()

	}

}

func TestHeadTailDoubleArrayQueue(t *testing.T) {

	var queue *DoubleArrayQueue[float32] = NewDoubleArrayQueue[float32]()

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
	queue = NewDoubleArrayQueueFromSlice([]float32{1.3, -2.5, 3.0, -4.0})
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

func TestPopDoubleArrayQueue(t *testing.T) {

	var queue DoubleQueue[float32] = NewDoubleArrayQueue[float32](1.3, 3, -2.5)

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
	queue = NewDoubleArrayQueue[float32]()
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

func TestEqualsDoubleArrayQueue(t *testing.T) {

	var queue *DoubleArrayQueue[float32] = NewDoubleArrayQueue[float32](1.3, -2.5)

	if !queue.Equals(NewDoubleArrayQueue[float32](1.3, -2.5)) {

		t.Log("queues are not equals")
		t.Fail()

	}
	if queue.Equals(NewDoubleLinkedQueue[float32](1.3, -2.5, -1)) {

		t.Log("queues are equals")
		t.Fail()

	}

}
