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
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	T := 1
	// T = readLine()[0]
	const C = 18
	for ; T > 0; T-- {
		arr := readLine()
		N, Q := arr[0], arr[1]
		H := readLine()
		f := make([][]int, N + 1)
		for i := 0; i < N; i++ {
			f[i] = make([]int, C)
		}
		lg := make([]int, N + 1)
		cnt := make([]int, N + 1)
		query := func(l, r int) int {
			k := lg[r - l + 1]
			return max(f[l][k], f[r - (1 << k) + 1][k])
		}
		for i := 2; i <= N; i++ {
			lg[i] = lg[i >> 1] + 1
		}
		for j := 0; j < C; j++ {
			for i := 0; i + (1 << j) - 1 < N; i++ {
				if j == 0 {
					f[i][j] = H[i]
				} else {
					f[i][j] = max(f[i][j - 1], f[i + (1 << (j - 1))][j - 1])
				}
			}
		}
		stk := make([]int, 0)
		for i := N - 1; i >= 0; i-- {
			for len(stk) > 0 && stk[len(stk) - 1] <= H[i] {
				stk = stk[: len(stk) - 1]
			}
			stk = append(stk, H[i])
			cnt[i] = len(stk)
		}

		for i := 1; i <= Q; i++ {
			arr = readLine()
			l, r := arr[0] - 1, arr[1] - 1
			if r >= N - 1 {
				Fprintln(out, 0)
			} else {
				t := query(l + 1, r)
				low, high := r + 1, N - 1
				for low < high {
					mid := (low + high) / 2
					if query(r + 1, mid) >= t {
						high = mid
					} else {
						low = mid + 1
					}
				}
				if query(r + 1, high) < t {
					Fprintln(out, 0)
				} else {
					Fprintln(out, cnt[high])
				}
			}
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }
