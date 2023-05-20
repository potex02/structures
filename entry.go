package structures

// Entry is a component of a linked structure.
//
// An entry is linked to the previous and the next entries.
// However, this type is used for both double and single linked structures, simply setting the prev entry at nil.
type Entry[T any] struct {
	// contains filtered or unexported fields
	element T
	prev    *Entry[T]
	next    *Entry[T]
}

// NewEntry returns a new [Entry].
//
// Element is the value of the entry.
// Prev and next are the entries to which the entry is linked.
func NewEntry[T any](element T, prev *Entry[T], next *Entry[T]) *Entry[T] {

	return &Entry[T]{element: element, prev: prev, next: next}

}

// NewEntrySlice creates a series of linked entries which containing the elements of e.
// The first and the last entries of the series are returned.
func NewEntrySlice[T any](e []T) (*Entry[T], *Entry[T]) {

	if len(e) == 0 {

		return nil, nil

	}
	first := NewEntry(e[0], nil, nil)
	current := first
	for i := 1; i != len(e); i++ {

		current.next = NewEntry(e[i], current, nil)
		current = current.next

	}
	return first, current

}

// NewEntrySingle returns a new [Entry] for a single linked structure.
//
// Element is the value of the entry.
// Next is the entry to which the entry is linked.
func NewEntrySingle[T any](element T, next *Entry[T]) *Entry[T] {

	return &Entry[T]{element: element, prev: nil, next: next}

}

// NewEntrySliceSingle creates a series of single linked entries which containing the elements of e.
// The first and the last entries of the series are returned.
func NewEntrySliceSingle[T any](e []T) (*Entry[T], *Entry[T]) {

	if len(e) == 0 {

		return nil, nil

	}
	first := NewEntrySingle(e[len(e)-1], nil)
	current := first
	for i := len(e) - 2; i >= 0; i-- {

		current.next = NewEntrySingle(e[i], nil)
		current = current.next

	}
	return first, current

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
