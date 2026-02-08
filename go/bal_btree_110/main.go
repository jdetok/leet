package main

import (
	"fmt"
)

func main() {
	arr := intPtrs(1, 2, 2, 3, 4, nil, 5, 6, 9, nil, 10, 11, nil, 11)
	printTree(getTreeNode(arr, 0), "", false)

	arr2 := intPtrs(1, 2, 3, nil, nil, 4, 5)
	printTree(getTreeNode(arr2, 0), "", false)
	
	arr3 := intPtrs(1, 2, 3, 4, 5, 6, 7, nil, 9)
	printTree(getTreeNode(arr3, 0), "", false)
}

type treeNode struct {
	val int
	left *treeNode
	right *treeNode
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
