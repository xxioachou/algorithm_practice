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
		var n int
		Fscan(in, &n)
		t := make([]int, n)
		s := make([]int, n)
		vis := make([]int, n)
		for i := range(t) {
			Fscan(in, &t[i])
			if t[i] == 1 || t[i] == 3 {
				Fscan(in, &s[i])
			}
			if t[i] == 5 {
				s[i] = 301
			}
		}
		cnt6, ans := 0, 0
		for i := 0; i < n; i++ {
			if t[i] == 2 {
				ans += cnt6
			}
			if t[i] == 4 {
				cnt6 = 0
			}
			if t[i] == 6 {
				cnt6 ++
			}
		}
		for i := 0; i < n;  {
			if t[i] != 1 {
				i ++
				continue
			}
			j := i + 1
			for j < n && t[j] != 1 {
				if t[j] % 2 == 1 && s[j] < s[i] {
					vis[j] = 1
				}
				j ++
			}
			i = j
		}
		stk := make([]int, 0)
		for i := 0; i < n; i++ {
			if t[i] % 2 == 0 {
				continue
			}

			if t[i] != 1 {
				stk = append(stk, i)
			} else {
				for len(stk) > 0 && s[stk[len(stk) - 1]] < s[i] {
					vis[stk[len(stk) - 1]] = 1
					stk = stk[:len(stk) - 1]
				}
			}
		}
		for i := 0; i < n; i++ {
			ans += vis[i]
		}

		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
