package list

import (
	"reflect"
)

// Stream provides aggregate operations for a [List].
type Stream[T any] struct {
	// contains filtered or unexported fields
	objects     []T
	constructor reflect.Value
}

// NewStream returns a new [Stream] associated at the list parameter.
//
// Constructor a [reflect.Value] rapresenting the function that create the resulting list from the stream.
// This function must have no parameters or must be a variadic function and must returns a List[T].
func NewStream[T any](list List[T], constructor reflect.Value) *Stream[T] {
	objects := make([]T, 0)
	list.Each(func(_ int, element T) {
		objects = append(objects, element)
	})
	return &Stream[T]{objects: objects, constructor: constructor}
}

// Map executes fun for all elements of s and returns a [Stream] containing the resulting elements.
//
// This method modifies the state of s, so it is not necessary to assign to resulting stream, but it
// can be useful for concatenate operations in a single instruction.
func (s *Stream[T]) Map(fun func(index int, element T) T) *Stream[T] {
	result := make([]T, 0)
	for i := range s.objects {
		result = append(result, fun(i, s.objects[i]))
	}
	s.objects = result
	return s
}

// Filter returns a [Stream] containing the elements that satisfy fun.
//
// This method modifies the state of s, so it is not necessary to assign to resulting stream, but it
// can be useful for concatenate operations in a single instruction.
func (s *Stream[T]) Filter(fun func(index int, element T) bool) *Stream[T] {
	result := make([]T, 0)
	for i := range s.objects {
		if fun(i, s.objects[i]) {
			result = append(result, s.objects[i])
		}
	}
	s.objects = result
	return s
}

// FilterMap executes fun for all elements of s and returns a [Stream] containing the resulting elements that satisfy fun.
//
// This method modifies the state of s, so it is not necessary to assign to resulting stream, but it
// can be useful for concatenate operations in a single instruction.
func (s *Stream[T]) FilterMap(fun func(index int, element T) (T, bool)) *Stream[T] {
	result := make([]T, 0)
	for i := range s.objects {
		if element, ok := fun(i, s.objects[i]); ok {
			result = append(result, element)
		}
	}
	s.objects = result
	return s
}

// Any returns true if at least one element of s satisfies fun.
func (s *Stream[T]) Any(fun func(index int, element T) bool) bool {
	for i := range s.objects {
		if fun(i, s.objects[i]) {
			return true
		}
	}
	return false
}

// All returns true if all elements of s satisfy fun.
func (s *Stream[T]) All(fun func(index int, element T) bool) bool {
	for i := range s.objects {
		if !fun(i, s.objects[i]) {
			return false
		}
	}
	return true
}

// None returns true if none of the elements of s satisfies fun.
func (s *Stream[T]) None(fun func(index int, element T) bool) bool {
	for i := range s.objects {
		if fun(i, s.objects[i]) {
			return false
		}
	}
	return true
}

// Count returns the number of elements that satisfy fun.
func (s *Stream[T]) Count(fun func(index int, element T) bool) int {
	result := 0
	for i := range s.objects {
		if fun(i, s.objects[i]) {
			result++
		}
	}
	return result
}

// Collect returns a [List] from s.
//
// the effective type of the result is the same the constructor.
//
// This method panics if constructor have wrong parameters or not returns a List[T].
func (s *Stream[T]) Collect() List[T] {
	result := s.constructor.Call([]reflect.Value{})[0].Interface().(List[T])
	for _, i := range s.objects {
		result.Add(i)
	}
	return result
}
