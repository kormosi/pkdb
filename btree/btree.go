// Tidbits from wiki and gipiti:

// Regarding BTrees:

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
	keys     []int // length of this cannot exceed K;
	children []*Node
	parent   *Node
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

func (node *Node) insert(val int) *Node {
	if node.hasFreeRoom() {
		node.keys = appendAndSort(node.keys, val)
		return node.absoluteRoot()
	} else {
		// TODO major confusion: newLeftNode and node point to the same object.. rename where appropriate
		newLeftNode, newRightNode, separationValue := splitNode(node, val)
		splitKeysBetweenNodes(newLeftNode, newRightNode, separationValue)
		if newLeftNode.parent == nil {
			newRoot := Node{keys: []int{separationValue}, children: []*Node{newLeftNode, newRightNode}}
			newRoot.createChildParentPointers()
			return &newRoot
		} else {
			node.parent.insertChild(newRightNode)
			node.parent.createChildParentPointers() // TODO no need to do this for the newRightNode?
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
	allKeys := appendAndSort(node.keys, val)
	medianIndex := K / 2 // TODO: What if K is an even number? How to choose median index then?
	separationValue := allKeys[medianIndex]
	node.keys = allKeys[:medianIndex] // The original node becomes the new left node
	newRightNode := Node{keys: allKeys[medianIndex+1:], children: []*Node{}}
	return node, &newRightNode, separationValue
}

func splitKeysBetweenNodes(leftNode *Node, rightNode *Node, separationValue int) {
	indexesToRemove := []int{}
	for i, v := range leftNode.children {
		// TODO should we only ever access the 0th index in v.keys?
		if v.keys[0] > separationValue {
			rightNode.children = append(rightNode.children, leftNode.children[i])
			rightNode.createChildParentPointers()
			indexesToRemove = append(indexesToRemove, i)
		}
	}
	// Remove the redundant leftNode children
	indexModifier := 0
	for _, idx := range indexesToRemove {
		leftNode.children = append(leftNode.children[:idx+indexModifier], leftNode.children[idx+1+indexModifier:]...)
		indexModifier-- // Without this we would get `slice bounds out of range` error in this loop
	}
}

func (node *Node) createChildParentPointers() {
	for _, child := range node.children {
		child.parent = node
	}
}

func (node *Node) insertChild(nodeToInsert *Node) {
	// TODO unsure as to why we index into keys[0] here -> investigate
	for idx, child := range node.children {
		if child.keys[0] > nodeToInsert.keys[0] {
			node.children = slices.Insert(node.children, idx, nodeToInsert)
			return
		}
	}
	node.children = append(node.children, nodeToInsert)
}

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

func (node Node) containsValue(val int) bool {
	return slices.Contains(node.keys, val)
}

func buildExampleBTree(values []int) Node {
	root := createEmptyBTree()
	for _, val := range values {
		root = root.insert(val)
	}
	return *root
}

func createEmptyBTree() *Node {
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

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
