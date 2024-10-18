[![GoDoc](https://godoc.org/github.com/potex02/structures?status.svg)](https://godoc.org/github.com/potex02/structures)
# Structures
An implementation of the data structures in Go using generics types.<br/>
The module provides the Structure interface which is implemented by all the the defined data structures:
```go
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
```
The module is available through the go get command:
```
go get github.com/potex02/structures
```
## Available structures
For now, the only available structures are the:
- Lists:
	- ArrayList;
	- LinkedList (double linked list with a pointer to the root and one to the tail);
- Stack;
- Queues:
	- Queue;
	- PriorityQueue.
	- DoubleQueue;
- Tables:
	- HashTable;
	- TreeTable;
- MultiTables:
	- MultiHashTable;
	- MultiTreeTable;
- Sets:
	- HashSet;
	- TreeSet;
- MultiSets:
	- MultiHashSet;
	- MultiTreeSet;
- Trees:
	- BinaryTree;
	- N-aryTtree.