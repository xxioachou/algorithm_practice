package main

import "fmt"

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	ans := ""
	for i := range word {
		t := make([]byte, 0)
		ok := false
		for j := i; j < len(word) && len(word)-(j-i+1) >= numFriends-1; j++ {
			if j-i < len(ans) && word[j] < ans[j-i] && !ok {
				break
			}
			if j-i >= len(ans) || word[j] > ans[j-i] {
				ok = true
			}
			t = append(t, word[j])
		}
		if ok {
			ans = string(t)
		}
	}
	return ans
}

func main() {
	var word string
	var num int
	fmt.Scan(&word, &num)
	fmt.Println(answerString(word, num))
}
