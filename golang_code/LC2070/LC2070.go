package main

import (
	"slices"
)

func maximumBeauty(items [][]int, queries []int) []int {
	m := len(queries)
	idx := make([]int, m)
	ans := make([]int, m)
	for i := range idx {
		idx[i] = i
	}
	// 按照询问的大小，从小到大排序
	slices.SortStableFunc(idx, func(i, j int) int {
		return queries[i] - queries[j]
	})
	// 按照 price 的大小，从小到大排序
	slices.SortStableFunc(items, func(a, b []int) int {
		return a[0] - b[0]
	})
	maxv := 0
	for i, j := 0, 0; i < m; i ++ {
		for j < len(items) && items[j][0] <= queries[idx[i]] {
			maxv = max(maxv, items[j][1])
			j ++
		}
		ans[idx[i]] = maxv
	}
	return ans
}
