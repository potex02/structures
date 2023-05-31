package list

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/potex02/structures"
)

// ArrayList provides a generic list implemented with a slice.
//
// It implements the interface [List].
type ArrayList[T any] struct {
	// contains filtered or unexported fields
	objects []T
}

// NewArrayList returns a new [ArrayList] containing the elements c.
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

// Len returns the length of l.
func (l *ArrayList[T]) Len() int {

	return len(l.objects)

}

// IsEmpty returns a bool which indicate if l is empty or not.
func (l *ArrayList[T]) IsEmpty() bool {

	return len(l.objects) == 0

}

// Contains returns if e is present in l.
func (l *ArrayList[T]) Contains(e T) bool {

	return l.IndexOf(e) >= 0

}

// IndexOf returns the first position of e in l.
// If e is not present, the result is -1.
func (l *ArrayList[T]) IndexOf(e T) int {

	for i := range l.objects {

		if reflect.DeepEqual(l.objects[i], e) {

			return i

		}

	}
	return -1

}

// LastIndexOf returns the last position of e in l.
// If e is not present, the result is -1.
func (l *ArrayList[T]) LastIndexOf(e T) int {

	for i := len(l.objects) - 1; i != -1; i-- {

		if reflect.DeepEqual(l.objects[i], e) {

			return i

		}

	}
	return -1

}

// ToSLice returns a slice which contains all elements of l.
func (l *ArrayList[T]) ToSlice() []T {

	slice := make([]T, len(l.objects))
	copy(slice, l.objects)
	return slice

}

// Get returns the elements at the specifies index.
// It returns an error if the the index is out of bounds.
func (l *ArrayList[T]) Get(index int) (T, error) {

	if !rangeCheck[T](l, index) {

		var result T

		return result, errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(len(l.objects)))

	}
	return l.objects[index], nil

}

// Set sets the value of element at the specified index and returns the overwritten value.
// It returns an error if the the index is out of bounds.
func (l *ArrayList[T]) Set(index int, e T) (T, error) {

	var result T

	if index == len(l.objects) {

		l.Add(e)
		return result, nil

	}
	if !rangeCheck[T](l, index) {

		return result, errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(len(l.objects)))

	}
	result = l.objects[index]
	l.objects[index] = e
	return result, nil

}

// Add adds the elements e at the end of l.
func (l *ArrayList[T]) Add(e ...T) {

	l.AddSlice(e)

}

// AddAtIndex adds the elements e at the specified index.
// It returns an error if the the index is out of bounds.
func (l *ArrayList[T]) AddAtIndex(index int, e ...T) error {

	return l.AddSliceAtIndex(index, e)

}

// AddSlice adds the elements of e at the end of l.
func (l *ArrayList[T]) AddSlice(e []T) {

	l.objects = append(l.objects, e...)

}

// AddSliceAtIndex adds the elements of e at the specified index.
// It returns an error if the the index is out of bounds.
func (l *ArrayList[T]) AddSliceAtIndex(index int, e []T) error {

	if index > len(l.objects) || index < 0 {

		return errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(len(l.objects)))

	}
	if index == len(l.objects) {

		l.AddSlice(e)
		return nil

	}
	elements := make([]T, len(l.objects))
	copy(elements, l.objects)
	l.objects = append(append(l.objects[:index], e...), elements[index:]...)
	return nil

}

// Remove removes the element at specified index and return the removed value.
// It returns an error if the the index is out of bounds.
func (l *ArrayList[T]) Remove(index int) (T, error) {

	var result T

	if !rangeCheck[T](l, index) {

		return result, errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(len(l.objects)))

	}
	result = l.objects[index]
	l.objects = append(l.objects[:index], l.objects[index+1:]...)
	return result, nil

}

// RemoveElement removes the element e from l if it is present.
// In that case, the method returns true, otherwhise it returns false.
func (l *ArrayList[T]) RemoveElement(e T) bool {

	for i := 0; i != len(l.objects); i++ {

		if reflect.DeepEqual(l.objects[i], e) {

			l.Remove(i)
			return true

		}

	}
	return false

}

// Clear removes all element from l.
func (l *ArrayList[T]) Clear() {

	l.objects = []T{}

}

// Iter returns a chan which permits to iterate an [ArrayList] with the range keyword.
//
//	for i := range l.Iter() {
//		// code
//	}
//
// This method can only be used to iterate an [ArrayList] if the index is not needed.
// if you need to iterate an [ArrayList] with the index there are two options:
//
//	for i := 0; i < list.Len(); i++ {
//		element, err := list.Get(i)
//		// Code
//	}
//
//	j := 0
//	for i := range l.Iter() {
//		// code
//		j++
//	}
func (l *ArrayList[T]) Iter() chan T {

	obj := make(chan T)
	go func() {

		defer close(obj)
		for _, i := range l.objects {

			obj <- i

		}

	}()
	return obj

}

// IterReverse returns a chan which permits to iterate an [ArrayList] in reverse order with the range keyword.
//
//	for i := range l.IterReverse() {
//		// code
//	}
//
// This method can only be used to iterate an [ArrayList] if the index is not needed.
// if you need to iterate an [ArrayList] in reverse order with the index there are two options:
//
//	for i := list.Len() - 1; i >= 0; i-- {
//		element, err := list.Get(i)
//		// Code
//	}
//
//	j := l.Len() -1
//	for i := range l.Iter() {
//		// code
//		j--
//	}
func (l *ArrayList[T]) IterReverse() chan T {

	obj := make(chan T)
	go func() {

		defer close(obj)
		for i := len(l.objects) - 1; i >= 0; i-- {

			obj <- l.objects[i]

		}

	}()
	return obj

}

// Equal returns true if l and st are both lists and their elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is a [LinkedList],
// but the elements of l and the elements of st are equals, this method returns anyway true.
func (l *ArrayList[T]) Equal(st structures.Structure[T]) bool {

	list, ok := st.(List[T])
	return ok && reflect.DeepEqual(l.ToSlice(), list.ToSlice())

}

// Copy returns a list containing a copy of the elements of l.
// The result of this method is of type [List], but the effective list which is created is an [ArrayList].
func (l *ArrayList[T]) Copy() List[T] {

	return NewArrayListFromSlice(l.ToSlice())

}

// String returns a rapresentation of l in the form of a string.
func (l *ArrayList[T]) String() string {

	return fmt.Sprintf("ArrayList[%T]%v", *new(T), l.objects)

}
