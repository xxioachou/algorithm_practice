package main

func minValidStrings(words []string, target string) int {
	idx, M := 0, 0
	for i := range(words) {
		M += len(words[i])
	}
	tr := make([][]int, M + 1)
	for i := range(tr) {
		tr[i] = make([]int, 26)
	}

	// 将前缀串插入 trie （最多 100000 个前缀）
	for _, str := range(words) {
		n := len(str)
		u := 0
		for i := 0; i < n; i++ {
			v := int(str[i] - 'a')
			if tr[u][v] == 0 {
				tr[u][v] = idx + 1
				idx ++
			}	
			u = tr[u][v]
		}
	}

	n := len(target)
	f := make([]int, n + 1)

	const inf = 0x3f3f3f3f
	f[0] = 0
	for i := 1; i <= n; i++ {
		f[i] = inf
	}

	for i := 0; i + 1 <= n; i++ {
		u := 0
		for j := i + 1; j <= n; j++ {
			v := int(target[j - 1] - 'a')
			if tr[u][v] == 0{
				break
			}

			f[j] = min(f[j], f[i] + 1)
			u = tr[u][v]
		}
	}
	ans := f[n]
	if ans >= inf {
		ans = -1
	}
	return ans
} 

func main() {

}