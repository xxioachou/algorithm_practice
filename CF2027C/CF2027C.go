package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T := 1
	for Fscan(in, &T); T > 0; T-- {
		n := 1
		Fscan(in, &n)
		a := make([]int, n + 1)
		mp := make(map[int][]int, 0)
		f := make([]int, n + 1)
		for i := 1; i <= n; i++ {
			f[i] = -1;
		}
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			if mp[a[i] + i] == nil {
				mp[a[i] + i] = make([]int, 0)
			}
			mp[a[i] + i] = append(mp[a[i] + i], i)
		}	
		var dfs func(int) int
		dfs = func(i int) int {
			if f[i] == -1 {
				res := i - 1
				idx := a[i] + i * 2 - 1
				if mp[idx] != nil {
					for _, j := range mp[idx] {
						if j != i {
							res = max(res, dfs(j) + i - 1)
						}
					}
				}
				f[i] = res
			}
			return f[i]
		}
		ans := 0
		for i := 1; i <= n; i++ {
			if a[i] == n - i + 1 {
				ans = max(ans, dfs(i))
			}
		}
		Fprintln(out, ans + n)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
