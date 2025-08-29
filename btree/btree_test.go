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

	node := buildExampleBTree()

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
			result := node.determineChildIndex(tt.value)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}

}

func TestCreateChildParentPointers(t *testing.T) {
	node := buildExampleBTree()
	node.createChildParentPointers()
	for _, child := range node.children {
		if child.parent != &node {
			t.Errorf("got %p, want %p", child.parent, &node)
		}
	}
}

// TODO make this test more robust
func TestFindSuitableNodeForInsertion(t *testing.T) {
	// Not an exhaustive test but we'll hope it's enough for now
	// Could be made better by creating trees of different levels here as params
	btree := buildExampleBTree()
	expected := []int{7}
	nodeFound := findSuitableNodeForInsertion(&btree, 8)
	if !reflect.DeepEqual(nodeFound.keys, expected) {
		t.Errorf("got %d, want %d", nodeFound.keys, expected)
	}
}

func TestInsertWithoutNodeSplitting(t *testing.T) {
	emptyBtree := buildEmptyBTree()
	expected := []int{1, 2}

	emptyBtree.findAndInsert(2)
	emptyBtree.findAndInsert(1)

	if !reflect.DeepEqual(emptyBtree.keys, expected) {
		t.Errorf("got %d, want %d", emptyBtree.keys, expected)
	}
}

func TestInsertWithNodeSplitting(t *testing.T) {
	// TODO test-case pre každý frame obrázku tohto obrázku:
	// https://upload.wikimedia.org/wikipedia/commons/3/33/B_tree_insertion_example.png

	// akonáhle spravím test pre každú úroveň z toho obrázku
	// tak môžem urobiť pomocou insert metódy aký strom ja chcem
	// a tým pádom rýchlo otestovať TestFindSuitableNodeForInsertion - aj keď
	// je to taký chicken-egg problém, lebo insert samotný používa TestFindSuitableNodeForInsertion

	btree := buildEmptyBTree()
	btree = btree.findAndInsert(1)
	expectedKeys := []int{1}
	if !reflect.DeepEqual(btree.keys, expectedKeys) {
		t.Errorf("got %d, want %d", btree.keys, expectedKeys)
	}

	btree = btree.findAndInsert(2)
	expectedKeys = []int{1, 2}
	if !reflect.DeepEqual(btree.keys, expectedKeys) {
		t.Errorf("got %d, want %d", btree.keys, expectedKeys)
	}

	// This should trigger the splitting of the node
	btree = btree.findAndInsert(3)
	expectedRoot := []int{2}
	expectedLeftChild := []int{1}
	expectedRightChild := []int{3}
	if !reflect.DeepEqual(btree.keys, expectedRoot) {
		t.Errorf("got %d, want %d", btree.keys, expectedRoot)
	}
	if !reflect.DeepEqual(btree.children[0].keys, expectedLeftChild) {
		t.Errorf("got %d, want %d", btree.children[0].keys, expectedLeftChild)
	}
	if !reflect.DeepEqual(btree.children[1].keys, expectedRightChild) {
		t.Errorf("got %d, want %d", btree.children[1].keys, expectedRightChild)
	}

	btree = btree.findAndInsert(4)
	expectedRoot = []int{2}
	expectedLeftChild = []int{1}
	expectedRightChild = []int{3, 4}
	if !reflect.DeepEqual(btree.keys, expectedRoot) {
		t.Errorf("got %d, want %d", btree.keys, expectedRoot)
	}
	if !reflect.DeepEqual(btree.children[0].keys, expectedLeftChild) {
		t.Errorf("got %d, want %d", btree.children[0].keys, expectedLeftChild)
	}
	if !reflect.DeepEqual(btree.children[1].keys, expectedRightChild) {
		t.Errorf("got %d, want %d", btree.children[1].keys, expectedRightChild)
	}

	btree = btree.findAndInsert(5)
	expectedRoot = []int{2, 4}
	expectedLeftChild = []int{1}
	expectedCenterChild := []int{3}
	expectedRightChild = []int{5}

	if !reflect.DeepEqual(btree.keys, expectedRoot) {
		t.Errorf("got %d, want %d", btree.keys, expectedRoot)
	}
	if !reflect.DeepEqual(btree.children[0].keys, expectedLeftChild) {
		t.Errorf("got %d, want %d", btree.children[0].keys, expectedLeftChild)
	}
	if !reflect.DeepEqual(btree.children[1].keys, expectedCenterChild) {
		t.Errorf("got %d, want %d", btree.children[1].keys, expectedCenterChild)
	}
	if !reflect.DeepEqual(btree.children[2].keys, expectedRightChild) {
		t.Errorf("got %d, want %d", btree.children[2].keys, expectedRightChild)
	}

	btree = btree.findAndInsert(6)
	expectedRoot = []int{2, 4}
	expectedLeftChild = []int{1}
	expectedCenterChild = []int{3}
	expectedRightChild = []int{5, 6}

	if !reflect.DeepEqual(btree.keys, expectedRoot) {
		t.Errorf("got %d, want %d", btree.keys, expectedRoot)
	}
	if !reflect.DeepEqual(btree.children[0].keys, expectedLeftChild) {
		t.Errorf("got %d, want %d", btree.children[0].keys, expectedLeftChild)
	}
	if !reflect.DeepEqual(btree.children[1].keys, expectedCenterChild) {
		t.Errorf("got %d, want %d", btree.children[1].keys, expectedCenterChild)
	}
	if !reflect.DeepEqual(btree.children[2].keys, expectedRightChild) {
		t.Errorf("got %d, want %d", btree.children[2].keys, expectedRightChild)
	}

	btree = btree.findAndInsert(7)
	expectedRoot = []int{4}
	// expectedLeftChild = []int{1}
	// expectedCenterChild = []int{3}
	// expectedRightChild = []int{5, 6}

	if !reflect.DeepEqual(btree.keys, expectedRoot) {
		t.Errorf("got %d, want %d", btree.keys, expectedRoot)
	}
	// if !reflect.DeepEqual(btree.children[0].keys, expectedLeftChild) {
	// 	t.Errorf("got %d, want %d", btree.children[0].keys, expectedLeftChild)
	// }
	// if !reflect.DeepEqual(btree.children[1].keys, expectedCenterChild) {
	// 	t.Errorf("got %d, want %d", btree.children[1].keys, expectedCenterChild)
	// }
	// if !reflect.DeepEqual(btree.children[2].keys, expectedRightChild) {
	// 	t.Errorf("got %d, want %d", btree.children[2].keys, expectedRightChild)
	// }
}
