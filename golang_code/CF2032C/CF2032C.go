package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

func solve(_r io.Reader, out io.Writer) {
	T := 1
	in := bufio.NewReader(_r)
	for Fscan(in, &T); T > 0; T-- {
		n := 1
		Fscan(in, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		ans := 0x3f3f3f3f
		for i, j := 0, 1; i+1 < n; i++ {
			for ; j < n && a[i]+a[i+1] > a[j]; j++ {
			}
			ans = min(ans, i+n-j)
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
