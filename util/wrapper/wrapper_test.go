package wrapper

import (
	"fmt"
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
	if wrapper.Hash() != fmt.Sprintf("%v", wrapper.ToValue()) {

		t.Log("hash is not the value")
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
	var hash = func(w Wrapper[test]) string {

		return fmt.Sprintf("%v", Float32(w.ToValue().a)*w.ToValue().b)

	}
	var builder WrapperBuilder[test] = NewWrapperBuilder[test](equal, compare, hash)
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
	if wrapper.Hash() != "-4.5" {

		t.Log("hash is not \"-4.5\"")
		t.Fail()

	}

}
