package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(_r io.Reader, out io.Writer) {
	T := 1
	in := bufio.NewReader(_r)
	// Fscan(in, &T)
	for ; T > 0; T-- {
		N := 1
		Fscan(in, &N)
		res := 0
		for i, last, t, v := 0, 0, 0, 0; i < N; i++ {
			Fscan(in, &t, &v)
			if last > 0 {
				res -= t - last
				if res < 0 {
					res = 0
				}
			}
			// Fprintln(out, last, t, v, res)
			res += v
			last = t
		}
		Fprintln(out, res)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
