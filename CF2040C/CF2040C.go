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
	cmp := func(n, k int) bool {
		t := int(math.Ceil(math.Log2(float64(k))))
		return t <= n - 1
	}
	for ; T > 0; T-- {
		var n, k int
		Fscan(in, &n, &k)
		if !cmp(n, k) {
			Fprintln(out, -1)
			continue
		}
		l, r := 0, n - 1
		ans := make([]int, n)
		m := n
		for i := 1; m > 1; i, m = i + 1, m - 1{
			if cmp(m - 1, k) {
				ans[l] = i
				l ++
			} else {
				k -= 1 << (m - 2)
				ans[r] = i
				r --
			}
		}
		ans[l] = n
		for i := range(ans) {
			Fprint(out, ans[i], " ")
		}
		Fprintln(out, "")
	}
}

func main() { solve(os.Stdin, os.Stdout) }
