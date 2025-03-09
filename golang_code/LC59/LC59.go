package main

func generateMatrix(n int) [][]int {
    r1, r2 := 0, n - 1
	c1, c2 := 0, n - 1
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for x := 1; r1 < r2 && c1 < c2; {
		for i := c1; i <= c2; i ++ {
			ans[r1][i] = x
			x ++
		}
		for j := r1 + 1; j + 1 <= r2; j ++ {
			ans[j][c2] = x
			x ++
		}
		for i := c2; i >= c1; i -- {
			ans[r2][i] = x
			x ++
		}
		for j := r2 - 1; j > r1; j -- {
			ans[j][c1] = x
			x ++
		}
		r1, r2 = r1 + 1, r2 - 1
		c1, c2 = c1 + 1, c2 - 1
	}
	if n % 2 == 1 {
		ans[n / 2][n / 2] = n * n
	}
	return ans
}

func main() {

}