package table

import (
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ structures.Structure[int] = NewHashTable[wrapper.Int, int]()
var _ Table[wrapper.Int, int] = NewHashTable[wrapper.Int, int]()

// HashTable provides a generic table implemented through hashing.
//
// It implements the interface [Table].
type HashTable[K util.Hasher, T any] struct {
	// contains filtered or unexported fields
	objects map[string]list.List[*Entry[K, T]]
}

// NewHashTable returns a new empty [HashTable] containing the elements c.
func NewHashTable[K util.Hasher, T any]() *HashTable[K, T] {

	return &HashTable[K, T]{objects: map[string]list.List[*Entry[K, T]]{}}

}

// NewHashTableFromSlice returns a new [HashTable] containing the elements of slice c.
// it panics if key and c have different lengths.
func NewHashTableFromSlice[K util.Hasher, T any](key []K, c []T) *HashTable[K, T] {

	table := NewHashTable[K, T]()
	if len(c) != 0 {

		table.PutSlice(key, c)

	}
	return table

}

// Len returns the length of t.
func (t *HashTable[K, T]) Len() int {

	result := 0
	for _, i := range t.objects {

		result += i.Len()

	}
	return result

}

// IsEmpty returns a bool which indicates if t is empty or not.
func (t *HashTable[K, T]) IsEmpty() bool {

	return t.Len() == 0

}

// ContainsKey returns true if the key is present on t.
func (t *HashTable[K, T]) ContainsKey(key K) bool {

	hash := key.Hash()
	for i := range t.objects {

		if i == hash {

			for j := t.objects[i].Iter(); !j.End(); j = j.Next() {

				if key.Compare(j.Element().Key()) == 0 {

					return true

				}

			}

		}

	}
	return false

}

// ContainsElement returns true if the element e is present on t.
func (t *HashTable[K, T]) ContainsElement(e T) bool {

	element, ok := interface{}(e).(util.Equaler)
	for _, i := range t.objects {

		for j := i.Iter(); !j.End(); j = j.Next() {

			if ok {

				if element.Equal(j.Element().Element()) {

					return true

				}

			} else if reflect.DeepEqual(e, j.Element().Element()) {

				return true

			}

		}

	}
	return false

}

// Keys returns a [list.List] which contains all keys of t.
func (t *HashTable[K, T]) Keys() list.List[K] {

	list := list.NewArrayList[K]()
	for _, i := range t.objects {

		for j := i.Iter(); !j.End(); j = j.Next() {

			list.Add(j.Element().Key())

		}

	}
	return list

}

// Elements returns a [list.List] which contains all elements of t.
func (t *HashTable[K, T]) Elements() list.List[T] {

	list := list.NewArrayList[T]()
	for _, i := range t.objects {

		for j := i.Iter(); !j.End(); j = j.Next() {

			list.Add(j.Element().Element())

		}

	}
	return list

}

// ToSLice returns a slice which contains all elements of t.
func (t *HashTable[K, T]) ToSlice() []T {

	return t.Elements().ToSlice()

}

// Get returns the element associated at the key.
// The method returns false if the key is not found.
func (t *HashTable[K, T]) Get(key K) (T, bool) {

	var result T

	hash := key.Hash()
	for i := range t.objects {

		if i == hash {

			for j := t.objects[i].Iter(); !j.End(); j = j.Next() {

				if key.Compare(j.Element().Key()) == 0 {

					return j.Element().Element(), true

				}

			}

		}

	}
	return result, false

}

// Put set the element e at the key and returns the overwritten value, if present.
// If the element is not present, the method returns false.
func (t *HashTable[K, T]) Put(key K, e T) (T, bool) {

	var result T

	hash := key.Hash()
	for i := range t.objects {

		if i == hash {

			for j := t.objects[i].Iter(); !j.End(); j = j.Next() {

				if key.Compare(j.Element().Key()) == 0 {

					result = j.Element().Element()
					j.Element().SetElement(e)
					return result, true

				}

			}
			t.objects[i].Add(NewEntry(key, e))
			return result, false

		}

	}
	list := list.NewLinkedList(NewEntry(key, e))
	t.objects[hash] = list
	return result, false

}

// PutSlice adds the elements of e at t.
// it panics if key and e have different lengths.
func (t *HashTable[K, T]) PutSlice(key []K, e []T) {

	if len(key) != len(e) {

		panic("Different lengths for keys and elements")

	}
	for i := 0; i != len(key); i++ {

		t.Put(key[i], e[i])

	}

}

// Remove removes the key from t and returns the value associated at the key.
// It returns false if the the key does not exists.
func (t *HashTable[K, T]) Remove(key K) (T, bool) {

	var result T

	hash := key.Hash()
	for i := range t.objects {

		if i == hash {

			for j := t.objects[i].Iter(); !j.End(); j = j.Next() {

				if key.Compare(j.Element().Key()) == 0 {

					result = j.Element().Element()
					t.objects[i].RemoveElement(j.Element())
					if t.objects[i].IsEmpty() {

						delete(t.objects, i)

					}
					return result, true

				}

			}

		}

	}
	return result, false

}

// Clear removes all element from t.
func (t *HashTable[K, T]) Clear() {

	t.objects = map[string]list.List[*Entry[K, T]]{}

}

// Iter returns an [Iterator] which permits to iterate a [HashTable].
//
//	for i := t.Iter(); !i.End(); i = i.Next() {
//		key := i.Key()
//		element := i.Element()
//		// Code
//	}
func (t *HashTable[K, T]) Iter() Iterator[K, T] {

	return NewHashTableIterator(t)

}

// Equal returns true if t and st are both [Table] and their keys and elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is an [TreeTable],
// but the elements of t and the elements of st are equals, this method returns anyway true.
func (t *HashTable[K, T]) Equal(st any) bool {

	table, ok := st.(Table[K, T])
	if ok && t != nil && table != nil {

		if t.Len() != table.Len() {

			return false

		}
		for i := t.Keys().Iter(); !i.End(); i = i.Next() {

			e1, _ := t.Get(i.Element())
			other, found := table.Get(i.Element())
			if !found {

				return false

			}
			element, ok := interface{}(e1).(util.Equaler)
			if ok {

				if !element.Equal(other) {

					return false

				}

			} else if !reflect.DeepEqual(e1, other) {

				return false

			}

		}
		return true

	}
	return false

}

// Compare returns 0 if t and st have the same length,
// -1 if t is shorten than st,
// 1 if t is longer than st,
// -2 if st is not a [Table] or if one between t and st is nil.
func (t *HashTable[K, T]) Compare(st any) int {

	table, ok := st.(Table[K, T])
	if ok && t != nil && table != nil {

		if t.Len() < table.Len() {

			return -1

		}
		if t.Len() > table.Len() {

			return 1

		}
		return 0

	}
	return -2

}

// Hash returns the hash code of t.
func (t *HashTable[K, T]) Hash() string {

	check := []string{reflect.TypeOf(new(K)).String(), reflect.TypeOf(new(T)).String()}
	return fmt.Sprintf("%v%v%v", check[0][1:], check[1][1:], t.Len())

}

// Copy returns a table containing a copy of the elements of t.
// The result of this method is of type [Table], but the effective table which is created is an [HashTable].
func (t *HashTable[K, T]) Copy() Table[K, T] {

	table := NewHashTable[K, T]()
	for _, i := range t.objects {

		for j := i.Iter(); !j.End(); j = j.Next() {

			table.Put(j.Element().Key(), j.Element().Element())

		}

	}
	return table

}

// String returns a rapresentation of t in the form of a string.
func (t *HashTable[K, T]) String() string {

	check := []string{reflect.TypeOf(new(K)).String(), reflect.TypeOf(new(T)).String()}
	result := fmt.Sprintf("HashTable[%T, %T][", check[0][1:], check[1][1:])
	first := true
	for _, i := range t.objects {

		for j := i.Iter(); !j.End(); j = j.Next() {

			if !first {

				result += ", "

			}
			result += fmt.Sprintf("%v: %v", j.Element().Key(), j.Element().Element())
			first = false

		}

	}
	result += "]"
	return result

}
