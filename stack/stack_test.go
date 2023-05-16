package stack

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
)

func TestNew(t *testing.T) {

	var stack structures.Structure[float32] = New[float32]()

	if stack == nil {

		t.Log("stack is nil")
		t.Fail()

	}
	if stack.Len() != 0 {

		t.Log("length is not 0")
		t.Fail()

	}

}

func TestNewFromElements(t *testing.T) {

	var stack *Stack[float64] = NewFromElements(1.3, -2.5, 3.0, -4.0)

	if stack == nil {

		t.Log("stack is nil")
		t.Fail()

	}
	if stack.Len() != 4 {

		t.Log("length is not 4")
		t.Fail()

	}
	if !reflect.DeepEqual(stack.objects, []float64{1.3, -2.5, 3.0, -4.0}) {

		t.Log("stack objects are", stack.objects)
		t.Fail()

	}

}
func TestNewFromSlice(t *testing.T) {

	var stack *Stack[float32] = NewFromSlice([]float32{1.3, -2.5, 3.0, -4.0})

	if stack == nil {

		t.Log("stack is nil")
		t.Fail()

	}
	if stack.Len() != 4 {

		t.Log("length is not 4")
		t.Fail()

	}
	if !reflect.DeepEqual(stack.objects, []float32{1.3, -2.5, 3.0, -4.0}) {

		t.Log("stack objects are", stack.objects)
		t.Fail()

	}

}
func TestTop(t *testing.T) {

	var stack *Stack[float32] = New[float32]()

	if stack.Top() != nil {

		t.Log("top is not nil")
		t.Fail()

	}
	stack = NewFromSlice([]float32{1.3, -2.5, 3.0, -4.0})
	if top := stack.Top(); *top != -4.0 {

		t.Log("top is", *top)
		t.Fail()

	}

}
func TestRemove(t *testing.T) {

	var stack *Stack[float32] = NewFromSlice([]float32{1.3, -2.5})

	e, err := stack.Remove()
	if err != nil {

		t.Log("err is", err)
		t.Fail()

	}
	if e != -2.5 {

		t.Log("e is", e)
		t.Fail()

	}
	if stack.Len() != 1 {

		t.Log("size is not 1")
		t.Fail()

	}
	e, err = stack.Remove()
	if err != nil {

		t.Log("err is", err)
		t.Fail()

	}
	if e != 1.3 {

		t.Log("e is", e)
		t.Fail()

	}
	if !stack.IsEmpty() {

		t.Log("stack not empty")
		t.Fail()

	}
	_, err = stack.Remove()
	if err == nil {

		t.Log("err is nil")
		t.Fail()

	}

}
func TestEquals(t *testing.T) {

	var stack *Stack[float32] = NewFromSlice([]float32{1.3, -2.5})

	if !stack.Equals(NewFromElements[float32](1.3, -2.5)) {

		t.Log("stacks are not equals")
		t.Fail()

	}
	if stack.Equals(NewFromElements[float32](1.3, -2.5, -1)) {

		t.Log("stacks are equals")
		t.Fail()

	}

}
