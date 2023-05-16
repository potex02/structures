### Note: this module is still under heavy development.
### Future releases may completely break compatibility with previous ones.
# Structures
An implementation of the data structures in Go using generics types.<br/>
The module provides the Structure interface which is implemented by all the the defined data structures:
```go
type Structure[T any] interface {
	// Len returns the numbers of elements in the structure.
	Len() int
	// IsEmpty returns a bool which indicate if the structure is empty or not.
	IsEmpty() bool
	// ToSLice returns a slice which contains all elements of the structure.
	ToSlice() []T
	// Clear removes all element from the structure.
	Clear()
	// Equals returns true if the structure and st are the same type of structure and their elements are equals.
	// In any other case, it returns false.
	Equals(s Structure[T]) bool
	// String returns a rapresentation of the structure in the form of a string.
	String() string
}
```
The module is aviable through the go get command:
```
go get github.com/potex02/structures/list
```
## Aviable structures
For now, the only aviable structures are the:
- List: an interface implemented by ArrayList and LinkedList;
- Stack;
- Queue.
The ArrayList, Stack and Queue are implemented through a slice,<br/>
while the LinkedList is a double linked list with a pointer to the root and one to the tail.