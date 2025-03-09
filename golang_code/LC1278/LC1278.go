package main

const inf = 0x3f3f3f3f

func palindromePartition(s string, k int) int {
	n := len(s)
	c := make([][]int, n + 1)
	f := make([][]int, n + 1)
	for i := range c {
		c[i] = make([]int, n + 1)
		f[i] = make([]int, n + 1)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	for i := n; i > 0; i -- {
		for j := 1; j <= n; j ++ {
			if i >= j {
				c[i][j] = 0
			} else {
				val := 0
				if s[i - 1] != s[j - 1] {
					val = 1
				}
				c[i][j] = c[i + 1][j - 1] + val
			}
		}
	}
	f[0][0] = 0
	for i := 1; i <= n; i ++ {
		for j := 1; j <= i && j <= k; j ++ {
			for u := 1; u <= i; u ++ {
				f[i][j] = min(f[i][j], f[i - u][j - 1] + c[i - u + 1][i])
			}
		}
	}
	return f[n][k]
}
