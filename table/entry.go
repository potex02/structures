package table

import (
	"github.com/potex02/structures/util"
)

type Entry[K util.Hasher[K], T any] struct {
	key     K
	element T
}

func NewEntry[K util.Hasher[K], T any](Key K, element T) *Entry[K, T] {

	return &Entry[K, T]{key: Key, element: element}

}
func (e *Entry[K, T]) Key() K {

	return e.key

}
func (e *Entry[K, T]) Element() T {

	return e.element

}
func (e *Entry[K, T]) SetElement(element T) {

	e.element = element

}
