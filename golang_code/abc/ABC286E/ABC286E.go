package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	
	T := 1
	const inf = 0x3f3f3f3f
	// T = readLine()[0]
	for ; T > 0; T-- {
		var N int
		Fscan(in, &N)
		A := make([]int, N)
		d := make([][]int, N)
		v := make([][]int, N)
		for i := range A {
			d[i] = make([]int, N)
			v[i] = make([]int, N)
			for j := range d[i] {
				d[i][j] = inf
			}
			d[i][i] = 0
			v[i][i] = A[i]
		}
		for i := range A {
			Fscan(in, &A[i])
		}
		for i := range A {
			var s string
			Fscan(in, &s)
			for j := range s {
				if s[j] == 'Y' {
					d[i][j] = 1
					v[i][j] = A[i] + A[j]
				}
			}
		}
		for k := 0; k < N; k ++ {
			for i := 0; i < N; i ++ {
				for j := 0; j < N; j ++ {
					if d[i][j] > d[i][k] + d[k][j] {
						d[i][j] = d[i][k] + d[k][j]
						v[i][j] = v[i][k] + v[k][j] - A[k]
					} else if d[i][j] == d[i][k] + d[k][j] {
						v[i][j] = max(v[i][j], v[i][k] + v[k][j] - A[k])
					}
				}
			}
		}
		var Q int
		Fscan(in, &Q)
		for i := 0; i < Q; i ++ {
			var U, V int
			Fscan(in, &U, &V)
			U, V = U - 1, V - 1
			if d[U][V] >= inf {
				Fprintln(out, "Impossible")
			} else {
				Fprintln(out, d[U][V], v[U][V])
			}
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }