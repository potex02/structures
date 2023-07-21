package table

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/list"
	"github.com/potex02/structures/tree"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ structures.Structure[int] = NewMultiTreeTable[wrapper.Int, int]()
var _ BaseTable[wrapper.Int, int] = NewMultiTreeTable[wrapper.Int, int]()
var _ MultiTable[wrapper.Int, int] = NewMultiTreeTable[wrapper.Int, int]()

// MultiTreeTable provides a generic table with duplicate keys implemented through a [tree.BinaryTree].
// It maintains the order of the keys.
//
// It implements the interface [MultiTable].
type MultiTreeTable[K util.Comparer, T any] struct {
	// contains filtered or unexported fields
	objects *tree.BinaryTree[*Entry[K, T]]
}

// NewMultiTreeTable returns a new empty [MultiTreeTable] containing the elements c.
func NewMultiTreeTable[K util.Comparer, T any]() *MultiTreeTable[K, T] {

	return &MultiTreeTable[K, T]{objects: tree.NewBinaryTree[*Entry[K, T]]()}

}

// NewMultiTreeTableFromSlice returns a new [MultiTreeTable] containing the elements of slice c.
// It panics if key and c have different lengths.
func NewMultiTreeTableFromSlice[K util.Comparer, T any](key []K, c []T) *MultiTreeTable[K, T] {

	table := NewMultiTreeTable[K, T]()
	if len(c) != 0 {

		table.PutSlice(key, c)

	}
	return table

}

// Len returns the length of t.
func (t *MultiTreeTable[K, T]) Len() int {

	return t.objects.Len()

}

// IsEmpty returns a bool which indicates if t is empty or not.
func (t *MultiTreeTable[K, T]) IsEmpty() bool {

	return t.objects.IsEmpty()

}

// Contains returns true if the key is present in on t associated with the element e.
func (t *MultiTreeTable[K, T]) Contains(key K, e T) bool {

	fun := util.EqualFunction(e)
	return t.objects.Any(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) bool {
		return key.Compare(i.Element().Key()) == 0 && fun(i.Element().Element())
	})

}

// ContainsKey returns true if the key is present on t.
func (t *MultiTreeTable[K, T]) ContainsKey(key K) bool {

	return t.objects.Any(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) bool {
		return key.Compare(i.Element().Key()) == 0
	})

}

// ContainsElement returns true if the element e is associated at any key of t.
func (t *MultiTreeTable[K, T]) ContainsElement(e T) bool {

	fun := util.EqualFunction(e)
	return t.objects.Any(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) bool {
		return fun(i.Element().Element())
	})

}

// Keys returns a [list.List] which contains all keys of t.
func (t *MultiTreeTable[K, T]) Keys() list.List[K] {

	list := list.NewArrayList[K]()
	t.objects.Each(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) { list.Add(i.Element().Key()) })
	return list

}

// Elements returns a [list.List] which contains all elements of t.
func (t *MultiTreeTable[K, T]) Elements() list.List[T] {

	list := list.NewArrayList[T]()
	t.objects.Each(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) { list.Add(i.Element().Element()) })
	return list

}

// ToSlice returns a slice which contains all elements of t.
func (t *MultiTreeTable[K, T]) ToSlice() []T {

	return t.Elements().ToSlice()

}

// Get returns a slice cotaining the elements associated at the key.
func (t *MultiTreeTable[K, T]) Get(key K) []T {

	result := make([]T, 0)
	t.objects.All(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) bool {
		check := key.Compare(i.Element().Key())
		if check == 0 {
			result = append(result, i.Element().Element())
		}
		return check >= 0
	})
	return result

}

// Put add the elements of e at the key.
func (t *MultiTreeTable[K, T]) Put(key K, e ...T) {

	for _, i := range e {

		t.objects.Add(NewEntry(key, i))

	}

}

// PutSlice adds the elements of e at the table.
// It panics if key and e have different lengths.
func (t *MultiTreeTable[K, T]) PutSlice(key []K, e []T) {

	if len(key) != len(e) {

		panic("Different lengths for keys and elements")

	}
	for i := 0; i != len(key); i++ {

		t.Put(key[i], e[i])

	}

}

// Replace replace all elements associated at the key with e and returns the slice of overwritten values.
func (t *MultiTreeTable[K, T]) Replace(key K, e ...T) []T {

	result := t.RemoveKey(key)
	t.Put(key, e...)
	return result

}

