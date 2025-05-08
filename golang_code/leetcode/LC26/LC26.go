package main

func removeDuplicates(nums []int) int {
    left, right := 0, 1
	for right < len(nums) {
		if nums[right] != nums[left] {
			nums[left + 1] = nums[right]
			left ++
		}
		right ++
	}
	return left + 1
}

func main() {

}