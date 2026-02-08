package main

import (
	"fmt"
	"math"
)

func main() {

	arrs := [][]*int{
		intPtrs(3, 9, 20, nil, nil, 15, 7),
		intPtrs(1, 2, 2, 3, 3, nil, nil, 4, 4), 
		// intPtrs(1, 2, 3, nil, nil, 4, 5),
		// intPtrs(1, 2, 3, 4, 5, 6, 7, nil, 9),
	}

	for i, arr := range arrs {
		fmt.Printf("building tree %d/%d...\n", i+1, len(arrs))
		tree := getTreeNode(arr, 0)
		printTree(tree, "", false)
		fmt.Printf("balanced: %v | max depth: %d\n", tree.isBalanced(), tree.getDepth())
		if i < len(arrs) - 1 { fmt.Println() }
	}
}

type treeNode struct {
	val int
	left *treeNode
	right *treeNode
}

func (n *treeNode) getDepth() int {
	if n == nil { 
		return 0
	}

	return (int(math.Max(float64(n.left.getDepth()), float64(n.right.getDepth()))) + 1)	
}

func (n *treeNode) isBalanced() bool {
	if n == nil { 
		return true
	}

	ld := n.left.getDepth() 
	rd := n.right.getDepth()

	if ld > (rd + 1) || ld < (rd - 1) || rd > (ld + 1) || rd < (ld - 1) {
		return false
	}

	return n.left.isBalanced() && n.right.isBalanced()
}

// pass array once, recursively builds tree structure
func getTreeNode(arr []*int, i int) *treeNode {
	if i >= len(arr) || arr[i] == nil { return nil }
	return &treeNode{
		*arr[i],
		getTreeNode(arr, 2 * i + 1),
		getTreeNode(arr, 2 * i + 2),
	}
}

// print the tree structure recursively, formatted as follows
// |--- 1
//      |--- 2
//      |--- 3
//           |--- 4
//           |--- 5
func printTree(n *treeNode, prefix string, isLeft bool) {
	if n == nil { return }

	fmt.Printf("%s|--- %v\n", prefix, n.val)

	printTree(n.left, prefix + func() string {
		if isLeft { return "|    " }
		return "     "
	}(), true)
	
	printTree(n.right, prefix + func() string {
		if isLeft { return "|    " }
		return "     "
	}(), false)
}

// pass int literals or nil, get slice of int ptrs
func intPtrs(arr ...any) []*int {
	var arrPtrs []*int
	for _, n := range arr {
		if n == nil {
			arrPtrs = append(arrPtrs, nil)
		}

		nInt, ok := n.(int)
		if ok { arrPtrs = append(arrPtrs, &nInt) }
	}
	return arrPtrs
}
