package main

import (
	"fmt"
	"slices"
)

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	slices.Sort(nums)
	vis := make([]bool, n)
	tmp := make([]int, 0)
	ans := make([][]int, 0)
	
	var dfs func(dep int)
	dfs = func(dep int) {
		if dep == n {
			t := make([]int, len(tmp))
			copy(t, tmp)
			fmt.Println(tmp)
			ans = append(ans, t)
			return
		}

		for i := 0; i < n; i++ {
			if vis[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i - 1] && !vis[i - 1] {
				continue
			}
			vis[i] = true
			tmp = append(tmp, nums[i])
			dfs(dep + 1)
			vis[i] = false
			tmp = tmp[: len(tmp) - 1]
		}
	}
	dfs(0)
	return ans
}

func main() {

}