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
		n, m := t[0], t[1]
		a := readLine()
		last := make([]int, n + m + 1)
		app := make([]int, n + m + 1)
		for i := 1; i < len(last); i++ {
			last[i] = -1
		}
		for _, v := range a {
			last[v] = 0
		}

		for i := 1; i <= m; i ++ {
			t = readLine()
			p, v := t[0], t[1]
			u := a[p - 1]
			if u != v {
				app[u] += i - last[u]
				last[u] = -1
				a[p - 1] = v
				last[v] = i
			}
		}
		for i := 1; i < len(last); i ++ {
			if last[i] != -1 {
				app[i] += m + 1 - last[i]
			}
		}
		ans := (1 + m) * m * n
		for i := 1; i < len(app); i++ {
			// Fprintln(out, "app: ", i, app[i])
			ans -= app[i] * (app[i] - 1) / 2
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
