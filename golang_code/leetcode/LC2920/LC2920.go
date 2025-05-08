package main

func maximumPoints(edges [][]int, coins []int, k int) int {
	const C = 16
    var dfs func(u, j, fa int) int
	n := len(coins)
	e := make([][]int, n)
	for _, v := range edges {
		e[v[0]] = append(e[v[0]], v[1])
		e[v[1]] = append(e[v[1]], v[0])
	}
	pow := func(x int) int {
		return 1 << min(x, C)
	}
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, C + 1)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}

	dfs = func(u, j, fa int) int {
		if ans[u][j] != -1 {
			return ans[u][j]
		}

		sum1 := coins[u] / pow(j) - k
		sum2 := coins[u] / pow(j + 1)
		for _, v := range e[u] {
			if v == fa {
				continue
			}
			sum1 += dfs(v, j, u)
			sum2 += dfs(v, min(j + 1, C), u)
		}

		ans[u][j] = max(sum1, sum2)
		return ans[u][j]
	}
	return dfs(0, 0, -1)
}

func main() {

}