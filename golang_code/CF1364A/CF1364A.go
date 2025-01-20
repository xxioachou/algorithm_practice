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
	T = readLine()[0]
	for ; T > 0; T-- {
		t := readLine()
		n, x := t[0], t[1]
		a := readLine()
		idx := -1
		for i, sum := 0, 0; i < n; i++ {
			sum += a[i]
			if sum % x != 0 {
				idx = i + 1
				break
			}
		}
		if idx == -1 {
			Fprintln(out, -1)
			continue
		}

		ans := 0
		for i, sum := 0, 0; i < n; i++ {
			sum += a[i]
			if sum % x == 0 {
				ans = max(ans, i + 1 - idx)
			} else {
				ans = max(ans, i + 1)
			}
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
