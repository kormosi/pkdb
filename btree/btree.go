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

const K = 3 // Maximum number of potential search keys for each node in a B-tree

// TODO do samostatného súboru to dať
type Node struct {
	keys     []int   // length of this cannot exceed K; (it's called keys, but it's also values for now)
	children []*Node // length
	// leaf bool	 // possibly needed in future?
}

func (node Node) hasFreeRoom() bool {
	return node.hasValue(-1)
}

func (node Node) hasValue(val int) bool {
	return slices.Contains(node.keys, val)
}

func (node Node) determineChild(val int) int {
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

func buildEmptyBTree(K int) Node {
	children := make([]*Node, K)
	keys := make([]int, K)
	for i := range keys {
		keys[i] = -1
	}
	root := Node{keys: keys, children: children}
	fmt.Println(children)
	return root
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
			childToSearch := node.determineChild(val)
			return isInBTree(*node.children[childToSearch], val)
		}
	}
	return false
}

func insert(node Node, val int) {
	// All insertions start at a leaf node.

	// To insert a new element, search the tree to find the leaf node where the new element should be added.
	if node.hasFreeRoom() {
		for idx, el := range node.keys {
			if el == -1 {
				node.keys[idx] = val
			}
		}

		// Then sort

	}

	// Insert the new element into that node with the following steps:

	// If the node contains fewer than the maximum allowed number of elements, then there is room for the new element. Insert the new element in the node, keeping the node's elements ordered.
}

// Pri insertoch budem musieť enforcovať veľkosť, ale to spravím potom na maine
// mergnúť túto branchu do main

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

	buildEmptyBTree(3)
}

// napísať testy - zistiť, ako sa to robí.
// žeby samostatný pkdb -> test adresár?
