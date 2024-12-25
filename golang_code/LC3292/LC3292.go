package main

import "fmt"

type pair struct {
	h1 int
	h2 int
}

type node struct {
	v int
	minv int
}

func minValidStrings(words []string, target string) int {
	m := len(words)
	const P1 = 131
	const P2 = 13331
	const mod = 1e9 + 7
	const inf = 0x3f3f3f3f
	h := make([][]pair, m)	
	for i, str := range(words) {
		n := len(str)
		h[i] = make([]pair, n + 1)
		for j := 0; j < n; j++ {
			ch := int(str[j] - 'a' + 1)
			h[i][j + 1].h1 = (h[i][j].h1 * P1 % mod + ch) % mod
			h[i][j + 1].h2 = (h[i][j].h2 * P2 % mod + ch) % mod
		}
	}
	n := len(target)
	ha := make([]pair, n + 1)
	p1 := make([]int, n + 1)
	p2 := make([]int, n + 1)
	p1[0] = 1
	p2[0] = 1
	for i := 0; i < n; i++ {
		ch := int(target[i] - 'a' + 1)
		ha[i + 1].h1 = (ha[i].h1 * P1 % mod + ch) % mod
		ha[i + 1].h2 = (ha[i].h2 * P2 % mod + ch) % mod
		p1[i + 1] = p1[i] * P1 % mod
		p2[i + 1] = p2[i] * P2 % mod
	}
	f := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		f[i] = inf
	}	
	tr := make([]node, (n + 10) * 4)
	var build func (int, int, int)
	var update func (int, int, int, int, int, int)
	var query func (int, int, int, int) int
	pushup := func(u int) {
		tr[u].v = min(tr[u << 1].v, tr[u << 1 | 1].v)
	}
	pushdown := func(u int) {
		if tr[u].minv != inf {
			tr[u << 1].v = min(tr[u << 1].v, tr[u].minv)
			tr[u << 1 | 1].v = min(tr[u << 1 | 1].v, tr[u].minv)
			tr[u << 1].minv = min(tr[u << 1].minv, tr[u].minv)
			tr[u << 1 | 1].minv = min(tr[u << 1 | 1].minv, tr[u].minv)
			tr[u].minv = inf
		}
	}
	build = func(u, l, r int) {
		tr[u] = node{ f[r], inf }
		if l == r {
			return
		}
		mid := (l + r) / 2
		build(u << 1, l, mid)
		build(u << 1 | 1, mid + 1, r)
		pushup(u)
	}
	update = func(u, l, r, L, R, v int) {
		if L <= l && r <= R {
			tr[u].v = min(tr[u].v, v)
			tr[u].minv = min(tr[u].minv, v)
			return
		}
		pushdown(u)
		mid := (l + r) / 2
		if L <= mid {
			update(u << 1, l, mid, L, R, v)
		}
		if R > mid {
			update(u << 1 | 1, mid + 1, r, L, R, v)
		}
		pushup(u)
	}
	query = func(u, l, r, x int) int {
		if l == r {
			return tr[u].v
		}
		pushdown(u)
		mid := (l + r) / 2
		if x <= mid {
			return query(u << 1, l, mid, x)
		}
		return query(u << 1 | 1, mid + 1, r, x)
	}
	check := func(l, r int) bool {
		length := r - l + 1
		tmp := pair{
			(ha[r].h1 - ha[l - 1].h1 * p1[length] % mod + mod) % mod,
			(ha[r].h2 - ha[l - 1].h2 * p2[length] % mod + mod) % mod}
		for i := 0; i < m; i++ {
			if len(words[i]) >= length &&
				tmp.h1 == h[i][length].h1 && 
				tmp.h2 == h[i][length].h2 {
				return true
			}
		}
		return false
	}
	build(1, 0, n)
	for i := 1; i <= n; i++ {
		l, r := i, n
		for l < r {
			mid := (l + r + 1) / 2
			if (check(i, mid)) {
				l = mid
			} else {
				r = mid - 1
			}
		}
		if check(i, r) {
			// fmt.Println(i, r, query(1, 0, n, i - 1))
			update(1, 0, n, i, r, query(1, 0, n, i - 1) + 1)
		}
		// fmt.Println(":", i, query(1, 0, n, i))
	}
	ans := query(1, 0, n, n)
	if ans >= inf {
		ans = -1
	}
	return ans
} 

func main() {
	var m int
	fmt.Scan(&m)
	words := make([]string, m)
	for i := range(words) {
		fmt.Scan(&words[i])
	}
	var target string
	fmt.Scan(&target)
	fmt.Println(minValidStrings(words, target))
}