package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	suf := make([]int, n + 1)
	for i := n - 1; i >= 0; i -- {
		suf[i] = suf[i + 1] | nums[i]
	}
	pre, ans := 0, 0
	for i := range nums {
		ans = max(ans, pre | suf[i + 1] | (nums[i] << k))
	}
	return int64(ans)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	parts := strings.Split(str, " ")
	a := make([]int, len(parts))
	for i := range a {
		a[i], _ = strconv.Atoi(parts[i])
	}
	str, _ = in.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	k, _ := strconv.Atoi(str)

	fmt.Println("a:", a)
	fmt.Println("k:", k)
	fmt.Println(maximumOr(a, k))
}