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
		var n int
		Fscan(in, &n)
		a := make([]int, n + 1)
		f := make([]int, n + 1)
		mp := make([]int, n + 1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			mp[i] = math.MinInt
		}
		for i := 1; i <= n; i++ {
			f[i] = max(f[i - 1], mp[a[i]] + i + 1)
			mp[a[i]] = max(mp[a[i]], f[i - 1] - i)
		}
		Fprintln(out, f[n])
	}
}

func main() { solve(os.Stdin, os.Stdout) }
