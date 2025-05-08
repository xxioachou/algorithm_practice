package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func solve(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T := 1
	for ; T > 0; T-- {
		n := 1
		Fscan(in, &n)
		st := make([][]int, n * (n - 1) / 2)
		mp := make([][]int, 201)
		vis := make([]bool, 201)
		for i, k := 0, 0; i < len(st); i++ {
			Fscan(in, &k)
			st[i] = make([]int, k)
			for j := range st[i] {
				Fscan(in, &st[i][j])
				if (mp[st[i][j]] == nil) {
					mp[st[i][j]] = make([]int, 0)
				}
				mp[st[i][j]] = append(mp[st[i][j]], i)
			}
		}

		if n == 2 {
			Fprintln(out, 1, st[0][0])
			Fprint(out, len(st[0]) - 1)
			for i := 1; i < len(st[0]); i++ {
				Fprint(out, " ", st[0][i])
			}
			Fprintln(out, "")
			continue
		}

		for i := 1; i <= 200; i++ {
			if mp[i] == nil || vis[i] {
				continue
			}

			res := make([]string, 0)
			for j := 1; j <= 200; j++ {
				if mp[j] == nil || vis[j] || len(mp[j]) != len(mp[i]) {
					continue;
				}
				// Fprintln(out, ":", i, j)
				ok := true
				for k, x := range mp[i] {
					if x != mp[j][k] {
						ok = false
						break
					}
				}
				if ok {
					vis[j] = true
					res = append(res, strconv.Itoa(j))
				}
			}
			Fprintln(out, strconv.Itoa(len(res)), strings.Join(res, " "))
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }
