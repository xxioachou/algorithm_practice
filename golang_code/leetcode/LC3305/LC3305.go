package main

import (
	"fmt"
	"os"
	"strconv"
)

func countOfSubstrings(word string, k int) int {
	n := len(word)
    cnt := make([]int, 26)
	s := make([]int, n + 1)
    sum := 0
    check1 := func(ch byte) bool {
        return ch == 'a' || ch == 'e' || 
            ch == 'i' || ch == 'o' ||
            ch == 'u'
    }
    check2 := func (i, j int) bool {
        return cnt[int('a' - 'a')] > 0 && cnt[int('e' - 'a')] > 0 && cnt[int('i' - 'a')] > 0 && 
        cnt[int('o' - 'a')] > 0 && cnt[int('u' - 'a')] > 0 && j - i - sum >= k
    }
	next := func(i int) int {
		l, r := i + 1, n
		for l < r {
			mid := (l + r + 1) / 2
			if s[mid] == s[i + 1] {
				l = mid
			} else {
				r = mid - 1
			}
		}
		return r - 1
	}
	for i := 0; i < n; i ++ {
		s[i + 1] = s[i]
		if !check1(word[i]) {
			s[i + 1] ++
		}
	}

	ans := 0
    for i, j := 0, 0; i < n; i ++ {
        for j < n && !check2(i, j) {
            if check1(word[j]) {
                sum ++
            }
            cnt[int(word[j] - 'a')] ++
            j ++
        }
		// fmt.Println(i, j)
        if check2(i, j) && j - i - sum == k {
			// fmt.Println(":", next(j - 1), j - 1)
            ans += next(j - 1) - (j - 1) + 1
        }
		cnt[int(word[i] - 'a')] --
		if check1(word[i]) {
			sum --
		}
    }
	return ans
}

func main() {
	w := os.Args[1]
	k, _ := strconv.Atoi(os.Args[2])
	fmt.Println(w, k)
	fmt.Println(countOfSubstrings(w, k))
}