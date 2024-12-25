package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(u *TreeNode, maxv, minv *int) bool {

	lmax := math.MinInt32
	rmax := math.MinInt32
	lmin := math.MaxInt32
	rmin := math.MaxInt32

	if u.Left != nil {
		if !dfs(u.Left, &lmax, &lmin) || lmax >= u.Val {
			return false
		}
	}

	if u.Right != nil {
		if !dfs(u.Right, &rmax, &rmin) || rmin <= u.Val {
			return false
		}
	}

	if maxv != nil {
		*maxv = max(*maxv, max(u.Val, lmax, rmax))
	}
	if minv != nil {
		*minv = min(*minv, min(u.Val, lmin, rmin))
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, nil, nil)
}

func main() {

}
