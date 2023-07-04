// Package structures implements the most common used data structures.
package structures

import (
	"fmt"

	"github.com/potex02/structures/util"
)

// Structure defines commons methods for all data structures.
//
// A Structure is a generic that can be used with any type T.
type Structure[T any] interface {
	fmt.Stringer
	util.Equaler
	util.Hasher
	// Len returns the numbers of elements in the structure.
	Len() int
	// IsEmpty returns a bool which indicates if the structure is empty or not.
	IsEmpty() bool
	// ToSlice returns a slice which contains all elements of the structure.
	ToSlice() []T
	// Clear removes all element from the structure.
	Clear()
}
