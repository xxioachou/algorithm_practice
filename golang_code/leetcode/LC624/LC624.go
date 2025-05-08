package main

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxDistance(arrays [][]int) int {
	maxv1, maxv2 := math.MinInt, math.MinInt
	idx1 := -1
	for i := range arrays {
		n := len(arrays[i])
		x := arrays[i][n - 1]
		if maxv1 <= x {
			maxv2 = maxv1
			maxv1 = x
			idx1 = i
		} else if maxv2 < x {
			maxv2 = x
		}
	}
	ans := math.MinInt
	for i := range arrays {
		x := arrays[i][0]
		if i == idx1 {
			ans = max(ans, abs(x - maxv2))
		} else {
			ans = max(ans, abs(x - maxv1))
		}
	}
	return ans
}