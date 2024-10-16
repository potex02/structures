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

var _ structures.Structure[int] = NewMultiHashTable[wrapper.Int, int]()
var _ BaseTable[wrapper.Int, int] = NewMultiHashTable[wrapper.Int, int]()
var _ MultiTable[wrapper.Int, int] = NewMultiHashTable[wrapper.Int, int]()

// MultiHashTable provides a generic table with duplicate keys implemented through hashing.
//
// It implements the interface [MultiTable].
type MultiHashTable[K util.Hasher, T any] struct {
	// contains filtered or unexported fields
	objects map[uint64]list.List[*Entry[K, T]]
}

// NewHashTable returns a new empty [MultiHashTable].
func NewMultiHashTable[K util.Hasher, T any]() *MultiHashTable[K, T] {
	return &MultiHashTable[K, T]{objects: map[uint64]list.List[*Entry[K, T]]{}}
}

// NewMultiHashTableFromSlice returns a new [MultiHashTable] containing the elements of slice c.
// It panics if key and c have different lengths.
func NewMultiHashTableFromSlice[K util.Hasher, T any](key []K, c []T) *MultiHashTable[K, T] {
	table := NewMultiHashTable[K, T]()
	if len(c) != 0 {
		table.PutSlice(key, c)
	}
	return table
}

// Len returns the length of t.
func (t *MultiHashTable[K, T]) Len() int {
	result := 0
	for _, i := range t.objects {
		result += i.Len()
	}
	return result
}

// IsEmpty returns a bool which indicates if t is empty or not.
func (t *MultiHashTable[K, T]) IsEmpty() bool {
	return t.Len() == 0
}

// Contains returns true if the key is present in on t associated with the element e.
func (t *MultiHashTable[K, T]) Contains(key K, e T) bool {
	fun := util.EqualFunction(e)
	hash := t.objects[key.Hash()]
	if hash == nil {
		return false
	}
	for i := hash.Iter(); !i.End(); i = i.Next() {
		if key.Compare(i.Element().Key()) == 0 && fun(i.Element().Element()) {
			return true
		}
	}
	return false
}

