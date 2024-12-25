package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	arr := make([]int, 0)
	mp := make(map[int]*TreeNode, 0)

	var dfs func(*TreeNode)
	dfs = func(u *TreeNode) {
		if u == nil {
			return
		}
		dfs(u.Left)
		arr = append(arr, u.Val)
		mp[u.Val] = u
		dfs(u.Right)
	}
	dfs(root)

	b := make([]int, len(arr))
	copy(b, arr)
	slices.Sort(b)

	var p, q *TreeNode
	for i := 0; i < len(arr); i++ {
		if b[i] != arr[i] {
			if p == nil {
				p = mp[b[i]]
			} else {
				q = mp[b[i]]
			}
		}
	}
	p.Val, q.Val = q.Val, p.Val
}

func main() {

}
