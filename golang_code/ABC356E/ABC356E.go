package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func min(x, y int) int {
	if x < y  {
		return x
	}
	return y
}

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
	for ; T > 0; T -- {
		_ = readLine()[0]
		const MAXN = 1000010
		A := readLine()
		sort.Ints(A)
		
		cnt := make([]int, MAXN)
		s := make([]int, MAXN)
		for _, v := range A {
			cnt[v] ++
		}
		for i := 1; i < MAXN; i ++ {
			s[i] = s[i - 1] + cnt[i]
		}
		ans := 0
		for i := 1; i < MAXN; i ++ {
			for j := i; j < MAXN; j += i {
				l := j
				r := min(j + i - 1, MAXN - 1)
				ans += (s[r] - s[l - 1]) * (j / i) * cnt[i]
				// if t := (s[r] - s[l - 1]) * (j / i); t > 0 {
				// 	Fprintln(out, i, l, r, t)
				// }
			}
		}
		// cnt[i] * cnt[i] - C(cnt[i], 2)
		for i := 1; i < MAXN; i ++ {
			ans -= cnt[i] * (cnt[i] + 1) / 2
		}		
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }