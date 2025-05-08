package main

import (
	"container/heap"
)

type pair struct {
	x int
	y int
}
type que []pair
func (q que) Swap(x, y int) {
	q[x], q[y] = q[y], q[x]
}
func (q que) Len() int {
	return len(q)
}
func(q que) Less(i, j int) bool {
	if q[i].x != q[j].x {
		return q[i].x < q[j].x
	}
	return q[i].y < q[j].y
}
func (q *que) Push(x any) {
	*q = append(*q, x.(pair))
}
func (q *que) Pop() any {
	old := *q
	*q = old[:len(old) - 1]
	return old[len(old) - 1]
}

func jump(nums []int) int {
	n := len(nums)
	h := make(que, 0)
	stk := make(que, 0)	
	ban := make([]bool, n)
	heap.Init(&h)
	heap.Init(&stk)

	heap.Push(&stk, pair{0 + nums[0], 0})
	heap.Push(&h, pair{0, 0})
	ans := 0
	for i := 1; i < n; i++ {
		for stk.Len() > 0 && stk[0].x < i{
			ban[stk[0].y] = true
			heap.Pop(&stk)
		}
		for h.Len() > 0 && ban[h[0].y] {
			heap.Pop(&h)
		}

		ans = h[0].x + 1
		heap.Push(&h, pair{ans, i})
		heap.Push(&stk, pair{i + nums[i], i})
	}
	return ans
}

func main() {

}