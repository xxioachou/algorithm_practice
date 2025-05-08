package main

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func countPairs(nums []int, k int) int {
	const N = 110
	cnt := make([][]int, N)
	ans := 0
	for i, x := range nums {
		g := gcd(k, i)
		t := k / g
		for j := 0; j < i; j += t {
			if len(cnt[j]) == 0 {
				continue
			}
			ans += cnt[j][x]
		}
		if len(cnt[i]) == 0 {
			cnt[i] = make([]int, N)
		}
		cnt[i][x]++
	}
	return ans
}
