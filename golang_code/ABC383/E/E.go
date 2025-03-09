package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

type pair struct {
	x int
	y int
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type edge struct {
	u int
	v int
	w int
}

type ByW []edge
func (a ByW) Len() int {
	return len(a)
}
func (a ByW) Swap(i, j int) { 
	a[i], a[j] = a[j], a[i]
}
func (a ByW) Less(i, j int) bool {
	return a[i].w < a[j].w
}

func solve(_r io.Reader, out io.Writer) {
	T := 1
	in := bufio.NewReader(_r)
	// Fscan(in, &T)
	for ; T > 0; T-- {
		var N, M, K int
		Fscan(in, &N, &M, &K)
		E := make([]edge, M)
		cnta := make([]int, N + N)
		cntb := make([]int, N + N)
		for i := range(E) {
			var u, v, w int
			Fscan(in, &u, &v, &w)
			E[i] = edge{u, v, w}
		}
		for i, a := 0, 0; i < K; i++ {
			Fscan(in, &a)
			cnta[a] ++
		}
		for i, b := 0, 0; i < K; i++ {
			Fscan(in, &b)
			cntb[b] ++
		}
				
		// 根据边权从小到达排序
		sort.Sort(ByW(E))

		p := make([]int, N + N)
		e := make([][]int, N + N)
		wg := make([]int, N + N)
		for i := 1; i < len(e); i++ {
			e[i] = make([]int, 0)
		}
		for i := 1; i < N + N; i++ {
			p[i] = i
		}
		var find func(int) int
		find = func(x int) int {
			if p[x] != x {
				p[x] = find(p[x])
			}
			return p[x]
		}

		// 建立 kruskal 重构树
		idx := N
		for i := range(E) {
			u, v, w := E[i].u, E[i].v, E[i].w
			// Fprintln(out, u, v, w)
			u = find(u)
			v = find(v)
			if u != v {
				fa := idx + 1
				idx ++
				wg[fa] = w
				e[fa] = append(e[fa], u, v)
				e[u] = append(e[u], fa)
				e[v] = append(e[v], fa)
				p[u] = fa
				p[v] = fa
			}
		}

		var dfs func(int, int) int
		dfs = func(u, fa int) int {
			res := 0
			for _, v := range(e[u]) {
				if v != fa {
					res += dfs(v, u)
					cnta[u] += cnta[v]
					cntb[u] += cntb[v]
				}
			}
			t := min(cnta[u], cntb[u])
			// Fprintln(out, u, cnta[u], cntb[u], wg[u])
			cnta[u] -= t
			cntb[u] -= t
			res += t * wg[u]
			return res
		}
		Fprintln(out, dfs(idx, 0))
	}
}

func main() { solve(os.Stdin, os.Stdout) }
