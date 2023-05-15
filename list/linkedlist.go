package list

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/potex02/structures"
)

// Entry is a component of a LinkedList.
// An entry is linked to the previous and the next entries.
type Entry[T any] struct {
	// contains filtered or unexported fields
	element T
	prev    *Entry[T]
	next    *Entry[T]
}

// NewEntry returns a new [Entry].
//
// element is the value of the entry.
// prev and next are thr entries to wich the entry is linked.
func NewEntry[T any](element T, prev *Entry[T], next *Entry[T]) *Entry[T] {

	return &Entry[T]{element: element, prev: prev, next: next}

}

// Element returns the element of e.
func (e *Entry[T]) Element() T {

	return e.element

}

// Element sets the element of e.
func (e *Entry[T]) SetElement(element T) {

	e.element = element

}

// Prev returns a pointer at the entry previous to e.
func (e *Entry[T]) Prev() *Entry[T] {

	return e.prev

}

// SetPrev sets the entry previous to e.
func (e *Entry[T]) SetPrev(prev *Entry[T]) {

	e.prev = prev

}

// Next returns a pointer at the entry next to e.
func (e *Entry[T]) Next() *Entry[T] {

	return e.next

}

// SetNext sets the entry next to e.
func (e *Entry[T]) SetNext(next *Entry[T]) {

	e.next = next

}

// LinkedList provides a generic double linked list.
//
// The list is implemented through a series of linked [Entry].
type LinkedList[T any] struct {
	// contains filtered or unexported fields
	root *Entry[T]
	tail *Entry[T]
	len  int
}

// NewLinkedList returns a new empty [LinkedList] containing the elements c.
//
// if no argument is passed, it will be created an empty [LinkedList].
func NewLinkedList[T any](c ...T) *LinkedList[T] {

	return NewLinkedListFromSlice(c)

}

// NewLinkedListFromSlice returns a new [LinkedList] containing the elements of slice c.
func NewLinkedListFromSlice[T any](c []T) *LinkedList[T] {

	list := &LinkedList[T]{root: nil, tail: nil, len: 0}
	if len(c) != 0 {

		list.AddSlice(c)

	}
	return list

}

// NewLinkedListFromStructure is a wrapper for NewLinkedListFromSlice(c.ToSlice()).
func NewLinkedListFromStructure[T any](c structures.Structure[T]) *LinkedList[T] {

	return NewLinkedListFromSlice(c.ToSlice())

}

// Len returns the length of l.
func (l *LinkedList[T]) Len() int {

	return l.len

}

// IsEmpty returns a bool wich indicate if l is empty or not.
func (l *LinkedList[T]) IsEmpty() bool {

	return l.len == 0

}

// Contains returns if e is present in l.
func (l *LinkedList[T]) Contains(e T) bool {

	return l.IndexOf(e) >= 0

}

// IndexOf returns the last position of e in l.
// If e is not present, the result is -1.
func (l *LinkedList[T]) IndexOf(e T) int {

	for i, j := 0, l.root; j != nil; i, j = i+1, j.Next() {

		if reflect.DeepEqual(j.Element(), e) {

			return i

		}

	}
	return -1

}

// LastIndexOf returns the last position of e in l.
// If e is not present, the result is -1.
func (l *LinkedList[T]) LastIndexOf(e T) int {

	for i, j := l.len-1, l.tail; j != nil; i, j = i-1, j.Prev() {

		if reflect.DeepEqual(j.Element(), e) {

			return i

		}

	}
	return -1

}

// ToSLice returns a slice wich contains all elements of l.
func (l *LinkedList[T]) ToSlice() []T {

	slice := make([]T, 0)
	for i := l.root; i != nil; i = i.Next() {

		slice = append(slice, i.Element())

	}
	return slice

}

// Get returns the elements at the specifies index.
// It returns an error if the the index is out of bounds.
func (l *LinkedList[T]) Get(index int) (T, error) {

	if !l.rangeCheck(index) {

		var result T

		return result, errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(l.len))

	}
	return l.getElementAtIndex(index).Element(), nil

}

// Set sets the value of element at the specified index and returns the overwritten value.
// It returns an error if the the index is out of bounds.
func (l *LinkedList[T]) Set(index int, e T) (T, error) {

	var result T

	if index == l.len {

		l.Add(e)
		return result, nil

	}
	if !l.rangeCheck(index) {

		return result, errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(l.len))

	}
	entry := l.getElementAtIndex(index)
	result = entry.Element()
	entry.SetElement(e)
	return result, nil

}

// Add adds the element e at the end of l.
func (l *LinkedList[T]) Add(e T) {

	entry := NewEntry(e, l.tail, nil)
	if l.len == 0 {

		l.tail = entry
		l.root = entry
		l.len++
		return

	}
	l.tail.SetNext(entry)
	entry.SetPrev(l.tail)
	l.tail = entry
	l.len++

}

