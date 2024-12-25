package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	if left != nil {
		return left
	}
	right := lowestCommonAncestor(root.Right, p, q)
	if right != nil {
		return right
	}
	return root
}

func main() {

}
