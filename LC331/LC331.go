package main

import (
	"strings"
)

func check(str []string, flag *bool) int {
	if len(str) == 0 {
		*flag = false
		return 0
	}
	// empty node has no child, no recursion
	if (str[0] == "#") {
		return 1
	}
	
	le := check(str[1:], flag)
	ri := check(str[le + 1:], flag)
    // fmt.Println(str, *flag, le, ri)
	return le + 1 + ri
}

func isValidSerialization(preorder string) bool {
    str := strings.Split(preorder, ",")
	flag := true
	cnt := check(str, &flag)
	if cnt != len(str) {
		flag = false
	}
	return flag
} 

func main() {

}