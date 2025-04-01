package main

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dfs func(l1 *ListNode, l2 *ListNode, add int) *ListNode
	dfs = func(l1, l2 *ListNode, add int) *ListNode {
		var ne *ListNode
		if l1 != nil || l2 != nil {
			if l1 != nil {
				add += l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				add += l2.Val
				l2 = l2.Next
			}
	
			ne = dfs(l1, l2, add / 10)
			add %= 10
		}
	
		if ne == nil && add == 0 {
			return nil
		}
		return &ListNode{Val: add, Next: ne}
	}
	
	res := dfs(l1, l2, 0)
	if res == nil {
		return &ListNode{Val: 0, Next: nil}
	}
	return res
}
