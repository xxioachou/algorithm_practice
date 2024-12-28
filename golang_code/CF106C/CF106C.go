package main

import "fmt"

func solve() {
	var n, m, a0, b0, c0, d0 int
	fmt.Scan(&n, &m, &c0, &d0)
	arr := make([][]int, 0)
	arr = append(arr, []int{0, 0, c0, d0})
	for i := 1; i <= m; i++ {
		fmt.Scan(&a0, &b0, &c0, &d0)
		arr = append(arr, []int{a0, b0, c0, d0})
	}
	m = len(arr)
	f := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]int, n+1)
	}
	for i, tmp := range arr {
		a, b, c, d := tmp[0], tmp[1], tmp[2], tmp[3]
		for j := 0; j <= n; j++ {
			for k := 0; k*c <= j && k*b <= a; k++ {
				f[i+1][j] = max(f[i+1][j], f[i][j-k*c]+k*d)
			}
		}
	}
	var ans int
	for i := 0; i <= n; i++ {
		ans = max(ans, f[m][i])
	}
	fmt.Println(ans)
}

func main() {
	t := 1
	// fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		solve()
	}
}
