package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) ||
		(p != nil && q == nil) ||
		p.Val != q.Val {
		return false
	}

	if !helper(p.Right, q.Left) ||
		!helper(p.Left, q.Right) {
		return false
	}

	return true
}

func isSymmetric(root *TreeNode) bool {
	return helper(root.Left, root.Right)
}

func main() {

}
