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
		t := readLine()
		_, K := t[0], t[1]
		A := readLine()	
		cnt := make([][]int, 60)
		for i := range cnt {
			cnt[i] = make([]int, 2)
		}
		for i := 59; i >= 0; i -- {
			for j := range A {
				cnt[i][(A[j] >> i & 1)] ++
			}
		}

		ans, X := 0, 0
		for i := 59; i >= 0; i -- {
			if cnt[i][0] >= cnt[i][1] && X + (1 << i) <= K {
				ans += (1 << i) * cnt[i][0]
				X += (1 << i)
			} else {
				ans += (1 << i) * cnt[i][1]
			}
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }