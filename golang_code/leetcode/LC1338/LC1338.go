package main

import (
	"slices"
	"sort"
)

func minSetSize(arr []int) int {
	slices.Sort(arr)
	c := make([]int, 0)
	for i := 0; i < len(arr); {
		j := i + 1
		for j < len(arr) && arr[j] == arr[i] {
			j ++
		}
		c = append(c, j - i)
		i = j
	}
	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})
	i, sum := 0, 0
	for sum < len(arr) / 2 {
		sum += c[i]
		i ++
	}
	return i
} 

func main() {

}