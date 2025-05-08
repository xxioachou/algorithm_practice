package main

import (
	"bufio"
	"fmt"
	"os"
)

func canCompleteCircuit(gas []int, cost []int) int {
	const K = 18
	n := len(gas)
	s := make([]int, n * 2 + 1)
	p := make([]int, n * 2 + 1)
	f := make([][]int, n * 2 + 1)
	lg := make([]int, n * 2 + 1)
	for i := range f {
		f[i] = make([]int, K)
	}
	for i := 2; i <= n * 2; i ++ {
		lg[i] = lg[i >> 1] + 1
	}
	for i := 0; i < n; i ++ {
		s[i + 1] = s[i] + gas[i]
		p[i + 1] = p[i] + cost[i]
		f[i + 1][0] = s[i + 1] - p[i] - cost[i]
		// fmt.Println(s[i + 1], gas[i], f[i + 1][0])
	}
	for i := 0; i < n; i ++ {
		s[n + i + 1] = s[n + i] + gas[i]
		p[n + i + 1] = p[n + i] + cost[i]
		f[n + i + 1][0] = s[n + i + 1] - p[n + i] - cost[i]
		// fmt.Println(s[n + i + 1], gas[i], f[n + i + 1][0])
	}
	
	for j := 1; j < K; j ++ {
		for i := 1; i + (1 << j) - 1 <= n * 2; i ++ {
			f[i][j] = min(f[i][j - 1], f[i + (1 << (j - 1))][j - 1])
		}
	}
	query := func(l, r int) int {
		k := lg[r - l + 1]
		return min(f[l][k], f[r - (1 << k) + 1][k])
	}
	for i := 1; i <= n; i ++ {
		if query(i, i + n - 1) - s[i - 1] + p[i - 1] >= 0 {
			return i
		}
	}
	return -1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	g := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i ++ {
		fmt.Fscan(in, &g[i])
	}
	for i := range c {
		fmt.Fscan(in, &c[i])
	}

	fmt.Println(g)
	fmt.Println(c)
	fmt.Println(canCompleteCircuit(g, c))
}