// AddAtIndex adds the element e at the specified index.
// It returns an error if the the index is out of bounds.
func (l *LinkedList[T]) AddAtIndex(index int, e T) error {

	if index > l.len || index < 0 {

		return errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(l.len))

	}
	if index == l.len {

		l.Add(e)
		return nil

	}
	prev := l.getElementAtIndex(index - 1)
	l.append(prev, e)
	return nil

}

// AddElements is a wrapper for AddSlice(e).
func (l *LinkedList[T]) AddElements(e ...T) {

	l.AddSlice(e)

}

// AddElementsAtIndex is a wrapper for AddSliceAtIndex(index, e).
func (l *LinkedList[T]) AddElementsAtIndex(index int, e ...T) error {

	return l.AddSliceAtIndex(index, e)

}

// AddSlice adds the elements of e at the end of l.
func (l *LinkedList[T]) AddSlice(e []T) {

	for _, i := range e {

		l.Add(i)

	}

}

// AddSliceAtIndex adds the elements of e at the specified index.
// It returns an error if the the index is out of bounds.
func (l *LinkedList[T]) AddSliceAtIndex(index int, e []T) error {

	if index > l.len || index < 0 {

		return errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(l.len))

	}
	if index == l.len {

		l.AddSlice(e)
		return nil

	}
	prev := l.getElementAtIndex(index - 1)
	for _, i := range e {

		l.append(prev, i)
		if prev == nil {

			prev = l.root

		} else {

			prev = prev.Next()

		}

	}
	return nil

}

// Remove removes the element at specified index and return the removed value.
// It returns an error if the the index is out of bounds.
func (l *LinkedList[T]) Remove(index int) (T, error) {

	var result T

	if !l.rangeCheck(index) {

		return result, errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(l.len))

	}
	entry := l.getElementAtIndex(index)
	result = entry.Element()
	if entry.Prev() == nil {

		l.root = entry.Next()

	} else {

		entry.Prev().SetNext(entry.Next())

	}
	entry.Next().SetPrev(entry.Prev())
	l.len--
	return result, nil

}

// RemoveElement removes the element e from l if it is present.
// In that case, the method returns true, otherwhise it returns false.
func (l *LinkedList[T]) RemoveElement(e T) bool {

	for i, j := 0, l.root; j != nil; i, j = i+1, j.Next() {

		if reflect.DeepEqual(j.Element(), e) {

			l.Remove(i)
			return true

		}

	}
	return false

}

// Clear removes all element from l.
func (l *LinkedList[T]) Clear() {

	l.root = nil
	l.tail = nil
	l.len = 0

}

// Iter returns a chan wich permits to iterate l with the range keyword.
// This method can only be used to iterate a [LinkedList] if the index is not needed.
// For now, the only way to iterate a [LinkedList] with the index is the following code:
//
//	for i := 0; i < l.Len(); i++ {
//		element, err := l.Get(i)
//		// Code
//	}
//
// The code above should be used only if the index is really needed, because can be very expensive.
func (l *LinkedList[T]) Iter() chan T {

	obj := make(chan T)
	go func() {

		for i := l.root; i != nil; i = i.Next() {

			obj <- i.Element()

		}
		close(obj)

	}()
	return obj

}

// Equals returns true if l and st are both lists and their elements are equals.
// In any other case, it returns false.
//
// Equals does not take into account the effective type of st. This means that if st is an [ArrayList],
// but the elements of l and the elements of st are equals, this methods returns anyway true.
func (l *LinkedList[T]) Equals(st structures.Structure[T]) bool {

	list, ok := st.(List[T])
	return ok && reflect.DeepEqual(l.ToSlice(), list.ToSlice())

}

// Copy returns a list containing a copy of the elements of l.
// The result of this method is of type [List], but the effective list wich is created is a [LinkedList].
func (l *LinkedList[T]) Copy() List[T] {

	return NewLinkedListFromSlice(l.ToSlice())

}

// String returns a rapresentation of l in the form of a string.
func (l *LinkedList[T]) String() string {

	return fmt.Sprintf("LinkedList[%T]%v", *new(T), l.ToSlice())

}
func (l *LinkedList[T]) getElementAtIndex(index int) *Entry[T] {

	if index <= l.len/2 {

		result := l.root
		for i := 0; i != index; i++ {

			if result == nil {

				return nil

			}
			result = result.Next()

		}
		return result

	}
	result := l.tail
	for i := l.len - 1; i != index; i-- {

		if result == nil {

			return nil

		}
		result = result.Prev()

	}
	return result

}
func (l *LinkedList[T]) append(prev *Entry[T], e T) {

	entry := NewEntry(e, prev, nil)
	l.len++
	if prev == nil {

		entry.SetNext(l.root)
		l.root.SetPrev(entry)
		l.root = entry
		return

	}
	entry.SetNext(prev.Next())
	prev.Next().SetPrev(entry)
	prev.SetNext(entry)

}
func (l *LinkedList[T]) rangeCheck(index int) bool {

	return index >= 0 && index < l.len

}
