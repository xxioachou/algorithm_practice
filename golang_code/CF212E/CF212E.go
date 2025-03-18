package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	
	const inf = 0x3f3f3f3f
	T := 1
	// T = readLine()[0]
	for ; T > 0; T-- {
		var n, x, y int
		Fscan(in, &n)
		e := make([][]int, n)
		ans := make([]bool, n + 1)
		for i := range e {
			e[i] = make([]int, 0)
		}

		for i := 0; i + 1 < n; i ++ {
			Fscan(in, &x, &y)
			x, y = x - 1, y - 1
			e[x] = append(e[x], y)
			e[y] = append(e[y], x)
		}
		var dfs func(u, fa int) int
		dfs = func(u, fa int) int {
			a := make([]int, 0)
			sz := 1
			for _, v := range e[u] {
				if v != fa {
					t := dfs(v, u)
					sz += t
					a = append(a, t)
				}
			}
			if n - sz > 0 {
				a = append(a, n - sz)
			}
			if len(a) < 2 {
				return sz
			}

			f := make([]bool, n + 1)
			f[0] = true
			for _, x := range a {
				for j := n; j >= x; j -- {
					f[j] = f[j] || f[j - x]
				}
			}
			// Fprintf(out, "u %v, f %v\n", u, f)

			for j := 1; j <= n; j ++ {
				ans[j] = ans[j] || f[j]
			}

			return sz
		}
	
		dfs(0, -1)
		cnt := 0
		for i := 1; i < n - 1; i ++ {
			if ans[i] {
				cnt ++
			}
		}
		Fprintln(out, cnt)
		for i := 1; i < n - 1; i ++ {
			if ans[i] {
				Fprintln(out, i, n - 1 - i)
			}
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }