package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int{
	if a < b {
		return b
	}
	return a
}

func solve(_r io.Reader, out io.Writer) {
	T := 1
	in := bufio.NewReader(_r)
	Fscan(in, &T)
	for ; T > 0; T-- {
		var n, q int
		Fscan(in, &n, &q)
		a := make([]int, n + 1)
		s := make([]int, n + 1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			s[i] = s[i - 1] + a[i]
		}
		f := func(i int) int {
			if i == 0 {
				return 0
			}
			
			th := (i - 1) / n + 1
			c := (i - 1) % n + 1
			res := s[n] * (th - 1)
			l, r := th, th + c - 1
			if r <= n {
				res += s[r] - s[l - 1]
			} else {
				res += s[n] - s[l - 1] + s[r - n]
			}
			// Fprintln(out, ":", l, r)
			return res
		}
		for i, l, r := 0, 0, 0; i < q; i++ {
			Fscan(in, &l, &r)
			Fprintln(out, f(r) - f(l - 1))
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }
