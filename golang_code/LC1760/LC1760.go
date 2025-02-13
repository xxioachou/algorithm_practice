package main

import "math"

func minimumSize(nums []int, maxOperations int) int {
	maxv := math.MinInt
	for _, v := range nums {
		maxv = max(maxv, v)
	}
	check := func(x int) bool {
		need := 0
		for _, v := range nums {
			need += (v + x - 1) / x - 1
			if need > maxOperations {
				return false
			}
		}
		return true
	}
    low, high := 1, maxv
	for low < high {
		mid := (low + high) / 2
		if check(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}

func main() {

}