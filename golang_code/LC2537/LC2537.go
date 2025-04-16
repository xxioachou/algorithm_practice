package main

func C(x int) int {
	return x * (x - 1) / 2
}

func countGood(nums []int, k int) int64 {
	pairs := 0
	cnt := make(map[int]int)
	ans := 0
	for i, j := 0, 0; i < len(nums); i++ {
		for j < len(nums) && pairs < k {
			pairs = pairs - C(cnt[nums[j]]) + C(cnt[nums[j]]+1)
			cnt[nums[j]]++
			j++
		}
		if pairs >= k {
			ans += len(nums) - j + 1
		}
		pairs = pairs - C(cnt[nums[i]]) + C(cnt[nums[i]]-1)
		cnt[nums[i]]--
	}
	return int64(ans)
}
