package btree

import "testing"

func TestBTreeSearch(t *testing.T) {
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
	keys := [2]int{
		2, 5,
	}

	tests := []struct {
		name     string
		value    int
		expected int
	}{
		{"1", 1, 0},
		{"3", 3, 1},
		{"6", 6, 2},
		// {"8", 8, 3},
		// {"12", 12, 4},
		// {"17", 17, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := determineChild(keys, tt.value)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}

}

func TestInsert(t *testing.T) {
	btree := buildEmptyBTree()

	insert(btree, 1)

}
