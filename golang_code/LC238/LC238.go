package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
    ans := make([]int, n)

	ans[0] = 1
	for i := 0; i + 1 < n; i ++ {
		ans[i + 1] = ans[i] * nums[i]
	}
	r := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= r
		r *= nums[i]
	}
	return ans
}

func main() {

}