package main

import "container/heap"

type que []int

func (q que) Less(i int, j int) bool {
	return q[i] < q[j]
}
func (q que) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q *que) Pop() any {
	t := *q
	n := len(t)
	x := t[n-1]
	*q = t[:n-1]
	return x
}
func (q *que) Push(x any) {
	*q = append(*q, x.(int))
}
func (q que) Len() int {
	return len(q)
}

func minOperations(nums []int, k int) int {
	h := make(que, 0)
	heap.Init(&h)
	for _, x := range nums {
		heap.Push(&h, x)
	}
	ans := 0
	for {
		x := heap.Pop(&h).(int)
		if x >= k {
			break
		}
		y := heap.Pop(&h).(int)
		heap.Push(&h, x * 2 + y)
		ans ++
	}
	return ans
}

func main() {

}
