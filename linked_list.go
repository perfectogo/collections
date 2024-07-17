package collections

import "errors"

// Node represents a single node in the linked list
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList represents the linked list
type LinkedList[T any] struct {
	HeadNode *Node[T]
	length   int
}

// NewNode creates a new node
func NewNode[T any](value T) *Node[T] {
	return &Node[T]{Value: value}
}

// NewLinkedList creates a new empty linked list
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// func (l *LinkedList[T]) Append(value T) {

// 	var newNode = NewNode(value)

// 	for node := l.HeadNode; node != nil; node = node.Next {

// 	}
// }

func (l *LinkedList[T]) PreAppend() {

}

// GetLength returns the length of the linked list
func (l *LinkedList[T]) GetLength() int {
	if l != nil {
		return l.length
	}
	return 0
}

func (ll *LinkedList[T]) Print() error {
	if ll == nil {
		return errors.New("the linked list is nil")
	}

	for node := ll.HeadNode; node != nil; node = node.Next {
		println(node.Value)
	}

	return nil
}
