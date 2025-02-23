package _89

/*
   	Given two integer arrays, preorder and postorder where preorder is the preorder traversal of a binary tree of distinct values
 	and postorder is the postorder traversal of the same tree, reconstruct and return the binary tree.
 	If there exist multiple answers, you can return any of them.

	Example 1:
	Input: preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
	Output: [1,2,3,4,5,6,7]

	URL: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	num := len(preorder) - 1
	return inorder(0, num, 0, preorder, postorder)
}

func inorder(preStart int, preEnd int, postStart int, pre []int, post []int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	if preStart == preEnd {
		node := &TreeNode{}
		node.Val = pre[preStart]
		return node
	}

	leftRoot := pre[preStart+1]
	numOfNodesInLeft := 1
	for post[postStart+numOfNodesInLeft-1] != leftRoot {
		numOfNodesInLeft++
	}

	node := &TreeNode{}
	node.Val = pre[preStart]
	node.Left = inorder(preStart+1, preStart+numOfNodesInLeft, postStart, pre, post)
	node.Right = inorder(preStart+numOfNodesInLeft+1, preEnd, postStart+numOfNodesInLeft, pre, post)
	return node
}
