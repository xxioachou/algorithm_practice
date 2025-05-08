package main

import (
	"math"
)

func maxValue(nums []int, k int) int {
	const P = 1 << 7
	n := len(nums)
	f := make([][][]int, n + 2)
	g := make([][][]int, n + 2)
	for i := 0; i <= n + 1; i++ {
		f[i] = make([][]int, n + 2)
		g[i] = make([][]int, n + 2)
		for j := 0; j <= n + 1; j++ {
			f[i][j] = make([]int, P)
			g[i][j] = make([]int, P)
		}
	}
	f[0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			for x := 0; x < P; x++ {
				f[i + 1][j][x] |= f[i][j][x];
				f[i + 1][j + 1][x | nums[i]] |= f[i][j][x];
			}
		}
	}
	g[n + 1][0][0] = 1
	for i := n + 1; i > 1; i-- {
		for j := 0; j <= n - i + 1; j++ {
			for x := 0; x < P; x++ {
				g[i - 1][j][x] |= g[i][j][x];
				g[i - 1][j + 1][x | nums[i - 2]] |= g[i][j][x];
			}
		}
	}

	ans := math.MinInt
	for i := k; i <= n - k; i++ {
		for j, x := range f[i][k] {
			if x == 1 {
				for t, y := range g[i + 1][k] {
					if y == 1 {
						ans = max(ans, j ^ t)
					}
				}
			}
		}
	}
	return ans
}

func main() {

}
