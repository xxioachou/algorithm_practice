package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type bitTree struct {
	n int
	tr []int
}
func constructor(sz int) bitTree {
	return bitTree{
		sz,
		make([]int, sz + 1),
	}
}
func lowbit(x int) int {
	return x & -x
}
func (b bitTree) add(x, v int) {
	for x <= b.n {
		b.tr[x] += v
		x += lowbit(x)
	}
}
func (b bitTree) sum(x int) int {
	res := 0
	for x > 0 {
		res += b.tr[x]
		x -= lowbit(x)
	}
	return res
}

func solve(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T := 1
	// Fscan(in, &T)
	for ; T > 0; T-- {
		var n, m int
		Fscan(in, &n, &m)
		a := make([]int, n + 1)
		v := make([][]int, m + 1)
		app := make([]int, m + 1)
		for i := 1; i <= m; i++ {
			v[i] = make([]int, 0)
		}
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			v[a[i]] = append(v[a[i]], i)
		}
		tr := constructor(n)
		
		for _, i := range(v[1]) {
			tr.add(i, 1)
			app[1] = 1
		}

		ans := 0
		for i := 2; i <= m; i++ {
			app[i] = app[i - 1]
			if len(v[i]) == 0 {
				ans += app[i - 1]
			} else {
				app[i] ++
				tlen := len(v[i])
				for j := 0; j < tlen; {
					k := j + 1
					for k < tlen && v[i][k] == v[i][k - 1] + 1 {
						k ++
					}

					l := 0
					if j > 0 {
						l = v[i][j - 1]
					}
					ans += tr.sum(v[i][j] - 1) - tr.sum(l)
					ans += i - 1

					j = k
				}
				ans += tr.sum(n) - tr.sum(v[i][tlen - 1])

				for _, j := range(v[i]) {
					tr.add(j, 1)
				}
			}
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }
