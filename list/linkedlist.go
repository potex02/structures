package list

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util"
)

var _ structures.Structure[int] = NewLinkedList[int]()
var _ List[int] = NewLinkedList[int]()

// LinkedList provides a generic double linked list.
// The list is implemented through a series of linked [structures.Entry].
//
// It implements the interface [List].
type LinkedList[T any] struct {
	// contains filtered or unexported fields
	root *structures.Entry[T]
	tail *structures.Entry[T]
	len  int
}

// NewLinkedList returns a new [LinkedList] containing the elements c.
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

// IsEmpty returns a bool which indicates if l is empty or not.
func (l *LinkedList[T]) IsEmpty() bool {

	return l.len == 0

}

// Contains returns if e is present in l.
func (l *LinkedList[T]) Contains(e T) bool {

	return l.IndexOf(e) >= 0

}

// IndexOf returns the first position of e in l.
// If e is not present, the result is -1.
func (l *LinkedList[T]) IndexOf(e T) int {

	element, ok := interface{}(e).(util.Equaler)
	for i, j := 0, l.root; j != nil; i, j = i+1, j.Next() {

		if ok {

			if element.Equal(j.Element()) {

				return i

			}

		} else if reflect.DeepEqual(j.Element(), e) {

			return i

		}

	}
	return -1

}

// LastIndexOf returns the last position of e in l.
// If e is not present, the result is -1.
func (l *LinkedList[T]) LastIndexOf(e T) int {

	element, ok := interface{}(e).(util.Equaler)
	for i, j := l.len-1, l.tail; j != nil; i, j = i-1, j.Prev() {

		if ok {

			if element.Equal(j.Element()) {

				return i

			}

		} else if reflect.DeepEqual(j.Element(), e) {

			return i

		}

	}
	return -1

}

// ToSLice returns a slice which contains all elements of l.
func (l *LinkedList[T]) ToSlice() []T {

	slice := make([]T, l.len)
	j := 0
	for i := l.root; i != nil; i = i.Next() {

		slice[j] = i.Element()
		j++

	}
	return slice

}

// Get returns the elements at the specifies index.
// It returns an error if the the index is out of bounds.
func (l *LinkedList[T]) Get(index int) (T, error) {

	if !rangeCheck[T](l, index) {

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
	if !rangeCheck[T](l, index) {

		return result, errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(l.len))

	}
	entry := l.getElementAtIndex(index)
	result = entry.Element()
	entry.SetElement(e)
	return result, nil

}

// Add adds the elements e at the end of l.
func (l *LinkedList[T]) Add(e ...T) {

	l.AddSlice(e)

}

// AddAtIndex adds the elements e at the specified index.
// It returns an error if the the index is out of bounds.
func (l *LinkedList[T]) AddAtIndex(index int, e ...T) error {

	return l.AddSliceAtIndex(index, e)

}

// AddSlice adds the elements of e at the end of l.
func (l *LinkedList[T]) AddSlice(e []T) {

	if len(e) == 0 {

		return

	}
	first, last := structures.NewEntrySlice(e)
	if first == nil || last == nil {

		return

	}
	if l.len == 0 {

		l.root = first
		l.tail = last
		l.len = len(e)
		return

	}
	l.tail.SetNext(first)
	first.SetPrev(l.tail)
	l.tail = last
	l.len += len(e)

}

// AddSliceAtIndex adds the elements of e at the specified index.
// It returns an error if the the index is out of bounds.
func (l *LinkedList[T]) AddSliceAtIndex(index int, e []T) error {

	if len(e) == 0 {

		return nil

	}
	if index > l.len || index < 0 {

		return errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(l.len))

	}
	if index == l.len {

		l.AddSlice(e)
		return nil

	}
	first, last := structures.NewEntrySlice(e)
	if index == 0 {

		l.root.SetPrev(last)
		last.SetNext(l.root)
		l.root = first
		l.len += len(e)
		return nil

	}
	prev := l.getElementAtIndex(index - 1)
	next := prev.Next()
	prev.SetNext(first)
	first.SetPrev(prev)
	last.SetNext(next)
	next.SetPrev(last)
	l.len += len(e)
	return nil

}

// Remove removes the element at specified index and return the removed value.
// It returns an error if the the index is out of bounds.
func (l *LinkedList[T]) Remove(index int) (T, error) {

	var result T

	if !rangeCheck[T](l, index) {

		return result, errors.New("Index " + strconv.Itoa(index) + " for size " + strconv.Itoa(l.len))

	}
	entry := l.getElementAtIndex(index)
	result = entry.Element()
	l.removeEntry(entry)
	return result, nil

}

