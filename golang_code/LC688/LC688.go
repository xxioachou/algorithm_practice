package main


func knightProbability(n int, k int, row int, column int) float64 {
    f := make([][][]float64, k + 1)
	for i := 0; i <= k; i++ {
		f[i] = make([][]float64, n * 2)
		for j := 0; j < n * 2; j++ {
			f[i][j] = make([]float64, n * 2)
		}
	}
	dx := []int{-1, -2, -2, -1, 1, 2, 2, 1}
	dy := []int{-2, -1, 1, 2, 2, 1, -1, -2}
	// [-n + 1, n) + n
	f[0][n][n] = 1.0
	lr, rr := -row + n, n -row + n
	lc, rc := -column + n, n - column + n
	for i := 1; i <= k; i++ {
		for j := 0; j < n * 2; j++ {
			for u := 0; u < n * 2; u++ {
				for v := 0; v < 8; v++ {
					x := j - dx[v]
					y := u - dy[v]
					if x < lr || x >= rr || y < lc || y >= rc {
						continue
					}
					f[i][j][u] += f[i - 1][x][y] * 1.0 / 8;
				}
			}
		}
	}
	ans := 0.0
	for i := lr; i < rr; i++ {
		for j := lc; j < rc; j++ {
			ans += f[k][i][j]
		}
	}
	return ans
}

func main() {

}