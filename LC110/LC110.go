package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	ok := true

	var dfs func(*TreeNode) int
	dfs = func(u *TreeNode) int {
		if u == nil {
			return 0
		}
		lh := dfs(u.Left)
		rh := dfs(u.Right)
		if max(lh, rh)-min(lh, rh) > 1 {
			ok = false
		}
		return max(lh, rh) + 1
	}
	dfs(root)
	return ok
}

func main() {

}
