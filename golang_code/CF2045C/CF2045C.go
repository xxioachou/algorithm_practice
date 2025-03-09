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
		var S, T string
		Fscan(in, &S, &T)
		v := make([]int, 26)
		for i := range v {
			v[i] = -1
		}
		for i, c := range T {
			if i + 1 < len(T) {
				v[int(c - 'a')] = i
			}
		}
		minv := 0x3f3f3f3f
		l, r := -1, -1
		for i := 0; i + 1 < len(S); i++ {
			val := int(S[i + 1] - 'a')
			if v[val] != -1 {
				if i + 1 + len(T) - v[val] < minv {
					minv = i + 1 + len(T) - v[val]
					l = i
					r = v[val]
				}
			}
		}
		if minv >= 0x3f3f3f3f {
			Fprintln(out, -1)
		} else {
			// Fprintln(out, l + 1, r + 1)
			Fprintln(out, S[:l + 1] + T[r:])
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }
