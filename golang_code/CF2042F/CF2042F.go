package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type node struct {
	s int
	add int 		// lazy tag
	v1 int			// s[i] + b[i]
	idx1 int 		// v1 对应的下标
	v2 int			// s[i]	- b[i + 1]
	idx2 int		// v2 对应的下标
	v3 int			// v1 - v2
	l int
	r int			// [l, r] v3 对应的区间
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

	const inf = 1000000000000000
	var build func(int, int, int)
	var update func(int, int, int, int, int, int)
	var change1 func(int, int, int, int)
	var change2 func(int, int, int, int)
	var query func(int, int, int, int, int) node
	T := 1
	// T = readLine()[0]
	for ; T > 0; T-- {
		n := readLine()[0]
		s := make([]int, n + 1)
		a := readLine()
		b := readLine()
		for i := 0; i < n; i++ {
			s[i + 1] = s[i] + a[i]
		}
		tr := make([]node, (n + 10) * 4)
		pushup1 := func(root, lson, rson *node) {
			if lson.v1 < rson.v1 {
				root.v1 = rson.v1
				root.idx1 = rson.idx1
			} else {
				root.v1 = lson.v1
				root.idx1 = lson.idx1
			}
			if lson.v2 < rson.v2 {
				root.v2 = lson.v2
				root.idx2 = lson.idx2
			} else {
				root.v2 = rson.v2
				root.idx2 = rson.idx2
			}
			if lson.v3 < rson.v3 {
				root.v3 = rson.v3
				root.l, root.r = rson.l, rson.r
			} else {
				root.v3 = lson.v3
				root.l, root.r = lson.l, lson.r
			}
			if root.v3 < rson.v1 - lson.v2 {
				root.v3 = rson.v1 - lson.v2
				root.l, root.r = lson.idx2, rson.idx1
			}
		}
		pushup := func(u int) {
			pushup1(&tr[u], &tr[u << 1], &tr[u << 1 | 1])
		}
		build = func(u, l, r int) {
			tr[u].s = s[r]
			tr[u].add = 0
			if r > 0 {
				tr[u].v1 = s[r] + b[r - 1]
				tr[u].idx1 = r
			} else {
				tr[u].v1 = -inf
				tr[u].idx1 = -1
			}
			if r < n {
				tr[u].v2 = s[r] - b[r]
				tr[u].idx2 = r
			} else {
				tr[u].v2 = inf
				tr[u].idx2 = -1
			}
			tr[u].v3 = tr[u].v1 - tr[u].v2
			tr[u].l, tr[u].r = tr[u].idx2, tr[u].idx1
			// Fprintln(out, "u, l, r, s, add, v1, v2, v3\n", u, l, r, tr[u].s, tr[u].add, tr[u].v1, tr[u].v2, tr[u].v3)
			if l == r {
				return
			}
			mid := (l + r) / 2
			build(u << 1, l, mid)
			build(u << 1 | 1, mid + 1, r)
			pushup(u)
		}
		pushdown := func(u int) {
			if tr[u].add != 0 {
				tr[u << 1].s += tr[u].add
				tr[u << 1].add += tr[u].add
				tr[u << 1 | 1].s += tr[u].add
				tr[u << 1 | 1].add += tr[u].add
				tr[u].add = 0
			}
		}
		update = func(u, l, r, L, R, v int) {
			if L <= l && r <= R {
				tr[u].s += v
				tr[u].add += v
				return
			}
			pushdown(u)
			mid := (l + r) / 2
			if L <= mid {
				update(u << 1, l, mid, L, R, v)
			}
			if R > mid {
				update(u << 1 | 1, mid + 1, r, L, R, v)
			}
			pushup(u)
		}
		change1 = func(u, l, r, x int) {
			if l == r {
				tr[u].v1 = tr[u].s + b[r - 1]
				tr[u].v3 = tr[u].v1 - tr[u].v2
				return
			}
			pushdown(u)
			mid := (l + r) / 2
			if x <= mid {
				change1(u << 1, l, mid, x)
			} else {
				change1(u << 1 | 1, mid + 1, r, x)
			}
			pushup(u)
		}
		change2 = func(u, l, r, x int) {
			if l == r {
				tr[u].v2 = tr[u].s - b[r]
				tr[u].v3 = tr[u].v1 - tr[u].v2
				return
			}
			pushdown(u)
			mid := (l + r) / 2
			if x <= mid {
				change2(u << 1, l, mid, x)
			} else {
				change2(u << 1 | 1, mid + 1, r, x)
			}
			pushup(u)
		}
		query = func(u, l, r, L, R int) node {
			if L <= l && r <= R {
				return tr[u]
			}
			pushdown(u)
			mid := (l + r) / 2
			if R <= mid {
				return query(u << 1, l, mid, L, R)
			}
			if L > mid {
				return query(u << 1 | 1, mid + 1, r, L, R)
			}
			var res node
			res1 := query(u << 1, l, mid, L, R)
			res2 := query(u << 1 | 1, l, mid, L, R)
			pushup1(&res, &res1, &res2)
			return res
		}
		build(1, 0, n)

		q := readLine()[0]
		for i := 0; i < q; i++ {
			tmp := readLine()
			if tmp[0] == 1 {
				p, x := tmp[1], tmp[2]
				update(1, 0, n, p, n, -a[p - 1] + x)
				a[p - 1] = x
			} else if tmp[0] == 2 {
				p, x := tmp[1], tmp[2]
				b[p - 1] = x
				change1(1, 0, n, p)
				change2(1, 0, n, p - 1)
			} else {
				l, r := tmp[1], tmp[2]
				mid := (l + r) / 2
				res := query(1, 0, n, l, r)
				res1 := query(1, 0, n, l, mid)
				res2 := query(1, 0, n, mid + 1, r)
				if res.v3 == res1.v3 || res.v3 == res2.v3 {
					Fprintln(out, res1.v3 + res2.v3)
				}
			}
		}
	}
}

func main() { solve(os.Stdin, os.Stdout) }
