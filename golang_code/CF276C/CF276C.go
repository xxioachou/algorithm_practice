package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

func solve(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T := 1
	// Fscan(in, &T)
	for ; T > 0; T-- {
		var n, q int
		Fscan(in, &n, &q)
		a := make([]int, n + 1)
		b := make([]int, n + 2)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		for i, l, r := 0, 0, 0; i < q; i++ {
			Fscan(in, &l, &r)
			b[l] ++
			b[r + 1] --
		}
		for i := 1; i <= n; i++ {
			b[i] += b[i - 1]
		}
		slices.Sort(b[1:n+1])
		slices.Sort(a[1:n+1])
		ans := 0
		for i := n; i >= 1; i-- {
			ans += b[i] * a[i]
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
