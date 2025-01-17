package main

import "container/heap"


type ListNode struct {
    Val int
    Next *ListNode
}
type que []*ListNode
func(q que) Len() int {
	return len(q)
}
func(q que) Less(i, j int) bool {
	return q[i].Val < q[j].Val
}
func(q que) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func(q *que) Push(x any) {
	*q = append(*q, x.(*ListNode))
}
func(q *que) Pop() any {
	old := *q
	n := len(old)
	x := old[n - 1]
	*q = old[: n - 1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
    h := make(que, 0)
	heap.Init(&h)
	for _, v := range lists {
		if v != nil {
			heap.Push(&h, v)
		}
	}
	head := ListNode{-1, nil}
	cur := &head

	for h.Len() > 0 {
		x := heap.Pop(&h).(*ListNode)
		cur.Next = x
		cur = x
		if x.Next != nil {
			heap.Push(&h, x.Next)
		}
	}
	return head.Next
}

func main() {

}
