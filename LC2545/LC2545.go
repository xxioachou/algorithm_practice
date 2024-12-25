package main

import "slices"

func sortTheStudents(score [][]int, k int) [][]int {
    slices.SortFunc(score, func(a, b []int) int {
		return a[k] - b[k]
	})
	return score
}

func main() {

}