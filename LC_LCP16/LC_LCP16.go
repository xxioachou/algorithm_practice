package main

import (
	"slices"
)

func maxWeight(edges [][]int, value []int) int {
	n := len(value)
	e := make([][]int, n)
	for i := 0; i < n; i++ {
		e[i] = make([]int, 0)
	}
	for _, v := range(edges) {
		e[v[0]] = append(e[v[0]], v[1])	
		e[v[1]] = append(e[v[1]], v[0])
	}
	vis := make([]int, n)
	ans := 0
	get := func(a, b, c, d int) int{
		arr := []int{a, b, c, d}
		slices.Sort(arr)
		res := 0
		for i := 0; i < len(arr); i++ {
			if i != 0 && arr[i] == arr[i - 1] {
				continue
			}
			res += value[arr[i]]
		}
		return res
	}
	for i := 0; i < n; i++ {
		for _, j := range(e[i]) {
			vis[j] = 1
		}
		mp := make(map[int]bool, 0)
		for _, j := range(e[i]) {
			for _, k := range(e[j]) {
				if vis[k] == 1 {
					key1 := j * n + k
					key2 := k * n + j
					if key1 > key2 {
						key1, key2 = key2, key1
					}
					key := key1 * n * n + key2
					mp[key] = true
				}
			}
		}
		if len(mp) > 0 {
			res := 0
			for k1 :=  range(mp) {
				key1 := k1 / (n * n)
				x1, y1 := key1 / n, key1 % n
				for k2 := range(mp) {
					key2 := k2 / (n * n)
					x2, y2 := key2 / n, key2 % n
					res = max(res, get(x1, y1, x2, y2))
				}
			}
			ans = max(ans, value[i] + res)
		}
		for _, j := range(e[i]) {
			vis[j] = 0
		}
	}
	return ans
} 

func main() {

}