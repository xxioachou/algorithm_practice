package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
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

	rev := func(a []int) {
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	}
	for i := 1; i < len(ans); i += 2 {
		rev(ans[i])
	}
	return ans
}

func main() {

}
