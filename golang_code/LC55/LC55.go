package main

func canJump(nums []int) bool {
	pos := 0
	for i := range nums {
		if pos < i {
			break
		}
		pos = max(pos, i + nums[i])
	}
	return pos >= len(nums) - 1
}

func main() {

}