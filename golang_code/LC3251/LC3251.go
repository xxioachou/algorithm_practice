package main

func countOfPairs(nums []int) int {
	n := len(nums)
	m := 1000
	const mod = 1000000007
	f := make([]int, m + 1)
	// [0, m] -> [1, m + 1]
	s := make([]int, m + 2)
	for j := 0; j <= nums[0]; j ++ {
		f[j] = 1
	}
	for j := 0; j <= m; j ++ {
		s[j + 1] = s[j]
		if j <= nums[0] {
			s[j + 1] = (s[j + 1] + f[j]) % mod
		}
	}

	for i := 2; i <= n; i++ {

		g := make([]int, m + 1)

		for j := 0; j <= nums[i - 1]; j ++ {
			k := nums[i - 1] - j
			x := nums[i - 2]
			l := max(x - j, k)
			r := min(m, x)
			// [l, r] 的和 -> [l + 1, r + 1] 的和
			sum := 0
			if l <= r {
				sum = ((s[r + 1] - s[l]) % mod + mod) % mod
			}
			g[k] = (g[k] + sum) % mod
		}

		copy(f, g)
		for j := 0; j <= m; j ++ {
			s[j + 1] = s[j]
			if j <= nums[i - 1] {
				s[j + 1] = (s[j + 1] + f[j]) % mod
			}
		}
	}
	ans := 0
	for j := 0; j <= nums[n-1]; j ++ {
		ans = (ans + f[j]) % mod
	}
	return ans
} 

func main() {

}