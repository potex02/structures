package set

import (
	"reflect"

	"github.com/potex02/structures/util"
)

// Stream provides aggregate operations for a [BaseSet].
type Stream[T util.Comparer] struct {
	// contains filtered or unexported fields
	objects     []T
	constructor reflect.Value
}

// NewStream returns a new [Stream] for a [BaseSet] associated at the set parameter.
//
// Constructor a [reflect.Value] rapresenting the function that create the resulting set from the stream.
// This function must have no parameters or must be a variadic function and must returns a BaseSet[T].
func NewStream[T util.Comparer](set BaseSet[T], constructor reflect.Value) *Stream[T] {
	objects := make([]T, 0)
	set.Each(func(element T) {
		objects = append(objects, element)
	})
	return &Stream[T]{objects: objects, constructor: constructor}
}

// Map executes fun for all elements of s and returns a [Stream] containing the resulting elements.
//
// This method modifies the state of s, so it is not necessary to assign to resulting stream, but it
// can be useful for concatenate operations in a single instruction.
func (s *Stream[T]) Map(fun func(element T) T) *Stream[T] {
	result := make([]T, 0)
	for _, i := range s.objects {
		result = append(result, fun(i))
	}
	s.objects = result
	return s
}

// Filter returns a [Stream] containing the elements that satisfy fun.
//
// This method modifies the state of s, so it is not necessary to assign to resulting stream, but it
// can be useful for concatenate operations in a single instruction.
func (s *Stream[T]) Filter(fun func(element T) bool) *Stream[T] {
	result := make([]T, 0)
	for _, i := range s.objects {
		if fun(i) {
			result = append(result, i)
		}
	}
	s.objects = result
	return s
}

// FilterMap executes fun for all elements of s and returns a [Stream] containing the resulting elements that satisfy fun.
//
// This method modifies the state of s, so it is not necessary to assign to resulting stream, but it
// can be useful for concatenate operations in a single instruction.
func (s *Stream[T]) FilterMap(fun func(element T) (T, bool)) *Stream[T] {
	result := make([]T, 0)
	for _, i := range s.objects {
		if element, ok := fun(i); ok {
			result = append(result, element)
		}
	}
	s.objects = result
	return s
}

// Union returns a [Stream] that is the result of the union between s and set.
//
// The result of a union of two sets is a set which contains all elements of both sets.
func (s *Stream[T]) Union(set BaseSet[T]) *Stream[T] {
	s.objects = append(s.objects, set.ToSlice()...)
	return s
}

// Intersection returns a [Stream] that is the result of the intersection between s and set.
//
// The result of a union of two sets is a set which contains only the elements present in both sets.
func (s *Stream[T]) Intersection(set BaseSet[T]) *Stream[T] {
	result := make([]T, 0)
	for _, i := range s.objects {
		if set.Contains(i) {
			result = append(result, i)
		}
	}
	s.objects = result
	return s
}

// Difference returns a [Stream] that is the result of the difference between s and set.
//
// The result of a union of two sets is a set which contains only the elements present the first set
// that are not present in the second set.
func (s *Stream[T]) Difference(set BaseSet[T]) *Stream[T] {
	result := make([]T, 0)
	for _, i := range s.objects {
		if !set.Contains(i) {
			result = append(result, i)
		}
	}
	s.objects = result
	return s
}

// Any returns true if at least one element of s satisfies fun.
func (s *Stream[T]) Any(fun func(element T) bool) bool {
	for _, i := range s.objects {
		if fun(i) {
			return true
		}
	}
	return false
}

// All returns true if all elements of s satisfy fun.
func (s *Stream[T]) All(fun func(element T) bool) bool {
	for _, i := range s.objects {
		if !fun(i) {
			return false
		}
	}
	return true
}

// None returns true if none of the elements of s satisfies fun.
func (s *Stream[T]) None(fun func(element T) bool) bool {
	for _, i := range s.objects {
		if fun(i) {
			return false
		}
	}
	return true
}

// Count returns the number of elements that satisfy fun.
func (s *Stream[T]) Count(fun func(element T) bool) int {
	result := 0
	for _, i := range s.objects {
		if fun(i) {
			result++
		}
	}
	return result
}

// Collect returns a [BaseSet] from s.
//
// the effective type of the result is the same the constructor.
//
// This method panics if constructor have wrong parameters or not returns a BaseSet[T].
func (s *Stream[T]) CollectBase() BaseSet[T] {
	result := s.constructor.Call([]reflect.Value{})[0].Interface().(BaseSet[T])
	for _, i := range s.objects {
		result.Add(i)
	}
	return result
}

// Collect returns a [Set] from s.
//
// the effective type of the result is the same the constructor.
//
// This method panics if constructor have wrong parameters or not returns a Set[T].
func (s *Stream[T]) Collect() Set[T] {
	result := s.constructor.Call([]reflect.Value{})[0].Interface().(Set[T])
	for _, i := range s.objects {
		result.Add(i)
	}
	return result
}

// CollectMulti returns a [MultiSet] from s.
//
// the effective type of the result is the same the constructor.
//
// This method panics if constructor have wrong parameters or not returns a MultiSet[T].
func (s *Stream[T]) CollectMulti() MultiSet[T] {
	result := s.constructor.Call([]reflect.Value{})[0].Interface().(MultiSet[T])
	for _, i := range s.objects {
		result.Add(i)
	}
	return result
}
