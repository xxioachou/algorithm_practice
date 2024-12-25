package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)
type pair struct {
	v int
	i int
}
type pque []pair
func (q pque) Len() int { return len(q) }
func (q pque) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q *pque) Push(x any) { *q = append(*q, x.(pair)) }
func (q pque) Less(i, j int) bool { 
	if q[i].v != q[j].v {
		return q[i].v > q[j].v
	}
	return q[i].i < q[j].i
}
func (q *pque) Pop() any {
	old := *q
	n := len(old)
	x := old[n - 1]
	*q = old[: n - 1]
	return x
}

func solve(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T := 1
	Fscan(in, &T)
	for ; T > 0; T-- {
		var n, k int
		var s string
		Fscan(in, &n, &k, &s)
		suf := make([]int, n + 2)
		a := make([]int, 0)
		for i := n; i > 0; i-- {
			suf[i] = suf[i + 1]
			if s[i - 1] == '1' {
				suf[i] ++
			} else {
				suf[i] --
			}
			if i > 1 {
				a = append(a, suf[i])
			}
		}
		slices.Sort(a)
		ans := 1
		for k > 0 && len(a) > 0 {
			k -= a[len(a) - 1]
			a = a[:len(a) - 1]
			ans ++
		}
		if k > 0 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
