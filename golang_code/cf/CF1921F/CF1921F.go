package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func solve(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T := 1
	Fscan(in, &T)
	for ; T > 0; T-- {
		n, q := 0, 0
		Fscan(in, &n, &q)
		m := int(math.Floor(math.Sqrt(float64(n))))
		a := make([]int, n + 1)
		sum := make([][]int, m + 1)
		p := make([][]int, m + 1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		for d := 1; d <= m; d++ {
			sum[d] = make([]int, n + 1)
			p[d] = make([]int, n + 1)
			for i := 1; i <= n; i++ {
				idx := (i - 1) / d + 1
				sum[d][i] = sum[d][max(i - d, 0)] + a[i]
				p[d][i] = p[d][max(i - d, 0)] + a[i] * idx
			}
		}

		for s, d, k, i := 0, 0, 0, 0; i < q; i++ {
			Fscan(in, &s, &d, &k)
			ans := 0
			if d > m {
				for j := 0; j < k; j++ {
					ans += a[s + d * j] * (j + 1)
				}
			} else {
				l, r := s, s + d * (k - 1)
				idx := (l - 1) / d + 1
				ans = p[d][r] - p[d][max(l - d, 0)] - (idx - 1) * (sum[d][r] - sum[d][max(l - d, 0)])
			}
			Fprint(out, ans, " ")
		}
		Fprintln(out, "")
	}
}

func main() { solve(os.Stdin, os.Stdout) }
