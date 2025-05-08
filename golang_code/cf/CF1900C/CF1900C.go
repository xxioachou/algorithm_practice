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
	Fscan(in, &T)
	for ; T > 0; T-- {
		var n int
		var s string
		Fscan(in, &n, &s)
		l := make([]int, n)
		r := make([]int, n)
		for i := 0; i < n; i ++ {
			Fscan(in, &l[i], &r[i])
			l[i] --
			r[i] --
		}
		var dfs func(u int) int
		dfs = func(u int) int {
			if l[u] < 0 && r[u] < 0 {
				return 0
			}
			lv, rv := inf, inf
			if l[u] >= 0 {
				lv = dfs(l[u])
			}
			if r[u] >= 0 {
				rv = dfs(r[u])
			}
			if s[u] != 'L' {
				lv ++
			}
			if s[u] != 'R' {
				rv ++
			}
			return min(lv, rv)
		}
		Fprintln(out, dfs(0))
	}
}

func main() { solve(os.Stdin, os.Stdout) }