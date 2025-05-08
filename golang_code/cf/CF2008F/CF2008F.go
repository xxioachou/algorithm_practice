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
	const mod = 1000000007
	qpow := func(a, b int) int {
		res := 1 % mod
		for b > 0 {
			if (b & 1) == 1 {
				res = res * a % mod
			}
			b >>= 1
			a = a * a % mod
		}
		return res
	}
	for ; T > 0; T-- {
		var n int
		Fscan(in, &n)
		a := make([]int, n + 1)
		s := make([]int, n + 1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			s[i] = (s[i - 1] + a[i]) % mod
		}
		ans := 0
		down := n * (n - 1) / 2 % mod
		for i := 1; i <= n; i++ {
			ans = (ans + a[i] * s[i - 1] % mod) % mod
		}
		ans = ans * qpow(down, mod - 2) % mod
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
