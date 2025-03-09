package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	x int
	y int
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
	for ; T > 0; T-- {
		t := readLine()
		n, _ := t[0], t[1]
		x := readLine()
		ans := make(map[int]int)
		for i := range x {
			ans[(i + 1) * (n - i) - 1] += 1
		}
		for i := 1; i < n; i++ {
			ans[i * (n - i)] += x[i] - x[i - 1] - 1
		}

		k := readLine()
		for _, v := range k {
			Fprint(out, ans[v], " ")
		}		
		Fprintln(out, "")
	}
}

func main() { solve(os.Stdin, os.Stdout) }
