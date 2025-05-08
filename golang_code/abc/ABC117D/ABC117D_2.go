package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func min(x, y int) int {
	if x < y  {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
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
		t := readLine()
		K := t[1]
		A := readLine()
		cnt := make([][]int, 2)
		cnt[0] = make([]int, 60)
		cnt[1] = make([]int, 60)
		suf := make([]int, 60)
		for i := 0; i < 60; i ++ {
			for j := range A {
				cnt[A[j] >> i & 1][i] ++
			}
			suf[i] = max(cnt[0][i], cnt[1][i]) << i
		}
		for i := 1; i < 60; i ++ {
			suf[i] += suf[i - 1]
		}
		K ++
		
		ans := 0
		for i, pre := 59, 0; i >= 0; i -- {
			if (K >> i & 1) == 1 {
				// X 这一位选 0, 后面的位随便选
				res := pre + (cnt[1][i] << i)
				if i > 0 {
					res += suf[i - 1]
				}
				ans = max(ans, res)
				// X 这一位选 1
				pre += (cnt[0][i] << i)
			} else {
				// X 这一位只能选 0
				pre += (cnt[1][i] << i)
			}
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }