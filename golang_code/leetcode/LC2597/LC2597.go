package main

import "slices"


func beautifulSubsets(nums []int, k int) int {
	var dfs func(dep, cnt int) int
	slices.Sort(nums)
	a := make([]int, 0)
	dfs = func(dep, cnt int) int {
		if dep == len(nums) {
			if cnt > 0 {
				for i, j := 0, 0; i < len(a); i ++ {
					for j < len(a) && a[j] < a[i] + k {
						j ++
					}
					if j < len(a) && a[j] == a[i] + k {
						return 0
					}
				}
				return 1
			}
			return 0
		}

		tmp := 0
		tmp += dfs(dep + 1, cnt)

	
		a = append(a, nums[dep])
		tmp += dfs(dep + 1, cnt + 1)
		a = a[:len(a) - 1]

		return tmp
	}
	return dfs(0, 0)	
}

