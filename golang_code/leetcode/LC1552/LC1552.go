package main

import (
	"slices"
)

func maxDistance(position []int, m int) int {
    slices.SortFunc(position, func(i, j int) int {
		return i - j
	})
	n := len(position)
	check := func(x int) bool {
		pos := position[0]
		used := 0
		for i := 0; used < m; {
			for i < n && position[i] < pos {
				i ++
			}
			if i >= n {
				break
			}
			if position[i] != pos {
				pos = position[i]
			}
			used ++
			pos += x
		}
		return used >= m
	}
	low, high := 0, position[n - 1] - position[0]
	for low < high {
		mid := (low + high + 1) / 2
		if check(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}	
	return high
}

func main() {

}