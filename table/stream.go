package table

import (
	"reflect"

	"github.com/potex02/structures/util"
)

// Stream provides aggregate operations for a [BaseTable].
type Stream[K util.Comparer, T any] struct {
	// contains filtered or unexported fields
	objects     []*Entry[K, T]
	constructor reflect.Value
}

// NewStream returns a new [Stream] for a [BaseTable] associated at the table parameter.
//
// Constructor a [reflect.Value] rapresenting the function that create the resulting table from the stream.
// This function must have no parameters or must be a variadic function and must returns a BaseTable[K, T].
func NewStream[K util.Comparer, T any](table BaseTable[K, T], constructor reflect.Value) *Stream[K, T] {

	objects := make([]*Entry[K, T], 0)
	table.Each(func(key K, element T) {
		objects = append(objects, NewEntry(key, element))
	})
	return &Stream[K, T]{objects: objects, constructor: constructor}

}

// Map executes fun for all elements of s and returns a [Stream] containing the resulting elements.
//
// This method modifies the state of s, so it is not necessary to assign to resulting stream, but it
// can be useful for concatenate operations in a single instruction.
func (s *Stream[K, T]) Map(fun func(key K, element T) T) *Stream[K, T] {

	result := make([]*Entry[K, T], 0)
	for _, i := range s.objects {

		result = append(result, NewEntry(i.Key(), fun(i.Key(), i.Element())))

	}
	s.objects = result
	return s

}

// Filter returns a [Stream] containing the elements that satisfy fun.
//
// This method modifies the state of s, so it is not necessary to assign to resulting stream, but it
// can be useful for concatenate operations in a single instruction.
func (s *Stream[K, T]) Filter(fun func(key K, element T) bool) *Stream[K, T] {

	result := make([]*Entry[K, T], 0)
	for _, i := range s.objects {

		if fun(i.Key(), i.Element()) {

			result = append(result, NewEntry(i.Key(), i.Element()))

		}

	}
	s.objects = result
	return s

}

// FilterMap executes fun for all elements of s and returns a [Stream] containing the resulting elements that satisfy fun.
//
// This method modifies the state of s, so it is not necessary to assign to resulting stream, but it
// can be useful for concatenate operations in a single instruction.
func (s *Stream[K, T]) FilterMap(fun func(key K, element T) (T, bool)) *Stream[K, T] {

	result := make([]*Entry[K, T], 0)
	for _, i := range s.objects {

		if element, ok := fun(i.Key(), i.Element()); ok {

			result = append(result, NewEntry(i.Key(), element))

		}

	}
	s.objects = result
	return s

}

// Any returns true if at least one element of s satisfies fun.
func (s *Stream[K, T]) Any(fun func(key K, element T) bool) bool {

	for _, i := range s.objects {

		if fun(i.Key(), i.Element()) {

			return true

		}

	}
	return false

}

// All returns true if all elements of s satisfy fun.
func (s *Stream[K, T]) All(fun func(key K, element T) bool) bool {

	for _, i := range s.objects {

		if !fun(i.Key(), i.Element()) {

			return false

		}

	}
	return true

}

// None returns true if none of the elements of s satisfies fun.
func (s *Stream[K, T]) None(fun func(key K, element T) bool) bool {

	for _, i := range s.objects {

		if fun(i.Key(), i.Element()) {

			return false

		}

	}
	return true

}

// Count returns the number of elements that satisfy fun.
func (s *Stream[K, T]) Count(fun func(key K, element T) bool) int {

	result := 0
	for _, i := range s.objects {

		if fun(i.Key(), i.Element()) {

			result++

		}

	}
	return result

}

// Collect returns a [Table] from s.
//
// the effective type of the result is the same the constructor.
//
// This method panics if constructor have wrong parameters or not returns a Table[K, T].
func (s *Stream[K, T]) Collect() Table[K, T] {

	result := s.constructor.Call([]reflect.Value{})[0].Interface().(Table[K, T])
	for _, i := range s.objects {

		result.Put(i.Key(), i.Element())

	}
	return result

}

// Collect returns a [MultiTable] from s.
//
// the effective type of the result is the same the constructor.
//
// This method panics if constructor have wrong parameters or not returns a MultiTable[K, T].
func (s *Stream[K, T]) CollectMulti() MultiTable[K, T] {

	result := s.constructor.Call([]reflect.Value{})[0].Interface().(MultiTable[K, T])
	for _, i := range s.objects {

		result.Put(i.Key(), i.Element())

	}
	return result

}
