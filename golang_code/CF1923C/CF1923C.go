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
		n, q := t[0], t[1]
		c := readLine()
		s := make([]int, n + 1)
		cnt := make([]int, n + 1)
		for i := 0; i < n; i++ {
			s[i + 1] = s[i] + c[i]
			cnt[i + 1] = cnt[i]
			if c[i] == 1 {
				cnt[i + 1] ++
			}
		}
		for i := 0; i < q; i++ {
			t = readLine()
			l, r := t[0], t[1]
			sum := s[r] - s[l - 1]
			tlen := r - l + 1
			cnt1 := cnt[r] - cnt[l - 1]
			if tlen == 1 || cnt1 * 2 + (tlen - cnt1) > sum {
				Fprintln(out, "NO")
			} else {
				Fprintln(out, "YES")
			}
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }
