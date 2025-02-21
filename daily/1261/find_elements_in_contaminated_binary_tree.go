package _261

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	vals map[int]*TreeNode
}

func Constructor(root *TreeNode) FindElements {
	vals := make(map[int]*TreeNode)
	ele := &FindElements{vals: vals}
	recover(root, 0, ele)
	return *ele
}

func recover(root *TreeNode, val int, ele *FindElements) {
	if root == nil {
		return
	}
	root.Val = val
	ele.vals[val] = root
	recover(root.Left, val*2+1, ele)
	recover(root.Right, val*2+2, ele)
}

func (this *FindElements) Find(target int) bool {
	_, ok := this.vals[target]
	return ok
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
