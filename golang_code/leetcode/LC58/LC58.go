package main

func lengthOfLastWord(s string) int {
    i := len(s) - 1
	ans := 0
	for i >= 0 {
		if s[i] == ' ' {
			if ans > 0 {
				break
			}
			i -- 
			continue
		}

		ans ++
		i --
	}
	return ans
}