package main

import "slices"

func minimumMoney(transactions [][]int) int64 {
    slices.SortFunc(transactions, func(a, b []int) int {
		return (b[0] - b[1]) - (a[0] - a[1])
	})
	maxv := 0
	ans := 0
	for _, v := range transactions {
		maxv = max(maxv, maxv + v[0] - v[1])
	}
	for _, v := range transactions {
		if v[0] - v[1] < 0 {
			ans = max(ans, v[0] + maxv)
		} else {
			ans = max(ans, v[0] + maxv - v[0] + v[1])
		}
	}
	return int64(ans)
}

func main() {

}