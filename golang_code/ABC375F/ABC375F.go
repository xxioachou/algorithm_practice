package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	x int
	y int
}

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	readLine := func() []int {
		str, _ := in.ReadString('\n')
		parts := strings.Fields(str)
		nums := make([]int, len(parts))
		for i := range parts {
			nums[i], _ = strconv.Atoi(parts[i])
		}
		return nums
	}
	
	const inf = 1e15
	T := 1
	// T = readLine()[0]
	for ; T > 0; T-- {
		t := readLine()
		N, M, Q := t[0], t[1], t[2]
		d := make([][]int, N + 1)
		e := make([][]pair, N + 1)
		A := make([]int, M + 1)
		B := make([]int, M + 1)
		C := make([]int, M + 1)
		for i := 1; i <= N; i++ {
			d[i] = make([]int, N + 1)
			e[i] = make([]pair, N + 1)
		}
		for i := 1; i <= M; i++ {
			t = readLine()
			a, b, c := t[0], t[1], t[2]
			e[a][b] = pair{c, i}
			e[b][a] = pair{c, i}
			A[i], B[i], C[i] = a, b, c
		}
		ban := make([]bool, M + 1)
		op := make([]int, Q + 1)
		x := make([]int, Q + 1)
		y := make([]int, Q + 1)
		for i := 1; i <= Q; i++ {
			t = readLine()
			op[i] = t[0]
			if op[i] == 1 {
				x[i] = t[1]
				ban[x[i]] = true
			} else {
				x[i], y[i] = t[1], t[2]
			}
		}

		for i := 1; i <= N; i++ {
			d[i][i] = 0
		}
		for i := 1; i <= N; i++ {
			for j := i + 1; j <= N; j++ {
				w, idx := e[i][j].x, e[i][j].y
				d[i][j], d[j][i] = inf, inf
				if idx > 0 && !ban[idx] {
					d[i][j], d[j][i] = w, w
				}
			}
		}
		for k := 1; k <= N; k++ {
			for i := 1; i <= N; i++ {
				for j := 1; j <= N; j++ {
					if d[i][j] > d[i][k] + d[k][j] {
						d[i][j] = d[i][k] + d[k][j]
					}
				}
			}
		}
	 	calc := func(u, v, w int) {
			for i := 1; i <= N; i++ {
				for j := 1; j <= N; j++ {
					if d[i][j] > d[i][u] + d[v][j] + w {
						d[i][j] = d[i][u] + d[v][j] + w
					}
					if d[i][j] > d[i][v] + d[u][j] + w {
						d[i][j] = d[i][v] + d[u][j] + w
					}
				}
			}
		}

		ans := make([]int, Q + 1)
		for i := Q; i > 0; i-- {
			if op[i] == 1 {
				ban[x[i]] = false
				calc(A[x[i]], B[x[i]], C[x[i]])
			} else {
				if d[x[i]][y[i]] >= inf {
					ans[i] = -1
				} else {
					ans[i] = d[x[i]][y[i]]
				}
			}
		}
		for i := 1; i <= Q; i++ {
			if op[i] == 2 {
				Fprintln(out, ans[i])
			}
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }