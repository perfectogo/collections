package collections

import (
	"errors"
	"fmt"
	"strings"
)

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

// GetLength returns the length of the linked list
func (l *LinkedList[T]) GetLength() int {
	if l != nil {
		return l.length
	}
	return 0
}

// ...
func (ll *LinkedList[T]) Print() error {

	if ll == nil {
		return errors.New("the linked list is nil")
	}

	var builder strings.Builder
	builder.WriteString("[")

	for node := ll.HeadNode; node != nil; node = node.Next {
		builder.WriteString(fmt.Sprintf("%v", node.Value))
		if node.Next != nil {
			builder.WriteString(" ")
		}
	}

	builder.WriteString("]")

	fmt.Println(builder.String())
	return nil
}

// this method for ....
func (ll *LinkedList[T]) Append(value T) error {
	var newNode = &Node[T]{
		Value: value,
		Next:  nil,
	}

	if ll.HeadNode == nil {
		ll.HeadNode = newNode
	} else {
		currentNode := ll.HeadNode
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}

	ll.length++
	return nil
}

// this method for ....
func (ll *LinkedList[T]) PreAppend(value T) error {
	var newNode = &Node[T]{
		Value: value,
	}

	if ll.HeadNode == nil {
		ll.HeadNode = newNode
	} else {
		newNode.Next = ll.HeadNode
		ll.HeadNode = newNode
	}

	ll.length++
	return nil
}

func (ll *LinkedList[T]) Insert(position int, value T) error {

	if position < 1 || position > ll.length+1 {
		return fmt.Errorf("insert: Index out of bounds")
	}

	newNode := &Node[T]{Value: value}
	var prevNode, currentNode = &Node[T]{}, ll.HeadNode

	for position > 1 {
		prevNode = currentNode
		currentNode = currentNode.Next
		position = position - 1
	}

	if prevNode != nil {
		prevNode.Next = newNode
		newNode.Next = currentNode
	} else {
		newNode.Next = currentNode
		ll.HeadNode = newNode
	}

	ll.length++

	return nil
}
