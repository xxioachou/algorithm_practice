package main

import (
	"slices"
)

func getKth(lo int, hi int, k int) int {
	ans := make(map[int]int)
	ans[1] = 0
	var dfs func(int) int
	dfs = func(x int) int {
		if _, ok := ans[x]; ok {
			return ans[x]
		}

		if x % 2 == 1 {
			ans[x] = dfs(x * 3 + 1) + 1
		} else {
			ans[x] = dfs(x / 2) + 1
		}
		return ans[x]
	}
	idx := make([]int, hi - lo + 1)
	for i := 0; i < hi - lo + 1; i++ {
		idx[i] = lo + i
		// fmt.Println(dfs(idx[i]))
	}
	// fmt.Println("")
	slices.SortFunc(idx, func(i, j int) int {
		if dfs(i) == dfs(j) {
			return i - j
		}
		return dfs(i) - dfs(j)
	})
	return idx[k - 1]
} 

func main() {

}