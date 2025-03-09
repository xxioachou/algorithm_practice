package main

import (
	"container/heap"
)

type node struct {
	val int
	idx int
}
type que []node
func (q que) Len() int {
	return len(q)
}
func (q que) Less(i, j int) bool {
	if q[i].val != q[j].val {
		return q[i].val < q[j].val
	}
	return q[i].idx < q[j].idx
}
func (q que) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q *que) Push(x any) {
	*q = append(*q, x.(node))
}
func (q *que) Pop() any {
	old := *q
	n := len(old)
	x := old[n - 1]
	*q = old[0 : n - 1]
	return x
}

func getFinalState(nums []int, k int, multiplier int) []int {
	q := make(que, 0)
    for i := 0; i < len(nums); i++ {
		heap.Push(&q, node{nums[i], i})
	}
	heap.Init(&q)
	for i := 0; i < k; i++ {
		t := heap.Pop(&q).(node)
		// fmt.Println(i, t.idx, t.val)
		heap.Push(&q, node{t.val * multiplier, t.idx})
	}
	for q.Len() > 0 {
		t := heap.Pop(&q).(node)
		nums[t.idx] = t.val
	}
	return nums
} 

func main() {

}