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
