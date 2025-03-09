package main

func combinationSum2(candidates []int, target int) [][]int {
    const N = 31
	cnt := make([]int, N)
	for _, v := range candidates {
		if v < N {
			cnt[v] ++
		}
	}
	var dfs func(u, cur int)
	ans := make([][]int, 0)
	t := make([]int, 0)
	dfs = func(u, cur int) {
		if cur >= target {
			if cur == target {
				ans = append(ans, t)
			}
			return
		}
		if u > target {
			return
		}
		dfs(u + 1, cur)

		for i := 1; i <= cnt[u]; i++ {
			t = append(t, u)
			dfs(u + 1, cur + u)
		}
		for i := 1; i <= cnt[u]; i++ {
			t = t[: len(t) - 1]
		}
	}
	dfs(1, 0)
	return ans
}

func main() {

}