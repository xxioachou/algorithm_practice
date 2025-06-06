package main

import (
	"fmt"
	"strings"
)

func robotWithString(s string) string {
	stk := make([]int, 0)
	f := make([]int, len(s)+1)
	f[len(s)] = 26
	for i := len(s) - 1; i >= 0; i-- {
		f[i] = min(f[i+1], int(s[i]-'a'))
	}
	var ans strings.Builder
	for i := 0; i < len(s); {
		if len(stk) > 0 && stk[len(stk)-1] <= f[i] {
			ans.WriteByte(byte(stk[len(stk)-1] + 'a'))
			stk = stk[:len(stk)-1]
			continue
		}

		stk = append(stk, int(s[i]-'a'))
		i++
	}
	for len(stk) > 0 {
		ans.WriteByte(byte(stk[len(stk)-1] + 'a'))
		stk = stk[:len(stk)-1]
	}
	return ans.String()
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(robotWithString(s))
}
