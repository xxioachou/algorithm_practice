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
		n, q := 1, 1
		Fscan(in, &n)
		a := make([]int, n + 1)
		s := make([]int, n + 1)
		ss := make([]int, n + 1)
		sb := make([]int, n + 1)
		sum := make([]int, n + 1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			s[i] = s[i - 1] + a[i]
			ss[i] = ss[i - 1] + s[i]
			sb[i] = sb[i - 1] + n - i + 1
		}
		for i := 1; i <= n; i++ {
			sum[i] = sum[i - 1] + (ss[n] - ss[i - 1]) - s[i - 1] * (n - i + 1);
		}
		Fscan(in, &q)
		get := func (x int) int {
			l, r := 0, n
			for l < r {
				mid := (l + r + 1) / 2
				if (sb[mid] < x) {
					l = mid
				} else {
					r = mid - 1
				}
			}
			return r
		}

		for l, r := 0, 0; q > 0; q-- {
			Fscan(in, &l, &r)
			i, j := get(l), get(r)
			ans := 0
			// Fprintln(out, i, j)
			if i == j {
				l = i + l - sb[i]
				r = i + r - sb[i]
				ans = ss[r] - ss[l - 1] - s[i] * (r - l + 1)
				// Fprintln(out, ":", l, r)
			} else {
				l1, r1 := i + l - sb[i], n
				l2, r2 := j + 1, j + r - sb[j]
				// Fprintln(out, "::", l1, r1, l2, r2)

				ans = sum[j] - sum[i + 1];
				ans += ss[r1] - ss[l1 - 1] - s[i] * (r1 - l1 + 1)
				ans += ss[r2] - ss[l2 - 1] - s[j] * (r2 - l2 + 1)
			}
			Fprintln(out, ans)
		}		
	}
}

func main() { solve(os.Stdin, os.Stdout) }
