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
	const mul = 1000000
	get := func (x, y, z int) int {
		return x * mul * mul + y * mul + z
	}
	for ; T > 0; T-- {
		var n int
		var s string
		Fscan(in, &n, &s)
		sum := make([][]int, 3)
		for j := 0; j < 3; j++ {
			sum[j] = make([]int, n + 1)
			for i := 0; i < n; i++ {
				sum[j][i + 1] = sum[j][i];
				if s[i] == byte(j + '0') {
					sum[j][i + 1] ++
				}
			}
		}
		mp := make(map[int]int)
		mp[get(0, 0, 0)] ++
		ans := 0
		for i := 1; i <= n; i++ {
			t := get(
				sum[0][i] - sum[1][i], 
				sum[1][i] - sum[2][i], 
				sum[0][i] - sum[2][i],
			)
			ans += mp[t]
			mp[t] ++
		}
		Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
