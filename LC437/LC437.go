package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	mp := make(map[int64]int, 0)
	var dfs func(*TreeNode, int64)
	ans := 0

	mp[0] = 1
	dfs = func(u *TreeNode, s int64) {
		if u == nil {
			return
		}
		s += int64(u.Val)
		ans += mp[s-int64(targetSum)]
		mp[s]++
		dfs(u.Left, s)
		dfs(u.Right, s)
		mp[s]--
	}
	dfs(root, 0)
	return ans
}

func main() {

}
