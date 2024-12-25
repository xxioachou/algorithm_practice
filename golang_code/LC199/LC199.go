package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var dfs func(*TreeNode, int)
	ans := make([]int, 0)
	dfs = func(u *TreeNode, dep int) {
		if u == nil {
			return
		}
		if dep == len(ans) {
			ans = append(ans, u.Val)
		}
		dfs(u.Right, dep+1)
		dfs(u.Left, dep+1)
	}
	dfs(root, 0)
	return ans
}

func main() {

}
