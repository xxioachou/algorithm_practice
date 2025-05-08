package main

import (
	"container/heap"
)

type FoodRatings struct {
	ratings 	map[string]int
	cuisines 	map[string]string
    st 			map[string]*que
}


func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
    fr := FoodRatings{
		ratings: make(map[string]int),
		cuisines: make(map[string]string),
		st: make(map[string]*que),
	}
	for i := range foods {
		if _, ok := fr.st[cuisines[i]]; !ok {
			fr.st[cuisines[i]] = &que{}
			heap.Init(fr.st[cuisines[i]])
		}
		fr.ratings[foods[i]] = ratings[i]
		fr.cuisines[foods[i]] = cuisines[i]
		heap.Push(fr.st[cuisines[i]], pair{rating: ratings[i], name: foods[i]})
	}
	return fr
}


func (fr *FoodRatings) ChangeRating(food string, newRating int)  {
    fr.ratings[food] = newRating
	heap.Push(fr.st[fr.cuisines[food]], pair{rating: newRating, name: food})
}


func (fr *FoodRatings) HighestRated(cuisine string) string {
	q := fr.st[cuisine]
    for q.Len() > 0 && fr.ratings[(*q)[0].name] != (*q)[0].rating {
		heap.Pop(q)
	}
	return (*q)[0].name
}

type pair struct {
	rating 	int
	name 	string
}

type que []pair

func (q que) Len() int {
	return len(q)
}

func (q que) Less(i, j int) bool {
	if q[i].rating != q[j].rating {
		return q[i].rating > q[j].rating
	}
	return q[i].name < q[j].name
}

func (q que) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *que) Push(x any) {
	*q = append(*q, x.(pair))
}

func (q *que) Pop() any {
	old := *q
	*q = old[:len(old) - 1]
	return old[len(old) - 1]
}
