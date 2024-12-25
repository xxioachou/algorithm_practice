package main

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	const inf = 0x3f3f3f3f
	ans := make([][][][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([][][]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = make([][]int, m + 1)
			for k := 1; k <= m; k++ {
				ans[i][j][k] = make([]int, n + 1)
				for u := 1; u <= n; u++ {
					ans[i][j][k][u] = -1
				}
			}
		}
	}
    var dfs func(int, int, int, int) int
	dfs = func(x, y, m1, n1 int) int {
		if m1 == 1 && n1 == 1 {
			return 0
		}
		if ans[x][y][m1][n1] >= 0 {
			return ans[x][y][m1][n1]
		}
		res := inf
		for i := x; i < len(horizontalCut) && i - x + 1 < m1; i++ {
			res = min(res, horizontalCut[i] + dfs(x, y, i - x + 1, n1) + dfs(i + 1, y, m1 - (i - x + 1), n1))
		}
		for j := y; j < len(verticalCut) && j - y + 1 < n1; j++ {
			res = min(res, verticalCut[j] + dfs(x, y, m1, j - y + 1) + dfs(x, j + 1, m1, n1 - (j - y + 1)))
		} 
		ans[x][y][m1][n1] = res
		return res
	}
	return dfs(0, 0, m, n)
} 

func main() {

}