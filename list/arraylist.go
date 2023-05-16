package list

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/potex02/structures"
)

// ArrayList provides a generic list implemented with a slice.
type ArrayList[T any] struct {
	// contains filtered or unexported fields
	objects []T
}

// NewArrayList returns a new empty [ArrayList] containing the elements c.
//
// if no argument is passed, it will be created an empty [ArrayList].
func NewArrayList[T any](c ...T) *ArrayList[T] {

	return NewArrayListFromSlice(c)

}

// NewArrayListFromSlice returns a new [ArrayList] containing the elements of slice c.
func NewArrayListFromSlice[T any](c []T) *ArrayList[T] {

	return &ArrayList[T]{objects: c}

}

// NewArrayListFromStructure is a wrapper for NewArrayListFromSlice(c.ToSlice()).
func NewArrayListFromStructure[T any](c structures.Structure[T]) *ArrayList[T] {

	return NewArrayListFromSlice(c.ToSlice())

}

// Len returns the length of a.
func (a *ArrayList[T]) Len() int {

	return len(a.objects)

}

// IsEmpty returns a bool which indicate if a is empty or not.
func (a *ArrayList[T]) IsEmpty() bool {

	return len(a.objects) == 0

}

// Contains returns if e is present in a.
func (a *ArrayList[T]) Contains(e T) bool {

	return a.IndexOf(e) >= 0

}

// IndexOf returns the last position of e in a.
// If e is not present, the result is -1.
func (a *ArrayList[T]) IndexOf(e T) int {

	for i := range a.objects {

		if reflect.DeepEqual(a.objects[i], e) {

			return i

		}

	}
	return -1

}

// LastIndexOf returns the last position of e in a.
// If e is not present, the result is -1.
func (a *ArrayList[T]) LastIndexOf(e T) int {

	for i := len(a.objects) - 1; i != -1; i-- {

		if reflect.DeepEqual(a.objects[i], e) {

			return i

		}

	}
	return -1

}

// ToSLice returns a slice which contains all elements of a.
func (a *ArrayList[T]) ToSlice() []T {

	slice := make([]T, len(a.objects))
	copy(slice, a.objects)
	return slice

}

// Get returns the elements at the specifies index.
// It returns an error if the the index is out of bounds.
func (a *ArrayList[T]) Get(index int) (T, error) {

	if !a.rangeCheck(index) {

		var result T

		return result, errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(len(a.objects)))

	}
	return a.objects[index], nil

}

// Set sets the value of element at the specified index and returns the overwritten value.
// It returns an error if the the index is out of bounds.
func (a *ArrayList[T]) Set(index int, e T) (T, error) {

	var result T

	if index == len(a.objects) {

		a.Add(e)
		return result, nil

	}
	if !a.rangeCheck(index) {

		return result, errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(len(a.objects)))

	}
	result = a.objects[index]
	a.objects[index] = e
	return result, nil

}

// Add adds the elements e at the end of a.
func (a *ArrayList[T]) Add(e ...T) {

	a.AddSlice(e)

}

// AddAtIndex adds the elements e at the specified index.
// It returns an error if the the index is out of bounds.
func (a *ArrayList[T]) AddAtIndex(index int, e ...T) error {

	return a.AddSliceAtIndex(index, e)

}

// AddSlice adds the elements of e at the end of a.
func (a *ArrayList[T]) AddSlice(e []T) {

	a.objects = append(a.objects, e...)

}

// AddSliceAtIndex adds the elements of e at the specified index.
// It returns an error if the the index is out of bounds.
func (a *ArrayList[T]) AddSliceAtIndex(index int, e []T) error {

	if index > len(a.objects) || index < 0 {

		return errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(len(a.objects)))

	}
	if index == len(a.objects) {

		a.AddSlice(e)
		return nil

	}
	elements := make([]T, len(a.objects))
	copy(elements, a.objects)
	a.objects = append(append(a.objects[:index], e...), elements[index:]...)
	return nil

}

// Remove removes the element at specified index and return the removed value.
// It returns an error if the the index is out of bounds.
func (a *ArrayList[T]) Remove(index int) (T, error) {

	var result T

	if !a.rangeCheck(index) {

		return result, errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(len(a.objects)))

	}
	result = a.objects[index]
	a.objects = append(a.objects[:index], a.objects[index+1:]...)
	return result, nil

}

// RemoveElement removes the element e from a if it is present.
// In that case, the method returns true, otherwhise it returns false.
func (a *ArrayList[T]) RemoveElement(e T) bool {

	for i := 0; i != len(a.objects); i++ {

		if reflect.DeepEqual(a.objects[i], e) {

			a.Remove(i)
			return true

		}

	}
	return false

}

// Clear removes all element from a.
func (a *ArrayList[T]) Clear() {

	a.objects = []T{}

}

// Iter returns a chan which permits to iterate a with the range keyword.
//
// This method can only be used to iterate an [ArrayList] if the index is not needed.
// For now, the only way to iterate an [ArrayList] with the index is the following code:
//
//	for i := 0; i < a.Len(); i++ {
//		element, err := a.Get(i)
//		// Code
//	}
func (a *ArrayList[T]) Iter() chan T {

	obj := make(chan T)
	go func() {

		for _, i := range a.objects {

			obj <- i

		}
		close(obj)

	}()
	return obj

}

// IterReverse returns a chan which permits to iterate a in reverse order with the range keyword.
//
// This method can only be used to iterate a [ArrayList] if the index is not needed.
// For now, the only way to iterate a [ArrayList] in reverse order with the index is the following code:
//
//	for i := a.Len() - 1; i >= 0; i-- {
//		element, err := a.Get(i)
//		// Code
//	}
func (a *ArrayList[T]) IterReverse() chan T {

	obj := make(chan T)
	go func() {

		for i := len(a.objects) - 1; i >= 0; i-- {

			obj <- a.objects[i]

		}
		close(obj)

	}()
	return obj

}

// Equals returns true if a and st are both lists and their elements are equals.
// In any other case, it returns false.
//
// Equals does not take into account the effective type of st. This means that if st is a [LinkedList],
// but the elements of a and the elements of st are equals, this methods returns anyway true.
func (a *ArrayList[T]) Equals(st structures.Structure[T]) bool {

	list, ok := st.(List[T])
	return ok && reflect.DeepEqual(a.ToSlice(), list.ToSlice())

}

// Copy returns a list containing a copy of the elements of a.
// The result of this method is of type [List], but the effective list which is created is an [ArrayList].
func (a *ArrayList[T]) Copy() List[T] {

	return NewArrayListFromSlice(a.ToSlice())

}

// String returns a rapresentation of a in the form of a string.
func (a *ArrayList[T]) String() string {

	return fmt.Sprintf("ArrayList[%T]%v", *new(T), a.objects)

}
func (a *ArrayList[T]) rangeCheck(index int) bool {

	return index >= 0 && index < len(a.objects)

}
