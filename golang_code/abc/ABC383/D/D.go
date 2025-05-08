package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

type pair struct {
	x int
	y int
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(_r io.Reader, out io.Writer) {
	T := 1
	in := bufio.NewReader(_r)
	// Fscan(in, &T)
	for ; T > 0; T-- {
		var N int
		Fscan(in, &N)
		// Fprintln(out, N)
		M := int(math.Floor(math.Sqrt(float64(N)))) + 10
		pri := make([]int, 0)
		vis := make([]bool, M + 1)
		for i := 2; i <= M; i++ {
			if !vis[i] {
				pri = append(pri, i)
			}
			for j := 0; pri[j] <= M / i; j++ {
				vis[pri[j] * i] = true
				if i % pri[j] == 0 {
					break
				}
			}
		}

		ans := 0
		for p := 2; p * p * p * p * p * p * p * p <= N; p++ {
			if !vis[p] {
				ans ++
				// Fprintln(out, ":", p)
			}
		}
		a := make([]int, 0)
		for p := 2; p * p <= N; p++ {
			if !vis[p] {
				a = append(a, p)
			}
		}
		for i := range(a) {
			for j := i + 1; j < len(a); j++ {
				if a[i] * a[i] <= N / (a[j] * a[j]) {
					ans ++
					// Fprintln(out, "::", a[i], a[j])
				} else {
					break
				}
			}
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
