package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T, n := 1, 1
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n + 1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}

		ans := 1
		if n % 2 == 0 {
			for i := 1; i + 1 <= n; i += 2 {
				ans = max(ans, a[i + 1] - a[i])
			}
		} else {
			if n == 1 {
				ans = 1
			} else if n == 3 {
				ans = min(a[2] - a[1], a[3] - a[2])
			} else {
				pre, suf := make([]int, n + 2), make([]int, n + 2)
				for i := 1; i <= n; i ++ {
					pre[i] = pre[i - 1]
					if (i % 2 == 0) {
						pre[i] = max(pre[i], a[i] - a[i - 1])
					}
				}
				for i := n; i >= 1; i-- {
					suf[i] = suf[i + 1]
					if (i % 2 == 0) {
						suf[i] = max(suf[i], a[i + 1] - a[i])
					}
				}
				ans = int(2e18)
				for i := 1; i <= n; i += 2 {
					// Fprintln(out, i, pre[i - 1], suf[i + 1])
					ans = min(ans, max(pre[i - 1], suf[i + 1]));
				}
			}
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
