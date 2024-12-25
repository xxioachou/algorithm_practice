package main

import "fmt"

func networkDelayTime(times [][]int, n int, k int) int {
	dis := make([]int, n)
	vis := make([]int, n)
	e := make([][][]int, n)
	for i := 0; i < n; i++ {
		e[i] = make([][]int, 0)
	}

	for _, arr := range times {
		u, v, w := arr[0], arr[1], arr[2]
		fmt.Println(u, v, w)
		e[u-1] = append(e[u-1], []int{v - 1, w})
	}
	const inf = 0x3f3f3f3f
	for i := 0; i < n; i++ {
		dis[i] = inf
	}
	dis[k-1] = 0
	for i := 0; i < n; i++ {
		var t int
		for j := 0; j < n; j++ {
			if vis[j] == 0 && (t == -1 || dis[t] > dis[j]) {
				t = j
			}
		}
		vis[t] = 1
		fmt.Println(t)
		for _, arr := range e[t] {
			v, w := arr[0], arr[1]
			fmt.Println(t+1, v+1)
			dis[v] = min(dis[v], dis[t]+w)
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Println(i, dis[i])
		ans = max(ans, dis[i])
	}
	if ans >= inf {
		ans = -1
	}
	return ans
}

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	times := make([][]int, 0)
	for i := 1; i <= m; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)
		times = append(times, []int{u, v, w})
	}
	fmt.Println(networkDelayTime(times, n, k))
}
