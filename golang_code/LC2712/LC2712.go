package main

func minimumCost(s string) int64 {
	n := len(s)
	ans := 0
	for i := 0; i + 1 < n; i ++ {
		if s[i + 1] != s[i] {
			ans += min(i + 1, n - (i + 1))
		}
	}
	return int64(ans)
}

