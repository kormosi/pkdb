// Tidbits from wiki and gipiti:

// Regarding BTrees:

// Usually, the number of keys is chosen to vary between d and 2d,
// where d is the minimum number of keys, and d+1 is the minimum
// branching factor of the tree.

// Great link that shows the node structure (with data records)
// https://en.wikipedia.org/wiki/B-tree#Node_structure
// I will need to re-work my implementation according to this.

// Regarding go modules and packages

// If it’s meant to be a library (not executable),
//	don’t use go run; instead, import it from a main package.

package btree

import (
	"fmt"
	"slices"
)

// Keep this constant odd (see the `insert` function)
const K = 3 // Maximum number of potential search keys for each node in a B-tree

// TODO do samostatného súboru to dať
type Node struct {
	keys     []int // length of this cannot exceed K; (it's called keys, but it's also values for now)
	children []*Node
	parent   *Node
	// leaf bool	 // possibly needed in future?
}

func (node Node) hasFreeRoom() bool {
	return len(node.keys) < K-1
}

func (node Node) hasValue(val int) bool {
	return slices.Contains(node.keys, val)
}

func (node Node) determineChildIndex(val int) int {
	for idx, key := range node.keys {
		if key == -1 {
			return idx // maybe return -1 in this case and thus end the search?
		}
		if val < key {
			return idx
		}
	}
	return len(node.keys)
}

func (node Node) hasValidChildren() bool {
	for _, child := range node.children {
		if child != nil {
			return true
		}
	}
	return false
}

func (node *Node) insert(val int) *Node {
	// All insertions start at a leaf node.
	// TODO: the above is important. We should never start at the root node.
	// We need an algorithm for traversing the tree and finding the value.

	// To insert a new element, search the tree to find the leaf node where the new element should be added.

	// If the node contains fewer than the maximum allowed number of elements, then there is room for the new element.
	// Insert the new element in the node, keeping the node's elements ordered.
	if node.hasFreeRoom() {
		node.keys = append(node.keys, val)
		slices.Sort(node.keys)
		return node
	} else {
		// Otherwise the node is full, evenly split it into two nodes so:
		// 1: A single median is chosen from among the leaf's elements and the new element that is being inserted.
		// Patrik's notes: choosing a single median will be easy if there's 2 keys + 1 new key (3 total, so just choose index 1, that is, the second element)
		// This won't work for non-odd K values, but I guess it's good for now.

		// Choose the median
		node.keys = append(node.keys, val) // opportunity for refactor,
		slices.Sort(node.keys)             // in the above `if` we do the same
		medianIndex := K / 2
		temporaryKeySlice := node.keys

		// 2: Values less than the median are put in the new left node, and values greater than the median are put in the new right node, with the median acting as a separation value.
		node.keys = temporaryKeySlice[:medianIndex] // The original node becomes the new left node
		newRightNode := Node{keys: temporaryKeySlice[medianIndex+1:], children: []*Node{}}

		// 3: The separation value is inserted in the node's parent...
		if node.parent == nil {
			newRootNode := Node{keys: []int{temporaryKeySlice[medianIndex]}, children: []*Node{node, &newRightNode}}
			return &newRootNode
		} else {
			// ... which may cause it to be split, and so on. If the node has no parent (i.e., the node was the root), create a new root above this node (increasing the height of the tree).
			// TODO not yet implemented, returning empty node pointer
			return &Node{keys: []int{}, children: []*Node{}}
		}

	}
}

func buildBTree() Node {
	// lowest level - left
	lowest_l_l := Node{keys: []int{1, -1}, children: []*Node{}}
	lowest_l_r := Node{keys: []int{3, -1}, children: []*Node{}}
	// lowest level - right
	lowest_r_l := Node{keys: []int{5, -1}, children: []*Node{}}
	lowest_r_r := Node{keys: []int{7, -1}, children: []*Node{}}

	// mid level
	mid_l := Node{keys: []int{2, -1}, children: []*Node{&lowest_l_l, &lowest_l_r}}
	mid_r := Node{keys: []int{6, -1}, children: []*Node{&lowest_r_l, &lowest_r_r}}

	// top level
	root := Node{keys: []int{4, -1}, children: []*Node{&mid_l, &mid_r}}

	return root
}

func buildEmptyBTree() *Node {
	children := []*Node{}
	keys := []int{}
	root := Node{keys: keys, children: children}
	return &root
}

func printBTree(root Node) {
	// TODO write a complementary function "traverseBtree"
	// that will parse and (optionally) print each and every node?
	fmt.Println(root)
	fmt.Println()
	fmt.Println(*root.children[0])
	fmt.Println()
	fmt.Println(*root.children[0].children[0])
	fmt.Println(*root.children[0].children[1])
	fmt.Println()
	fmt.Println(*root.children[1])
	fmt.Println()
	fmt.Println(*root.children[1].children[0])
	fmt.Println(*root.children[1].children[1])
}

func isInBTree(node Node, val int) bool {
	if node.hasValue(val) {
		return true
	} else {
		if node.hasValidChildren() {
			childToSearchIndex := node.determineChildIndex(val)
			return isInBTree(*node.children[childToSearchIndex], val)
		}
	}
	return false
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	// root := buildBTree()
	// printBTree(root)

	// fmt.Println()

	// TODO otestovať to s viacerými hodnotami
	// fmt.Println(isInBTree(root, 4))
	// fmt.Println(isInBTree(root, 2))
	// fmt.Println(isInBTree(root, 1))
	// fmt.Println(isInBTree(root, 3))
	// fmt.Println(isInBTree(root, 6))
	// fmt.Println(isInBTree(root, 5))
	// fmt.Println(isInBTree(root, 7))

	// fmt.Println(isInBTree(root, 8))

	// buildEmptyBTree()
}