// ReplaceSlice replace all elements associated at the key with e and returns the slice of overwritten values.
func (t *MultiTreeTable[K, T]) ReplaceSlice(key []K, e []T) []T {

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
func (t *MultiTreeTable[K, T]) Remove(key K, e T) bool {

	fun := util.EqualFunction(e)
	found := false
	return !t.objects.None(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) bool {
		if key.Compare(i.Element().Key()) == 0 && fun(i.Element().Element()) {
			t.objects.RemoveFunc(i.Element(), func(i *Entry[K, T], other *tree.Node[*Entry[K, T]]) bool {
				return i.Key().Compare(other.Element().Key()) == 0 && util.EqualFunction(i.Element())(other.Element().Element())
			})
			found = true
		}
		return found
	})

}

// RemoveKey remove all elements associated at the key and returns the slice of removed values.
func (t *MultiTreeTable[K, T]) RemoveKey(key K) []T {

	result := make([]T, 0)
	for i := t.objects.Iter(); !i.End(); i = i.Next() {

		check := key.Compare(i.Element().Key())
		for !i.End() && check == 0 {
			result = append(result, i.Element().Element())
			i = i.Remove()
			check = key.Compare(i.Element().Key())
		}
		if check == -1 {

			return result

		}

	}
	return result

}

// Each executes fun for all elements of t.
//
// This method should be used to remove elements. Use Iter insted.
func (t *MultiTreeTable[K, T]) Each(fun func(key K, element T)) {

	t.objects.Each(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) {
		fun(i.Element().Key(), i.Element().Element())
	})

}

// Stream returns a [MultiStream] rapresenting t.
func (t *MultiTreeTable[K, T]) Stream() *Stream[K, T] {

	return NewStream[K, T](t, reflect.ValueOf(NewMultiTreeTable[K, T]))

}

// Clear removes all element from t.
func (t *MultiTreeTable[K, T]) Clear() {

	t.objects.Clear()

}

// Iter returns an [Iterator] which permits to iterate a [MultiTreeTable].
//
//	for i := t.Iter(); !i.End(); i = i.Next() {
//		key := i.Key()
//		element := i.Element()
//		// Code
//	}
func (t *MultiTreeTable[K, T]) Iter() Iterator[K, T] {

	return NewMultiTreeTableIterator(t)

}

// Equal returns true if t and st are both multitables and their keys and elements are equals.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is an [MultiHashTable],
// but the elements of t and the elements of st are equals, this method returns anyway true.
func (t *MultiTreeTable[K, T]) Equal(st any) bool {

	table, ok := st.(MultiTable[K, T])
	if ok && t != nil && table != nil {

		if t.Len() != table.Len() {

			return false

		}
		return t.objects.All(t.objects.Root(), func(i *tree.Node[*Entry[K, T]]) bool {
			return table.Contains(i.Element().Key(), i.Element().Element())
		})

	}
	return false

}

// Compare returns 0 if t and st have the same length,
// -1 if t is shorten than st,
// 1 if t is longer than st,
// -2 if st is not a [MultiTable] or if one between t and st is nil.
func (t *MultiTreeTable[K, T]) Compare(st any) int {

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
func (t *MultiTreeTable[K, T]) Hash() string {

	check := []string{reflect.TypeOf(new(K)).String(), reflect.TypeOf(new(T)).String()}
	return fmt.Sprintf("%v%v%v", check[0][1:], check[1][1:], t.Len())

}

// Copy returns a multitable containing a copy of the elements of t.
// The result of this method is of type [MultiTable], but the effective table which is created is an [MultiTreeTable].
//
// This method uses [util.Copy] to make copies of the elements.
func (t *MultiTreeTable[K, T]) Copy() MultiTable[K, T] {

	slice := t.objects.ToSlice()
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	result := NewMultiTreeTable[K, T]()
	for _, i := range slice {

		result.Put(i.Key(), util.Copy(i.Element()))

	}
	return result

}

// String returns a rapresentation of t in the form of a string.
func (t *MultiTreeTable[K, T]) String() string {

	check := []string{reflect.TypeOf(new(K)).String(), reflect.TypeOf(new(T)).String()}
	result := fmt.Sprintf("MultiTreeTable[%v, %v][", check[0][1:], check[1][1:])
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
