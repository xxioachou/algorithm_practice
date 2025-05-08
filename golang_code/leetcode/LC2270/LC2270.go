package main

func waysToSplitArray(nums []int) int {
    sum := 0
	for _, v := range nums {
		sum += v
	}
	cnt := 0
	for i, t := 0, 0; i < len(nums) - 1; i ++ {
		t += nums[i]
		if t >= sum - t {
			cnt ++
		}
	}
	return cnt
}

func main() {

}
