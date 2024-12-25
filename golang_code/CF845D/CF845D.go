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
		for i := range(t) {
			Fscan(in, &t[i])
			if t[i] == 1 || t[i] == 3 {
				Fscan(in, &s[i])
			}
		}
		cnt6, ans := 0, 0
		cnt := make([]int, 301)
		for i := 0; i < n; i++ {
			if t[i] == 1 {
				for j := 1; j < s[i]; j++ {
					ans += cnt[j]
				}
			}
			if t[i] == 2 {
				ans += cnt6
			}
			if t[i] == 3 {
				for j := range(cnt) {
					cnt[j] = 0
				}
				cnt[s[i]] ++
			}
			if t[i] == 4 {
				cnt6 = 0
			}
			if t[i] == 5 {
				for j := range(cnt) {
					cnt[j] = 0
				}
			}
			if t[i] == 6 {
				cnt6 ++
			}
		}
		for j := range(cnt) {
			cnt[j] = 0
		}
		for i := n - 1; i >= 0; i-- {
			if t[i] == 1 {
				for j := 1; j < s[i]; j++ {
					ans += cnt[j]
				}
			}
			if t[i] == 3 {
				cnt[s[i]] ++
			}
			if t[i] == 5 {
				for j := range(cnt) {
					cnt[j] = 0
				}
			}
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
