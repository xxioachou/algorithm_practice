package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type node struct {
	lson int
	rson int
	cnt int
}

func solve(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T := 1
	Fscan(in, &T)

	const N = 100010
	tr := make([]node, N * 4 + 17 * N)
	root := make([]int, N)
	dfn := make([]int, N)
	outn := make([]int, N)
	p := make([]int, N)
	idx, timestamp := 0, 0
	var insert func(int, int, int, int) int
	var query func(int, int, int, int, int, int) int
	insert = func(p, l, r, x int) int {
		idx ++
		q := idx
		if l == r {
			tr[q].cnt ++
			return q
		}
		tr[q] = tr[p]
		mid := (l + r) / 2
		if x <= mid {
			tr[q].lson = insert(tr[p].lson, l, mid, x)
		} else {
			tr[q].rson = insert(tr[p].rson, mid + 1, r, x)
		}
		tr[q].cnt = tr[tr[q].lson].cnt + tr[tr[q].rson].cnt
		return q
	}
	query = func(p, q, l, r, L, R int) int {
		if L <= l && r <= R {
			return tr[q].cnt - tr[p].cnt;
		}
		mid := (l + r) / 2
		res := 0
		if L <= mid {
			res += query(tr[p].lson, tr[q].lson, l, mid, L, R)
		}
		if R > mid {
			res += query(tr[p].rson, tr[q].rson, mid + 1, r, L, R)
		}
		return res
	}

	for ; T > 0; T-- {
		for i := 1; i < idx; i++ {
			tr[i].cnt, tr[i].lson, tr[i].rson = 0, 0, 0
		}
		idx, timestamp = 0, 0

		var n, q int
		Fscan(in, &n, &q)
		e := make([][]int, n + 1)
		for i := 1; i <= n; i++ {
			e[i] = make([]int, 0)
		}
		for i, u, v := 1, 0, 0; i < n; i++ {
			Fscan(in, &u, &v)
			e[u] = append(e[u], v)
			e[v] = append(e[v], u)
		}
		for i := 1; i <= n; i++ {
			Fscan(in, &p[i])
		}
		var dfs func(int, int)
		dfs = func(u, fa int) {
			timestamp ++
			dfn[u] = timestamp
			for _, v := range(e[u]) {
				if v != fa {
					dfs(v, u)
				}
			}
			outn[u] = timestamp
			// Fprintln(out, u, dfn[u], outn[u])
		}
		dfs(1, 0)
		for i := 1; i <= n; i++ {
			root[i] = insert(root[i - 1], 1, n, dfn[p[i]])
		}
		for i, l, r, x := 1, 0, 0, 0; i <= q; i++ {
			Fscan(in, &l, &r, &x)
			res := query(root[l - 1], root[r], 1, n, dfn[x], outn[x])
			if res > 0 {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		}
	}	
}

func main() { solve(os.Stdin, os.Stdout) }
