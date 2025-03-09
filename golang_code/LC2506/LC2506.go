package main

func similarPairs(words []string) int {
	ans := 0
	mp := make(map[string]int)
	cnt := make([]int, 26)
	get := func(str string) string {
		for i := range cnt {
			cnt[i] = 0
		}	
		res := make([]byte, 26)
		for _, ch := range str {
			i := int(ch - 'a')
			cnt[i] ++
			if cnt[i] == 1 {
				res[i] = 1
			}
		}
		return string(res)
	}
	for i := range words {
		t := get(words[i])
		if v, ok := mp[t]; ok {
			ans += v
		}
		mp[t] ++
	}
	return ans
}

