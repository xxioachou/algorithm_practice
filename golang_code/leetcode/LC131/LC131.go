package main


func partition(s string) [][]string {
	n := len(s)
	p := make([][]int, n)
	for i := 0; i < n; i ++ {
		p[i] = make([]int, 0)
	}
	check := func(str []byte) bool {
		for i := 0; i < len(str) / 2; i ++ {
			if str[i] != str[len(str) - 1 - i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < n; i ++ {
		str := make([]byte, 0)
		for j := i; j < n; j ++ {
			str = append(str, s[j])
			if check(str) {
				p[i] = append(p[i], j)
			}
		}
	}
	var dfs func(dep int)
	res := make([]string, 0)
	ans := make([][]string, 0)
	dfs = func(dep int) {
		if dep == n {
			if len(res) != 0 {
				tmp := make([]string, len(res))
				copy(tmp, res)
				ans = append(ans, tmp)
			}
			return
		}
		for _, x := range p[dep] {
			res = append(res, s[dep : x + 1])
			dfs(x + 1)
			res = res[:len(res) - 1]
		}
	}
	dfs(0)
	return ans
}

