package main

func numberOfWays(startPos int, endPos int, k int) int {
	const mod = 1e9 + 7
	f := make(map[int]map[int]int)
	var dfs func(pos, step int) int
	abs := func(x int) int {
		if x >= 0 {
			return x
		}
		return -x
	}
	dfs = func(pos, step int) int {
		if abs(pos - endPos) > step {
			return 0
		}		
		if step == 0 {
			return 1
		}
		if val, ok := f[pos][step]; ok {
			return val
		}
		val := (dfs(pos + 1, step - 1) + dfs(pos - 1, step - 1)) % mod
		if _, ok := f[pos]; !ok {
			f[pos] = make(map[int]int)
		}
		f[pos][step] = val
		return val
	}
	return dfs(startPos, k)
}

func main() {

}
