package main

import "fmt"

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)
	p := startPos + 1
	m := max(fruits[n - 1][0] + 1, p)

	s := make([]int, m + 1)
	for _, v := range fruits {
		s[v[0] + 1] += v[1]
	}
	for i := 1; i <= m; i ++ {
		s[i] += s[i - 1]
	}

	ans := s[p] - s[p - 1]
	for x := 0; x <= k && p - x > 0; x ++ {
		res := s[p] - s[p - x - 1]
		r := min(p - x + k - x, m)
		if r >= p {
			res += s[r] - s[p]
		}
		ans = max(ans, res)
	}
	for x := 0; x <= k && p + x <= m; x ++ {
		res := s[p + x] - s[p - 1]
		l := max(1, p + x - (k - x))
		if l <= p - 1 {
			res += s[p - 1] - s[l - 1]
		}
		ans = max(ans, res)
	}
	return ans
}

func main() {
	var n, p, k int
	fmt.Scan(&n, &p, &k)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, 2)
		fmt.Scan(&f[i][0], &f[i][1])
	}
	fmt.Println(maxTotalFruits(f, p, k))
}