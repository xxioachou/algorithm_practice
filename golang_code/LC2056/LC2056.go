package main

type pair struct {
	dx int
	dy int
	step int
}

func check(sch []pair, positions [][]int) bool {
	n := len(positions)
	lo := make([][]int, n)
	for i := range(positions) {
		lo[i] = make([]int, len(positions[i]))
		copy(lo[i], positions[i])	
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if (lo[j][0] == lo[k][0] && lo[j][1] == lo[k][1]) {
					return false
				}
			}
		}
		for j := 0; j < n; j++ {
			dx, dy, step := sch[j].dx, sch[j].dy, sch[j].step
			if i + 1 > step {
				continue
			}
			lo[j][0] += dx
			lo[j][1] += dy
		}
	}
	return true
}

func countCombinations(pieces []string, positions [][]int) int {
	n := len(pieces)
	t := make([]int, n)
	for i, x := range(pieces) {
		if x == "rook" {
			t[i] = 0
		} else if x == "queen" {
			t[i] = 1
		} else {
			t[i] = 2
		}
	}
	sch := make([]pair, n)
	dx := [3][]int {
		{-1, 1, 0, 0},					// 上下左右
		{-1, -1, -1, 0, 1, 1, 1, 0},	// 左上、正上、右上、正右、右下、正下、左下、正左
		{-1, -1, 1, 1}}					// 左上、右上、右下、左下
	dy := [3][]int {
		{0, 0, -1, 1},
		{-1, 0, 1, 1, 1, 0, -1, -1},
		{-1, 1, 1, -1}}
	var dfs func(int) int
	dfs = func(dep int) int {
		if dep == n {
			if (check(sch, positions)) {
				return 1
			}
			return 0
		}

		x, y := positions[dep][0], positions[dep][1]
		res := 0
		for i := 0; i < len(dx[t[dep]]); i++ {		// 枚举方向
			dirx := dx[t[dep]][i]
			diry := dy[t[dep]][i]
			j := 1
			if i == 0 {
				j = 0
			}
			for ; j < 8; j++ {				// 枚举步数
				if x + dirx * j < 1 || x + dirx * j > 8 || 
					y + diry * j < 1 || y + diry * j > 8 {
						break
				}
				sch[dep] = pair{dirx, diry, j}
				res += dfs(dep + 1)
			}
		}
		return res
	}
	return dfs(0)
} 

func main() {

}