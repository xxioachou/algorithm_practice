package main

func validSubstringCount(word1 string, word2 string) int64 {
    cnt := make([]int, 26)
	cur := make([]int, 26)
	for _, c := range word2 {
		cnt[int(c - 'a')] ++
	}
	ans := 0
	for i, j := 0, 0; j < len(word1); j++ {
		cur[int(word1[j] - 'a')] ++
		for i < j && cur[int(word1[i] - 'a')] - 1 >= cnt[int(word1[i] - 'a')]{
			cur[int(word1[i] - 'a')] --
			i ++
		}
		ok := true
		for k := 0; k < 26; k++ {
			if cur[k] < cnt[k] {
				ok = false
				break
			}
		}
		if ok {
			ans += i + 1
		}
	}
	return int64(ans)
}

func main() {

}
