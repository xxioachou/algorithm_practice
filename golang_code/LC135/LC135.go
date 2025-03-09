package main

import "container/list"

func candy(ratings []int) int {
	const inf = 0x3f3f3f3f
	n := len(ratings)
	e := make([][]int, n + 1)
	din := make([]int, n + 1)
	for i := 0; i < n; i ++ {
		if i + 1 < n && ratings[i] > ratings[i + 1] {
			// i + 1 --> i
			e[i + 1 + 1] = append(e[i + 1 + 1], i + 1)
			din[i + 1] ++
		}
		if i - 1 >= 0 && ratings[i] > ratings[i - 1] {
			// i - 1 --> i
			e[i - 1 + 1] = append(e[i - 1 + 1], i + 1)
			din[i + 1] ++
		}
	}
	for i := 1; i <= n; i ++ {
		e[0] = append(e[0], i)
	}
	
	f := make([]int, n + 1)
	q := list.New()
	q.PushBack(0)
	for q.Len() > 0 {
		u := q.Front().Value.(int)
		q.Remove(q.Front())

		for _, v := range e[u] {
			f[v] = max(f[u] + 1, f[v])
			din[v] --
			if din[v] == 0 {
				q.PushBack(v)
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i ++ {
		ans += f[i]
	}
	return ans
}