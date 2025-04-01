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

		a := make([][]int, n)
		s := make([]int, n)
		idx := make([]int, n)
		for i := range a {
			a[i] = readLine()
			// Fprintln(out, "a[i]:", a[i])
			idx[i] = i
			for _, x := range a[i] {
				s[i] += x
			}
			// Fprintln(out, "s[i]:", s[i])
		}
		sort.Slice(idx, func(i, j int) bool {
			return s[idx[i]] > s[idx[j]]
		})
		// Fprintln(out, "idx:", idx)

		sum, ans := 0, 0
		for _, x := range idx {
			for _, v := range a[x] {
				sum += v
				ans += sum
			}
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }