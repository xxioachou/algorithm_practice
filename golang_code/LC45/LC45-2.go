package main

func jump(nums []int) int {
	n := len(nums)
	curRight := 0
	nextRight := 0
	ans := 0
	for i := 0; i < n - 1; i ++ {
		nextRight = max(nextRight, i + nums[i])
		if i == curRight {
			ans ++
			curRight = nextRight
		}
	}
	return ans
}

func main() {

}