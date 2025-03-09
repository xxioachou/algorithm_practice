package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func min(x, y int) int {
	if x < y  {
		return x
	}
	return y
}

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	
	const inf = 1e15
	T := 1
	// T = readLine()[0]
	for ; T > 0; T -- {
		var N, M int
		Fscan(in, &N, &M)
		f := make([]int, M + 1)
		for i := range f {
			f[i] = -inf
		}
		f[0] = 0
		for i := 1; i <= N; i ++ {
			var x, y, z int
			Fscan(in, &x, &y, &z)

			for j := min(M, i); j > 0; j -- {
				f[j] = max(f[j], f[j - 1] + abs(x) + abs(y) + abs(z))
			}
		}
		Fprintln(out, f[M])
	}
}

func main() { solve(os.Stdin, os.Stdout) }