package table

import (
	"fmt"
	"hash/fnv"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ structures.Structure[int] = NewHashTable[wrapper.Int, int]()
var _ BaseTable[wrapper.Int, int] = NewHashTable[wrapper.Int, int]()
var _ Table[wrapper.Int, int] = NewHashTable[wrapper.Int, int]()

// HashTable provides a generic table implemented through hashing.
//
// It implements the interface [Table].
type HashTable[K util.Hasher, T any] struct {
	// contains filtered or unexported fields
	objects map[uint64]list.List[*Entry[K, T]]
}

// NewHashTable returns a new empty [HashTable].
func NewHashTable[K util.Hasher, T any]() *HashTable[K, T] {
	return &HashTable[K, T]{objects: map[uint64]list.List[*Entry[K, T]]{}}
}

// NewHashTableFromSlice returns a new [HashTable] containing the elements of slice c.
// It panics if key and c have different lengths.
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
	hash := t.objects[key.Hash()]
	if hash == nil {
		return false
	}
	for i := hash.Iter(); !i.End(); i = i.Next() {
		if key.Compare(i.Element().Key()) == 0 {
			return true
		}
	}
	return false
}

// ContainsElement returns true if the element e is present on t.
func (t *HashTable[K, T]) ContainsElement(e T) bool {
	fun := util.EqualFunction(e)
	for _, i := range t.objects {
		for j := i.Iter(); !j.End(); j = j.Next() {
			if fun(j.Element().Element()) {
				return true
			}
		}
	}
	return false
}

// Keys returns a [list.List] which contains all keys of t.
func (t *HashTable[K, T]) Keys() list.List[K] {
	list := list.NewArrayList[K]()
	t.Each(func(key K, _ T) {
		list.Add(key)
	})
	return list
}

// Elements returns a [list.List] which contains all elements of t.
func (t *HashTable[K, T]) Elements() list.List[T] {
	list := list.NewArrayList[T]()
	t.Each(func(_ K, element T) {
		list.Add(element)
	})
	return list
}

// ToSlice returns a slice which contains all elements of t.
func (t *HashTable[K, T]) ToSlice() []T {
	return t.Elements().ToSlice()
}

// Get returns the element associated at the key.
// The method returns false if the key is not found.
func (t *HashTable[K, T]) Get(key K) (T, bool) {

	var result T

	hash := t.objects[key.Hash()]
	if hash == nil {
		return result, false
	}
	for i := hash.Iter(); !i.End(); i = i.Next() {
		if key.Compare(i.Element().Key()) == 0 {
			return i.Element().Element(), true
		}
	}
	return result, false
}

// Put set the element e at the key and returns the overwritten value, if present.
// If the element is not present, the method returns false.
func (t *HashTable[K, T]) Put(key K, e T) (T, bool) {

	var result T

	hash := t.objects[key.Hash()]
	if hash == nil {
		list := list.NewLinkedList(NewEntry(key, e))
		t.objects[key.Hash()] = list
		return result, false
	}
	for i := hash.Iter(); !i.End(); i = i.Next() {
		if key.Compare(i.Element().Key()) == 0 {

			result = i.Element().Element()
			i.Element().SetElement(e)
			return result, true
		}
	}
	hash.Add(NewEntry(key, e))
	return result, false
}

// PutSlice adds the elements of e at t.
// It panics if key and e have different lengths.
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

	hash := t.objects[key.Hash()]
	if hash == nil {
		return result, false
	}
	for i := hash.Iter(); !i.End(); i = i.Next() {
		if key.Compare(i.Element().Key()) == 0 {
			result = i.Element().Element()
			hash.RemoveElement(i.Element())
			if hash.IsEmpty() {
				delete(t.objects, key.Hash())
			}
			return result, true
		}
	}
	return result, false
}

// Each executes fun for all elements of t.
//
// This method should be used to remove elements. Use Iter insted.
func (t *HashTable[K, T]) Each(fun func(key K, element T)) {
	for _, i := range t.objects {
		i.Each(func(_ int, element *Entry[K, T]) {
			fun(element.Key(), element.Element())
		})
	}
}

// Stream returns a [Stream] rapresenting t.
func (t *HashTable[K, T]) Stream() *Stream[K, T] {
	return NewStream[K, T](t, reflect.ValueOf(NewHashTable[K, T]))
}

// Clear removes all element from t.
func (t *HashTable[K, T]) Clear() {
	t.objects = map[uint64]list.List[*Entry[K, T]]{}
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
			if !found || !util.EqualFunction(e1)(other) {
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
func (t *HashTable[K, T]) Hash() uint64 {
	h := fnv.New64()
	for _, i := range t.objects {
		h.Write([]byte(fmt.Sprintf("%v", i.Hash())))
	}
	return h.Sum64()
}

// Copy returns a table containing a copy of the elements of t.
// The result of this method is of type [Table], but the effective table which is created is an [HashTable].
//
// This method uses [util.Copy] to make copies of the elements.
func (t *HashTable[K, T]) Copy() Table[K, T] {
	table := NewHashTable[K, T]()
	t.Each(func(key K, element T) {
		table.Put(key, util.Copy(element))
	})
	return table
}

// String returns a rapresentation of t in the form of a string.
func (t *HashTable[K, T]) String() string {
	check := []string{reflect.TypeOf(new(K)).String(), reflect.TypeOf(new(T)).String()}
	result := fmt.Sprintf("HashTable[%v, %v][", check[0][1:], check[1][1:])
	first := true
	t.Each(func(key K, element T) {
		if !first {
			result += ", "
		}
		result += fmt.Sprintf("%v: %v", key, element)
		first = false
	})
	result += "]"
	return result
}
