package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
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
	for ; T > 0; T -- {
		_ = readLine()[0]
		a := readLine()
		v1, v2 := 0, math.MaxInt
		for _, v := range a {
			v1 |= v
			v2 &= v
		}
		x := v1 ^ v2
		Fprintln(out, (x & -x) << 1)

		// ans := (x & -x) << 1
		// mp := make(map[int]int)
		// for _, v := range a {
		// 	mp[v % ans] ++
		// }
		// if len(mp) != 2 {
		// 	Fprintln(out, "WA!!!!")
		// }
	}
}

func main() { solve(os.Stdin, os.Stdout) }