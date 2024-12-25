package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

func solve(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T := 1
	Fscan(in, &T)
	const N = 200010
	const M = 19
	lg := make([]int, N)
	for i := 2; i < N; i++ {
		lg[i] = lg[i >> 1] + 1
	}
	for ; T > 0; T-- {
		var n, q int
		Fscan(in, &n, &q)
		a := make([]int, n + 1)
		b := make([]int, n)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		for i := 1; i + 1 <= n; i++ {
			b[i] = a[i] - a[i + 1]
			if a[i] < a[i + 1] {
				b[i] = -b[i]
			}
			// Fprint(out, b[i], " ")
		}
		// Fprintln(out, "")
		f := make([][]int, n)
		for i := 1; i < n; i++ {
			f[i] = make([]int, M)
		}
		for j := 0; j < M; j++ {
			for i := 1; i + (1 << j) - 1 < n; i++ {
				if j != 0 {
					f[i][j] = gcd(f[i][j - 1], f[i + (1 << (j - 1))][j - 1])					
				} else {
					f[i][j] = b[i]
				}
			}
		}
		query := func(l, r int) int {
			k := lg[r - l + 1]
			return gcd(f[l][k], f[r - (1 << k) + 1][k])
		}
		for l, r, i := 0, 0, 0; i < q; i++ {
			Fscan(in, &l, &r)
			if l == r {
				Fprint(out, 0, " ")
			} else {
				Fprint(out, query(l, r - 1), " ")
			}
		}
		Fprintln(out, "")
	}
}

func main() { solve(os.Stdin, os.Stdout) }
