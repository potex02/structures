package table

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
	"github.com/potex02/structures/tree"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ structures.Structure[int] = NewTreeTable[wrapper.Int, int]()
var _ BaseTable[wrapper.Int, int] = NewTreeTable[wrapper.Int, int]()
var _ Table[wrapper.Int, int] = NewTreeTable[wrapper.Int, int]()

// TreeTable provides a generic table implemented through a [tree.BinaryTree].
// It maintains the order of the keys.
//
// It implements the interface [Table].
type TreeTable[K util.Comparer, T any] struct {
	// contains filtered or unexported fields
	objects *tree.BinaryTree[*Entry[K, T]]
}

// NewTreeTable returns a new empty [TreeTable].
func NewTreeTable[K util.Comparer, T any]() *TreeTable[K, T] {
	return &TreeTable[K, T]{objects: tree.NewBinaryTree[*Entry[K, T]]()}
}

// NewTreeTableFromSlice returns a new [TreeTable] containing the elements of slice c.
// It panics if key and c have different lengths.
func NewTreeTableFromSlice[K util.Comparer, T any](key []K, c []T) *TreeTable[K, T] {
	table := NewTreeTable[K, T]()
	if len(c) != 0 {
		table.PutSlice(key, c)
	}
	return table
}

// Len returns the length of t.
func (t *TreeTable[K, T]) Len() int {
	return t.objects.Len()
}

// IsEmpty returns a bool which indicates if t is empty or not.
func (t *TreeTable[K, T]) IsEmpty() bool {
	return t.objects.IsEmpty()
}

// ContainsKey returns true if the key is present on t.
func (t *TreeTable[K, T]) ContainsKey(key K) bool {
	return t.objects.Contains(NewEntry[K, T](key, *new(T)))
}

// ContainsElement returns true if the element e is present on t.
func (t *TreeTable[K, T]) ContainsElement(e T) bool {
	fun := util.EqualFunction(e)
	return t.objects.Any(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) bool {
		return fun(i.Element().Element())
	})
}

// Keys returns a [list.List] which contains all keys of t.
func (t *TreeTable[K, T]) Keys() list.List[K] {
	list := list.NewArrayList[K]()
	t.objects.Each(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) { list.Add(i.Element().Key()) })
	return list
}

// Elements returns a [list.List] which contains all elements of t.
func (t *TreeTable[K, T]) Elements() list.List[T] {
	list := list.NewArrayList[T]()
	t.objects.Each(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) { list.Add(i.Element().Element()) })
	return list
}

// ToSlice returns a slice which contains all elements of t.
func (t *TreeTable[K, T]) ToSlice() []T {
	slice := make([]T, t.Len())
	t.objects.Each(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) { slice = append(slice, i.Element().Element()) })
	return slice
}

// Get returns the element associated at the key.
// The method returns false if the key is not found.
func (t *TreeTable[K, T]) Get(key K) (T, bool) {

	var result T

	return result, t.objects.Any(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) bool {
		if key.Compare(i.Element().Key()) == 0 {
			result = i.Element().Element()
			return true
		}
		return false
	})
}

// Put set the element e at the key and returns the overwritten value, if present.
// If the element is not present, the method returns false.
func (t *TreeTable[K, T]) Put(key K, e T) (T, bool) {

	var result T

	found := t.objects.Any(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) bool {
		if check := key.Compare(i.Element().Key()); check == 0 {
			result = i.Element().Element()
			i.Element().SetElement(e)
			return true
		}
		return false
	})
	if !found {
		t.objects.Add(NewEntry(key, e))
	}
	return result, found
}

// PutSlice adds the elements of e at t.
// It panics if key and e have different lengths.
func (t *TreeTable[K, T]) PutSlice(key []K, e []T) {
	if len(key) != len(e) {
		panic("Different lengths for keys and elements")
	}
	for i := 0; i != len(key); i++ {
		t.Put(key[i], e[i])
	}
}

// Remove removes the key from t and returns the value associated at the key.
// It returns false if the the key does not exists.
func (t *TreeTable[K, T]) Remove(key K) (T, bool) {

	var result T

	return result, t.objects.Any(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) bool {
		if check := key.Compare(i.Element().Key()); check == 0 {
			result = i.Element().Element()
			t.objects.Remove(i.Element())
			return true
		}
		return false
	})
}

// Each executes fun for all elements of t.
//
// This method should be used to remove elements. Use Iter insted.
func (t *TreeTable[K, T]) Each(fun func(key K, element T)) {
	t.objects.Each(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) {
		fun(i.Element().Key(), i.Element().Element())
	})
}

// Stream returns a [Stream] rapresenting t.
func (t *TreeTable[K, T]) Stream() *Stream[K, T] {
	return NewStream[K, T](t, reflect.ValueOf(NewTreeTable[K, T]))
}

// Clear removes all element from t.
func (t *TreeTable[K, T]) Clear() {
	t.objects.Clear()
}

// Iter returns an [Iterator] which permits to iterate a [TreeTable].
//
//	for i := t.Iter(); !i.End(); i = i.Next() {
//		key := i.Key()
//		element := i.Element()
//		// Code
//	}
func (t *TreeTable[K, T]) Iter() Iterator[K, T] {
	return NewTreeTableIterator(t)
}

// RangeIter returns a function that allows to iterate a [TreeTable] using the range keyword.
//
//	for i := range t.RangeIter() {
//		// Code
//	}
//
// Unlike [TreeTable.Iter], it doesn't allow to remove elements during the iteration.
func (t *TreeTable[K, T]) RangeIter() func(yield func(K, T) bool) {
	return func(yield func(K, T) bool) {
		for i := range t.objects.RangeIter() {
			if !yield(i.Key(), i.Element()) {
				return
			}
		}
	}
}

// Equal returns true if t and st are both [Table] and their keys and elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is a [HashTable],
// but the elements of t and the elements of st are equals, this method returns anyway true.
func (t *TreeTable[K, T]) Equal(st any) bool {
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
			if !util.EqualFunction(e1)(other) {
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
func (t *TreeTable[K, T]) Compare(st any) int {
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
func (t *TreeTable[K, T]) Hash() uint64 {
	h := fnv.New64()
	t.objects.Each(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) {
		h.Write([]byte(fmt.Sprintf("%v", i.Hash())))
	})
	return h.Sum64()
}

// Copy returns a table containing a copy of the elements of t.
// The result of this method is of type [Table], but the effective table which is created is an [TreeTable].
//
// This method uses [util.Copy] to make copies of the elements.
func (t *TreeTable[K, T]) Copy() Table[K, T] {
	slice := t.objects.ToSlice()
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	result := NewTreeTable[K, T]()
	for _, i := range slice {
		result.Put(i.Key(), util.Copy(i.Element()))
	}
	return result
}

// String returns a rapresentation of t in the form of a string.
func (t *TreeTable[K, T]) String() string {
	check := []string{reflect.TypeOf(new(K)).String(), reflect.TypeOf(new(T)).String()}
	result := fmt.Sprintf("TreeTable[%v, %v][", check[0][1:], check[1][1:])
	first := true
	t.objects.Each(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) {
		if !first {
			result += ", "
		}
		result += fmt.Sprintf("%v: %v", i.Element().Key(), i.Element().Element())
		first = false
	})
	result += "]"
	return result
}
