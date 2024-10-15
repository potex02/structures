package set

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/tree"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ structures.Structure[wrapper.Int] = NewMultiTreeSet[wrapper.Int]()
var _ BaseSet[wrapper.Int] = NewMultiTreeSet[wrapper.Int]()
var _ MultiSet[wrapper.Int] = NewMultiTreeSet[wrapper.Int]()

// MultiTreeSet provides a generic set ith duplicate elements implemented through a [tree.BinaryTree].
// It maintains the order of the elements.
//
// It implements the interface [MultiSet].
type MultiTreeSet[T util.Comparer] struct {
	// contains filtered or unexported fields
	objects *tree.BinaryTree[T]
}

// NewMultiTreeSet returns a new [MultiTreeSet] containing the elements c.
//
// if no argument is passed, it will be created an empty [MultiTreeSet].
func NewMultiTreeSet[T util.Comparer](c ...T) *MultiTreeSet[T] {
	return NewMultiTreeSetFromSlice(c)
}

// NewMultiTreeSetFromSlice returns a new [MultiTreeSet] containing the elements of slice c
func NewMultiTreeSetFromSlice[T util.Comparer](c []T) *MultiTreeSet[T] {
	set := &MultiTreeSet[T]{objects: tree.NewBinaryTree[T]()}
	if len(c) != 0 {
		set.AddSlice(c)
	}
	return set
}

// Len returns the length of s.
func (s *MultiTreeSet[T]) Len() int {
	return s.objects.Len()
}

// IsEmpty returns a bool which indicates if s is empty or not.
func (s *MultiTreeSet[T]) IsEmpty() bool {
	return s.objects.IsEmpty()
}

// Contains returns if e is present in s.
func (s *MultiTreeSet[T]) Contains(e T) bool {
	return s.objects.Contains(e)
}

// ToSlice returns a slice which contains all elements of s.
func (s *MultiTreeSet[T]) ToSlice() []T {
	return s.objects.ToSlice()
}

// Add adds the elements e at s.
func (s *MultiTreeSet[T]) Add(e ...T) {
	s.AddSlice(e)
}

// AddSlice adds the elements of e at s.
func (s *MultiTreeSet[T]) AddSlice(e []T) {
	for _, i := range e {
		s.objects.Add(i)
	}
}

// Remove removes the element e from s if it is present.
// In that case, the method returns true, otherwhise it returns false.
func (s *MultiTreeSet[T]) Remove(e T) bool {
	return s.objects.Remove(e)
}

// RemoveAll removes all occurrences of e from s.
func (s *MultiTreeSet[T]) RemoveAll(e T) {
	for i := s.objects.Iter(); !i.End(); i = i.Next() {
		check := e.Compare(i.Element())
		for !i.End() && check == 0 {
			i = i.Remove()
			check = e.Compare(i.Element())
		}
		if check == -1 {
			return
		}
	}
}

// Each executes fun for all elements of s.
//
// This method should be used to remove elements. Use Iter insted.
func (s *MultiTreeSet[T]) Each(fun func(element T)) {
	s.objects.Each(s.objects.Root(), func(i *tree.Node[T]) {
		fun(i.Element())
	})
}

// Count returns the number of occurrences of e in s.
func (s *MultiTreeSet[T]) Count(e T) int {
	result := 0
	s.objects.All(s.objects.Root(), func(i *tree.Node[T]) bool {
		check := e.Compare(i.Element())
		if check == 0 {
			result++
		}
		return check >= 0
	})
	return result
}

// ToSet returns a [TreeSet] containing the elements of the s.
func (s *MultiTreeSet[T]) ToSet() Set[T] {
	slice := s.ToSlice()
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return NewTreeSetFromSlice(slice)
}

// Stream returns a [Stream] rapresenting s.
func (s *MultiTreeSet[T]) Stream() *Stream[T] {
	return NewStream[T](s, reflect.ValueOf(NewMultiTreeSet[T]))
}

// Clear removes all element from s.
func (s *MultiTreeSet[T]) Clear() {
	s.objects.Clear()
}

// Iter returns an [Iterator] which permits to iterate a [MultiHashSet].
//
//	for i := s.Iter(); !i.End(); i = i.Next() {
//		element := i.Element()
//		// Code
//	}
func (s *MultiTreeSet[T]) Iter() Iterator[T] {
	return NewMultiTreeSetIterator(s)
}

// Equal returns true if s and st are both multisets and have the same lengtha nd contains the same elements.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st. This means that if st is a [MultiHashSet],
// but the elements of s and the elements of st are equals, this method returns anyway true.
func (s *MultiTreeSet[T]) Equal(st any) bool {
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
func (s *MultiTreeSet[T]) Compare(st any) int {
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
func (s *MultiTreeSet[T]) Hash() string {
	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("%v%v", check[1:], s.Len())
}

// Copy returns a set containing a copy of the elements of s.
// The result of this method is of type [Set], but the effective table which is created is an [MultiTreeSet].
//
// This method uses [util.Copy] to make copies of the elements.
func (s *MultiTreeSet[T]) Copy() MultiSet[T] {
	slice := s.ToSlice()
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	result := NewMultiTreeSet[T]()
	for _, i := range slice {
		result.Add(util.Copy(i))
	}
	return result
}

// String returns a rapresentation of s in the form of a string.
func (s *MultiTreeSet[T]) String() string {
	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("MultiTreeSet[%v]%v", check[1:], s.ToSlice())
}
