package main

import "math"

func minimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}
	
	sums := make([]int, 6)
	res := 0
	ans := math.MaxInt
    for i, j := 0, 0; i < len(nums); i ++ {
		for j < len(nums) && res < k {
			for b := 5; b >= 0; b -- {
				if (nums[j] >> b & 1) == 1 {
					if sums[b] == 0 {
						res |= 1 << b
					}
					sums[b] ++
				}
			}
			j ++
		}
		if res >= k {
			ans = min(ans, j - i)
		}
		for b := 5; b >= 0; b -- {
			if (nums[i] >> b & 1) == 1 {
				if sums[b] == 1 {
					res ^= 1 << b;
				}
				sums[b] --
			}
		}
	}
	if ans >= math.MaxInt {
		ans = -1
	}
	return ans
}

func main() {

}
