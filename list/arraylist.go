package list

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util"
)

var _ structures.Structure[int] = NewArrayList[int]()
var _ List[int] = NewArrayList[int]()

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

// IsEmpty returns a bool which indicates if l is empty or not.
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

	element, ok := interface{}(e).(util.Equaler)
	for i := range l.objects {

		if ok {

			if element.Equal(l.objects[i]) {

				return i

			}

		} else if reflect.DeepEqual(l.objects[i], e) {

			return i

		}

	}
	return -1

}

// LastIndexOf returns the last position of e in l.
// If e is not present, the result is -1.
func (l *ArrayList[T]) LastIndexOf(e T) int {

	element, ok := interface{}(e).(util.Equaler)
	for i := len(l.objects) - 1; i != -1; i-- {

		if ok {

			if element.Equal(l.objects[i]) {

				return i

			}

		} else if reflect.DeepEqual(l.objects[i], e) {

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

	element, ok := interface{}(e).(util.Equaler)
	for i := 0; i != len(l.objects); i++ {

		if ok {

			if element.Equal(l.objects[i]) {

				l.Remove(i)
				return true

			}

		} else if reflect.DeepEqual(l.objects[i], e) {

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

// Iter returns a chan which permits to iterate an [ArrayList].
//
//	for i := l.Iter(); !i.End(); i = i.Next() {
//		element := i.Element()
//		index := i.Index()
//		// Code
//	}
func (l *ArrayList[T]) Iter() Iterator[T] {

	return NewArrayListIterator(l)

}

// Iter returns a chan which permits to iterate an [ArrayList] in reverse order.
//
//	for i := l.IterReverse(); !i.End(); i = i.Prev() {
//		element := i.Element()
//		index := i.Index()
//		// Code
//	}
func (l *ArrayList[T]) IterReverse() Iterator[T] {

	return NewArrayListReverseIterator(l)

}

// Equal returns true if l and st are both lists and their elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is a [LinkedList],
// but the elements of l and the elements of st are equals, this method returns anyway true.
func (l *ArrayList[T]) Equal(st any) bool {

	list, ok := st.(List[T])
	if ok && l != nil && list != nil {

		if l.Len() != list.Len() {

			return false

		}
		other := list.Iter()
		for _, i := range l.objects {

			element, ok := interface{}(i).(util.Equaler)
			if ok {

				if !element.Equal(other.Element()) {

					return false

				}

			} else if !reflect.DeepEqual(i, other.Element()) {

				return false

			}
			other = other.Next()

		}
		return true

	}
	return false

}

// Compare returns 0 if l and st are equals,
// -1 if l is shorten than st,
// 1 if l is longer than st,
// -2 if st is not a [List] or if one between l and st is nil.
//
// If l and st have the same length, the result is the comparison
// between the first different element of the two lists if T implemets [util.Comparer],
// otherwhise the result is 0.
func (l *ArrayList[T]) Compare(st any) int {

	list, ok := st.(List[T])
	if ok && l != nil && list != nil {

		if l.Len() < list.Len() {

			return -1

		}
		if l.Len() > list.Len() {

			return 1

		}
		other := list.Iter()
		for _, i := range l.objects {

			element, ok := interface{}(i).(util.Comparer)
			if !ok {

				return 0

			}
			if result := element.Compare(other.Element()); result != 0 {

				return result

			}
			other = other.Next()

		}
		return 0

	}
	return -2

}

// Hash returns the hash code of l.
func (l *ArrayList[T]) Hash() string {

	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("%v%v", check[1:], l.Len())

}

// Copy returns a list containing a copy of the elements of l.
// The result of this method is of type [List], but the effective list which is created is an [ArrayList].
func (l *ArrayList[T]) Copy() List[T] {

	return NewArrayListFromSlice(l.ToSlice())

}

// String returns a rapresentation of l in the form of a string.
func (l *ArrayList[T]) String() string {

	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("ArrayList[%v]%v", check[1:], l.objects)

}
