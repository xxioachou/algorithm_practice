package main

import "sort"

type RangeFreqQuery struct {
	mp 	map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	res := RangeFreqQuery{mp: make(map[int][]int)}
	for i, x := range arr {
		if _, ok := res.mp[x]; !ok {
			res.mp[x] = make([]int, 0)
		}
		res.mp[x] = append(res.mp[x], i)
	}
	return res
}

func (r RangeFreqQuery) lowerBound(value, x int) int {
	return sort.Search(len(r.mp[value]), func(i int) bool {
		return r.mp[value][i] >= x
	})
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	if _, ok := this.mp[value]; !ok {
		return 0
	}
	return this.lowerBound(value, right + 1) - this.lowerBound(value, left)
}

func main() {

}
