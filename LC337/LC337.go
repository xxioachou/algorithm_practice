package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(*TreeNode) []int
	dfs = func(u *TreeNode) []int {
		if u == nil {
			return []int{0, 0}
		}
		a, b := u.Val, 0

		l := dfs(u.Left)
		r := dfs(u.Right)

		a += l[1]
		a += r[1]
		b += max(l[0], l[1])
		b += max(r[0], r[1])
		return []int{a, b}
	}
	return max(dfs(root)[0], dfs(root)[1])
}

func main() {

}
