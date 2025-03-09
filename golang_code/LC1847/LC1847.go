package main

import (
	"fmt"
	"slices"
)

type pair struct {
	size int 
	id int
}
type node struct {
	maxv int
	minv int
}

func closestRoom(rooms [][]int, queries [][]int) []int {
    n := len(rooms)
	a := make([]pair, 1)
	for i := 1; i <= n; i++ {
		a = append(a, pair{ rooms[i - 1][1], rooms[i - 1][0] });
	}
	slices.SortFunc(a[1:], func (x, y pair) int {
		return x.size - y.size
	})
	// for i := 1; i <= n; i++ {
	// 	fmt.Println(i, a[i].size, a[i].id)
	// }
	tr := make([]node, (n + 10) * 4)
	const inf = 1000000010
	var Build func(int, int, int)
	var QueryMax func(int, int, int, int, int, int) int
	var QueryMin func(int, int, int, int, int, int) int
	Pushup := func (u int)  {
		tr[u].maxv = max(tr[u << 1].maxv, tr[u << 1 | 1].maxv)
		tr[u].minv = min(tr[u << 1].minv, tr[u << 1 | 1].minv)
	}
	Build = func (u, l, r int) {
		tr[u] = node{ a[r].id, a[r].id }
		if l == r {
			return
		}
		mid := (l + r) / 2
		Build(u << 1, l, mid)
		Build(u << 1 | 1, mid + 1, r)
		Pushup(u)
	}
	QueryMin = func(u, l, r, L, R, v int) int {
		if L <= l && r <= R {
			if tr[u].minv >= v {
				return tr[u].minv
			}
			if tr[u].maxv < v {
				return inf
			}
		}

		mid := (l + r) / 2
		res := inf
		if L <= mid {
			res = min(res, QueryMin(u << 1, l, mid, L, R, v))
		}
		if R > mid {
			res = min(res, QueryMin(u << 1 | 1, mid + 1, r, L, R, v))
		}
		return res
	}
	QueryMax = func(u, l, r, L, R, v int) int {
		if L <= l && r <= R {
			if tr[u].maxv < v {
				return tr[u].maxv
			}
			if tr[u].minv >= v {
				return -inf
			}
		}

		mid := (l + r) / 2
		res := -inf
		if L <= mid {
			res = max(res, QueryMax(u << 1, l, mid, L, R, v))
		}
		if R > mid {
			res = max(res, QueryMax(u << 1 | 1, mid + 1, r, L, R, v))
		}
		return res
	}
	GetPos := func (size int) int {
		l, r := 1, n
		for l < r {
			mid := (l + r) / 2
			if a[mid].size >= size {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if a[r].size < size {
			return -1
		}
		return r
	}
	Build(1, 1, n)
	ans := make([]int, 0)
	for _, v := range(queries) {
		p, ms := v[0], v[1]
		// >= ms 的起始位置
		l := GetPos(ms)
		// fmt.Println("l:", l)
		if l == -1 {
			ans = append(ans, -1)
			continue
		}
		// >= p 的最小值
		minv := QueryMin(1, 1, n, l, n, p)
		// < p 的最大值
		maxv := QueryMax(1, 1, n, l, n, p)
		// fmt.Println(":", minv, maxv)

		if minv - p >= p - maxv {
			ans = append(ans, maxv)
		} else {
			ans = append(ans, minv)
		}
	}
	return ans
} 

func main() {
	var n int
	fmt.Scan(&n)
	rooms := make([][]int, n)
	for i := range(rooms) {
		rooms[i] = make([]int, 2)
		fmt.Scan(&rooms[i][0], &rooms[i][1])
	}
	var k int
	fmt.Scan(&k)
	queries := make([][]int, k)
	for i := range(queries) {
		queries[i] = make([]int, 2)
		fmt.Scan(&queries[i][0], &queries[i][1])
	}
	ans := closestRoom(rooms, queries)
	for _, v := range(ans) {
		fmt.Print(v, " ")
	}
	fmt.Println("")
}