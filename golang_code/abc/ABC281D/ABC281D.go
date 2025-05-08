package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	// T = readLine()[0]
	const inf = 1e12
	for ; T > 0; T -- {
		t := readLine()
		N, K, D := t[0], t[1], t[2]
		a := readLine()
		f := make([][][]int, N + 1)
		for i := range f {
			f[i] = make([][]int, K + 1)
			for j := range f[i] {
				f[i][j] = make([]int, D + 1)
				for k := range f[i][j] {
					f[i][j][k] = -inf
				}
			}
		}
		f[0][0][0] = 0
		for i := 0; i < N; i ++ {
			for j := 0; j <= i + 1 && j <= K; j ++ {
				for k := 0; k < D; k ++ {
					f[i + 1][j][k] = f[i][j][k]
					if j > 0 {
						r := ((k - a[i]) % D + D) % D
						f[i + 1][j][k] = max(f[i + 1][j][k], f[i][j - 1][r] + a[i])
					}
				}
			}
		}
		ans := f[N][K][0]
		if ans < 0 {
			ans = -1
		}
		Fprintln(out, ans)
	}

}

func main() { solve(os.Stdin, os.Stdout) }