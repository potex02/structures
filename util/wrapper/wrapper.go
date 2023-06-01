// package wrapper provides a simply way to implement the [util] interfaces for any type.
package wrapper

import "github.com/potex02/structures/util"

// Wrapper is an interface which implements the equality, comparison and hashing operations for type T.
//
// An implementation of Wrapper for primitive types is already defined.
type Wrapper[T any] interface {
	util.Equaler
	util.Hasher
	// ToValue returns the wrapped value.
	ToValue() T
}
