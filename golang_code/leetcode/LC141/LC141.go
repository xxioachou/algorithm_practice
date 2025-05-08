package main

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	l, r := head, head
	for i := 0; i < 10005; i ++ {
		if l == nil || r == nil || r.Next == nil {
			return false
		}
		l = l.Next
		r = r.Next.Next
		if l == r {
			return true
		}
	}
	return false
}
