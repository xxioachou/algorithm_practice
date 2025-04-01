package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	readLine := func() []int {
		str, _ := in.ReadString('\n')
		parts := strings.Fields(str)
		nums := make([]int, len(parts))
		for i := range parts {
			nums[i], _ = strconv.Atoi(parts[i])
		}
		return nums
	}
	
	T := 1
	// T = readLine()[0]
	for ; T > 0; T-- {
		n := readLine()[0]
		a := readLine()
		r := make([]int, n)
		l := make([]int, n)
		ans := 0
		for i := 0; i < n; i ++ {
			if i > 0 && a[i - 1] < a[i] {
				l[i] = l[i - 1]
			} else {
				l[i] = i
			}
			ans = max(ans, i - l[i] + 1)
		}
		for i := n - 1; i >= 0; i -- {
			if i + 1 < n && a[i] < a[i + 1] {
				r[i] = r[i + 1]
			} else {
				r[i] = i
			}
		}
		if n >= 2 {
			ans = max(ans, r[1] - 0 + 1)
			ans = max(ans, n - 1 - l[n - 2] + 1)
			for i := 1; i + 1 < n; i ++ {
				if a[i + 1] - a[i - 1] > 1 {
					ans = max(ans, r[i + 1] - l[i - 1] + 1)
				} else {
					ans = max(ans, r[i + 1] - i + 1)
					ans = max(ans, i - l[i - 1] + 1)
				}
			}
		}
		
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }