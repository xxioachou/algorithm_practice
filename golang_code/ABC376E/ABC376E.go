package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func min(x, y int) int {
	if x < y {
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
	for ; T > 0; T-- {
		t := readLine()
		N, K := t[0], t[1]
		A := readLine()
		B := readLine()
		idx := make([]int, N)
		for i := range idx {
			idx[i] = i
		}
		sort.SliceStable(idx, func (i, j int) bool {
			return A[idx[i]] < A[idx[j]]
		})
		// Fprintln(out, idx)

		h := make(que, 0)
		heap.Init(&h)
		sumB := 0
		ans := math.MaxInt
		for i := 0; i < N; i++ {
			if i >= K - 1 {
				ans = min(ans, A[idx[i]] * (B[idx[i]] + sumB))
			}

			heap.Push(&h, B[idx[i]])
			sumB += B[idx[i]]
			if h.Len() > K - 1 {
				x := heap.Pop(&h).(int)
				// Fprintln(out, "i:", i, "Pop:", x)
				sumB -= x
			}
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }

type que []int
func (q que) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q que) Len() int {
	return len(q)
}
func (q que) Less(i, j int) bool {
	return q[i] > q[j]
}
func (q *que) Push(x any) {
	*q = append(*q, x.(int))
}
func (q *que) Pop() any {
	old := *q
	n := len(old)
	*q = old[: n - 1]
	return old[n - 1]
}