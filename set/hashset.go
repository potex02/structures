package set

import (
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/table"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

const obj uint8 = 0

var _ structures.Structure[wrapper.Int] = NewHashSet[wrapper.Int]()
var _ Set[wrapper.Int] = NewHashSet[wrapper.Int]()

// HashSet provides a generic set implemented through a [table.HashTable].
//
// It implements the interface [Set].
type HashSet[T util.Hasher] struct {
	// contains filtered or unexported fields
	objects table.Table[T, uint8]
}

// NewHashSet returns a new [HashSet] containing the elements c.
//
// if no argument is passed, it will be created an empty [HashSet].
func NewHashSet[T util.Hasher](c ...T) *HashSet[T] {

	return NewHashSetFromSlice(c)

}

// NewHashSetFromSlice returns a new [HashSet] containing the elements of slice c
func NewHashSetFromSlice[T util.Hasher](c []T) *HashSet[T] {

	set := &HashSet[T]{objects: table.NewHashTable[T, uint8]()}
	set.AddSlice(c)
	return set

}

// Len returns the length of s.
func (s *HashSet[T]) Len() int {

	return s.objects.Len()

}

// IsEmpty returns a bool which indicates if s is empty or not.
func (s *HashSet[T]) IsEmpty() bool {

	return s.objects.IsEmpty()

}

// Contains returns if e is present in s.
func (s *HashSet[T]) Contains(e T) bool {

	return s.objects.ContainsKey(e)

}

// ToSlice returns a slice which contains all elements of s.
func (s *HashSet[T]) ToSlice() []T {

	return s.objects.Keys().ToSlice()

}

// Add adds the elements e at s.
func (s *HashSet[T]) Add(e ...T) {

	s.AddSlice(e)

}

// AddSlice adds the elements of e at s.
func (s *HashSet[T]) AddSlice(e []T) {

	for _, i := range e {

		s.objects.Put(i, obj)

	}

}

// Remove removes the element e from s if it is present.
// In that case, the method returns true, otherwhise it returns false.
func (s *HashSet[T]) Remove(e T) bool {

	_, ok := s.objects.Remove(e)
	return ok

}

// Each executes fun for all elements of s.
//
// This method should be used to remove elements. Use Iter insted.
func (s *HashSet[T]) Each(fun func(element T)) {

	for i := s.objects.Iter(); !i.End(); i = i.Next() {

		fun(i.Key())

	}

}

// Stream returns a [Stream] rapresenting s.
func (s *HashSet[T]) Stream() *Stream[T] {

	return NewStream[T](s, reflect.ValueOf(NewHashSet[T]))

}

// Clear removes all element from s.
func (s *HashSet[T]) Clear() {

	s.objects.Clear()

}

// Iter returns an [Iterator] which permits to iterate a [HashSet].
//
//	for i := s.Iter(); !i.End(); i = i.Next() {
//		element := i.Element()
//		// Code
//	}
func (s *HashSet[T]) Iter() Iterator[T] {

	return NewHashSetIterator(s)

}

// Equal returns true if s and st are both sets and have the same lengtha nd contains the same elements.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st.
func (s *HashSet[T]) Equal(st any) bool {

	set, ok := st.(Set[T])
	if ok && s != nil && set != nil {

		if s.Len() != set.Len() {

			return false

		}
		for i := s.objects.Iter(); !i.End(); i = i.Next() {

			if !set.Contains(i.Key()) {

				return false

			}

		}
		return true

	}
	return false

}

// Compare returns 0 if s and st have the same length,
// -1 if s is shorten than st,
// 1 if s is longer than st,
// -2 if st is not a [Set] or if one between s and st is nil.
func (s *HashSet[T]) Compare(st any) int {

	set, ok := st.(Set[T])
	if ok && s != nil && set != nil {

		if s.Len() < set.Len() {

			return -1

		}
		if s.Len() > set.Len() {

			return 1

		}
		return 0

	}
	return -2

}

// Hash returns the hash code of s.
func (s *HashSet[T]) Hash() string {

	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("%v%v", check[1:], s.Len())

}

// Copy returns a set containing a copy of the elements of s.
// The result of this method is of type [Set], but the effective table which is created is an [HashSet].
func (s *HashSet[T]) Copy() Set[T] {

	return NewHashSetFromSlice(s.ToSlice())

}

// String returns a rapresentation of s in the form of a string.
func (s *HashSet[T]) String() string {

	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("HashSet[%v]%v", check[1:], s.ToSlice())

}
