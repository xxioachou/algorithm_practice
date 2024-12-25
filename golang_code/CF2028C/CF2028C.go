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
		var n, m, v int
		Fscan(in, &n, &m, &v)
		a := make([]int, n + 1)
		s := make([]int, n + 1)
		pre := make([]int, n + 2)
		suf := make([]int, n + 2)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		for i := 1; i <= n; i++ {
			s[i] = s[i - 1] + a[i]
		}
		for i, j := 1, 0; i <= n; i++ {
			for j + 1 < i && s[i] - s[j + 1] >= v {
				j ++
			}
			t := 0
			if s[i] - s[j] >= v {
				t = 1
			}
			pre[i] = pre[j] + t
		}
		for i, j := n, n + 1; i >= 1; i-- {
			for j - 1 > i && s[j - 2] - s[i - 1] >= v {
				j --
			}
			t := 0
			if s[j - 1] - s[i - 1] >= v {
				t = 1
			}
			suf[i] = suf[j] + t
		}
		ans := math.MinInt
		for l, r := 0, 1; r <= n; r++ {
			for l <= r && pre[l] < m - suf[r + 1] {
				l++;
			}
			if l <= r {
				ans = max(ans, s[r] - s[l])
			}
		}
		if ans < 0 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
