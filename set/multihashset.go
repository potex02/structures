package set

import (
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/table"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ structures.Structure[wrapper.Int] = NewMultiHashSet[wrapper.Int]()
var _ BaseSet[wrapper.Int] = NewMultiHashSet[wrapper.Int]()
var _ MultiSet[wrapper.Int] = NewMultiHashSet[wrapper.Int]()

// MultiHashSet provides a generic set ith duplicate elements implemented through a [table.MultiHashTable].
//
// It implements the interface [MultiSet].
type MultiHashSet[T util.Hasher] struct {
	// contains filtered or unexported fields
	objects table.MultiTable[T, uint8]
}

// NewMultiHashSet returns a new [MultiHashSet] containing the elements c.
//
// if no argument is passed, it will be created an empty [MultiHashSet].
func NewMultiHashSet[T util.Hasher](c ...T) *MultiHashSet[T] {
	return NewMultiHashSetFromSlice(c)
}

// NewMultiHashSetFromSlice returns a new [MultiHashSet] containing the elements of slice c
func NewMultiHashSetFromSlice[T util.Hasher](c []T) *MultiHashSet[T] {
	set := &MultiHashSet[T]{objects: table.NewMultiHashTable[T, uint8]()}
	if len(c) != 0 {
		set.AddSlice(c)
	}
	return set
}

// Len returns the length of s.
func (s *MultiHashSet[T]) Len() int {
	return s.objects.Len()
}

// IsEmpty returns a bool which indicates if s is empty or not.
func (s *MultiHashSet[T]) IsEmpty() bool {
	return s.objects.IsEmpty()
}

// Contains returns if e is present in s.
func (s *MultiHashSet[T]) Contains(e T) bool {
	return s.objects.ContainsKey(e)
}

// ToSlice returns a slice which contains all elements of s.
func (s *MultiHashSet[T]) ToSlice() []T {
	return s.objects.Keys().ToSlice()
}

// Add adds the elements e at s.
func (s *MultiHashSet[T]) Add(e ...T) {
	s.AddSlice(e)
}

// AddSlice adds the elements of e at s.
func (s *MultiHashSet[T]) AddSlice(e []T) {
	for _, i := range e {
		s.objects.Put(i, obj)
	}
}

// Remove removes the element e from s if it is present.
// In that case, the method returns true, otherwhise it returns false.
func (s *MultiHashSet[T]) Remove(e T) bool {
	return s.objects.Remove(e, obj)
}

// RemoveAll removes all occurrences of e from s.
func (s *MultiHashSet[T]) RemoveAll(e T) {
	for i := s.objects.Iter(); !i.End(); i = i.Next() {
		for !i.End() && e.Compare(i.Key()) == 0 {
			i = i.Remove()
		}
	}
}

// Each executes fun for all elements of s.
//
// This method should be used to remove elements. Use Iter insted.
func (s *MultiHashSet[T]) Each(fun func(element T)) {
	s.objects.Each(func(key T, _ uint8) {
		fun(key)
	})
}

// Count returns the number of occurrences of e in s.
func (s *MultiHashSet[T]) Count(e T) int {
	result := 0
	s.objects.Each(func(key T, _ uint8) {
		if e.Compare(key) == 0 {
			result++
		}
	})
	return result
}

// ToSet returns a [HashSet] containing the elements of the s.
func (s *MultiHashSet[T]) ToSet() Set[T] {
	return NewHashSetFromSlice(s.ToSlice())
}

// Stream returns a [Stream] rapresenting s.
func (s *MultiHashSet[T]) Stream() *Stream[T] {
	return NewStream[T](s, reflect.ValueOf(NewMultiHashSet[T]))
}

// Clear removes all element from s.
func (s *MultiHashSet[T]) Clear() {
	s.objects.Clear()
}

// Iter returns an [Iterator] which permits to iterate a [MultiHashSet].
//
//	for i := s.Iter(); !i.End(); i = i.Next() {
//		element := i.Element()
//		// Code
//	}
func (s *MultiHashSet[T]) Iter() Iterator[T] {
	return NewMultiHashSetIterator(s)
}

// Equal returns true if s and st are both multisets and have the same lengtha nd contains the same elements.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is a [MultiTreeSet],
// but the elements of s and the elements of st are equals, this method returns anyway true.
func (s *MultiHashSet[T]) Equal(st any) bool {
	set, ok := st.(MultiSet[T])
	if ok && s != nil && set != nil {
		if s.Len() != set.Len() {
			return false
		}
		for i := s.ToSet().Iter(); !i.End(); i = i.Next() {
			if s.Count(i.Element()) != set.Count(i.Element()) {
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
// -2 if st is not a [MultiSet] or if one between s and st is nil.
func (s *MultiHashSet[T]) Compare(st any) int {
	set, ok := st.(MultiSet[T])
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
func (s *MultiHashSet[T]) Hash() string {
	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("%v%v", check[1:], s.Len())
}

// Copy returns a set containing a copy of the elements of s.
// The result of this method is of type [Set], but the effective table which is created is an [MultiHashSet].
//
// This method uses [util.Copy] to make copies of the elements.
func (s *MultiHashSet[T]) Copy() MultiSet[T] {
	result := NewMultiHashSet[T]()
	s.Each(func(element T) {
		result.Add(util.Copy(element))
	})
	return result
}

// String returns a rapresentation of s in the form of a string.
func (s *MultiHashSet[T]) String() string {
	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("MultiHashSet[%v]%v", check[1:], s.ToSlice())
}
