package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type pair struct {
	str string
	l int
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
		var s string
		Fscan(in, &s)
		n := len(s)
		l, r := -1, -1
		for i := 1; i < n; i++ {
			tlen := n - i
			strs := make([]pair, 0)
			for j := 0; j + tlen - 1 < n; j++ {
				if s[j] != s[i] {
					strs = append(strs, pair{s[j : j + tlen], j})
				}
			}
			if len(strs) > 0 {
				slices.SortFunc(strs, func (a, b pair) int {
					for j := 0; j < tlen; j++ {
						if a.str[j] != b.str[j] {
							ch := int(s[n - tlen + j])
							return (int(b.str[j]) ^ ch) - (int(a.str[j]) ^ ch)
						}
					}
					return 0
				})
				// for _, str := range strs {
				// 	Fprintln(out, str)
				// }
				l, r = strs[0].l + 1, strs[0].l + tlen
				break
			}
		}
		if l == -1 {
			l, r = 1, 1
		}
		Fprintln(out, 1, n, l, r)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
