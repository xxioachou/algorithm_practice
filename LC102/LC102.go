package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var dfs func(*TreeNode, int)
	ans := make([][]int, 0)
	dfs = func(u *TreeNode, dep int) {
		if u == nil {
			return
		}
		if dep == len(ans) {
			ans = append(ans, make([]int, 0))
		}
		ans[dep] = append(ans[dep], u.Val)
		dfs(u.Left, dep+1)
		dfs(u.Right, dep+1)
	}
	dfs(root, 0)
	return ans
}

func main() {

}
