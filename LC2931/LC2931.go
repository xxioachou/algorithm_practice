package main

import "container/heap"
type Pair struct {
	val int
	idx int
}

type PriorityQueue []Pair
func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].val != pq[j].val {
		return pq[i].val < pq[j].val
	}
	return pq[i].idx < pq[j].idx
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x any) {
	pair := x.(Pair)
	*pq = append(*pq, pair)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	pair := old[n - 1]
	*pq = old[0 : n-1]
	return pair
}

func maxSpending(values [][]int) int64 {
    q := make(PriorityQueue, 0)
	heap.Init(&q)
	m := len(values)
	for i := 0; i < m; i++ {
		n := len(values[i])
		heap.Push(&q, Pair{values[i][n - 1], i})
		values[i] = values[i][: n - 1]
	}
	ans := int64(0)
	for d := 1; q.Len() > 0; d++ {
		tmp := heap.Pop(&q).(Pair)
		ans += int64(tmp.val) * int64(d)
		n := len(values[tmp.idx])
		if n > 0 {
			t := Pair{values[tmp.idx][n - 1], tmp.idx}
			values[tmp.idx] = values[tmp.idx][: n - 1]
			heap.Push(&q, t)
		}
	}
	return int64(ans)
} 

func main() {

}