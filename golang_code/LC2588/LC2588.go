package main

func beautifulSubarrays(nums []int) int64 {
	mp := make(map[int]int)
	mp[0] ++
	sum, ans := 0, 0
	for _, x := range nums {
		sum ^= x
		if v, ok := mp[sum]; ok {
			ans += v
		}
		mp[sum] ++
	}
	return int64(ans)
}