// RemoveElement removes the element e from l if it is present.
// In that case, the method returns true, otherwhise it returns false.
func (l *LinkedList[T]) RemoveElement(e T) bool {

	element, ok := interface{}(e).(util.Equaler)
	for i, j := 0, l.root; j != nil; i, j = i+1, j.Next() {

		if ok {

			if element.Equal(j.Element()) {

				l.Remove(i)
				return true

			}

		} else if reflect.DeepEqual(j.Element(), e) {

			l.Remove(i)
			return true

		}

	}
	return false

}

// Each executes fun for all elements of l.
func (l *LinkedList[T]) Each(fun func(index int, element T)) {

	for i, j := 0, l.root; j != nil; i, j = i+1, j.Next() {

		fun(i, j.Element())

	}

}

// Stream returns a [Stream] rapresenting l.
func (l *LinkedList[T]) Stream() *Stream[T] {

	return NewStream[T](l, reflect.ValueOf(NewLinkedList[T]))

}

// Sort sorts the elements of l.
//
// This method panics if T does not implement [util.Comparer]
func (l *LinkedList[T]) Sort() {

	other := l.tail.Next()
	for swapped := true; swapped; {

		swapped = false
		element := l.root
		for element.Next() != other {

			if interface{}(element.Element()).(util.Comparer).Compare(element.Next().Element()) > 0 {

				t := element.Element()
				element.SetElement(element.Next().Element())
				element.Next().SetElement(t)
				swapped = true

			}
			element = element.Next()

		}
		other = element

	}

}

// SortFunc sorts the elements of l as determined by the less function.
func (l *LinkedList[T]) SortFunc(less func(i T, j T) bool) {

	other := l.tail.Next()
	for swapped := true; swapped; {

		swapped = false
		element := l.root
		for element.Next() != other {

			if !less(element.Element(), element.Next().Element()) {

				t := element.Element()
				element.SetElement(element.Next().Element())
				element.Next().SetElement(t)
				swapped = true

			}
			element = element.Next()

		}
		other = element

	}

}

// Clear removes all element from l.
func (l *LinkedList[T]) Clear() {

	l.root = nil
	l.tail = nil
	l.len = 0

}

// Iter returns an [Iterator] which permits to iterate a [LinkedList].
//
//	for i := l.Iter(); !i.End(); i = i.Next() {
//		element := i.Element()
//		index := i.Index()
//		// Code
//	}
func (l *LinkedList[T]) Iter() Iterator[T] {

	return NewLinkedListIterator(l)

}

// IterReverse returns an [Iterator] which permits to iterate a [LinkedList] in reverse order.
//
//	for i := l.IterReverse(); !i.End(); i = i.Prev() {
//		element := i.Element()
//		index := i.Index()
//		// Code
//	}
func (l *LinkedList[T]) IterReverse() Iterator[T] {

	return NewLinkedListReverseIterator(l)

}

// Equal returns true if l and st are both lists and their elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is an [ArrayList],
// but the elements of l and the elements of st are equals, this method returns anyway true.
func (l *LinkedList[T]) Equal(st any) bool {

	list, ok := st.(List[T])
	if ok && l != nil && list != nil {

		if l.Len() != list.Len() {

			return false

		}
		other := list.Iter()
		for i := l.root; i != nil; i = i.Next() {

			element, ok := interface{}(i.Element()).(util.Equaler)
			if ok {

				if !element.Equal(other.Element()) {

					return false

				}

			} else if !reflect.DeepEqual(i.Element(), other.Element()) {

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
func (l *LinkedList[T]) Compare(st any) int {

	list, ok := st.(List[T])
	if ok && l != nil && list != nil {

		if l.Len() < list.Len() {

			return -1

		}
		if l.Len() > list.Len() {

			return 1

		}
		other := list.Iter()
		for i := l.root; i != nil; i = i.Next() {

			element, ok := interface{}(i.Element()).(util.Comparer)
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
func (l *LinkedList[T]) Hash() string {

	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("%v%v", check[1:], l.Len())

}

// Copy returns a list containing a copy of the elements of l.
// The result of this method is of type [List], but the effective list which is created is a [LinkedList].
func (l *LinkedList[T]) Copy() List[T] {

	return NewLinkedListFromSlice(l.ToSlice())

}

// String returns a rapresentation of l in the form of a string.
func (l *LinkedList[T]) String() string {

	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("LinkedList[%v]%v", check[1:], l.ToSlice())

}

func (l *LinkedList[T]) getElementAtIndex(index int) *structures.Entry[T] {

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

func (l *LinkedList[T]) removeEntry(entry *structures.Entry[T]) {

	if entry.Prev() == nil {

		l.root = entry.Next()

	} else {

		entry.Prev().SetNext(entry.Next())

	}
	if entry.Next() == nil {

		l.tail = entry.Prev()

	} else {

		entry.Next().SetPrev(entry.Prev())

	}
	l.len--

}
