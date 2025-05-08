package main

import "fmt"

const mod = 1000000007
const P = 131
const inf = 0x3f3f3f3f

func minCut(s string) int {
	n := len(s)
	h := make([]int, n + 2)	
	rh := make([]int, n + 2)	
	p := make([]int, n + 2)	
	p[0] = 1
	for i := 1; i <= n; i ++ {
		p[i] = p[i - 1] * P % mod
	}
	for i := 0; i < n; i ++ {
		h[i + 1] = (h[i] * P % mod + int(s[i] - 'a' + 1)) % mod
	}
	for i := n; i > 0; i -- {
		rh[i] = (rh[i + 1] * P % mod + int(s[i - 1] - 'a' + 1)) % mod
	}
	hash1 := func(l, r int) int {
		return (h[r] - h[l - 1] * p[r - l + 1] % mod + mod) % mod
	}
	hash2 := func(l, r int) int {
		return (rh[l] - rh[r + 1] * p[r - l + 1] % mod + mod) % mod
	}

	f := make([]int, n + 1)
	for i := range f {
		f[i] = inf
	}
	f[0] = 0
	for i := 1; i <= n; i ++ {
		for j := i; j > 0; j -- {
			if hash1(j, i) == hash2(j, i) {
				fmt.Println(j, i)
				f[i] = min(f[i], f[j - 1] + 1)
			}
		}
	}
	return f[n] - 1
}
