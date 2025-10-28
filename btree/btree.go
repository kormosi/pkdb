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

type Node struct {
	keys     []int // length of this cannot exceed K; (it's called keys, but it's also values for now)
	children []*Node
	parent   *Node
	// leaf bool	 // possibly needed in future?
}

// TODO tuto pokračovať, spraviť z toho clean-code novinový kód
func (root *Node) Insert(val int) *Node {
	// To insert a new element, search the tree to find the leaf node where the new element should be added.
	node := findNodeSuitableForInsertion(root, val)
	return node.insert(val)
}

// TODO what if value already is in the Btree?
// Should we concern ourselves with such a possiblity?
// Maybe it won't happen when used in DB, because of hashing.
// Then again, collisions can happen, so maybe it should be handled.
func findNodeSuitableForInsertion(node *Node, val int) *Node {
	if node.hasChildren() {
		childToSearchIndex := node.determineChildIndex(val)
		return findNodeSuitableForInsertion(node.children[childToSearchIndex], val)
	} else {
		return node
	}
}

func (node *Node) insert(val int) *Node {
	if node.hasFreeRoom() {
		node.keys = appendAndSort(node.keys, val)
		return node.absoluteRoot()
	} else {
		// Otherwise the node is full, evenly split it into two nodes so:
		// A single median is chosen from among the leaf's elements
		// and the new element that is being inserted.

		// Values less than the median are put in the new left node,
		// and values greater than the median are put in the new right node,
		// with the median acting as a separation value.

		newLeftNode, newRightNode, separationValue := splitNode(node, val)

		indexesToRemove := []int{}

		if len(newLeftNode.children) > 0 {
			for i, v := range newLeftNode.children {
				// TODO should we only ever access the 0th index in v.keys?
				if v.keys[0] > separationValue {
					newRightNode.children = append(newRightNode.children, node.children[i])
					indexesToRemove = append(indexesToRemove, i)
				}
			}
		}

		// Remove the redundant node children
		indexModifier := 0
		for _, idx := range indexesToRemove {
			node.children = append(node.children[:idx+indexModifier], node.children[idx+1+indexModifier:]...)
			indexModifier-- // Without this we would get `slice bounds out of range` error in this loop
		}

		if node.parent == nil {
			// If the node has no parent (i.e., the node was the root),
			// create a new root above this node (increasing the height of the tree).
			newRoot := Node{keys: []int{separationValue}, children: []*Node{node, newRightNode}}
			newRoot.createChildParentPointers()
			return &newRoot
		} else {
			if node.parent.hasRoomForChildren() {
				node.parent.children = append(node.parent.children, newRightNode)
			} else {
				node.insertChild(newRightNode)
			}
				node.parent.createChildParentPointers()
			return node.parent.insert(separationValue)
		}
	}
}

// TODO unit test for every such function
func (node Node) hasFreeRoom() bool {
	return len(node.keys) < K-1
}

func appendAndSort(slice []int, val int) []int {
	newKeys := append(slice, val)
	slices.Sort(newKeys)
	return newKeys
}

func (node *Node) absoluteRoot() *Node {
	if node.parent == nil {
		return node
	}
	return node.parent.absoluteRoot()
}

func splitNode(node *Node, val int) (*Node, *Node, int) {
	temporaryKeySlice := appendAndSort(node.keys, val)
	medianIndex := K / 2 // What if K is an even number? How to choose median index then?

	// Values less than the median are put in the new left node,
	// and values greater than the median are put in the new right node,
	// with the median acting as a separation value.
	node.keys = temporaryKeySlice[:medianIndex] // The original node becomes the new left node
	newRightNode := Node{keys: temporaryKeySlice[medianIndex+1:], children: []*Node{}}

	separationValue := temporaryKeySlice[medianIndex]

	return node, &newRightNode, separationValue
}

func (node Node) hasChildren() bool {
	for _, child := range node.children {
		if child != nil {
			return true
		}
	}
	return false
}

func (node Node) determineChildIndex(val int) int {
	for idx, key := range node.keys {
		if val < key {
			return idx
		}
	}
	return len(node.keys)
}

func (node Node) hasRoomForChildren() bool {
	return len(node.children) < K
}

func (node Node) containsValue(val int) bool {
	return slices.Contains(node.keys, val)
}

func (node *Node) createChildParentPointers() {
	for _, child := range node.children {
		child.parent = node
	}
}

func (node *Node) insertChild(nodeToInsert *Node) {
	// TODO unsure as to why we index into keys[0] here -> investigate
	for idx, child := range node.parent.children {
		if child.keys[0] > nodeToInsert.keys[0] {
			node.parent.children = slices.Insert(node.parent.children, idx, nodeToInsert)
			return
		}
	}
	node.parent.children = append(node.parent.children, nodeToInsert)
}

func buildExampleBTree() Node {
	// lowest level - left
	lowest_l_l := Node{keys: []int{1}, children: []*Node{}}
	lowest_l_r := Node{keys: []int{3}, children: []*Node{}}
	// lowest level - right
	lowest_r_l := Node{keys: []int{5}, children: []*Node{}}
	lowest_r_r := Node{keys: []int{7}, children: []*Node{}}

	// mid level
	mid_l := Node{keys: []int{2}, children: []*Node{&lowest_l_l, &lowest_l_r}}
	mid_r := Node{keys: []int{6}, children: []*Node{&lowest_r_l, &lowest_r_r}}

	// top level
	root := Node{keys: []int{4}, children: []*Node{&mid_l, &mid_r}}

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

// TODO maybe this function can be deprecated?
func isInBTree(node Node, val int) bool {
	if node.containsValue(val) {
		return true
	} else {
		if node.hasChildren() {
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
