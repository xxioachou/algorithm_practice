package main

import (
	"bufio"
	. "fmt"
	"io"
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
		H, W, D := 0, 0, 0
		Fscan(in, &H, &W, &D)
		s := make([]string, H)
		Fscanln(in, &s[0])
		for i := range(s) {
			Fscanln(in, &s[i])
			// Fprintln(out, s[i])	
		}
		a := make([]pair, 0)
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if s[i][j] == '.' {
					a = append(a, pair{i, j})
				}
			}
		}
		ans := 0
		for i := range(a) {
			for j := range(a) {
				if i == j {
					continue
				}
				x1, y1, x2, y2 := a[i].x, a[i].y, a[j].x, a[j].y
				res := 0
				for x := 0; x < H; x++ {
					for y := 0; y < W; y++ {
						if s[x][y] != '.' {
							continue
						}
						if abs(x - x1) + abs(y - y1) <= D ||
							abs(x - x2) + abs(y - y2) <= D {
								res ++
						}
					}
				}
				ans = max(ans, res)
			}	
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
