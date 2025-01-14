package main

import "slices"

func minOperations(nums []int, k int) int {
    slices.Sort(nums)
	for i := 0; i < len(nums); {
		if nums[i] >= k {
			return  i
		}

		j := i + 1
		for j < len(nums) && nums[j] == nums[i] {
			j ++
		}
		i = j
	}
	return -1
}

func main() {

}
