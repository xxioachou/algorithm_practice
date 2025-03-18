package main

import (
	"fmt"
)

func largestVariance(s string) int {
	ans, n := 0, len(s)
	pre := make([][]int, n + 1)
	for i := range pre {
		pre[i] = make([]int, 26)
	}

	for i, ch := range s {
		for j := range pre[i] {
			pre[i + 1][j] = pre[i][j]
		}
		pre[i + 1][int(ch - 'a')] ++
	}

	for x1 := 0; x1 < 26; x1 ++ {
		if pre[n][x1] == 0 {
			continue
		}
		for x2 := 0; x2 < 26; x2 ++ {
			if pre[n][x2] == 0 {
				continue
			}
			for l, r, minv := 0, 1, 0; r <= n; r ++ {
				for l + 1 < r && pre[l + 1][x1] < pre[r][x1] && pre[l + 1][x2] < pre[r][x2] {
					l ++
					minv = min(minv, pre[l][x1] - pre[l][x2])
				}
				if pre[l][x1] < pre[r][x1] && pre[l][x2] < pre[r][x2] {
					// fmt.Printf("x1 %v, x2 %v, l %v, r %v, minv %v\n", x1, x2, l, r, minv)
					ans = max(ans, pre[r][x1] - pre[r][x2] - minv)
				}
			}
		}
	}
	return ans
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Printf("%v\n", largestVariance(s))
}