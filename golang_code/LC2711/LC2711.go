package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func differenceOfDistinctValues(grid [][]int) [][]int {
	const V = 50;
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for k := range ans {
		ans[k] = make([]int, n)
	}

	for j := n - 1; j >= 0; j -- {
		x, y := 0, j
		l := make([]int, V + 1)
		r := make([]int, V + 1)
		lcnt, rcnt := 0, 0
		for x < m && y < n {
			r[grid[x][y]] ++
			if r[grid[x][y]] == 1 {
				rcnt ++
			}
			x, y = x + 1, y + 1
		}
		x, y = 0, j
		for x < m && y < n {
			r[grid[x][y]] --
			if r[grid[x][y]] == 0 {
				rcnt --
			}

			ans[x][y] = abs(lcnt - rcnt)

			l[grid[x][y]] ++
			if l[grid[x][y]] == 1 {
				lcnt ++
			}
			x, y = x + 1, y + 1
		}
	}
	for i := 1; i < m; i ++ {
		x, y := i, 0
		l := make([]int, V + 1)
		r := make([]int, V + 1)
		lcnt, rcnt := 0, 0
		for x < m && y < n {
			r[grid[x][y]] ++
			if r[grid[x][y]] == 1 {
				rcnt ++
			}
			x, y = x + 1, y + 1
		}

		x, y = i, 0
		for x < m && y < n {
			r[grid[x][y]] --
			if r[grid[x][y]] == 0 {
				rcnt --
			}
			
			ans[x][y] = abs(lcnt - rcnt)

			l[grid[x][y]] ++
			if l[grid[x][y]] == 1 {
				lcnt ++
			}
			x, y = x + 1, y + 1
		}
	}
	return ans
}
