package stack

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
)

func TestNewArrayStack(t *testing.T) {

	var stack structures.Structure[float32] = NewArrayStack[float32]()

	if stack == nil {

		t.Log("stack is nil")
		t.Fail()

	}
	if stack.Len() != 0 {

		t.Log("length is not 0")
		t.Fail()

	}

}
func TestNewArrayStackFromSlice(t *testing.T) {

	var stack *ArrayStack[float32] = NewArrayStackFromSlice([]float32{1.3, -2.5, 3.0, -4.0})

	if stack == nil {

		t.Log("stack is nil")
		t.Fail()

	}
	if stack.Len() != 4 {

		t.Log("length is not 4")
		t.Fail()

	}
	if !reflect.DeepEqual(stack.ToSlice(), []float32{1.3, -2.5, 3.0, -4.0}) {

		t.Log("stack objects are", stack.objects)
		t.Fail()

	}

}
func TestTopArrayStack(t *testing.T) {

	var stack Stack[float32] = NewArrayStack[float32]()

	if _, err := stack.Top(); err == nil {

		t.Log("the stack is not empty")
		t.Fail()

	}
	stack = NewArrayStackFromSlice([]float32{1.3, -2.5, 3.0, -4.0})
	top, err := stack.Top()
	if top != -4.0 {

		t.Log("top is", top)
		t.Fail()

	}
	if err != nil {

		t.Log("err is", err)
		t.Fail()

	}

}
func TestPopArrayStack(t *testing.T) {

	var stack *ArrayStack[float32] = NewArrayStack[float32](1.3, -2.5)

	e, err := stack.Pop()
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
	e, err = stack.Pop()
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
	if _, err := stack.Pop(); err == nil {

		t.Log("err is nil")
		t.Fail()

	}

}
func TestEqualArrayStack(t *testing.T) {

	var stack *ArrayStack[float64] = NewArrayStack(1.3, -2.5)

	if !stack.Equal(NewArrayStack(1.3, -2.5)) {

		t.Log("stacks are not equals")
		t.Fail()

	}
	if stack.Equal(NewLinkedStack[float64](1.3, -2.5, -1)) {

		t.Log("stacks are equals")
		t.Fail()

	}

}
