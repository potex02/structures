package table

import (
	"github.com/potex02/structures/util"
)

// Entry is a component of a hash structure.
type Entry[K util.Comparer, T any] struct {
	// contains filtered or unexported fields
	key     K
	element T
}

// NewEntry returns a new [Entry].
func NewEntry[K util.Comparer, T any](key K, element T) *Entry[K, T] {

	return &Entry[K, T]{key: key, element: element}

}

// Key returns the key of e.
func (e *Entry[K, T]) Key() K {

	return e.key

}

// Element returns the element of e.
func (e *Entry[K, T]) Element() T {

	return e.element

}

// Element sets the element of e.
func (e *Entry[K, T]) SetElement(element T) {

	e.element = element

}

// Compare returns the comparison between the key of e and o.
func (e *Entry[K, T]) Compare(o any) int {

	entry, ok := o.(*Entry[K, T])
	if ok && e != nil && entry != nil {

		return e.key.Compare(entry.Key())

	}
	return -2

}
