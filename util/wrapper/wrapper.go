// package wrapper provides a simply way to implement the [util] interfaces for any type.
package wrapper

import (
	"fmt"
	"reflect"

	"github.com/potex02/structures/util"
)

// Wrapper is an interface which implements the equality, comparison and hashing operations for type T.
//
// An implementation of Wrapper for primitive types is already defined.
type Wrapper[T any] interface {
	util.Equaler
	util.Hasher
	// ToValue returns the wrapped value.
	ToValue() T
}

// DefaultEqual is the default function used as method Equals by a [WrapperBuilder] to create wrappers.
//
// The result is reflect.DeepEqual between the values of w and o if o is a [Wrapper]
// otherwise is false.
func DefaultEqual[T any](w Wrapper[T], o any) bool {

	other, ok := o.(Wrapper[T])
	if ok {

		return reflect.DeepEqual(w.ToValue(), other.ToValue())

	}
	return false

}

// DefaultCompare is the default function used as method Compare by a [WrapperBuilder] to create wrappers.
//
// The result is 0 if o is a [Wrapper], otherwise is -2.
func DefaultCompare[T any](w Wrapper[T], o any) int {

	_, ok := o.(Wrapper[T])
	if !ok {
		return -2
	}
	return 0

}

// DefaultHash is the default function used as method Hash by a [WrapperBuilder] to create wrappers.
//
// The result is the string rapresentation of the wrapped value.
func DefaultHash[T any](w Wrapper[T]) string {

	return fmt.Sprintf("%v", w.ToValue())

}

// WrapperBuilder is a type which permits to create wrapppers of type T.
type WrapperBuilder[T any] struct {
	// contains filtered or unexported fields
	equal   func(r Wrapper[T], o any) bool
	compare func(r Wrapper[T], o any) int
	hash    func(r Wrapper[T]) string
}

// NewDefaultWrapperBuilder returns a new [WrapperBuilder].
//
// The methods Equal, Compare and Hash of the wrappers builded by the builder are
// respectively [DefaultEqual], [DefaultCompare] and [DefaultHash].
func NewDefaultWrapperBuilder[T any]() WrapperBuilder[T] {

	return WrapperBuilder[T]{
		equal:   DefaultEqual[T],
		compare: DefaultCompare[T],
		hash:    DefaultHash[T],
	}

}

// NewWrapperBuilder returns a new [WrapperBuilder] with
// the custom Equal, Compare and Hash methods.
//
// The w parameter of the function parameters is the receiver of the methods.
func NewWrapperBuilder[T any](equal func(w Wrapper[T], o any) bool, compare func(w Wrapper[T], o any) int, hash func(w Wrapper[T]) string) WrapperBuilder[T] {

	return WrapperBuilder[T]{
		equal:   equal,
		compare: compare,
		hash:    hash,
	}

}

// NewEqualWrapperBuilder returns a new [WrapperBuilder] with
// the custom Equal method.
//
// The w parameter of the function parameter is the receiver of the method.
func NewEqualWrapperBuilder[T any](equal func(w Wrapper[T], o any) bool) WrapperBuilder[T] {

	return WrapperBuilder[T]{
		equal:   equal,
		compare: DefaultCompare[T],
		hash:    DefaultHash[T],
	}

}

// NewCompareWrapperBuilder returns a new [WrapperBuilder] with
// the custom Compare method.
//
// The w parameter of the function parameter is the receiver of the method.
func NewCompareWrapperBuilder[T any](compare func(w Wrapper[T], o any) int) WrapperBuilder[T] {

	return WrapperBuilder[T]{
		equal:   DefaultEqual[T],
		compare: compare,
		hash:    DefaultHash[T],
	}

}

// NewHashWrapperBuilder returns a new [WrapperBuilder] with
// the custom Compare and Hash methods.
//
// The w parameter of the function parameters is the receiver of the methods.
func NewHashWrapperBuilder[T any](compare func(w Wrapper[T], o any) int, hash func(w Wrapper[T]) string) WrapperBuilder[T] {

	return WrapperBuilder[T]{
		equal:   DefaultEqual[T],
		compare: compare,
		hash:    hash,
	}

}

// Wrap returns a [Wrapper] that wraps the value.
func (b WrapperBuilder[T]) Wrap(value T) Wrapper[T] {

	return wrapperResult[T]{
		value:   value,
		equal:   b.equal,
		compare: b.compare,
		hash:    b.hash,
	}

}

type wrapperResult[T any] struct {
	value   T
	equal   func(r Wrapper[T], o any) bool
	compare func(r Wrapper[T], o any) int
	hash    func(r Wrapper[T]) string
}

func (r wrapperResult[T]) Equal(o any) bool {

	return r.equal(r, o)

}

func (r wrapperResult[T]) Compare(o any) int {

	return r.compare(r, o)

}

func (r wrapperResult[T]) Hash() string {

	return r.hash(r)

}

func (r wrapperResult[T]) ToValue() T {

	return r.value

}

func (r wrapperResult[T]) String() string {

	return fmt.Sprintf("%v", r.value)

}
