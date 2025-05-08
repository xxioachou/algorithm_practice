package main

import (
	"slices"
)

type pair struct {
	l int
	r int
}

type MyCalendarThree struct {
	points []pair
}

func Constructor() MyCalendarThree {
    return MyCalendarThree{
		make([]pair, 0),
	}
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	points := (*this).points
	points = append(points, pair{startTime, 1})
	points = append(points, pair{endTime, -1})
	slices.SortFunc(points, func(a, b pair) int {
		if a.l != b.l {
			return a.l - b.l
		}
		return a.r - b.r
	})
	k := 0
	cnt := 0
	for i := 0; i < len(points); {
		j := i + 1
		cnt += points[i].r
		for j < len(points) && points[j].l == points[i].l {
			cnt += points[j].r
			j ++			
		}
		k = max(k, cnt)
		i = j
	}
	(*this).points = points
	return k
} 

func main() {

}