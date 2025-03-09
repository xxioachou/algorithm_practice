package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}

	if !isSameTree(p.Left, q.Left) ||
		p.Val != q.Val ||
		!isSameTree(p.Right, q.Right) {
		return false
	}
	return true
}

func main() {

}
