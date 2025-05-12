package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(_in io.Reader, _out io.Writer) {
	in := bufio.NewReader(_in)
	out := bufio.NewWriter(_out)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}
	ans := 0
	// [0, l - 1] [l, r] [r + 1, n]
	for r, l, pre := n, 1, 0; r >= 0; r-- {
		for l <= r && pre < s[n]-s[r] {
			pre = s[l]
			l++
		}
		if pre == s[n]-s[r] {
			// Fprintln(out, l-1, r+1, pre)
			ans = pre
		}
	}

	Fprintln(out, ans)
}

func main() { solve(os.Stdin, os.Stdout) }
