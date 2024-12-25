package main

func totalNQueens(n int) int {
	var dfs func(int)
	col := make([]int, n)
	diag1 := make([]int, n * 2)
	diag2 := make([]int, n * 2)
	ans := 0
	dfs = func(dep int) {
		if dep == n {
			ans += 1
			return
		}
		for i := 0; i < n; i++ {
			if col[i] != 0 || diag1[dep + i] != 0 || 
				diag2[dep - i + n] != 0 {
					continue
			}
			col[i], diag1[dep + i], diag2[dep - i + n] = 1, 1, 1
			dfs(dep + 1)
			col[i], diag1[dep + i], diag2[dep - i + n] = 0, 0, 0
		}
	}
	dfs(0)
	return ans
}

func main() {

}