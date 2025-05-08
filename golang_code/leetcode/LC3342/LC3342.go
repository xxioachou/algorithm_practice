package main

import (
	"container/heap"
	"math"
)

func minTimeToReach(moveTime [][]int) int {
	n, m := len(moveTime), len(moveTime[0])
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, m)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = 0

	h := make(que, 0)
	heap.Init(&h)
	heap.Push(&h, node{dis[0][0], 0, 0})

	for h.Len() > 0 {
		no := heap.Pop(&h).(node)
		x, y, d, c := no.x, no.y, no.d, 1
		if d > dis[x][y] {
			continue
		}
		if (x+y)%2 == 1 {
			c = 2
		}

		for i := 0; i < 4; i++ {
			a, b := x+dx[i], y+dy[i]
			if a < 0 || a >= n || b < 0 || b >= m {
				continue
			}
			tmp := max(d, moveTime[a][b]) + c
			if dis[a][b] > tmp {
				dis[a][b] = tmp
				heap.Push(&h, node{dis[a][b], a, b})
			}
		}
	}
	return dis[n-1][m-1]
}

type node struct {
	d, x, y int
}

type que []node

func (q que) Len() int {
	return len(q)
}

func (q que) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q que) Less(i, j int) bool {
	return q[i].d < q[j].d
}

func (q *que) Push(x any) {
	*q = append(*q, x.(node))
}

func (q *que) Pop() any {
	old := *q
	*q = old[:len(old)-1]
	return old[len(old)-1]
}
