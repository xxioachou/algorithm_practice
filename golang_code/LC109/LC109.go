package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	a := make([]*ListNode, 0)
	for head != nil {
		a = append(a, head)
		head = head.Next
	}

	var dfs func([]*ListNode) *TreeNode
	dfs = func(u []*ListNode) *TreeNode {
		if len(u) == 0 {
			return nil
		}
		idx := len(u) / 2
		r := TreeNode{u[idx].Val, nil, nil}
		r.Left = dfs(u[:idx])
		r.Right = dfs(u[idx+1:])
		return &r
	}

	return dfs(a)
}

func main() {

}
