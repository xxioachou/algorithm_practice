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

	var n, q int
	var a []int
	const N = 2010
	s := make([][]int, N)
	p := make([][]int, N)
	t := make([][]int, N)
	for i := 0; i < N; i++ {
		s[i] = make([]int, N)
		p[i] = make([]int, N)
		t[i] = make([]int, N)
	}
	
	T := 1
	T = readLine()[0]
	for ; T > 0; T-- {
		a = readLine()
		n, q = a[0], a[1]

		for i := 1; i <= n; i++ {
			a = readLine()
			for j := 1; j <= n; j++ {
				s[i][j] = s[i][j - 1] + s[i - 1][j] - s[i - 1][j - 1] + a[j - 1]
				p[i][j] = p[i][j - 1] + p[i - 1][j] - p[i - 1][j - 1] + a[j - 1] * i
				t[i][j] = t[i][j - 1] + t[i - 1][j] - t[i - 1][j - 1] + a[j - 1] * j
			}
		}
		sum := func (arr [][]int, x1, y1, x2, y2 int) int {
			return arr[x2][y2] - arr[x1 - 1][y2] - arr[x2][y1 - 1] + arr[x1 - 1][y1 - 1]
		}
		for i:= 0; i < q; i++ {
			a = readLine()
			x1, y1, x2, y2 := a[0], a[1], a[2], a[3]	
			
			sum1 := sum(s, x1, y1, x2, y2)
			ans := sum(t, x1, y1, x2, y2) - sum1 * (y1 - 1)
			tmp := sum(p, x1, y1, x2, y2) - sum1 * x1
			ans += tmp * (y2 - y1 + 1)
			Fprint(out, ans, " ")
			// Fprintln(out, "sum1: ", sum1)
			// Fprintln(out, "sumt: ", sum(t, x1, y1, x2, y2))
			// Fprintln(out, "sump: ", sum(p, x1, y1, x2, y2))
			// Fprintln(out, "tmp: ", tmp)
		}
		Fprintln(out, "")
	}
}

func main() { solve(os.Stdin, os.Stdout) }
