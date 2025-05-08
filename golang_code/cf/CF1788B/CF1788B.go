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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	
	T := 1
	Fscan(in, &T)
	for ; T > 0; T -- {
		var n int
		Fscan(in, &n)
		x, y := 0, 0
		for t, f := 1, true; n > 0; t *= 10 {
			r := n % 10
			if r % 2 == 1 {
				if f {
					x += r / 2 * t
					y += (r - r / 2) * t
				} else {
					x += (r - r / 2) * t
					y += r / 2 * t
				}
				f = !f
			} else {
				x += r / 2 * t
				y += r / 2 * t
			}
			n /= 10
		}
		Fprintln(out, x, y)
	}
}

func main() { solve(os.Stdin, os.Stdout) }