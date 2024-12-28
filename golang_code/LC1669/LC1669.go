package main

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	tail := list2
	for tail.Next != nil {
		tail = tail.Next
	}
	l, r := list1, list1
	for i := 0; i < a - 1; i++ {
		l = l.Next
	}
	for i := 0; i < b + 1; i++ {
		r = r.Next
	}
	l.Next = list2
	tail.Next = r
	return list1
}	

func main() {

}