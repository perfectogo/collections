package collections

// DNode represents a node in a doubly linked list
type DNode[T any] struct {
	Val  T
	Prev *DNode[T]
	Next *DNode[T]
}

// DLinkedList represents a doubly linked list
type DLinkedList[T any] struct {
	Head   *DNode[T]
	Tail   *DNode[T]
	length int
}

// NewDNode creates a new doubly linked node
func NewDNode[T any](value T) *DNode[T] {
	return &DNode[T]{Val: value}
}

// NewDLinkedList creates a new empty doubly linked list
func NewDLinkedList[T any]() *DLinkedList[T] {
	return &DLinkedList[T]{}
}

// GetLength returns the length of the doubly linked list
func (l *DLinkedList[T]) GetLength() int {
	return l.length
}
