package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T := 1
	// T = readLine()[0]
	for ; T > 0; T-- {	
		s, _ := in.ReadString('\n')
		s = strings.TrimRight(s, "\r\n")
		x, y, z, u, v := 0, 0, 0, 0, 0
		for _, c := range s {
			if c == 'A' || c == 'E' || c == 'I' ||
				c == 'O' || c == 'U' {
				x ++
			} else if c == 'N' {
				u ++
			} else if c == 'G' {
				v ++
			} else if c == 'Y' {
				z ++
			} else {
				y ++
			}
		}
		ans := 0
		for i := 0; i <= z; i++ {
			for j := 0; j <= min(u, v); j++ {
				vow := x + i
				con := y + z - i + u - j + v - j + j
				t := min(vow, con / 2)
				res := 0
				if t * 2 <= j {
					res = t * 5
				} else {
					res = j / 2 * 5 + (t - j / 2) * 3 + (j % 2)
				}
				if ans < res {
					ans = res
				}
			}
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
