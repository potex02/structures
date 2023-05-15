### Note: this module is still under heavy development.
### Future releases may completely break compatibility with previous ones.
# Structures
An implementation of the data structures in Go using generics types.<br/>
The module provides the Structure interface wich is implemented by all the the defined data structures. 
## Aviable structures
For now, the only aviable strctures are the:
- List: an interface implemented by ArrayList and LinkedList;
- Stack;
- Queue.
The ArrayList, Stack and Queue are implemented through a slice,<br/>
while the LinkedList is a double linked list with a pointer to the root and one to the tail.