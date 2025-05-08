package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func same(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && same(p.Left, q.Left) && same(p.Right, q.Right)
}

func dfs(root *TreeNode, tmp *TreeNode, height *int, cnt *int) int {
	if root == nil {
		return 0
	}
	h := max(dfs(root.Left, tmp, height, cnt), dfs(root.Right, tmp, height, cnt)) + 1
	if h == *height {
		if same(root, tmp) {
			*cnt += 1
		}
	}
	return h
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var find func(*TreeNode) int
	ans := make([]*TreeNode, 0)
	mp := make(map[*TreeNode]int, 0)
	find = func(u *TreeNode) int {
		if u == nil {
			return 0
		}
		h := max(find(u.Left), find(u.Right)) + 1
		cnt := 0
		dfs(root, u, &h, &cnt)
		if cnt > 1 {
			ans = append(ans, u)
		}
		return h
	}
	find(root)
	return ans
}

func main() {

}
