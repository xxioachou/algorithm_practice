package main

import "fmt"

func removeDuplicates(nums []int) int {
    x := -10004
	cnt := 0
	i := 0
	for _, y := range nums {
		if x != y {
			x = y
			cnt = 0
		}
		cnt ++
		if cnt <= 2 {
			nums[i] = x
			i ++
		}
	}
	return i
}

func main() {

}