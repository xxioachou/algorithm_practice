package main

func countOfPairs(nums []int) int {
	n := len(nums)
	m := 50
	f := make([][]int, m + 1)
	s := make([]int, m + 2)
	// [0, m] -> [1, m + 1]
	const mod = 1000000007
	for i := 0; i <= m; i ++ {
		f[i] = make([]int, m + 1)
	}
	for j := 0; j <= nums[0]; j ++ {
		f[j][nums[0] - j] = 1
	}
	for j := 0; j <= m; j ++ {
		s[j + 1] = s[j]
		if j <= nums[0] {
			s[j + 1] = (s[j + 1] + f[nums[0] - j][j]) % mod
		}
	}

	for i := 2; i <= n; i++ {

		g := make([][]int, m + 1)
		for j := 0; j <= m; j ++ {
			g[j] = make([]int, m + 1)
		}

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
			g[j][k] = (g[j][k] + sum) % mod
			// for u := 0; u <= nums[i-2] && u <= j && nums[i-2]-u >= k; u++ {
			// 	v := nums[i-2] - u
			// 	g[j][k] = (g[j][k] + f[u][v]) % mod
			// }
		}

		copy(f, g)
		for j := 0; j <= m; j ++ {
			s[j + 1] = s[j]
			if j <= nums[i - 1] {
				s[j + 1] = (s[j + 1] + f[nums[i - 1] - j][j]) % mod
			}
		}
	}
	ans := 0
	for j := 0; j <= nums[n-1]; j ++ {
		ans = (ans + f[j][nums[n - 1] - j]) % mod
	}
	return ans
}

func main() {

}
