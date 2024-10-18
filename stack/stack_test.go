package stack

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
)

func TestNewStack(t *testing.T) {

	var stack structures.Structure[float32] = NewStack[float32]()

	if stack == nil {
		t.Log("stack is nil")
		t.Fail()
	}
	if stack.Len() != 0 {
		t.Log("length is not 0")
		t.Fail()
	}
}
func TestNewStackFromSlice(t *testing.T) {

	var stack *Stack[float32] = NewStackFromSlice([]float32{1.3, -2.5, 3.0, -4.0})

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
func TestTopStack(t *testing.T) {

	var stack *Stack[float32] = NewStack[float32]()

	if _, ok := stack.Top(); ok {
		t.Log("the stack is not empty")
		t.Fail()
	}
	stack = NewStackFromSlice([]float32{1.3, -2.5, 3.0, -4.0})
	top, ok := stack.Top()
	if top != -4.0 {
		t.Log("top is", top)
		t.Fail()
	}
	if !ok {
		t.Log("the stack is empty")
		t.Fail()
	}
}
func TestPopStack(t *testing.T) {

	var stack *Stack[float32] = NewStack[float32](1.3, -2.5)

	e, ok := stack.Pop()
	if !ok {
		t.Log("the stack is empty")
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
	e, ok = stack.Pop()
	if !ok {
		t.Log("the stack is empty")
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
	if _, ok := stack.Pop(); ok {
		t.Log("the stack is not empty")
		t.Fail()
	}
}
func TestEqualStack(t *testing.T) {

	var stack *Stack[float64] = NewStack(1.3, -2.5)

	if !stack.Equal(NewStack(1.3, -2.5)) {
		t.Log("stacks are not equals")
		t.Fail()
	}
	if stack.Equal(NewStack[float64](1.3, -2.5, -1)) {
		t.Log("stacks are equals")
		t.Fail()
	}
}
