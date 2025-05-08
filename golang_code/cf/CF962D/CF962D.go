package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"slices"
)

type que []int
func (q que) Len() int {
	return len(q)
}
func (q que) Less(i, j int) bool {
	return q[i] < q[j]
}
func (q que) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q *que) Push(x any) {
	*q = append(*q, x.(int))
}
func (q *que) Pop() any {
	old := *q
	n := len(old)
	x := old[n - 1]	
	*q = old[: n - 1]
	return x
}

func solve(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T := 1
	// Fscan(in, &T)
	for ; T > 0; T-- {
		var n int
		Fscan(in, &n)
		mp := make(map[int][]int)
		a := make([]int, n)
		h := make(que, 0)
		heap.Init(&h)
		for i := range(a) {
			Fscan(in, &a[i])
			if _, ex := mp[a[i]]; !ex {
				mp[a[i]] = make([]int, 0)
				heap.Push(&h, a[i])
			}	
			mp[a[i]] = append(mp[a[i]], i)
		}
		for h.Len() > 0 {
			x := heap.Pop(&h).(int)
			v := mp[x]
			slices.Sort(v)
			// Fprintln(out, "x: ", x)
			for i := 1; i < len(v); i += 2 {
				a[v[i - 1]] = -1
				a[v[i]] = x * 2
				if _, ex := mp[x * 2]; !ex {
					mp[x * 2] = make([]int, 0)
					heap.Push(&h, x * 2)
				}
				mp[x * 2] = append(mp[x * 2], v[i])
			}
			// for i := 0; i < n; i++ {
			// 	Fprint(out, a[i], " ")
			// }
			// Fprintln(out, "")
		}
		ans := make([]int, 0)
		for _, v := range(a) {
			if v != -1 {
				ans = append(ans, v)
			}
		}
		Fprintln(out, len(ans))
		for _, v := range(ans) {
			Fprint(out, v, " ")
		}
		Fprintln(out, "")
	}
}

func main() { solve(os.Stdin, os.Stdout) }
