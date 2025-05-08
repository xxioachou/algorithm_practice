package main

func longestCommonPrefix(strs []string) string {
	ans := make([]byte, 0)
	for i := 0; ; i ++ {
		ok := true
		var ch byte
		for j, str := range strs {
			if i >= len(str) {
				ok = false
				break
			}
			if j == 0 {
				ch = str[i]
			} else if ch != str[i] {
				ok = false
				break
			}
		}
		if !ok {
			break
		}
		ans = append(ans, ch)
	}
	return string(ans)
}
