package main

import (
	"slices"
)

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	slices.Sort(horizontalCut)
	slices.Sort(verticalCut)
	ans := 0
	i, j := 0, 0
	for i < len(horizontalCut) && j < len(verticalCut) {
		if horizontalCut[i] < verticalCut[j] {
			ans += horizontalCut[i] * (n - j)
			i ++
		} else {
			ans += verticalCut[j] * (m - i)
			j ++
		}
	}
	for i < len(horizontalCut) {
		ans += horizontalCut[i] * (n - j)
		i ++
	}
	for j < len(verticalCut) {
		ans += verticalCut[j] * (m - i)
		j ++
	}
	return int64(ans)
} 

func main() {

}