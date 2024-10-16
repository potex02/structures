package wrapper

import (
	"testing"
)

type test struct {
	a int
	b Float32
}

func TestNewDefaultWrapper(t *testing.T) {

	var builder WrapperBuilder[test] = NewDefaultWrapperBuilder[test]()
	var wrapper Wrapper[test] = builder.Wrap(test{a: 1, b: -4.5})

	if wrapper.ToValue().a != 1 {
		t.Log("a is not 1")
		t.Fail()
	}
	if wrapper.ToValue().b != -4.5 {
		t.Log("b is not -4.5")
		t.Fail()
	}
	if !wrapper.Equal(builder.Wrap(test{a: 1, b: -4.5})) {
		t.Log("wrappers are not equal")
		t.Fail()
	}
	if wrapper.Equal(builder.Wrap(test{a: 2, b: -4.5})) {
		t.Log("wrappers are equal")
		t.Fail()
	}
	if wrapper.Compare(nil) != -2 {
		t.Log("compare is not -2")
		t.Fail()
	}
	if wrapper.Compare(builder.Wrap(test{a: 1, b: -4})) != 0 {
		t.Log("compare is not 0")
		t.Fail()
	}
	if wrapper.Hash() != 4151711685334854028 {
		t.Log("hash is ", wrapper.Hash())
		t.Fail()
	}
	if !wrapper.Copy().Equal(wrapper) {
		t.Log("The copy is", wrapper.Copy())
		t.Fail()
	}
}
func TestNewWrapper(t *testing.T) {

	var equal = func(w Wrapper[test], o any) bool {
		other, ok := o.(Wrapper[test])
		if ok {
			return w.ToValue().a == other.ToValue().a
		}
		return false
	}
	var compare = func(w Wrapper[test], o any) int {
		other, ok := o.(Wrapper[test])
		if ok {
			return w.ToValue().b.Compare(other.ToValue().b)
		}
		return -2
	}
	var hash = func(w Wrapper[test]) uint64 {
		return (Float32(w.ToValue().a) * w.ToValue().b).Hash()
	}
	var copy = func(w Wrapper[test]) test {
		return test{a: w.ToValue().a, b: 0}
	}
	var builder WrapperBuilder[test] = NewWrapperBuilder[test](equal, compare, hash, copy)
	var wrapper Wrapper[test] = builder.Wrap(test{a: 1, b: -4.5})

	if wrapper.ToValue().a != 1 {
		t.Log("a is not 1")
		t.Fail()
	}
	if wrapper.ToValue().b != -4.5 {
		t.Log("b is not -4.5")
		t.Fail()
	}
	if !wrapper.Equal(builder.Wrap(test{a: 1, b: -4.5})) {
		t.Log("wrappers are not equal")
		t.Fail()
	}
	if !wrapper.Equal(builder.Wrap(test{a: 1, b: -4})) {
		t.Log("wrappers are not equal")
		t.Fail()
	}
	if wrapper.Equal(builder.Wrap(test{a: -1, b: -4.5})) {
		t.Log("wrappers are equal")
		t.Fail()
	}
	if wrapper.Compare(nil) != -2 {
		t.Log("compare is not -2")
		t.Fail()
	}
	if wrapper.Compare(builder.Wrap(test{a: 1, b: -4.5})) != 0 {
		t.Log("compare is not 0")
		t.Fail()
	}
	if wrapper.Compare(builder.Wrap(test{a: 1, b: -10})) != 1 {
		t.Log("compare is not 1")
		t.Fail()
	}
	if wrapper.Hash() != 499999996 {
		t.Log("hash is ", wrapper.Hash())
		t.Fail()
	}
	if !wrapper.Copy().Equal(builder.Wrap(test{a: 1, b: 0})) {
		t.Log("The copy is", wrapper.Copy())
		t.Fail()
	}
}
