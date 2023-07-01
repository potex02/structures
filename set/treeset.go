package set

import (
	"fmt"
	"reflect"

	"github.com/potex02/structures"
	"github.com/potex02/structures/tree"
	"github.com/potex02/structures/util"
	"github.com/potex02/structures/util/wrapper"
)

var _ structures.Structure[wrapper.Int] = NewHashSet[wrapper.Int]()
var _ Set[wrapper.Int] = NewTreeSet[wrapper.Int]()

// TreeSet provides a generic set implemented through a [tree.BinaryTree].
// It maintains the order of the elements.
//
// It implements the interface [Set].
type TreeSet[T util.Comparer] struct {
	// contains filtered or unexported fields
	objects *tree.BinaryTree[T]
}

// NewTreeSet returns a new [TreeSet] containing the elements c.
//
// if no argument is passed, it will be created an empty [TreeSet].
func NewTreeSet[T util.Comparer](c ...T) *TreeSet[T] {

	return NewTreeSetFromSlice(c)

}

// NewTreeSetFromSlice returns a new [TreeSet] containing the elements of slice c
func NewTreeSetFromSlice[T util.Comparer](c []T) *TreeSet[T] {

	set := &TreeSet[T]{objects: tree.NewBinaryTree[T]()}
	set.AddSlice(c)
	return set

}

// Len returns the length of s.
func (s *TreeSet[T]) Len() int {

	return s.objects.Len()

}

// IsEmpty returns a bool which indicates if s is empty or not.
func (s *TreeSet[T]) IsEmpty() bool {

	return s.objects.IsEmpty()

}

// Contains returns if e is present in s.
func (s *TreeSet[T]) Contains(e T) bool {

	return s.objects.Contains(e)

}

// ToSLice returns a slice which contains all elements of s.
func (s *TreeSet[T]) ToSlice() []T {

	return s.objects.ToSlice()

}

// Add adds the elements e at s.
func (s *TreeSet[T]) Add(e ...T) {

	s.AddSlice(e)

}

// AddSlice adds the elements of e at s.
func (s *TreeSet[T]) AddSlice(e []T) {

	for _, i := range e {

		found := s.objects.Any(s.objects.Root(), func(j *tree.Node[T]) bool { return i.Compare(j.Element()) == 0 })
		if !found {
			s.objects.Add(i)
		}

	}

}

// Remove removes the element e from s if it is present.
// In that case, the method returns true, otherwhise it returns false.
func (s *TreeSet[T]) Remove(e T) bool {

	ok := s.objects.Remove(e)
	return ok

}

// Each executes fun for all elements of s.
func (s *TreeSet[T]) Each(fun func(element T)) {

	s.objects.Each(s.objects.Root(), func(i *tree.Node[T]) {
		fun(i.Element())
	})

}

// Stream returns a [Stream] rapresenting s.
func (s *TreeSet[T]) Stream() *Stream[T] {

	return NewStream[T](s, reflect.ValueOf(NewTreeSet[T]))

}

// Clear removes all element from s.
func (s *TreeSet[T]) Clear() {

	s.objects.Clear()

}

// Iter returns an [Iterator] which permits to iterate a [TreeSet].
//
//	for i := s.Iter(); !i.End(); i = i.Next() {
//		element := i.Element()
//		// Code
//	}
func (s *TreeSet[T]) Iter() Iterator[T] {

	return NewTreeSetIterator(s)

}

// Equal returns true if s and st are both sets and have the same length and contains the same elements.
// In any other case, it returns false.
//
// Equal does not take into account the effective type of st.
func (s *TreeSet[T]) Equal(st any) bool {

	set, ok := st.(Set[T])
	if ok && s != nil && set != nil {

		if s.Len() != set.Len() {

			return false

		}
		for i := s.objects.Iter(); !i.End(); i = i.Next() {

			if !set.Contains(i.Element()) {

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
func (s *TreeSet[T]) Compare(st any) int {

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
func (s *TreeSet[T]) Hash() string {

	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("%v%v", check[1:], s.Len())

}

// Copy returns a set containing a copy of the elements of s.
// The result of this method is of type [Set], but the effective table which is created is an [TreeSet].
func (s *TreeSet[T]) Copy() Set[T] {

	return NewTreeSetFromSlice(s.ToSlice())

}

// String returns a rapresentation of s in the form of a string.
func (s *TreeSet[T]) String() string {

	check := reflect.TypeOf(new(T)).String()
	return fmt.Sprintf("TreeSet[%v]%v", check[1:], s.ToSlice())

}
