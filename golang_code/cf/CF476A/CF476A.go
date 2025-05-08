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
		var n, m int
		Fscan(in, &n, &m)
		f := make([][]int, n + 1)
		for i := range f {
			f[i] = make([]int, m)
			for j := range f[i] {
				f[i][j] = inf
			}
		}
		f[0][0] = 0
		f[1][1] = 1
		for i := 2; i <= n; i ++ {
			for j := 0; j < m; j ++ {
				f[i][j] = min(f[i - 1][(j - 1 + m) % m], f[i - 2][(j - 1 + m) % m]) + 1
			}
		}
		if f[n][0] < inf {
			Fprintln(out, f[n][0])
		} else {
			Fprintln(out, -1)
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }