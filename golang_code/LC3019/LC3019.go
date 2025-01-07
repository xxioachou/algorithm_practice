package main

import (
	"strings"
)

func countKeyChanges(s string) int {
	s = strings.ToLower(s)
	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i - 1] {
			ans++
		}
	}
	return ans
}

func main() {
	
}
