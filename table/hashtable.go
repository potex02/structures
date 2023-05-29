// package table implements dynamic tables.
package table

import (
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
)

// HashTable provides a generic table.
// The table is implemented through hashing.
//
// It implements the interface [structures.Structure].
type HashTable[K comparable, T any] struct {
	// contains filtered or unexported fields
	objects map[K]T
}

// NewHashTable returns a new [HashTable] containing the elements c.
//
// if no argument is passed, it will be created an empty [HashTable].
func NewHashTable[K comparable, T any]() *HashTable[K, T] {

	return &HashTable[K, T]{objects: map[K]T{}}

}

// NewHashTableFromSlice returns a new [HashTable] containing the elements of slice c.
// it panics if key and c have different lengths.
func NewHashTableFromSlice[K comparable, T any](key []K, c []T) *HashTable[K, T] {

	table := NewHashTable[K, T]()
	if len(c) != 0 {

		table.PutSlice(key, c)

	}
	return table

}

// Len returns the length of t.
func (t *HashTable[K, T]) Len() int {

	return len(t.objects)

}

// IsEmpty returns a bool which indicate if t is empty or not.
func (t *HashTable[K, T]) IsEmpty() bool {

	return len(t.objects) == 0

}

// ContainsKey returns true if the key is present on t.
func (t *HashTable[K, T]) ContainsKey(key K) bool {

	_, ok := t.objects[key]
	return ok

}

// ContainsElement returns true if the element e is present on t.
func (t *HashTable[K, T]) ContainsElement(e T) bool {

	for _, i := range t.objects {

		if reflect.DeepEqual(i, e) {

			return true

		}

	}
	return false

}

// Keys returns a [list.List] which contains all keys of t.
func (t *HashTable[K, T]) Keys() list.List[K] {

	list := list.NewArrayList[K]()
	j := 0
	for i := range t.objects {

		list.Add(i)
		j++

	}
	return list

}

// Keys returns a [list.List] which contains all elements of t.
func (t *HashTable[K, T]) Elements() list.List[T] {

	list := list.NewArrayList[T]()
	j := 0
	for _, i := range t.objects {

		list.Add(i)
		j++

	}
	return list

}

// ToSLice returns a slice which contains all elements of t.
func (t *HashTable[K, T]) ToSlice() []T {

	slice := make([]T, len(t.objects))
	j := 0
	for _, i := range t.objects {

		slice[j] = i
		j++

	}
	return slice

}

// Get returns the element associated at the key.
// The method returns false if the key is not found.
func (t *HashTable[K, T]) Get(key K) (T, bool) {

	result, ok := t.objects[key]
	return result, ok

}

// Put set the element e at the key and returns the overwritten value, if present.
// If the element is not present, the method returns false.
func (t *HashTable[K, T]) Put(key K, e T) (T, bool) {

	result, ok := t.objects[key]
	t.objects[key] = e
	return result, ok

}

// PutSlice adds the elements of e at t if not present.
// it panics if key and e have different lengths.
func (t *HashTable[K, T]) PutSlice(key []K, e []T) {

	if len(key) != len(e) {

		panic("Different lengths for keys and elements")

	}
	for i := 0; i != len(key); i++ {

		t.Put(key[i], e[i])

	}

}

// Remove removes the key from t and return the value associated at the key.
// It returns false if the the key does not exists.
func (t *HashTable[K, T]) Remove(key K) (T, bool) {

	result, ok := t.objects[key]
	if ok {

		delete(t.objects, key)

	}
	return result, ok

}

// Clear removes all element from t.
func (t *HashTable[K, T]) Clear() {

	t.objects = map[K]T{}

}

// Equals returns true if t and st are both [HashTable] and their keys and elements are equals.
// In any other case, it returns false.
func (t *HashTable[K, T]) Equals(st structures.Structure[T]) bool {

	return reflect.DeepEqual(t, st)

}

// Copy returns an [HashTable] containing a copy of the elements of t.
func (t *HashTable[K, T]) Copy() *HashTable[K, T] {

	table := NewHashTable[K, T]()
	for i, j := range t.objects {

		table.Put(i, j)

	}
	return table

}

// String returns a rapresentation of t in the form of a string
func (t *HashTable[K, T]) String() string {

	return fmt.Sprintf("HashTable[%T, %T]%v", *new(K), *new(T), t.objects)

}
