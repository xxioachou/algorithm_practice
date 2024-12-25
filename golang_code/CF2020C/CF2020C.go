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
	for ; T > 0; T-- {
		var b, c, d int
		Fscan(in, &b, &c, &d)
		ans := 0
		for i := 0; i < 60; i++ {
			x := b >> i & 1
			y := c >> i & 1
			z := d >> i & 1
			if (x ^ y) == 1 && (y ^ z) == 0 {
				ans = -1
				break
			}
			ans |= (y ^ z) << i
		}
		Fprintln(out, ans)
		// if (ans | b) - (ans & c) != d {
		// 	Fprintln(out, "NO")
		// } else {
		// 	Fprintln(out, "YES")
		// }
	}
}

func main() { solve(os.Stdin, os.Stdout) }
