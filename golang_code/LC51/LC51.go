package main

import "strings"

func solveNQueens(n int) [][]string {
	var dfs func(int)
	col := make([]int, n)
	diag1 := make([]int, n * 2)
	diag2 := make([]int, n * 2)
	mp := make([][]byte, n)
	ans := make([][]string, 0)
	for i := 0; i < n; i++ {
		mp[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			mp[i][j] = '.'
		}
	}
	dfs = func(dep int) {
		if dep == n {
			res := make([]string, 0)
			for i := 0; i < n; i++ {
				res = append(res, string(mp[i]))
			}
			ans = append(ans, res)
			return
		}
		for i := 0; i < n; i++ {
			if col[i] != 0 || diag1[dep + i] != 0 || 
				diag2[dep - i + n] != 0 {
					continue
			}
			col[i], diag1[dep + i], diag2[dep - i + n] = 1, 1, 1
			mp[dep][i] = 'Q'
			dfs(dep + 1)
			col[i], diag1[dep + i], diag2[dep - i + n] = 0, 0, 0
			mp[dep][i] = '.'
		}
	}
	dfs(0)
	return ans
}

func main() {

}
