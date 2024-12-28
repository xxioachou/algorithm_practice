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
		var H, W, D int
		Fscan(in, &H, &W, &D)
		D ++
		s := make([]string, H)
		Fscanln(in, &s[0])
		for i := range(s) {
			Fscanln(in, &s[i])
		}
		q := make([]pair, 0)
		dis := make([][]int, H)
		for i := 0; i < H; i++ {
			dis[i] = make([]int, W)
		}
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if s[i][j] == 'H' {
					q = append(q, pair{i, j})
					dis[i][j] = 1
				}
			}
		}
		ans := 0
		dx := []int{-1, 1, 0, 0}
		dy := []int{0, 0, -1, 1}
		for len(q) > 0 {
			x, y := q[0].x, q[0].y
			q = q[1:]
			ans ++
			if dis[x][y] + 1 <= D {
				for i := 0; i < 4; i++ {
					a, b := x + dx[i], y + dy[i]
					if a < 0 || a >= H || b < 0 || b >= W ||
					s[a][b] == '#' || dis[a][b] > 0 {
						continue
					}
					dis[a][b] = dis[x][y] + 1
					q = append(q, pair{a, b})
				}
			}
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
