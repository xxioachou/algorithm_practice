package main

import "math"

func sumOfBeauties(nums []int) int {
	n := len(nums)
	suf := make([]int, n + 1)
	suf[n] = math.MaxInt
	for i := n - 1; i >= 0; i -- {
		suf[i] = min(nums[i], suf[i + 1])
	}
	pre := nums[0]
	ans := 0
	for i := 1; i + 1 < n; i ++ {
		if pre < nums[i] && nums[i] < suf[i + 1] {
			ans += 2
		} else if nums[i - 1] < nums[i] && nums[i] < nums[i + 1] {
			ans ++
		}
		pre = max(pre, nums[i])
	}
	return ans
}