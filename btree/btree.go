// from gipiti:
// If it’s meant to be a library (not executable),
//
//	don’t use go run; instead, import it from a main package.
package main

import (
	"fmt"
)

//type node struct {
//	value int
//	child_l *node
//	child_r *node
//}

const N = 3

type node struct {
	keys [N - 1]int // should be len(children)-1 (it's called keys but it's also values)
	// value int  //  i'm thinking -1 for no value (when the node is a root or internal)
	children [N]*node //  an array of pointers? is this right? (size of this array will be subject to restrictions regarding order of the tree)
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

// TODO premenovať na "isValueInNode" alebo také niečo -
// je to viac domain specific

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
			return idx
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

// TODO poriešiť git štruktúru, takto:
// pkdb -> btree
// a nech git root je v pkdb

// napísať testy - zistiť, ako sa to robí.
// žeby samostatný pkdb -> test adresár?
