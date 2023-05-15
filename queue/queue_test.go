package queue

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
)

func TestNew(t *testing.T) {

	var queue structures.Structure[float32] = New[float32]()

	if queue == nil {

		t.Log("queue is nil")
		t.Fail()

	}
	if queue.Len() != 0 {

		t.Log("length is not 0")
		t.Fail()

	}

}
func TestNewFromSlice(t *testing.T) {

	var queue *Queue[float32] = NewFromSlice([]float32{1.3, -2.5, 3.0, -4.0})

	if queue == nil {

		t.Log("queue is nil")
		t.Fail()

	}
	if queue.Len() != 4 {

		t.Log("length is not 4")
		t.Fail()

	}
	if !reflect.DeepEqual(queue.objects, []float32{1.3, -2.5, 3.0, -4.0}) {

		t.Log("queue objects are", queue.objects)
		t.Fail()

	}

}
func TestHeadTail(t *testing.T) {

	var queue *Queue[float32] = New[float32]()

	if queue.Head() != nil {

		t.Log("head is not nil")
		t.Fail()

	}
	if queue.Tail() != nil {

		t.Log("tail is not nil")
		t.Fail()

	}
	queue = NewFromSlice([]float32{1.3, -2.5, 3.0, -4.0})
	if head := queue.Head(); *head != 1.3 {

		t.Log("Head is", *head)
		t.Fail()

	}
	if tail := queue.Tail(); *tail != -4.0 {

		t.Log("Head is", *tail)
		t.Fail()

	}

}
func TestRemove(t *testing.T) {

	var queue *Queue[float32] = NewFromSlice([]float32{1.3, -2.5})

	e, err := queue.Remove()
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
	e, err = queue.Remove()
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
	e, err = queue.Remove()
	if err == nil {

		t.Log("err is nil")
		t.Fail()

	}

}
func TestEquals(t *testing.T) {

	var queue *Queue[float32] = NewFromSlice([]float32{1.3, -2.5})

	if !queue.Equals(NewFromElements[float32](1.3, -2.5)) {

		t.Log("queues are not equals")
		t.Fail()

	}
	if queue.Equals(NewFromElements[float32](1.3, -2.5, -1)) {

		t.Log("queues are equals")
		t.Fail()

	}

}
