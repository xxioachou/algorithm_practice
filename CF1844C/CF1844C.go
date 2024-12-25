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
	for Fscan(in, &T); T > 0; T-- {
		n := 1
		Fscan(in, &n)
		c := make([]int64, n)
		ans := int64(math.MinInt64)
		for i := 0; i < n; i++ {
			Fscan(in, &c[i])
			ans = max(ans, c[i])
		}

		sum := int64(0)
		ok := false
		for i := 0; i < n; i += 2 {
			sum += max(c[i], 0)
			if c[i] >= 0 {
				ok = true
			}
		}
		if ok {
			ans = max(ans, sum)
		}

		ok = false
		sum = int64(0)
		for i := 1; i < n; i += 2 {
			sum += max(c[i], 0)
			if c[i] >= 0 {
				ok = true
			}
		}
		if ok {
			ans = max(ans, sum)
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
