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
		var n, m, k int
		Fscan(in, &n, &m, &k)
		x := make([][]int, n)
		t := make([]int, k + 1)
		lock := make([]bool, k + 1)
		ans := make([]int, n)
		for i := 0; i < n; i++ {
			x[i] = make([]int, m)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				Fscan(in, &x[i][j])
			}
		}
		for j := 0; j < m; j++ {
			for i := range(t) {
				t[i] = 0
			}
			for i := 0; i < n; i++ {
				if x[i][j] == 0 || ans[i] != 0 {
					continue
				}
				if !lock[x[i][j]] {
					t[x[i][j]] ++
				}
			}
			
			for i := 1; i <= k; i++ {
				if t[i] > 1 {
					lock[i] = true
				}
			}
			for i := 0; i < n; i++ {
				if x[i][j] != 0 && lock[x[i][j]] && ans[i] == 0 {
					ans[i] = j + 1
				}
			}
		}
		for i := 0; i < n; i++ {
			Fprintln(out, ans[i])
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }
