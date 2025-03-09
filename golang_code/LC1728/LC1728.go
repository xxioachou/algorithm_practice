package main

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	const MOUSE_TURN = 0
	const CAT_TURN = 1
	const MOUSE_WIN = 0
	const CAT_WIN = 1
	const MAX_TURNS = 100

	foodX, foodY := -1, -1
	mouseX, mouseY := -1, -1
	catX, catY := -1, -1
	n := len(grid)
	m := len(grid[0])
	dx := []int {-1, 1, 0, 0}
	dy := []int {0, 0, -1, 1}
	f := make([][][][][]int, MAX_TURNS)
	for i := range f {
		f[i] = make([][][][]int, n)
		for j := range f[i] {
			f[i][j] = make([][][]int, m)
			for k := range f[i][j] {
				f[i][j][k] = make([][]int, n)
				for u := range f[i][j][k] {
					f[i][j][k][u] = make([]int, m)
					for v := range f[i][j][k][u] {
						f[i][j][k][u][v] = -1
					}
				}
			}
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'F' {
				foodX, foodY = i, j
			} else if grid[i][j] == 'M' {
				mouseX, mouseY = i, j
			} else if grid[i][j] == 'C' {
				catX, catY = i, j
			}
		}
	}
	valid := func (x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < m && 
		grid[x][y] != '#'
	}

	var dfs func(dep, mx, my, cx, cy int) int
	dfs = func(dep, mx, my, cx, cy int) int {
		if mx == foodX && my == foodY {
			return MOUSE_WIN
		}
		if (cx == foodX && cy == foodY) ||
			(cx == mx && cy == my) ||
			(dep >= MAX_TURNS) {
			return CAT_WIN
		}
		if f[dep][mx][my][cx][cy] != -1 {
			return f[dep][mx][my][cx][cy]
		}

		if dep % 2 == MOUSE_TURN {
			for i := 0; i < 4 && f[dep][mx][my][cx][cy] == -1; i ++ {
				x, y := mx, my
				for j := 0; j <= mouseJump && f[dep][mx][my][cx][cy] == -1; j ++ {
					// 不能跳过墙体或者跳出图
					if !valid(x, y) {
						break
					}
					if dfs(dep + 1, x, y, cx, cy) == MOUSE_WIN {
						f[dep][mx][my][cx][cy] = MOUSE_WIN
					}
					x += dx[i]
					y += dy[i]
				}
			}
			if f[dep][mx][my][cx][cy] == -1 {
				f[dep][mx][my][cx][cy] = CAT_WIN
			}
			// fmt.Println(strings.Repeat(":", dep) + ",", dep, mx, my, cx, cy)
			// fmt.Println(strings.Repeat(":", dep) + "res:", f[dep][mx][my][cx][cy])
			return f[dep][mx][my][cx][cy]
		} 

		for i := 0; i < 4 && f[dep][mx][my][cx][cy] == -1; i ++ {
			x, y := cx, cy
			for j := 0; j <= catJump && f[dep][mx][my][cx][cy] == -1; j ++ {
				// 不能跳过墙体或者跳出图
				if !valid(x, y) {
					break
				}
				if dfs(dep + 1, mx, my, x, y) == CAT_WIN {
					f[dep][mx][my][cx][cy] = CAT_WIN
				}
				x += dx[i]
				y += dy[i]
			}
		}
		if f[dep][mx][my][cx][cy] == -1 {
			f[dep][mx][my][cx][cy] = MOUSE_WIN
		}
		// fmt.Println(strings.Repeat(":", dep) + ",", dep, mx, my, cx, cy)
		// fmt.Println(strings.Repeat(":", dep) + "res:", f[dep][mx][my][cx][cy])
		return f[dep][mx][my][cx][cy]
	}
	return dfs(0, mouseX, mouseY, catX, catY) == MOUSE_WIN
}
