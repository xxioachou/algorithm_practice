package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ceil(x, y int) int {
	if x >= 0 {
		return (x + y - 1) / y
	}
	return x / y
}

func minReverseOperations(n int, p int, banned []int, k int) []int {
	ans := make([]int, n)
	vis := make([]bool, n)
	for i := range ans {
		ans[i] = 0x3f3f3f3f
	}
	for _, x := range banned {
		vis[x] = true
	}
	

	q := make([]int, 0)
	q = append(q, p)
	ans[p] = 0
	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		l := max(0, u - k + 1)
		r := min(u, n - k)
		for L := l; L <= r; L ++ {
			v := L * 2 + k - 1 - u
		
			if !vis[v] && ans[v] > ans[u] + 1 {
				ans[v] = ans[u] + 1
				q = append(q, v)
			}
		}
	}
	for i := range ans {
		if ans[i] >= 0x3f3f3f3f {
			ans[i] = -1
		}
	}
	return ans
}

func main() {
	var n, p, k int
	fmt.Scan(&n, &p, &k)
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	// fmt.Println("str:", str)
	parts := strings.Split(str, " ")
	b := make([]int, len(parts))
	for i := range b {
		b[i], _ = strconv.Atoi(parts[i])
	}
	// fmt.Println("b:", b)
	fmt.Println(minReverseOperations(n, p, b, k))
}