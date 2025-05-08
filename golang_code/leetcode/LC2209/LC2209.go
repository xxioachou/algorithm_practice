package main

import (
	"fmt"
)

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	const inf = 0x3f3f3f3f
	n := len(floor)
	s := make([]int, n + 1)
	for i := range floor {
		x := 0
		if floor[i] == '1' {
			x ++
		}
		s[i + 1] = s[i] + x
	}
	f := make([][]int, n + 1)
	que := make([][]int, numCarpets + 1)
	for i := range f {
		f[i] = make([]int, numCarpets + 1)
	}
	for j := range f[0] {
		f[0][j] = -inf
	}

	f[0][0] = 0
	que[0] = append(que[0], 0)
	for i := 1; i <= n; i ++ {
		for j := 0; j <= i && j <= numCarpets; j ++ {
			f[i][j] = f[i - 1][j]
			if j > 0 {
				for len(que[j - 1]) > 0 && que[j - 1][0] < i - carpetLen {
					que[j - 1] = que[j - 1][1:]
				}
				k := que[j - 1][0]
				f[i][j] = max(f[i][j], f[k][j - 1] - s[k] + s[i])
				// for k := max(i - carpetLen, 0); k < i; k ++ {
					// 	f[i][j] = max(f[i][j], f[k][j - 1] + s[i] - s[k])
					// }
				// fmt.Printf("[i:%d, j:%d, k: %d] %d\n", i, j, k, f[i][j])
			} else {
				// fmt.Printf("[i:%d, j:%d] %d\n", i, j, f[i][j])
			}
		}
		for j := 0; j <= i && j <= numCarpets; j ++ {
			for len(que[j]) > 0 {
				k := que[j][len(que[j]) - 1]
				if f[k][j] - s[k] <= f[i][j] - s[i] {
					que[j] = que[j][:len(que[j]) - 1]
				} else {
					break
				}
			}
			que[j] = append(que[j], i)
		}
	}
	ans := 0
	for j := 0; j <= numCarpets; j ++ {
		ans = max(ans, f[n][j])
	}
	return s[n] - ans
}

func main() {
	var f string
	var nc, len int
	fmt.Scan(&f, &nc, &len)
	fmt.Println(minimumWhiteTiles(f, nc, len))
}