package main

import (
	"container/heap"
	"fmt"
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
	*q = old[: len(old) - 1]
	return old[len(old) - 1]
}

func jump(nums []int) int {
	n := len(nums)
	h := make(que, 0)
	heap.Init(&h)
	stk := make([]int, 0)	
	ban := make([]bool, n + 1)

	stk = append(stk, 0)
	heap.Push(&h, pair{0, 0})
	ans := 0
	for i := 1; i < n; i++ {
		for len(stk) > 0 {
			j := stk[len(stk) - 1]
			if nums[j] + j < i {
				stk = stk[: len(stk) - 1]
				ban[j] = true
			} else {
				break
			}
		}
		fmt.Println(":", ans)
		fmt.Println(":stk", stk)
		for h.Len() > 0 {
			p := heap.Pop(&h).(pair)
			if ban[p.y] {
				continue
			}
			ans = p.x + 1
			heap.Push(&h, p)
			break
		}
		fmt.Println(":", ans)
		heap.Push(&h, pair{ans, i})
		stk = append(stk, i)
		fmt.Println("stk", stk)
	}
	return ans
}

func main() {

}