package main

func findBall(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	ans := make([]int, n)
	valid := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

    for i := 0; i < n; i ++ {
		x, y := 0, i
		for x < m {
			t := grid[x][y]
			if !valid(x, y + t) || grid[x][y + t] != t {
				ans[i] = -1
				break
			}
			x, y = x + 1, y + t
		}
		if ans[i] != -1 {
			ans[i] = y
		}
	}
	return ans
}

func main() {

}