package main


func checkPartitioning(s string) bool {
	n := len(s)
	f := make([][]bool, n + 1)
	g := make([][]bool, n + 1)
	for i := range f {
		f[i] = make([]bool, 4)
		for j := range f[i] {
			f[i][j] = false
		}
	}
	for i := range g {
		g[i] = make([]bool, n + 1)
	}
	for i := n; i > 0; i -- {
		for j := 1; j <= n; j ++ {
			if i >= j {
				g[i][j] = true
			} else {
				g[i][j] = g[i + 1][j - 1] && (s[i - 1] == s[j - 1])
			}
		}
	}
	f[0][0] = true
	for i := 1; i <= n; i ++ {
		for j := 1; j <= i && j <= 3; j ++ {
			for k := 1; k <= i; k ++ {
				f[i][j] = f[i][j] || (f[i - k][j - 1] && g[i - k + 1][i])
			}
		}		
	}
	return f[n][3]
}
