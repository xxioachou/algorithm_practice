package main

import (
	"container/heap"
	"fmt"
)

type pair struct {
	t   int
	cnt int
}
type pq []pair

func (q pq) Less(i, j int) bool {
	if q[i].t != q[j].t {
		return q[i].t < q[j].t
	}
	return q[i].cnt > q[j].cnt
}
func (q pq) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q pq) Len() int {
	return len(q)
}
func (q *pq) Push(x any) {
	*q = append(*q, x.(pair))
}
func (q *pq) Pop() any {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

func eatenApples(apples []int, days []int) int {
	q := make(pq, 0)
	heap.Init(&q)
	n := len(apples)
	ans := 0
	for i := 1; ; i++ {
		if i <= n {
			heap.Push(&q, pair{i + days[i - 1], apples[i - 1]})
		}
		for q.Len() > 0 {
			tmp := heap.Pop(&q).(pair)
			if tmp.t <= i || tmp.cnt <= 0 {
				continue
			}
			heap.Push(&q, tmp)
			break
		}
		if q.Len() == 0 {
			if i > n {
				break
			}
			continue
		}
		ans ++
		t := heap.Pop(&q).(pair)
		heap.Push(&q, pair{t.t, t.cnt - 1})
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)
	apples := make([]int, n)
	days := make([]int, n)
	for i := range(apples) {
		fmt.Scan(&apples[i])
	}
	for i := range(days) {
		fmt.Scan(&days[i])
	}
	fmt.Println(eatenApples(apples, days))
}
