package main

import (
	"container/list"
	"math"
)

func minTimeToReach(moveTime [][]int) int {
	n, m := len(moveTime), len(moveTime[0])
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, m)
		for j := range d[i] {
			d[i][j] = math.MaxInt
		}
	}
	d[0][0] = 0
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	q := list.New()
	q.PushBack(0)
	for q.Len() > 0 {
		v := q.Front().Value.(int)
		q.Remove(q.Front())
		x, y := v/m, v%m
		for i := 0; i < 4; i++ {
			a, b := x+dx[i], y+dy[i]
			if a < 0 || a >= n || b < 0 || b >= m {
				continue
			}
			tmp := max(moveTime[a][b], d[x][y]) + 1
			if d[a][b] > tmp {
				d[a][b] = tmp
				q.PushBack(a*m + b)
			}
		}
	}
	return d[n-1][m-1]
}

func main() {

}
