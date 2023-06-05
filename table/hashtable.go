// package table implements dynamic tables.
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

// HashTable provides a generic table.
// The table is implemented through hashing.
//
// It implements the interface [structures.Structure].
//
// The check on the equality of the keys is done with the Compare method.
//
// The check on the equality of the elements is done with the Equal method if T implements [util.Equaler],
// otherwise it is done with [reflect.DeepEqual].
type HashTable[K util.Hasher, T any] struct {
	// contains filtered or unexported fields
	objects map[string]list.List[*Entry[K, T]]
}

// NewHashTable returns a new [HashTable] containing the elements c.
//
// if no argument is passed, it will be created an empty [HashTable].
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

			for j := range t.objects[i].Iter() {

				if key.Compare(j.Key()) == 0 {

					return true

				}

			}

		}

	}
	return false

}

// ContainsElement returns true if the element e is present on t.
func (t *HashTable[K, T]) ContainsElement(e T) bool {

	for _, i := range t.objects {

		for j := range i.Iter() {

			if reflect.DeepEqual(e, j.Element()) {

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

		for j := range i.Iter() {

			list.Add(j.Key())

		}

	}
	return list

}

// Keys returns a [list.List] which contains all elements of t.
func (t *HashTable[K, T]) Elements() list.List[T] {

	list := list.NewArrayList[T]()
	for _, i := range t.objects {

		for j := range i.Iter() {

			list.Add(j.Element())

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

			for j := range t.objects[i].Iter() {

				if key.Compare(j.Key()) == 0 {

					return j.Element(), true

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

			for j := range t.objects[i].Iter() {

				if key.Compare(j.Key()) == 0 {

					result = j.Element()
					j.SetElement(e)
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

	var result T

	hash := key.Hash()
	for i := range t.objects {

		if i == hash {

			for j := range t.objects[i].Iter() {

				if key.Compare(j.Key()) == 0 {

					result = j.Element()
					t.objects[i].RemoveElement(j)
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

// Equal returns true if t and st are both [HashTable] and their keys and elements are equals.
// In any other case, it returns false.
func (t *HashTable[K, T]) Equal(st any) bool {

	table, ok := st.(*HashTable[K, T])
	if ok && t != nil && table != nil {

		if t.Len() != table.Len() {

			return false

		}
		for i := range t.Keys().Iter() {

			e1, _ := t.Get(i)
			other, found := table.Get(i)
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
// -2 if st is not a [HashTable] or if one between t and st is nil.
func (t *HashTable[K, T]) Compare(st any) int {

	table, ok := st.(*HashTable[K, T])
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

// Copy returns an [HashTable] containing a copy of the elements of t.
func (t *HashTable[K, T]) Copy() *HashTable[K, T] {

	table := NewHashTable[K, T]()
	for _, i := range t.objects {

		for j := range i.Iter() {

			table.Put(j.Key(), j.Element())

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

		for j := range i.Iter() {

			if !first {

				result += ", "

			}
			result += fmt.Sprintf("%v: %v", j.Key(), j.Element())
			first = false

		}

	}
	result += "]"
	return result

}
