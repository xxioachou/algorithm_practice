package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T := 1
	// Fscan(in, &T)
	for ; T > 0; T-- {
		n := 1
		Fscan(in, &n)
		a := make([]int, n + 1)
		f := make([]int, n + 1)
		g := make([]int, n + 1)
		e := make([][]int, n + 1)
		for i := 1; i <= n; i++ {
			e[i] = make([]int, 0)
		}
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			if a[i] == 0 {
				a[i] = -1
			}
		}
		for i, u, v := 1, 0, 0; i < n; i++ {
			Fscan(in, &u, &v)
			e[u] = append(e[u], v)
			e[v] = append(e[v], u)
		}
		var dfs1 func(int, int)
		var dfs2 func(int, int)
		dfs1 = func(u, fa int) {
			f[u] = a[u]
			for _, v := range(e[u]) {
				if v == fa {
					continue
				}
				dfs1(v, u)
				f[u] += max(f[v], 0)
			}
		}
		dfs2 = func(u, fa int) {
			sum := 0
			for _, v := range(e[u]) {
				if v == fa {
					continue
				}
				sum += max(f[v], 0)
			}
			for _, v := range(e[u]) {
				if v == fa {
					continue
				}
				g[v] = max(a[v], a[v] + g[u] + sum - max(f[v], 0))
				dfs2(v, u)
			}
		}
		dfs1(1, 0)
		g[1] = a[1]
		dfs2(1, 0)
		// for i := 1; i <= n; i++ {
		// 	Fprintln(out, i, f[i], g[i], a[i])
		// }
		for i := 1; i <= n; i++ {
			Fprint(out, f[i] + g[i] - a[i], " ")
		}
		Fprintln(out, "")
	}
}

func main() { solve(os.Stdin, os.Stdout) }
