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

// TODO rename to btree?

package btree

import (
	"fmt"
)

const N = 3 // so-called order of the tree

type node struct {
	keys [N - 1]int // should be len(children)-1; (it's called keys but it's also values)
	// value int  //  i'm thinking -1 for no value (when the node is a root or internal)
	children [N]*node // (size of this array will be subject to restrictions regarding order of the tree)
}

func buildBTree() node {
	// lowest level - left
	lowest_l_l := node{keys: [N - 1]int{1, -1}, children: [N]*node{}}
	lowest_l_r := node{keys: [N - 1]int{3, -1}, children: [N]*node{}}
	// lowest level - right
	lowest_r_l := node{keys: [N - 1]int{5, -1}, children: [N]*node{}}
	lowest_r_r := node{keys: [N - 1]int{7, -1}, children: [N]*node{}}

	// mid level
	mid_l := node{keys: [N - 1]int{2, -1}, children: [N]*node{&lowest_l_l, &lowest_l_r}}
	mid_r := node{keys: [N - 1]int{6, -1}, children: [N]*node{&lowest_r_l, &lowest_r_r}}

	// top level
	root := node{keys: [N - 1]int{4, -1}, children: [N]*node{&mid_l, &mid_r}}

	return root
}

func printBTree(root node) {
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

func isValueInNode(array [N - 1]int, val int) bool {
	for _, el := range array {
		if el == val {
			return true
		}
	}
	return false
}

func hasValidChildren(node node) bool {
	for _, child := range node.children {
		if child != nil {
			return true
		}
	}
	return false
}

func determineChild(keys [N - 1]int, val int) int {
	for idx, key := range keys {
		if key == -1 {
			return idx // maybe return -1 in this case and thus end the search?
		}
		if val < key {
			return idx
		}
	}
	return N - 1
}

func isInBTree(node node, val int) bool {
	if isValueInNode(node.keys, val) {
		return true
	} else {
		if hasValidChildren(node) {
			childToSearch := determineChild(node.keys, val)
			return isInBTree(*node.children[childToSearch], val)
		}
	}
	return false
}

// TODO use array only on the highest-level
// in lower functions, use slices for better testability?
// candidate functions:  determineChild, isValueInNode (could use the .contains method)

func main() {
	root := buildBTree()
	printBTree(root)

	fmt.Println()

	// TODO otestovať to s viacerými hodnotami
	// fmt.Println(isInBTree(root, 4))
	// fmt.Println(isInBTree(root, 2))
	// fmt.Println(isInBTree(root, 1))
	// fmt.Println(isInBTree(root, 3))
	// fmt.Println(isInBTree(root, 6))
	// fmt.Println(isInBTree(root, 5))
	// fmt.Println(isInBTree(root, 7))

	fmt.Println(isInBTree(root, 8))
}

// napísať testy - zistiť, ako sa to robí.
// žeby samostatný pkdb -> test adresár?
