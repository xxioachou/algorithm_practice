package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	
	T := 1
	// T = readLine()[0]
	for ; T > 0; T-- {
		var N, M int
		Fscan(in, &N, &M)
		t, ans := 0, 0		
		// [0, M - 1] -> [1, M]
		tr1 := bitTree{n: M, tr: make([]int, M + 1)}
		sq := 0
		add := func(r int) {
			tr1.add(r + 1, 1)
		}

		add(0)
		for i, A := 0, 0; i < N; i ++ {
			Fscan(in, &A)
			t = (t + A) % M

			sum := tr1.sum(M)
			ans += t * sum - sq

			sum2 := tr1.sum(M) - tr1.sum(t + 1)
			ans += M * sum2

			x := tr1.sum(t + 1) - tr1.sum(t)
			sq -= x * t
			sq += (x + 1) * t

			add(t)
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }

type bitTree struct {
	n 	int
	tr	[]int
}

func lowbit(x int) int {
	return x & -x;
}

func (tr *bitTree) add(x, v int) {
	for x <= tr.n {
		tr.tr[x] += v
		x += lowbit(x)
	}
}

func (tr *bitTree) sum(x int) int {
	res := 0
	for x > 0 {
		res += tr.tr[x]
		x -= lowbit(x)
	}
	return res
}