package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

func solve(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T := 1
	Fscan(in, &T)
	for ; T > 0; T-- {
		var n int
		Fscan(in, &n)
		a := make([]int, n + 1)
		b := make([]int, n + 1)
		p := make([]int, n + 1)
		s := make([]int, n + 1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		for i := 1; i + 1 <= n; i++ {
			b[i] = gcd(a[i], a[i + 1])
		}
		b[n] = 0x3f3f3f3f
		for i := 1; i < n; i++ {
			p[i] = p[i - 1]
			if b[i] >= b[i - 1] {
				p[i] ++
			}
		}
		for i := n - 1; i >= 1; i-- {
			s[i] = s[i + 1]
			if (b[i] <= b[i + 1]) {
				s[i] ++
			}
		}
		ok := false
		if s[2] >= n - 2 {
			ok = true
		}
		if p[n - 2] >= n - 2 {
			ok = true
		}
		for i := 2; i < n; i++ {
			t := gcd(a[i - 1], a[i + 1])
			if p[i - 2] >= i - 2 && s[i + 1] >= n - (i + 1) && 
				b[i - 2] <= t && t <= b[i + 1] {
				ok = true
				break
			}
		}
		if ok {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }
