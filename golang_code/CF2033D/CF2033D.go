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
	Fscan(in, &T)
	for ; T > 0; T-- {
		var n int
		Fscan(in, &n)
		a := make([]int, n + 1)
		s := make([]int, n + 1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			s[i] = s[i - 1] + a[i]
		}
		f := make([]int, n + 1)
		mp := make(map[int]int)
		mp[s[0]] = 0
		for i := 1; i <= n; i++ {
			f[i] = f[i - 1]
			if _, ex := mp[s[i]]; ex {
				f[i] = max(f[i], f[mp[s[i]]] + 1)
			}
			mp[s[i]] = i
		}
		Fprintln(out, f[n])
	}
}

func main() { solve(os.Stdin, os.Stdout) }
