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
		str = strings.TrimRight(str, "\r\n")
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
		s := make([]string, n)
		e := make([][]int, 26)
		din := make([]int, 26)
		for i := range e {
			e[i] = make([]int, 0)
		}
		for i := range s {
			s[i], _ = in.ReadString('\n')
			s[i] = strings.TrimRight(s[i], "\r\n")
			// Fprintln(out, i, s[i])
		}
		for i := 0; i + 1 < n; i++ {
			for j := i + 1; j < n; j++ {
				ok := false
				for k := 0; k < min(len(s[i]), len(s[j])); k++ {
					if s[i][k] != s[j][k] {
						x := int(s[i][k] - 'a')
						y := int(s[j][k] - 'a')
						// Fprintln(out, i, j, k, x, y)
						e[x] = append(e[x], y)
						din[y] ++
						ok = true
						break
					}
				}
				if !ok && len(s[i]) > len(s[j]) {
					Fprintln(out, "Impossible")
					return
				}
			}
		}
		q := make([]int, 0)
		for i := range e {
			if din[i] == 0 {
				q = append(q, i)
			}
		}
		ans := make([]byte, 0)
		for len(q) > 0 {
			u := q[0]
			q = q[1:]

			ans = append(ans, byte(u + 'a'))
			for _, v := range e[u] {
				din[v] --
				if din[v] == 0 {
					q = append(q, v)
				}
			}
		}
		for i := range e {
			if din[i] > 0 {
				Fprintln(out, "Impossible")
				return
			}
		}
		Fprintln(out, string(ans))

		// b := make([]string, len(s))
		// copy(b, s)
		// slices.SortFunc(b, func (x, y string) int {
		// 	for i := 0; i < min(len(x), len(y)); i++ {
		// 		if x[i] != y[i] {
		// 			return strings.IndexByte(string(ans), x[i]) - strings.IndexByte(string(ans), y[i])
		// 		}
		// 	}
		// 	return len(x) - len(y)
		// })
		// for i := range s {
		// 	if s[i] != b[i] {
		// 		Fprintln(out, "WA!")
		// 		return
		// 	}
		// }
		// Fprintln(out, "AC!")
	}
}

func main() { solve(os.Stdin, os.Stdout) }
