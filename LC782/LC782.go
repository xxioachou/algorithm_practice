package main

import (
	"math"
)

func movesToChessboard(board [][]int) int {
    n := len(board)
	a := make([][]byte, 2)
	a[0] = make([]byte, n)
	a[1] = make([]byte, n)
	for i := 0; i < n; i++ {
		a[0][i] = byte(i % 2) + '0'
		a[1][i] = byte(1 - i % 2) + '0'
	}
	// fmt.Println(string(a[0]), string(a[1]))
	end1 := make([]byte, n * n)
	end2 := make([]byte, n * n)
	start := make([]byte, n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			end1[i * n + j] = a[i % 2][j]
			end2[i * n + j] = a[1 - i % 2][j]
			start[i * n + j] = byte(board[i][j]) + '0'
		}
	}
	ed1 := string(end1)
	ed2 := string(end2)

	// fmt.Println(string(end1), string(end2))
	q := make([]string, 0)
	dis := make(map[string]int, 0)
	q = append(q, string(start))
	dis[string(start)] = 0
	for len(q) > 0 {
		t := q[0]
		q = q[1 :]
		d := dis[t]
		if t == ed1 || t == ed2 {
			break
		}
		b := []byte(t)
		// 交换两行
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				// i * n, i * n + n - 1
				// j * n, j * n + n - 1
				for k := 0; k < n; k++ {
					b[i * n + k], b[j * n + k] = b[j * n + k], b[i * n + k]
				}
				tmp := string(b)
				if _, e := dis[tmp]; !e {
					dis[tmp] = d + 1
					q = append(q, tmp)
				}
				for k := 0; k < n; k++ {
					b[i * n + k], b[j * n + k] = b[j * n + k], b[i * n + k]
				}
			}
		}
		// 交换两列
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				// 0 * n + i (n - 1) * n + i
				// 0 * n + j (n - 1) * n + j
				for k := 0; k < n; k++ {
					b[k * n + i], b[k * n + j] = b[k * n + j], b[k * n + i]
				}
				tmp := string(b)
				if _, e := dis[tmp]; !e {
					dis[tmp] = d + 1
					q = append(q, tmp)
				}
				for k := 0; k < n; k++ {
					b[k * n + i], b[k * n + j] = b[k * n + j], b[k * n + i]
				}

			}
		}
	}

	ans := math.MaxInt
	if v, e := dis[ed1]; e {
		ans = min(ans, v)
	}
	if v, e := dis[ed2]; e {
		ans = min(ans, v)
	}
	if ans >= math.MaxInt {
		ans = -1
	}
	return ans
} 

func main() {

}