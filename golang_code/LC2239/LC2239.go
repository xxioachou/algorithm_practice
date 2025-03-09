package main

import "math"

func findClosestNumber(nums []int) int {
    minv := math.MaxInt
	ans := -1
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for _, v := range nums {
		if minv > abs(v) {
			minv = abs(v)
			ans = v
		} else if minv == abs(v) && ans < v {
			ans = v
		}
	}
	return ans
}

func main() {

}
