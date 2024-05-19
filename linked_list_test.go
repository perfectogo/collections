package collections

import (
	"testing"
)

// TestNewNode tests the NewNode function
func TestNewNode(t *testing.T) {
	node := NewNode(5)
	if node == nil {
		t.Fatalf("Expected non-nil node, got nil")
	}
	if node.Value != 5 {
		t.Errorf("Expected node value to be 5, got %d", node.Value)
	}
	if node.Next != nil {
		t.Errorf("Expected node next to be nil, got %v", node.Next)
	}

	nodeStr := NewNode("test")
	if nodeStr == nil {
		t.Fatalf("Expected non-nil node, got nil")
	}
	if nodeStr.Value != "test" {
		t.Errorf("Expected node value to be 'test', got %s", nodeStr.Value)
	}
	if nodeStr.Next != nil {
		t.Errorf("Expected node next to be nil, got %v", nodeStr.Next)
	}
}

// BenchmarkNewNode benchmarks the NewNode function
func BenchmarkNewNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewNode(i)
	}
}

// TestNewLinkedList tests the NewLinkedList function
func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	if list == nil {
		t.Fatalf("Expected non-nil linked list, got nil")
	}
	if list.HeadNode != nil {
		t.Errorf("Expected head node to be nil, got %v", list.HeadNode)
	}
	if list.GetLength() != 0 {
		t.Errorf("Expected length to be 0, got %d", list.GetLength())
	}

	listStr := NewLinkedList[string]()
	if listStr == nil {
		t.Fatalf("Expected non-nil linked list, got nil")
	}
	if listStr.HeadNode != nil {
		t.Errorf("Expected head node to be nil, got %v", listStr.HeadNode)
	}
	if listStr.GetLength() != 0 {
		t.Errorf("Expected length to be 0, got %d", listStr.GetLength())
	}
}
