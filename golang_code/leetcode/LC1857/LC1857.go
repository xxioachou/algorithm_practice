package main

import "container/list"

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	din := make([]int, n)
	f := make([]int, n)
	state := make([][]int, n)
	e := make([][]int, n)
	for i := range e {
		e[i] = make([]int, 0)
		state[i] = make([]int, 26)
	}
	for _, eg := range edges {
		u, v := eg[0], eg[1]
		din[v]++
		e[u] = append(e[u], v)
	}

	q1 := list.New()
	q2 := list.New()
	for i := 0; i < n; i++ {
		if din[i] == 0 {
			q1.PushBack(i)
			f[i] = 1
			state[i][int(colors[i]-'a')]++
			q2.PushBack(state[i])
		}
	}
	for q1.Len() > 0 {
		u := q1.Front().Value.(int)
		st := q2.Front().Value.([]int)
		q1.Remove(q1.Front())
		q2.Remove(q2.Front())

		for _, v := range e[u] {
			ne_st := make([]int, len(st))
			copy(ne_st, st)
			ne_st[int(colors[v]-'a')]++
			maxv := 0
			for _, x := range ne_st {
				maxv = max(maxv, x)
			}
			if maxv > f[v] {
				f[v] = maxv
				state[v] = ne_st
			}
			din[v]--
			if din[v] == 0 {
				q1.PushBack(v)
				q2.PushBack(state[v])
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		if din[i] != 0 {
			ans = -1
			break
		}
		ans = max(ans, f[i])
	}
	return ans
}
