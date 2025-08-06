package btree

import (
	"reflect"
	"testing"
)

func TestIsInBtree(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		expected bool
	}{
		{"1", 1, true},
		{"2", 2, true},
		{"3", 3, true},
		{"4", 4, true},
		{"5", 5, true},
		{"6", 6, true},
		{"7", 7, true},
		{"8", 8, false},
		{"9", 9, false},
	}

	node := buildBTree()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isInBTree(node, tt.value)
			if result != tt.expected {
				t.Errorf("got %t, want %t", result, tt.expected)
			}
		})
	}
}

func TestDetermineChild(t *testing.T) {
	node := Node{
		keys:     []int{2, 5, 7, 10, 15},
		children: []*Node{},
	}

	tests := []struct {
		name     string
		value    int
		expected int
	}{
		{"1", 1, 0},
		{"3", 3, 1},
		{"6", 6, 2},
		{"8", 8, 3},
		{"12", 12, 4},
		{"17", 17, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := node.determineChild(tt.value)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}

}

func TestInsertWithoutNodeSplitting(t *testing.T) {
	emptyBtree := buildEmptyBTree()
	expected := []int{1, 2}

	emptyBtree.insert(2)
	emptyBtree.insert(1)

	if !reflect.DeepEqual(emptyBtree.keys, expected) {
		t.Errorf("got %d, want %d", emptyBtree.keys, expected)
	}
}

func TestInsertWithNodeSplitting(t *testing.T) {
	emptyBtree := buildEmptyBTree()
	expectedRoot := []int{2}
	expectedLeftChild := []int{1}
	expectedRightChild := []int{3}

	emptyBtree.insert(2)
	emptyBtree.insert(1)
	// This should trigger the splitting of the node
	newRootNode := emptyBtree.insert(3)

	// TODO make these asserts more compact?
	// e.g. by for i,j in zip([root, lchild, rchild], [expectedRoot, elc, erc])
	if !reflect.DeepEqual(newRootNode.keys, expectedRoot) {
		t.Errorf("got %d, want %d", newRootNode.keys, expectedRoot)
	}
	if !reflect.DeepEqual(newRootNode.children[0].keys, expectedLeftChild) {
		t.Errorf("got %d, want %d", newRootNode.children[0].keys, expectedLeftChild)
	}
	if !reflect.DeepEqual(newRootNode.children[1].keys, expectedRightChild) {
		t.Errorf("got %d, want %d", newRootNode.children[1].keys, expectedRightChild)
	}
}
