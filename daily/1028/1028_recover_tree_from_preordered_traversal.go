package _028

import "strconv"

/*

	We run a preorder depth-first search (DFS) on the root of a binary tree.

	At each node in this traversal, we output D dashes (where D is the depth of this node), then we output the value of this node.  If the depth of a node is D, the depth of its immediate child is D + 1.  The depth of the root node is 0.

	If a node has only one child, that child is guaranteed to be the left child.

	Given the output traversal of this traversal, recover the tree and return its root.


	Example 1:


	Input: traversal = "1-2--3--4-5--6--7"
	Output: [1,2,5,3,4,6,7]

	URL: https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/description/

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(traversal string) *TreeNode {
	if traversal == "" {
		return nil
	}
	return parse(traversal)
}

func parse(traversal string) *TreeNode {
	stack := []*TreeNode{}
	idx := 0
	for idx < len(traversal) {
		level := 0
		for idx < len(traversal) && traversal[idx] == '-' {
			level++
			idx++
		}

		start := idx
		for idx < len(traversal) && traversal[idx] >= '0' && traversal[idx] <= '9' {
			idx++
		}
		val, _ := strconv.Atoi(traversal[start:idx])
		node := &TreeNode{Val: val}

		for len(stack) > level {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			parent := stack[len(stack)-1]
			if parent.Left == nil {
				parent.Left = node
			} else {
				parent.Right = node
			}
		}

		stack = append(stack, node)
	}
	return stack[0]
}
