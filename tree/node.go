package tree

// Node is a component of a tree structure.
type Node[T any] struct {
	// contains filtered or unexported fields
	element T
	parent  *Node[T]
	left    *Node[T]
	right   *Node[T]
}

// NewNode returns a new [Node].
func NewNode[T any](element T, parent *Node[T], left *Node[T], right *Node[T]) *Node[T] {

	return &Node[T]{element: element, parent: parent, left: left, right: right}

}

// Element returns the element of n.
func (n *Node[T]) Element() T {

	return n.element

}

// Element sets the element of n.
func (n *Node[T]) SetElement(element T) {

	n.element = element

}

// Parent returns the parent [Node] of n.
func (n *Node[T]) Parent() *Node[T] {

	return n.parent

}

// SetParent sets the the parent [Node] of n.
func (n *Node[T]) SetParent(parent *Node[T]) {

	n.parent = parent

}

// Left returns the left [Node] of n.
func (n *Node[T]) Left() *Node[T] {

	return n.left

}

// SetLeft sets the the left [Node] of n.
func (n *Node[T]) SetLeft(left *Node[T]) {

	n.left = left

}

// Right returns the right [Node] of n.
func (n *Node[T]) Right() *Node[T] {

	return n.right

}

// SetRight sets the the right [Node] of n.
func (n *Node[T]) SetRight(right *Node[T]) {

	n.right = right

}

// Max returns the node with the max element in the subtree
// with node as root.
func (n *Node[T]) Max() *Node[T] {

	if n.right == nil {

		return n

	}
	return n.Right().Max()

}

// Min returns the node with the min element in the subtree
// with node as root.
func (n *Node[T]) Min() *Node[T] {

	if n.left == nil {

		return n

	}
	return n.Left().Min()

}
