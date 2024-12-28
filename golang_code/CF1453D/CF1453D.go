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
	Fscan(in, &T)
	for ; T > 0; T-- {
		k := 1
		Fscan(in, &k)
		if k % 2 == 1 {
			Fprintln(out, -1)
		} else {
			a := make([]int, 1)
			a[0] = 1
			k -= 2

			for k > 0 {
				if a[len(a) - 1] == 1 {
					x := 0
					for ((1 << ((x + 1) + 2)) - 4) <= k {
						x ++
					}
					for i := 0; i < x; i++ {
						a = append(a, 0)
					}
					if x > 0 {
						k -= (1 << (x + 2)) - 4
					} else {
						a = append(a, 1)
						k -= 2
					}
				} else {
					a = append(a, 1)
					k -= 2
				}
			}

			Fprintln(out, len(a))
			for i:= range(a) {
				Fprint(out, a[i], " ")
			}
			Fprintln(out, "")
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }
