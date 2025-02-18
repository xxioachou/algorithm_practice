package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
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
	
	T := 1
	Fscan(in, &T)
	for ; T > 0; T -- {
		var n int
		var str string
		Fscan(in, &n, &str)
		idx := []int{-1, -1, -1, -1}
		chs := []byte{'N', 'S', 'E', 'W'}
		ans := make([]byte, n)
		
		turn := true
		for i, ch := range str {
			v := slices.Index(chs, byte(ch))
			if idx[v ^ 1] != -1 {
				j := idx[v ^ 1]
				if turn {
					ans[j], ans[i] = 'R', 'R'
				} else {
					ans[j], ans[i] = 'H', 'H'
				}
				idx[v ^ 1] = -1	
				turn = !turn
			} else if idx[v] != -1 {
				j := idx[v]
				ans[j], ans[i] = 'R', 'H'
				idx[v] = -1
			} else {
				idx[v] = i
			}
		}

		ok := true
		for i := 0; i < 4; i ++ {
			if idx[i] != -1  {
				ok = false
				break
			}
		}
		if slices.Index(ans, 'R') == -1 || slices.Index(ans, 'H') == -1 {
			ok = false
		}
		if ok {
			Fprintln(out, string(ans))
		} else {
			Fprintln(out, "NO")
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }