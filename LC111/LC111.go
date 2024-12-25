package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	res := math.MaxInt32
	if root.Left != nil {
		res = min(res, minDepth(root.Left)+1)
	}
	if root.Right != nil {
		res = min(res, minDepth(root.Right)+1)
	}
	return res
}

func main() {

}
