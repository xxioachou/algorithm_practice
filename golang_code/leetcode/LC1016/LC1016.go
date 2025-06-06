package main

import "strings"

func find(x int, p []int) int {
	if p[x] != x {
		p[x] = find(p[x], p)
	}
	return p[x]
}

func merge(x, y int, p []int) {
	x = find(x, p)
	y = find(y, p)
	if x != y {
		p[x] = y
	}
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	p := make([]int, 26)
	f := make([]int, 26)
	for i := range p {
		p[i], f[i] = i, i
	}
	n := len(s1)
	for i := 0; i < n; i++ {
		x, y := int(s1[i]-'a'), int(s2[i]-'a')
		merge(x, y, p)
	}
	for i := 1; i < 26; i++ {
		for j := 0; j < i; j++ {
			if find(i, p) == find(j, p) {
				f[i] = j
				break
			}
		}
	}
	var ans strings.Builder
	for _, ch := range baseStr {
		x := int(ch - 'a')
		ans.WriteByte(byte(f[x] + 'a'))
	}
	return ans.String()
}
