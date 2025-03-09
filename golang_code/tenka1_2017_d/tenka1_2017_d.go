// https://atcoder.jp/contests/tenka1-2017/tasks/tenka1_2017_d
package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func min(x, y int) int {
	if x < y  {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	
	T := 1
	// T = readLine()[0]
	for ; T > 0; T -- {
		var N, K int
		Fscan(in, &N, &K)
		A := make([]int, N)
		B := make([]int, N)
		for i := range A {
			Fscan(in, &A[i], &B[i])
		}
		K ++
		bitmask := (1 << 31) - 1
		ans := 0
		for i := 30; i >= 0; i -- {
			if (K >> i & 1)  == 1{
				// 或和 这一位是 0
				sum := 0
				for j := range A {
					if (A[j] & bitmask) == A[j] && (A[j] >> i & 1) == 0 {
						sum += B[j]
					}
				}
				ans  = max(ans, sum)
				// 或和 这一位是 1，继续看下一位
			} else {
				// 或和 这一位只能是 0
				bitmask ^= (1 << i)
			}
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }