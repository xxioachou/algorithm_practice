package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
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
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

func main() {

}