// ContainsKey returns true if the key is present on t.
func (t *MultiHashTable[K, T]) ContainsKey(key K) bool {
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

// ContainsElement returns true if the element e is associated at any key of t.
func (t *MultiHashTable[K, T]) ContainsElement(e T) bool {
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
func (t *MultiHashTable[K, T]) Keys() list.List[K] {
	list := list.NewArrayList[K]()
	t.Each(func(key K, _ T) {
		list.Add(key)
	})
	return list
}

// Elements returns a [list.List] which contains all elements of t.
func (t *MultiHashTable[K, T]) Elements() list.List[T] {
	list := list.NewArrayList[T]()
	t.Each(func(_ K, element T) {
		list.Add(element)
	})
	return list
}

// ToSlice returns a slice which contains all elements of t.
func (t *MultiHashTable[K, T]) ToSlice() []T {
	return t.Elements().ToSlice()
}

// Get returns a slice cotaining the elements associated at the key.
func (t *MultiHashTable[K, T]) Get(key K) []T {
	result := make([]T, 0)
	hash := t.objects[key.Hash()]
	if hash == nil {
		return result
	}
	hash.Each(func(_ int, element *Entry[K, T]) {
		if key.Compare(element.Key()) == 0 {
			result = append(result, element.Element())
		}
	})
	return result
}

// Put add the elements of e at the key.
func (t *MultiHashTable[K, T]) Put(key K, e ...T) {
	hash := t.objects[key.Hash()]
	if hash == nil {
		list := list.NewLinkedList[*Entry[K, T]]()
		for _, j := range e {
			list.Add(NewEntry(key, j))
		}
		t.objects[key.Hash()] = list
		return
	}
	for _, j := range e {
		hash.Add(NewEntry(key, j))
	}
}

// PutSlice adds the elements of e at the table.
// It panics if key and e have different lengths.
func (t *MultiHashTable[K, T]) PutSlice(key []K, e []T) {
	if len(key) != len(e) {
		panic("Different lengths for keys and elements")
	}
	for i := 0; i != len(key); i++ {
		t.Put(key[i], e[i])
	}
}

// Replace replace all elements associated at the key with e and returns the slice of overwritten values.
func (t *MultiHashTable[K, T]) Replace(key K, e ...T) []T {
	result := t.RemoveKey(key)
	t.Put(key, e...)
	return result
}

// ReplaceSlice replace all elements associated at the key with e and returns the slice of overwritten values.
func (t *MultiHashTable[K, T]) ReplaceSlice(key []K, e []T) []T {
	if len(key) != len(e) {
		panic("Different lengths for keys and elements")
	}
	result := make([]T, 0)
	for _, i := range key {
		result = append(result, t.RemoveKey(i)...)
	}
	t.PutSlice(key, e)
	return result
}

// Remove removes the key associated at e from t.
// It returns false if the the entry does not exists.
func (t *MultiHashTable[K, T]) Remove(key K, e T) bool {
	fun := util.EqualFunction(e)
	hash := t.objects[key.Hash()]
	if hash == nil {
		return false
	}
	for i := hash.Iter(); !i.End(); i = i.Next() {
		if key.Compare(i.Element().Key()) == 0 && fun(i.Element().Element()) {
			hash.RemoveElement(i.Element())
			if hash.IsEmpty() {
				delete(t.objects, key.Hash())
			}
			return true
		}
	}
	return false
}

// RemoveKey remove all elements associated at the key and returns the slice of removed values.
func (t *MultiHashTable[K, T]) RemoveKey(key K) []T {
	result := make([]T, 0)
	hash := t.objects[key.Hash()]
	if hash == nil {
		return result
	}
	for i := hash.Iter(); !i.End(); i = i.Next() {
		for !i.End() && key.Compare(i.Element().Key()) == 0 {
			result = append(result, i.Element().Element())
			i = i.Remove()
		}
	}
	if hash.IsEmpty() {
		delete(t.objects, key.Hash())
	}
	return result
}

// Each executes fun for all elements of t.
//
// This method should be used to remove elements. Use Iter insted.
func (t *MultiHashTable[K, T]) Each(fun func(key K, element T)) {
	for _, i := range t.objects {
		i.Each(func(_ int, element *Entry[K, T]) {
			fun(element.Key(), element.Element())
		})
	}
}

// Stream returns a [MultiStream] rapresenting t.
func (t *MultiHashTable[K, T]) Stream() *Stream[K, T] {
	return NewStream[K, T](t, reflect.ValueOf(NewMultiHashTable[K, T]))
}

// Clear removes all element from t.
func (t *MultiHashTable[K, T]) Clear() {
	t.objects = map[uint64]list.List[*Entry[K, T]]{}
}

// Iter returns an [Iterator] which permits to iterate a [MultiHashTable].
//
//	for i := t.Iter(); !i.End(); i = i.Next() {
//		key := i.Key()
//		element := i.Element()
//		// Code
//	}
func (t *MultiHashTable[K, T]) Iter() Iterator[K, T] {
	return NewMultiHashTableIterator(t)
}

// Equal returns true if t and st are both multitables and their keys and elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is an [MultiTreeTable],
// but the elements of t and the elements of st are equals, this method returns anyway true.
func (t *MultiHashTable[K, T]) Equal(st any) bool {
	table, ok := st.(MultiTable[K, T])
	if ok && t != nil && table != nil {
		if t.Len() != table.Len() {
			return false
		}
		for _, i := range t.objects {
			for j := i.Iter(); !j.End(); j = j.Next() {
				if !table.Contains(j.Element().Key(), j.Element().Element()) {
					return false
				}
			}
		}
		return true
	}
	return false
}

// Compare returns 0 if t and st have the same length,
// -1 if t is shorten than st,
// 1 if t is longer than st,
// -2 if st is not a [MultiTable] or if one between t and st is nil.
func (t *MultiHashTable[K, T]) Compare(st any) int {
	table, ok := st.(MultiTable[K, T])
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
func (t *MultiHashTable[K, T]) Hash() uint64 {
	h := fnv.New64()
	for _, i := range t.objects {
		h.Write([]byte(fmt.Sprintf("%v", i.Hash())))
	}
	return h.Sum64()
}

// Copy returns a multitable containing a copy of the elements of t.
// The result of this method is of type [MultiTable], but the effective table which is created is an [MultiHashTable].
//
// This method uses [util.Copy] to make copies of the elements.
func (t *MultiHashTable[K, T]) Copy() MultiTable[K, T] {
	table := NewMultiHashTable[K, T]()
	t.Each(func(key K, element T) {
		table.Put(key, util.Copy(element))
	})
	return table
}

// String returns a rapresentation of t in the form of a string.
func (t *MultiHashTable[K, T]) String() string {
	check := []string{reflect.TypeOf(new(K)).String(), reflect.TypeOf(new(T)).String()}
	result := fmt.Sprintf("MultiHashTable[%v, %v][", check[0][1:], check[1][1:])
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
