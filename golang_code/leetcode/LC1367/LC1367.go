package main


type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	var dfs func(*ListNode, *TreeNode, bool) bool
	dfs = func(h *ListNode, r *TreeNode, start bool) bool {
		if h == nil {
			return true
		}
		if r == nil {
			return false
		}
		if r.Val == h.Val {
			if dfs(h.Next, r.Left, true) {
				return true
			}
			if dfs(h.Next, r.Right, true) {
				return true
			}
		}
		if !start {
			return dfs(h, r.Left, false) || 
					dfs(h, r.Right, false)
		}
		return false
	}
	return dfs(head, root, false)
} 

func main() {

}